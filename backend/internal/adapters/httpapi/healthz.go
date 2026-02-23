package httpapi

import "net/http"

// healthz godoc
//
// @Summary Health check
// @Description Returns service liveness status.
// @Tags system
// @Produce json
// @Success 200 {object} statusResponse
// @Router /healthz [get]
func (h handlers) healthz(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]any{"status": "ok"})
}
