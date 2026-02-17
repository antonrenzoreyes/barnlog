// Package config loads and validates runtime configuration for the backend.
package config

import (
	"fmt"
	"log/slog"
	"os"
	"strings"
	"time"
)

// Config contains server and infrastructure settings sourced from environment variables.
type Config struct {
	Env             string
	HTTPAddr        string
	LogLevel        slog.Level
	ShutdownTimeout time.Duration
}

// LoadFromEnv builds Config from environment variables and defaults.
func LoadFromEnv() (Config, error) {
	cfg := Config{
		Env:             getenv("BARNLOG_ENV", "dev"),
		HTTPAddr:        getenv("BARNLOG_HTTP_ADDR", ":8080"),
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
