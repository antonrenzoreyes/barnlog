// Package httpapi provides HTTP route registration and handlers for backend APIs.
package httpapi

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// Routes builds the public HTTP router for backend endpoints.
func Routes(logger *slog.Logger) http.Handler {
	r := chi.NewRouter()
	h := newHandlers(logger)

	r.Get("/healthz", h.healthz)
	r.Get("/readyz", h.readyz)
	r.NotFound(h.notFound)

	return r
}
