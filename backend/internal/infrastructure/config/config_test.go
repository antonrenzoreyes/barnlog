package config

import (
	"log/slog"
	"strings"
	"testing"
	"time"
)

func TestLoadFromEnvDefaults(t *testing.T) {
	t.Setenv("BARNLOG_ENV", "")
	t.Setenv("BARNLOG_HTTP_ADDR", "")
	t.Setenv("BARNLOG_LOG_LEVEL", "")
	t.Setenv("BARNLOG_SHUTDOWN_TIMEOUT", "")

	cfg, err := LoadFromEnv()
	if err != nil {
		t.Fatalf("LoadFromEnv() error = %v", err)
	}

	if cfg.Env != "dev" {
		t.Fatalf("expected Env=dev, got %q", cfg.Env)
	}
	if cfg.HTTPAddr != ":8080" {
		t.Fatalf("expected HTTPAddr=:8080, got %q", cfg.HTTPAddr)
	}
	if cfg.LogLevel != slog.LevelInfo {
		t.Fatalf("expected LogLevel=info, got %v", cfg.LogLevel)
	}
	if cfg.ShutdownTimeout != 10*time.Second {
		t.Fatalf("expected ShutdownTimeout=10s, got %s", cfg.ShutdownTimeout)
	}
}

func TestLoadFromEnvCustomValues(t *testing.T) {
	t.Setenv("BARNLOG_ENV", "prod")
	t.Setenv("BARNLOG_HTTP_ADDR", ":9090")
	t.Setenv("BARNLOG_LOG_LEVEL", "debug")
	t.Setenv("BARNLOG_SHUTDOWN_TIMEOUT", "3s")

	cfg, err := LoadFromEnv()
	if err != nil {
		t.Fatalf("LoadFromEnv() error = %v", err)
	}

	if cfg.Env != "prod" {
		t.Fatalf("expected Env=prod, got %q", cfg.Env)
	}
	if cfg.HTTPAddr != ":9090" {
		t.Fatalf("expected HTTPAddr=:9090, got %q", cfg.HTTPAddr)
	}
	if cfg.LogLevel != slog.LevelDebug {
		t.Fatalf("expected LogLevel=debug, got %v", cfg.LogLevel)
	}
	if cfg.ShutdownTimeout != 3*time.Second {
		t.Fatalf("expected ShutdownTimeout=3s, got %s", cfg.ShutdownTimeout)
	}
}

func TestLoadFromEnvInvalidLogLevel(t *testing.T) {
	t.Setenv("BARNLOG_LOG_LEVEL", "not-a-level")

	_, err := LoadFromEnv()
	if err == nil {
		t.Fatalf("expected error for invalid log level")
	}
	if !strings.Contains(err.Error(), "BARNLOG_LOG_LEVEL") {
		t.Fatalf("expected BARNLOG_LOG_LEVEL in error, got %q", err.Error())
	}
}

func TestLoadFromEnvInvalidShutdownTimeout(t *testing.T) {
	t.Setenv("BARNLOG_SHUTDOWN_TIMEOUT", "definitely-not-a-duration")

	_, err := LoadFromEnv()
	if err == nil {
		t.Fatalf("expected error for invalid shutdown timeout")
	}
	if !strings.Contains(err.Error(), "BARNLOG_SHUTDOWN_TIMEOUT") {
		t.Fatalf("expected BARNLOG_SHUTDOWN_TIMEOUT in error, got %q", err.Error())
	}
}
