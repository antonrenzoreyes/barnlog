// Package config loads and validates runtime configuration for the backend.
package config

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"strings"
	"time"
)

// Config contains server and infrastructure settings sourced from environment variables.
type Config struct {
	Env             string
	HTTPAddr        string
	DBPath          string
	MigrationsPath  string
	AutoMigrate     bool
	LogLevel        slog.Level
	ShutdownTimeout time.Duration
}

// LoadFromEnv builds Config from environment variables and defaults.
func LoadFromEnv() (Config, error) {
	cfg := Config{
		Env:             getenv("BARNLOG_ENV", "dev"),
		HTTPAddr:        getenv("BARNLOG_HTTP_ADDR", ":8080"),
		DBPath:          getenv("BARNLOG_DB_PATH", "backend/db/dev.sqlite3"),
		MigrationsPath:  getenv("BARNLOG_MIGRATIONS_PATH", "backend/db/migrations"),
		AutoMigrate:     true,
		ShutdownTimeout: 10 * time.Second,
	}

	logLevel, err := parseLogLevel(getenv("BARNLOG_LOG_LEVEL", "info"))
	if err != nil {
		return Config{}, err
	}
	cfg.LogLevel = logLevel

	if timeout := strings.TrimSpace(os.Getenv("BARNLOG_SHUTDOWN_TIMEOUT")); timeout != "" {
		dur, err := time.ParseDuration(timeout)
		if err != nil {
			return Config{}, fmt.Errorf("parse BARNLOG_SHUTDOWN_TIMEOUT: %w", err)
		}
		cfg.ShutdownTimeout = dur
	}

	if raw := strings.TrimSpace(os.Getenv("BARNLOG_AUTO_MIGRATE")); raw != "" {
		enabled, err := strconv.ParseBool(raw)
		if err != nil {
			return Config{}, fmt.Errorf("parse BARNLOG_AUTO_MIGRATE: %w", err)
		}
		cfg.AutoMigrate = enabled
	}

	return cfg, nil
}

func parseLogLevel(raw string) (slog.Level, error) {
	var level slog.Level
	if err := level.UnmarshalText([]byte(strings.TrimSpace(strings.ToLower(raw)))); err != nil {
		return 0, fmt.Errorf("parse BARNLOG_LOG_LEVEL: %w", err)
	}
	return level, nil
}

func getenv(key, fallback string) string {
	if value := strings.TrimSpace(os.Getenv(key)); value != "" {
		return value
	}
	return fallback
}
