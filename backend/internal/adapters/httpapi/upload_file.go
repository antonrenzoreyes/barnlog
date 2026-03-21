package httpapi

import (
	"bytes"
	"errors"
	"io"
	"log/slog"
	"mime/multipart"
	"net/http"
	"path"
	"strings"
)

const (
	maxAnimalPhotoSizeBytes   int64 = 10 << 20 // 10 MiB
	maxMultipartOverheadBytes int64 = 1 << 20  // 1 MiB for multipart envelope
	fileFieldName                   = "file"
)

var animalPhotoAllowedContentTypes = map[string]struct{}{
	"image/jpeg": {},
	"image/png":  {},
	"image/webp": {},
	"image/gif":  {},
}

type uploadPolicy struct {
	maxFileSizeBytes     int64
	maxFilesPerUpload    int
	allowedContentTypes  map[string]struct{}
	unsupportedTypeError string
}

var animalPhotoUploadPolicy = uploadPolicy{
	maxFileSizeBytes:     maxAnimalPhotoSizeBytes,
	maxFilesPerUpload:    1,
	allowedContentTypes:  animalPhotoAllowedContentTypes,
	unsupportedTypeError: "unsupported_file_type",
}

// uploadAnimalPhoto godoc
//
// @Summary Upload animal photo
// @Description Uploads an animal photo (max 10 MiB; allowed MIME types: image/jpeg, image/png, image/webp, image/gif) and returns a generated file_id.
// @Tags uploads
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "Animal photo file (max 10 MiB; allowed MIME types: image/jpeg, image/png, image/webp, image/gif)"
// @Success 201 {object} uploadFileResponse
// @Failure 400 {object} errorResponse "invalid_multipart | file_required | multiple_files_not_allowed | invalid_file | unsupported_file_type"
// @Failure 413 {object} errorResponse "file_too_large"
// @Failure 500 {object} errorResponse "internal_error"
// @Router /uploads/animal-photos [post]
func (h uploadHandlers) uploadAnimalPhoto(w http.ResponseWriter, r *http.Request) {
	h.uploadWithPolicy(w, r, animalPhotoUploadPolicy)
}

func (h uploadHandlers) uploadWithPolicy(w http.ResponseWriter, r *http.Request, policy uploadPolicy) {
	if h.fileStore == nil {
		h.logger.Error("file store is nil")
		writeError(w, http.StatusInternalServerError, "internal_error")
		return
	}

	maxUploadRequestBytes := policy.maxFileSizeBytes + maxMultipartOverheadBytes
	r.Body = http.MaxBytesReader(w, r.Body, maxUploadRequestBytes)

	if err := r.ParseMultipartForm(policy.maxFileSizeBytes + 512); err != nil {
		if isRequestTooLarge(err) {
			writeError(w, http.StatusRequestEntityTooLarge, "file_too_large")
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

	totalFiles := countMultipartFiles(r.MultipartForm.File)
	if totalFiles == 0 {
		writeError(w, http.StatusBadRequest, "file_required")
		return
	}
	if totalFiles > policy.maxFilesPerUpload {
		writeError(w, http.StatusBadRequest, "multiple_files_not_allowed")
		return
	}

	fileHeaders := r.MultipartForm.File[fileFieldName]
	if len(fileHeaders) != policy.maxFilesPerUpload {
		writeError(w, http.StatusBadRequest, "file_required")
		return
	}

	fileHeader := fileHeaders[0]
	file, err := fileHeader.Open()
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid_file")
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
		writeError(w, http.StatusBadRequest, "invalid_file")
		return
	}
	if sniffBytesRead == 0 {
		writeError(w, http.StatusBadRequest, "invalid_file")
		return
	}

	sniffBuffer = sniffBuffer[:sniffBytesRead]
	contentType := http.DetectContentType(sniffBuffer)
	if _, ok := policy.allowedContentTypes[contentType]; !ok {
		writeError(w, http.StatusBadRequest, policy.unsupportedTypeError)
		return
	}

	fileID, totalBytes, err := h.fileStore.Save(
		r.Context(),
		io.MultiReader(bytes.NewReader(sniffBuffer), file),
		policy.maxFileSizeBytes,
	)
	if err != nil {
		if errors.Is(err, errFileTooLarge) {
			writeError(w, http.StatusRequestEntityTooLarge, "file_too_large")
			return
		}

		h.logger.Error("save file", slog.Any("error", err))
		writeError(w, http.StatusInternalServerError, "internal_error")
		return
	}

	writeJSON(w, http.StatusCreated, uploadFileResponse{
		FileID:      fileID,
		FileName:    sanitizeUploadedFileName(fileHeader.Filename),
		ContentType: contentType,
		SizeBytes:   totalBytes,
	})
}

func countMultipartFiles(files map[string][]*multipart.FileHeader) int {
	totalFiles := 0
	for _, fileHeaders := range files {
		totalFiles += len(fileHeaders)
	}
	return totalFiles
}

func sanitizeUploadedFileName(name string) string {
	trimmed := strings.TrimSpace(name)
	if trimmed == "" {
		return "file"
	}

	normalized := strings.ReplaceAll(trimmed, "\\", "/")
	baseName := path.Base(normalized)
	if baseName == "." || baseName == "/" || baseName == "" {
		return "file"
	}
	return baseName
}

func isRequestTooLarge(err error) bool {
	var maxBytesErr *http.MaxBytesError
	return errors.As(err, &maxBytesErr)
}
