// Package main boots the backend HTTP server and optional database migrations.
package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"barnlog/backend/docs"
	"barnlog/backend/internal/adapters/httpapi"
	"barnlog/backend/internal/infrastructure/config"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	gomigrate "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// @title Barnlog Backend API
// @version 1.0
// @description Barnlog backend HTTP API.
// @servers.url http://localhost:8080
func main() {
	if err := run(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "server failed: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	cfg, err := config.LoadFromEnv()
	if err != nil {
		return fmt.Errorf("load config: %w", err)
	}

	logger := newLogger(cfg)
	if cfg.AutoMigrate {
		if err := runMigrations(logger, cfg); err != nil {
			return err
		}
	} else {
		logger.Info("auto migration disabled")
	}
	srv := newHTTPServer(cfg, buildRouter(cfg, logger))

	logger.Info(
		"http server starting",
		slog.String("addr", cfg.HTTPAddr),
		slog.String("env", cfg.Env),
		slog.Any("log_level", cfg.LogLevel),
	)

	errCh := make(chan error, 1)
	go func() {
		if serveErr := srv.ListenAndServe(); serveErr != nil && !errors.Is(serveErr, http.ErrServerClosed) {
			errCh <- serveErr
		}
	}()

	if err := waitForShutdownSignalOrServerError(ctx, errCh); err != nil {
		return err
	}

	if err := shutdownServer(cfg, srv); err != nil {
		return err
	}

	logger.Info("server stopped")
	return nil
}

func buildRouter(cfg config.Config, logger *slog.Logger) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(30 * time.Second))
	r.Get("/swagger/openapi.json", openAPIDoc)
	r.Mount("/", httpapi.Routes(httpapi.RouteDeps{
		Logger:        logger,
		PhotoStoreDir: cfg.PhotoDir,
	}))
	return r
}

func openAPIDoc(w http.ResponseWriter, _ *http.Request) {
	doc := docs.SwaggerInfo.ReadDoc()

	var payload map[string]any
	if err := json.Unmarshal([]byte(doc), &payload); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	if version, ok := payload["openapi"].(string); ok && strings.HasPrefix(version, "3.1.") {
		payload["openapi"] = "3.0.3"
	}
	normalizeUploadPhotoRequestBody(payload)
	body, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_, _ = w.Write(body)
}

func normalizeUploadPhotoRequestBody(payload map[string]any) {
	paths, ok := payload["paths"].(map[string]any)
	if !ok {
		return
	}
	uploads, ok := paths["/uploads/photos"].(map[string]any)
	if !ok {
		return
	}
	post, ok := uploads["post"].(map[string]any)
	if !ok {
		return
	}
	requestBody, ok := post["requestBody"].(map[string]any)
	if !ok {
		return
	}
	content, ok := requestBody["content"].(map[string]any)
	if !ok {
		return
	}
	content["multipart/form-data"] = map[string]any{
		"schema": map[string]any{
			"type":     "object",
			"required": []string{"photo"},
			"properties": map[string]any{
				"photo": map[string]any{
					"type":        "string",
					"format":      "binary",
					"description": "Photo file to upload",
				},
			},
		},
		"example": map[string]any{
			"photo": "(binary file)",
		},
	}
	delete(content, "application/x-www-form-urlencoded")
}

func newHTTPServer(cfg config.Config, handler http.Handler) *http.Server {
	return &http.Server{
		Addr:              cfg.HTTPAddr,
		Handler:           handler,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       60 * time.Second,
	}
}

func waitForShutdownSignalOrServerError(ctx context.Context, errCh <-chan error) error {
	select {
	case <-ctx.Done():
		return nil
	case serveErr := <-errCh:
		return fmt.Errorf("http server failed: %w", serveErr)
	}
}

func shutdownServer(cfg config.Config, srv *http.Server) error {
	shutdownCtx, cancel := context.WithTimeout(context.Background(), cfg.ShutdownTimeout)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		return fmt.Errorf("graceful shutdown failed: %w", err)
	}
	return nil
}

func newLogger(cfg config.Config) *slog.Logger {
	handlerOpts := &slog.HandlerOptions{Level: cfg.LogLevel}
	if cfg.Env == "local" || cfg.Env == "dev" {
		return slog.New(slog.NewTextHandler(os.Stdout, handlerOpts))
	}
	return slog.New(slog.NewJSONHandler(os.Stdout, handlerOpts))
}

func runMigrations(logger *slog.Logger, cfg config.Config) error {
	dbPath, err := filepath.Abs(cfg.DBPath)
	if err != nil {
		return fmt.Errorf("resolve db path: %w", err)
	}
	migrationsPath, err := filepath.Abs(cfg.MigrationsPath)
	if err != nil {
		return fmt.Errorf("resolve migrations path: %w", err)
	}

	if err := os.MkdirAll(filepath.Dir(dbPath), 0o750); err != nil {
		return fmt.Errorf("create db directory: %w", err)
	}

	dbURL := (&url.URL{Scheme: "sqlite3", Path: dbPath}).String()
	srcURL := (&url.URL{Scheme: "file", Path: migrationsPath}).String()

	logger.Info("running migrations", slog.String("database", dbURL), slog.String("source", srcURL))

	m, err := gomigrate.New(srcURL, dbURL)
	if err != nil {
		return fmt.Errorf("initialize migrate: %w", err)
	}
	defer func() {
		srcErr, dbErr := m.Close()
		if srcErr != nil {
			logger.Warn("close migration source", slog.Any("error", srcErr))
		}
		if dbErr != nil {
			logger.Warn("close migration db", slog.Any("error", dbErr))
		}
	}()

	if err := m.Up(); err != nil {
		if errors.Is(err, gomigrate.ErrNoChange) {
			logger.Info("no pending migrations")
			return nil
		}
		return fmt.Errorf("run migrations: %w", err)
	}

	logger.Info("migrations applied")
	return nil
}
