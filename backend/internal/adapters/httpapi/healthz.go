package httpapi

import "net/http"

// healthz returns service liveness status.
func (h handlers) healthz(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, statusResponse{Status: "ok"})
}
