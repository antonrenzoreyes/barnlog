package httpapi

import (
	"net/http"
	"time"
)

func (h handlers) readyz(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]any{
		"status":    "ready",
		"timestamp": time.Now().UTC().Format(time.RFC3339),
	})
}
