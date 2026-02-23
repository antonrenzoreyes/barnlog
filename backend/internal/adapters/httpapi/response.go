package httpapi

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"barnlog/backend/internal/application"
)

func writeJSON(w http.ResponseWriter, status int, payload any) {
	body, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, _ = w.Write(append(body, '\n'))
}

func writeError(w http.ResponseWriter, status int, code string) {
	writeJSON(w, status, errorResponse{Error: code})
}

func writeBusinessError(w http.ResponseWriter, logger *slog.Logger, err error) bool {
	be, ok := application.AsBusinessError(err)
	if !ok {
		return false
	}

	switch be.Code {
	case application.CodeInvalidInput,
		application.CodeNameRequired,
		application.CodeSpeciesInvalid,
		application.CodeBirthdateInvalid,
		application.CodePhotoNotFound:
		writeError(w, http.StatusBadRequest, string(be.Code))
	case application.CodeConflict,
		application.CodeIdempotencyPayloadMismatch,
		application.CodeIdempotencyEventTypeMismatch:
		writeError(w, http.StatusConflict, string(be.Code))
	default:
		logger.Error("unknown business error code", slog.String("code", string(be.Code)), slog.Any("error", err))
		writeError(w, http.StatusInternalServerError, "internal_error")
	}
	return true
}
