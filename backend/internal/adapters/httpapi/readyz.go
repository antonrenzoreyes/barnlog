package httpapi

import (
	"net/http"
	"time"
)

// readyz returns service readiness status and current UTC timestamp.
func (h handlers) readyz(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, readyResponse{
		Status:    "ready",
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	})
}
