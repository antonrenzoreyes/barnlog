package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"barnlog/backend/internal/adapters/httpapi"
	"barnlog/backend/internal/infrastructure/config"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "server failed: %v\n", err)
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
	srv := newHTTPServer(cfg, buildRouter(logger))

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

func buildRouter(logger *slog.Logger) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(30 * time.Second))
	r.Mount("/", httpapi.Routes(logger))
	return r
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
