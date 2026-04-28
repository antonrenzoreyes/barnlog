package main

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"barnlog/backend/internal/infrastructure/config"
)

func TestBuildRouterServesOpenAPIDoc(t *testing.T) {
	t.Parallel()

	router := buildRouter(config.Config{FileDir: t.TempDir()}, testLogger())
	request := httptest.NewRequest(http.MethodGet, "/swagger/openapi.json", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, recorder.Code)
	}
	if got := recorder.Header().Get("Content-Type"); got != "application/json; charset=utf-8" {
		t.Fatalf("expected content-type application/json; charset=utf-8, got %q", got)
	}

	var payload map[string]any
	if err := json.Unmarshal(recorder.Body.Bytes(), &payload); err != nil {
		t.Fatalf("expected OpenAPI JSON body, got error: %v", err)
	}
	if _, ok := payload["openapi"].(string); !ok {
		t.Fatalf("expected openapi version in response body, got %v", payload["openapi"])
	}
}

func testLogger() *slog.Logger {
	return slog.New(slog.NewTextHandler(io.Discard, nil))
}
