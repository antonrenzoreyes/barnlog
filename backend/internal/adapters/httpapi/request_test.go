package httpapi

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestDecodeJSONRequest(t *testing.T) {
	t.Parallel()

	t.Run("request body too large", func(t *testing.T) {
		t.Parallel()

		oversizedName := strings.Repeat("N", maxJSONBodyBytes)
		body := `{"name":"` + oversizedName + `","species":"goat"}`
		req := httptest.NewRequest(http.MethodPost, "/animals", strings.NewReader(body))
		req.Header.Set("Content-Type", jsonContentType)
		rec := httptest.NewRecorder()

		var dst map[string]any
		status, code, ok := decodeJSONRequest(rec, req, &dst)
		if ok {
			t.Fatalf("expected decode to fail for oversized body")
		}
		if status != http.StatusRequestEntityTooLarge {
			t.Fatalf("expected status=%d, got %d", http.StatusRequestEntityTooLarge, status)
		}
		if code != "request_too_large" {
			t.Fatalf("expected code=request_too_large, got %q", code)
		}
	})

	t.Run("valid json", func(t *testing.T) {
		t.Parallel()

		req := httptest.NewRequest(http.MethodPost, "/animals", strings.NewReader(`{"name":"Nanny","species":"goat"}`))
		req.Header.Set("Content-Type", jsonContentType)
		rec := httptest.NewRecorder()

		var dst map[string]any
		status, code, ok := decodeJSONRequest(rec, req, &dst)
		if !ok {
			t.Fatalf("expected decode to succeed, got status=%d code=%q", status, code)
		}
		if dst["name"] != "Nanny" {
			t.Fatalf("expected name=Nanny, got %#v", dst["name"])
		}
	})
}
