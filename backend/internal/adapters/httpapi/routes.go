// Package httpapi provides HTTP route registration and handlers for backend APIs.
package httpapi

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// RouteDeps contains dependencies required to build HTTP routes.
type RouteDeps struct {
	Logger        *slog.Logger
	PhotoStoreDir string
}

// Routes builds the public HTTP router for backend endpoints.
func Routes(deps RouteDeps) http.Handler {
	r := chi.NewRouter()
	h := newHandlers(deps.Logger)
	upload := newUploadHandlers(deps.Logger, newFilePhotoStore(deps.PhotoStoreDir))

	r.Get("/healthz", h.healthz)
	r.Get("/readyz", h.readyz)
	r.Post("/uploads/photos", upload.uploadPhoto)
	r.NotFound(h.notFound)

	return r
}
