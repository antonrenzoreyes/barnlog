package httpapi

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRoutes(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		method         string
		path           string
		expectedStatus int
		assertBody     func(t *testing.T, payload map[string]any)
	}{
		{
			name:           "healthz",
			method:         http.MethodGet,
			path:           "/healthz",
			expectedStatus: http.StatusOK,
			assertBody: func(t *testing.T, payload map[string]any) {
				t.Helper()
				if payload["status"] != "ok" {
					t.Fatalf("expected status=ok, got %#v", payload["status"])
				}
			},
		},
		{
			name:           "readyz",
			method:         http.MethodGet,
			path:           "/readyz",
			expectedStatus: http.StatusOK,
			assertBody: func(t *testing.T, payload map[string]any) {
				t.Helper()
				if payload["status"] != "ready" {
					t.Fatalf("expected status=ready, got %#v", payload["status"])
				}
				timestamp, ok := payload["timestamp"].(string)
				if !ok {
					t.Fatalf("expected timestamp string, got %#v", payload["timestamp"])
				}
				if _, err := time.Parse(time.RFC3339, timestamp); err != nil {
					t.Fatalf("expected RFC3339 timestamp, got %q: %v", timestamp, err)
				}
			},
		},
		{
			name:           "not found",
			method:         http.MethodGet,
			path:           "/does-not-exist",
			expectedStatus: http.StatusNotFound,
			assertBody: func(t *testing.T, payload map[string]any) {
				t.Helper()
				if payload["error"] != "not_found" {
					t.Fatalf("expected error=not_found, got %#v", payload["error"])
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			rec := performRequest(t, tt.method, tt.path)
			assertJSONStatus(t, rec, tt.expectedStatus)

			var payload map[string]any
			decodeJSON(t, rec, &payload)
			tt.assertBody(t, payload)
		})
	}
}

func performRequest(t *testing.T, method, path string) *httptest.ResponseRecorder {
	t.Helper()

	h := Routes(RouteDeps{
		Logger:        testLogger(),
		PhotoStoreDir: t.TempDir(),
	})
	req := httptest.NewRequest(method, path, nil)
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)
	return rec
}

func assertJSONStatus(t *testing.T, rec *httptest.ResponseRecorder, expectedStatus int) {
	t.Helper()

	if rec.Code != expectedStatus {
		t.Fatalf("expected status %d, got %d", expectedStatus, rec.Code)
	}
	if got := rec.Header().Get("Content-Type"); got != "application/json" {
		t.Fatalf("expected content-type application/json, got %q", got)
	}
}

func decodeJSON(t *testing.T, rec *httptest.ResponseRecorder, out any) {
	t.Helper()

	if err := json.Unmarshal(rec.Body.Bytes(), out); err != nil {
		t.Fatalf("unmarshal response: %v", err)
	}
}

func testLogger() *slog.Logger {
	return slog.New(slog.NewTextHandler(io.Discard, nil))
}
