package httpapi

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestWriteJSONSuccess(t *testing.T) {
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
