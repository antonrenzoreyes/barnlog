package httpapi

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestWriteJSONSuccess(t *testing.T) {
	t.Parallel()

	rec := httptest.NewRecorder()
	writeJSON(rec, http.StatusCreated, map[string]any{"created": true})

	if rec.Code != http.StatusCreated {
		t.Fatalf("expected status %d, got %d", http.StatusCreated, rec.Code)
	}
	if got := rec.Header().Get("Content-Type"); got != "application/json" {
		t.Fatalf("expected content-type application/json, got %q", got)
	}
	body := strings.TrimSpace(rec.Body.String())
	if body != `{"created":true}` {
		t.Fatalf("unexpected body: %q", body)
	}
}

func TestWriteJSONMarshalError(t *testing.T) {
	t.Parallel()

	rec := httptest.NewRecorder()

	// channels cannot be marshaled by encoding/json
	writeJSON(rec, http.StatusOK, map[string]any{"invalid": make(chan int)})

	if rec.Code != http.StatusInternalServerError {
		t.Fatalf("expected status %d, got %d", http.StatusInternalServerError, rec.Code)
	}
	if got := strings.TrimSpace(rec.Body.String()); got != http.StatusText(http.StatusInternalServerError) {
		t.Fatalf("expected body %q, got %q", http.StatusText(http.StatusInternalServerError), got)
	}
}

func TestWriteError_UsesAllowedCode(t *testing.T) {
	t.Parallel()

	rec := httptest.NewRecorder()
	writeError(rec, http.StatusBadRequest, "not_found")

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, rec.Code)
	}

	var payload map[string]any
	if err := json.Unmarshal(rec.Body.Bytes(), &payload); err != nil {
		t.Fatalf("decode payload: %v", err)
	}
	if payload["error"] != "not_found" {
		t.Fatalf("expected error=not_found, got %#v", payload["error"])
	}
}

func TestWriteError_NormalizesUnknownCode(t *testing.T) {
	t.Parallel()

	rec := httptest.NewRecorder()
	writeError(rec, http.StatusBadRequest, "totally_unknown_code")
	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, rec.Code)
	}

	var payload map[string]any
	if err := json.Unmarshal(rec.Body.Bytes(), &payload); err != nil {
		t.Fatalf("decode payload: %v", err)
	}
	if payload["error"] != "internal_error" {
		t.Fatalf("expected error=internal_error, got %#v", payload["error"])
	}
}
