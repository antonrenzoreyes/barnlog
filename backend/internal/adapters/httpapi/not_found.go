package httpapi

import (
	"log/slog"
	"net/http"
)

// notFound godoc
//
// @Summary Route not found
// @Description Returned when route does not exist.
// @Tags system
// @Produce json
// @Failure 404 {object} errorResponse
func (h handlers) notFound(w http.ResponseWriter, r *http.Request) {
	h.logger.Warn("route not found", slog.String("method", r.Method), slog.String("path", r.URL.Path))
	writeJSON(w, http.StatusNotFound, map[string]any{"error": "not_found"})
}
