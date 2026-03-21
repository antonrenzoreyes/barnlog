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

func TestUploadAnimalPhoto(t *testing.T) {
	t.Parallel()

	fileDir := t.TempDir()
	router := Routes(RouteDeps{
		Logger:       testLogger(),
		FileStoreDir: fileDir,
	})

	t.Run("created", func(t *testing.T) {
		t.Parallel()

		rec := performUpload(t, router, "animal.png", samplePNGBytes())
		assertJSONStatus(t, rec, http.StatusCreated)

		var payload map[string]any
		decodeJSON(t, rec, &payload)

		fileID, ok := payload["file_id"].(string)
		if !ok || fileID == "" {
			t.Fatalf("expected non-empty file_id, got %#v", payload["file_id"])
		}

		fileName, ok := payload["file_name"].(string)
		if !ok || fileName != "animal.png" {
			t.Fatalf("expected file_name=animal.png, got %#v", payload["file_name"])
		}

		contentType, ok := payload["content_type"].(string)
		if !ok || contentType != "image/png" {
			t.Fatalf("expected content_type=image/png, got %#v", payload["content_type"])
		}

		if got := payload["size_bytes"]; got != float64(len(samplePNGBytes())) {
			t.Fatalf("expected size_bytes=%d, got %#v", len(samplePNGBytes()), payload["size_bytes"])
		}

		if _, err := os.Stat(filepath.Join(fileDir, fileID)); err != nil {
			t.Fatalf("expected saved file: %v", err)
		}
	})

	t.Run("file required", func(t *testing.T) {
		t.Parallel()

		rec := performUploadWithoutFile(t, router)
		assertJSONStatus(t, rec, http.StatusBadRequest)

		var payload map[string]any
		decodeJSON(t, rec, &payload)
		if payload["error"] != "file_required" {
			t.Fatalf("expected error=file_required, got %#v", payload["error"])
		}
	})

	t.Run("non-image file is rejected", func(t *testing.T) {
		t.Parallel()

		rec := performUpload(t, router, "animal.txt", []byte("not-an-image"))
		assertJSONStatus(t, rec, http.StatusBadRequest)

		var payload map[string]any
		decodeJSON(t, rec, &payload)
		if payload["error"] != "unsupported_file_type" {
			t.Fatalf("expected error=unsupported_file_type, got %#v", payload["error"])
		}
	})

	t.Run("multiple files are rejected", func(t *testing.T) {
		t.Parallel()

		rec := performUploadWithTwoFiles(t, router, samplePNGBytes(), samplePNGBytes())
		assertJSONStatus(t, rec, http.StatusBadRequest)

		var payload map[string]any
		decodeJSON(t, rec, &payload)
		if payload["error"] != "multiple_files_not_allowed" {
			t.Fatalf("expected error=multiple_files_not_allowed, got %#v", payload["error"])
		}
	})

	t.Run("wrong form field name is rejected", func(t *testing.T) {
		t.Parallel()

		rec := performUploadWithFieldName(t, router, "photo", "animal.png", samplePNGBytes())
		assertJSONStatus(t, rec, http.StatusBadRequest)

		var payload map[string]any
		decodeJSON(t, rec, &payload)
		if payload["error"] != "file_required" {
			t.Fatalf("expected error=file_required, got %#v", payload["error"])
		}
	})

	t.Run("request body too large is rejected", func(t *testing.T) {
		t.Parallel()

		oversized := append(samplePNGBytes(), bytes.Repeat([]byte{0x00}, int(maxAnimalPhotoSizeBytes+maxMultipartOverheadBytes)+1)...)
		rec := performUpload(t, router, "huge.png", oversized)
		assertJSONStatus(t, rec, http.StatusRequestEntityTooLarge)

		var payload map[string]any
		decodeJSON(t, rec, &payload)
		if payload["error"] != "file_too_large" {
			t.Fatalf("expected error=file_too_large, got %#v", payload["error"])
		}
	})
}

func performUpload(t *testing.T, router http.Handler, filename string, content []byte) *httptest.ResponseRecorder {
	t.Helper()

	body, contentType := buildMultipartBody(t, "file", filename, content)
	req := httptest.NewRequest(http.MethodPost, "/uploads/animal-photos", body)
	req.Header.Set("Content-Type", contentType)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)
	return rec
}

func performUploadWithoutFile(t *testing.T, router http.Handler) *httptest.ResponseRecorder {
	t.Helper()

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	if err := writer.Close(); err != nil {
		t.Fatalf("close multipart writer: %v", err)
	}

	req := httptest.NewRequest(http.MethodPost, "/uploads/animal-photos", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)
	return rec
}

func performUploadWithTwoFiles(t *testing.T, router http.Handler, first, second []byte) *httptest.ResponseRecorder {
	t.Helper()

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	firstPart, err := writer.CreateFormFile("file", "first.png")
	if err != nil {
		t.Fatalf("create first form file: %v", err)
	}
	if _, err := io.Copy(firstPart, bytes.NewReader(first)); err != nil {
		t.Fatalf("write first form file: %v", err)
	}

	secondPart, err := writer.CreateFormFile("file", "second.png")
	if err != nil {
		t.Fatalf("create second form file: %v", err)
	}
	if _, err := io.Copy(secondPart, bytes.NewReader(second)); err != nil {
		t.Fatalf("write second form file: %v", err)
	}

	if err := writer.Close(); err != nil {
		t.Fatalf("close multipart writer: %v", err)
	}

	req := httptest.NewRequest(http.MethodPost, "/uploads/animal-photos", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)
	return rec
}

func performUploadWithFieldName(
	t *testing.T,
	router http.Handler,
	fieldName string,
	filename string,
	content []byte,
) *httptest.ResponseRecorder {
	t.Helper()

	body, contentType := buildMultipartBody(t, fieldName, filename, content)
	req := httptest.NewRequest(http.MethodPost, "/uploads/animal-photos", body)
	req.Header.Set("Content-Type", contentType)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)
	return rec
}

func buildMultipartBody(t *testing.T, fieldName, filename string, content []byte) (*bytes.Buffer, string) {
	t.Helper()

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(fieldName, filename)
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
