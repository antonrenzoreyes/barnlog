package httpapi

import (
	"encoding/json"
	"net/http"
)

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
	writeJSON(w, status, newErrorResponse(code))
}
