package httpapi

import (
	"net/http"
	"time"
)

// readyz godoc
//
// @Summary Readiness check
// @Description Returns service readiness status and current UTC timestamp.
// @Tags system
// @Produce json
// @Success 200 {object} readyResponse
// @Router /readyz [get]
func (h handlers) readyz(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]any{
		"status":    "ready",
		"timestamp": time.Now().UTC().Format(time.RFC3339),
	})
}
