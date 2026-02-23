// Package httpapi provides HTTP route registration and handlers for backend APIs.
package httpapi

import (
	"log/slog"
	"net/http"

	"barnlog/backend/internal/application"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

// RouteDeps contains dependencies required to build HTTP routes.
type RouteDeps struct {
	Logger        *slog.Logger
	PhotoStoreDir string
	AnimalWriter  application.AnimalWriter
}

// Routes builds the public HTTP router for backend endpoints.
func Routes(deps RouteDeps) http.Handler {
	if deps.AnimalWriter == nil {
		panic("httpapi: AnimalWriter is required")
	}

	r := chi.NewRouter()
	r.Use(withRequestMeta)

	h := newHandlers(deps.Logger)
	animal := newAnimalHandlers(deps.Logger, deps.AnimalWriter)
	upload := newUploadHandlers(deps.Logger, newFilePhotoStore(deps.PhotoStoreDir))

	r.Get("/healthz", h.healthz)
	r.Get("/readyz", h.readyz)
	r.Post("/animals", animal.createAnimal)
	r.Post("/uploads/photos", upload.uploadPhoto)
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/openapi.json"),
	))
	r.NotFound(h.notFound)

	return r
}
