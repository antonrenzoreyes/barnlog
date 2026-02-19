package httpapi

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestUploadPhoto(t *testing.T) {
	t.Parallel()

	photoDir := t.TempDir()
	router := Routes(RouteDeps{
		Logger:        testLogger(),
		PhotoStoreDir: photoDir,
	})

	t.Run("created", func(t *testing.T) {
		t.Parallel()

		rec := performUpload(t, router, "animal.png", samplePNGBytes())
		assertJSONStatus(t, rec, http.StatusCreated)

		var payload map[string]any
		decodeJSON(t, rec, &payload)

		photoID, ok := payload["photo_id"].(string)
		if !ok || photoID == "" {
			t.Fatalf("expected non-empty photo_id, got %#v", payload["photo_id"])
		}

		contentType, ok := payload["content_type"].(string)
		if !ok || contentType == "" {
			t.Fatalf("expected non-empty content_type, got %#v", payload["content_type"])
		}

		if got := payload["size_bytes"]; got == nil {
			t.Fatalf("expected size_bytes, got nil")
		}

		if _, err := os.Stat(filepath.Join(photoDir, photoID)); err != nil {
			t.Fatalf("expected saved photo file: %v", err)
		}
	})

	t.Run("photo required", func(t *testing.T) {
		t.Parallel()

		rec := performUploadWithoutPhoto(t, router)
		assertJSONStatus(t, rec, http.StatusBadRequest)

		var payload map[string]any
		decodeJSON(t, rec, &payload)
		if payload["error"] != "photo_required" {
			t.Fatalf("expected error=photo_required, got %#v", payload["error"])
		}
	})

	t.Run("non-image file is rejected", func(t *testing.T) {
		t.Parallel()

		rec := performUpload(t, router, "animal.txt", []byte("not-an-image"))
		assertJSONStatus(t, rec, http.StatusBadRequest)

		var payload map[string]any
		decodeJSON(t, rec, &payload)
		if payload["error"] != "unsupported_photo_type" {
			t.Fatalf("expected error=unsupported_photo_type, got %#v", payload["error"])
		}
	})

	t.Run("request body too large is rejected", func(t *testing.T) {
		t.Parallel()

		oversized := append(samplePNGBytes(), bytes.Repeat([]byte{0x00}, int(maxUploadRequestBytes)+1)...)
		rec := performUpload(t, router, "huge.png", oversized)
		assertJSONStatus(t, rec, http.StatusRequestEntityTooLarge)

		var payload map[string]any
		decodeJSON(t, rec, &payload)
		if payload["error"] != "photo_too_large" {
			t.Fatalf("expected error=photo_too_large, got %#v", payload["error"])
		}
	})
}

func performUpload(t *testing.T, router http.Handler, filename string, content []byte) *httptest.ResponseRecorder {
	t.Helper()

	body, contentType := buildMultipartBody(t, filename, content)
	req := httptest.NewRequest(http.MethodPost, "/uploads/photos", body)
	req.Header.Set("Content-Type", contentType)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)
	return rec
}

func performUploadWithoutPhoto(t *testing.T, router http.Handler) *httptest.ResponseRecorder {
	t.Helper()

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	if err := writer.Close(); err != nil {
		t.Fatalf("close multipart writer: %v", err)
	}

	req := httptest.NewRequest(http.MethodPost, "/uploads/photos", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)
	return rec
}

func buildMultipartBody(t *testing.T, filename string, content []byte) (*bytes.Buffer, string) {
	t.Helper()

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(photoFieldName, filename)
	if err != nil {
		t.Fatalf("create form file: %v", err)
	}
	if _, err := io.Copy(part, bytes.NewReader(content)); err != nil {
		t.Fatalf("write form file: %v", err)
	}
	if err := writer.Close(); err != nil {
		t.Fatalf("close multipart writer: %v", err)
	}

	return body, writer.FormDataContentType()
}

func samplePNGBytes() []byte {
	return []byte{
		0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A,
		0x00, 0x00, 0x00, 0x0D, 0x49, 0x48, 0x44, 0x52,
		0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x01,
		0x08, 0x02, 0x00, 0x00, 0x00, 0x90, 0x77, 0x53,
		0xDE, 0x00, 0x00, 0x00, 0x0C, 0x49, 0x44, 0x41,
		0x54, 0x08, 0xD7, 0x63, 0xF8, 0xCF, 0xC0, 0x00,
		0x00, 0x03, 0x01, 0x01, 0x00, 0x18, 0xDD, 0x8D,
		0x18, 0x00, 0x00, 0x00, 0x00, 0x49, 0x45, 0x4E,
		0x44, 0xAE, 0x42, 0x60, 0x82,
	}
}
