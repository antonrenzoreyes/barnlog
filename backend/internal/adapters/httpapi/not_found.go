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
	writeError(w, http.StatusNotFound, "not_found")
}
