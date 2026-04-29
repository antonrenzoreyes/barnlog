package httpapi

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"barnlog/backend/internal/application"
)

var allowedErrorCodes = map[string]struct{}{
	"birthdate_invalid":               {},
	"conflict":                        {},
	"file_required":                   {},
	"file_too_large":                  {},
	"idempotency_event_type_mismatch": {},
	"idempotency_payload_mismatch":    {},
	"internal_error":                  {},
	"invalid_input":                   {},
	"invalid_file":                    {},
	"invalid_json":                    {},
	"invalid_multipart":               {},
	"multiple_files_not_allowed":      {},
	"name_required":                   {},
	"not_found":                       {},
	"photo_not_found":                 {},
	"species_invalid":                 {},
	"unsupported_media_type":          {},
	"unsupported_file_type":           {},
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	body, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	// #nosec G705 -- payload is JSON-encoded and served with application/json.
	_, _ = w.Write(append(body, '\n'))
}

func writeError(w http.ResponseWriter, status int, code string) {
	writeJSON(w, status, newErrorResponse(normalizeErrorCode(code)))
}

func normalizeErrorCode(code string) string {
	if _, ok := allowedErrorCodes[code]; ok {
		return code
	}

	return "internal_error"
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
