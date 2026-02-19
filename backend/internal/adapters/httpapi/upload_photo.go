package httpapi

import (
	"bytes"
	"errors"
	"io"
	"log/slog"
	"net/http"
)

const (
	maxPhotoSizeBytes               = 10 << 20 // 10 MiB
	maxMultipartOverheadBytes       = 1 << 20  // 1 MiB for multipart envelope
	maxUploadRequestBytes     int64 = maxPhotoSizeBytes + maxMultipartOverheadBytes
	photoFieldName                  = "photo"
)

var allowedPhotoContentTypes = map[string]struct{}{
	"image/jpeg": {},
	"image/png":  {},
	"image/webp": {},
	"image/gif":  {},
}

func (h uploadHandlers) uploadPhoto(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, maxUploadRequestBytes)

	if err := r.ParseMultipartForm(maxPhotoSizeBytes + 512); err != nil {
		if isRequestTooLarge(err) {
			writeJSON(w, http.StatusRequestEntityTooLarge, map[string]any{"error": "photo_too_large"})
			return
		}

		writeJSON(w, http.StatusBadRequest, map[string]any{"error": "invalid_multipart"})
		return
	}

	if r.MultipartForm != nil {
		defer func() {
			if err := r.MultipartForm.RemoveAll(); err != nil {
				h.logger.Warn("remove multipart temp files", slog.Any("error", err))
			}
		}()
	}

	if h.photoStore == nil {
		h.logger.Error("photo store is nil")
		writeJSON(w, http.StatusInternalServerError, map[string]any{"error": "internal_error"})
		return
	}

	file, _, err := r.FormFile(photoFieldName)
	if err != nil {
		if isRequestTooLarge(err) {
			writeJSON(w, http.StatusRequestEntityTooLarge, map[string]any{"error": "photo_too_large"})
			return
		}

		writeJSON(w, http.StatusBadRequest, map[string]any{"error": "photo_required"})
		return
	}
	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			h.logger.Warn("close uploaded file", slog.Any("error", closeErr))
		}
	}()

	sniffBuffer := make([]byte, 512)
	sniffBytesRead, sniffErr := io.ReadFull(file, sniffBuffer)
	if sniffErr != nil && !errors.Is(sniffErr, io.EOF) && !errors.Is(sniffErr, io.ErrUnexpectedEOF) {
		writeJSON(w, http.StatusBadRequest, map[string]any{"error": "invalid_photo"})
		return
	}
	if sniffBytesRead == 0 {
		writeJSON(w, http.StatusBadRequest, map[string]any{"error": "invalid_photo"})
		return
	}

	sniffBuffer = sniffBuffer[:sniffBytesRead]
	contentType := http.DetectContentType(sniffBuffer)
	if _, ok := allowedPhotoContentTypes[contentType]; !ok {
		writeJSON(w, http.StatusBadRequest, map[string]any{"error": "unsupported_photo_type"})
		return
	}

	photoID, totalBytes, err := h.photoStore.Save(
		r.Context(),
		io.MultiReader(bytes.NewReader(sniffBuffer), file),
		maxPhotoSizeBytes,
	)
	if err != nil {
		if errors.Is(err, errPhotoTooLarge) {
			writeJSON(w, http.StatusRequestEntityTooLarge, map[string]any{"error": "photo_too_large"})
			return
		}

		h.logger.Error("save photo", slog.Any("error", err))
		writeJSON(w, http.StatusInternalServerError, map[string]any{"error": "internal_error"})
		return
	}

	writeJSON(w, http.StatusCreated, map[string]any{
		"photo_id":     photoID,
		"content_type": contentType,
		"size_bytes":   totalBytes,
	})
}

func isRequestTooLarge(err error) bool {
	var maxBytesErr *http.MaxBytesError
	return errors.As(err, &maxBytesErr)
}
