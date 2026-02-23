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

// uploadPhoto godoc
//
// @Summary Upload photo
// @Description Uploads a photo file and returns a generated photo_id.
// @Tags uploads
// @Accept multipart/form-data
// @Produce json
// @Param photo formData file true "Photo file"
// @Success 201 {object} uploadPhotoResponse
// @Failure 400 {object} errorResponse
// @Failure 413 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /uploads/photos [post]
func (h uploadHandlers) uploadPhoto(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, maxUploadRequestBytes)

	if err := r.ParseMultipartForm(maxPhotoSizeBytes + 512); err != nil {
		if isRequestTooLarge(err) {
			writeError(w, http.StatusRequestEntityTooLarge, "photo_too_large")
			return
		}

		writeError(w, http.StatusBadRequest, "invalid_multipart")
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
		writeError(w, http.StatusInternalServerError, "internal_error")
		return
	}

	file, _, err := r.FormFile(photoFieldName)
	if err != nil {
		if isRequestTooLarge(err) {
			writeError(w, http.StatusRequestEntityTooLarge, "photo_too_large")
			return
		}

		writeError(w, http.StatusBadRequest, "photo_required")
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
		writeError(w, http.StatusBadRequest, "invalid_photo")
		return
	}
	if sniffBytesRead == 0 {
		writeError(w, http.StatusBadRequest, "invalid_photo")
		return
	}

	sniffBuffer = sniffBuffer[:sniffBytesRead]
	contentType := http.DetectContentType(sniffBuffer)
	if _, ok := allowedPhotoContentTypes[contentType]; !ok {
		writeError(w, http.StatusBadRequest, "unsupported_photo_type")
		return
	}

	photoID, totalBytes, err := h.photoStore.Save(
		r.Context(),
		io.MultiReader(bytes.NewReader(sniffBuffer), file),
		maxPhotoSizeBytes,
	)
	if err != nil {
		if errors.Is(err, errPhotoTooLarge) {
			writeError(w, http.StatusRequestEntityTooLarge, "photo_too_large")
			return
		}

		h.logger.Error("save photo", slog.Any("error", err))
		writeError(w, http.StatusInternalServerError, "internal_error")
		return
	}

	writeJSON(w, http.StatusCreated, uploadPhotoResponse{
		PhotoID:     photoID,
		ContentType: contentType,
		SizeBytes:   totalBytes,
	})
}

func isRequestTooLarge(err error) bool {
	var maxBytesErr *http.MaxBytesError
	return errors.As(err, &maxBytesErr)
}
