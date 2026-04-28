package httpapi

import (
	"encoding/json"
	"net/http"
)

var allowedErrorCodes = map[string]struct{}{
	"file_required":              {},
	"file_too_large":             {},
	"internal_error":             {},
	"invalid_file":               {},
	"invalid_json":               {},
	"invalid_multipart":          {},
	"multiple_files_not_allowed": {},
	"not_found":                  {},
	"unsupported_file_type":      {},
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
