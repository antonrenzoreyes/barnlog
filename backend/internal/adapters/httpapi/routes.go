// Package httpapi provides HTTP route registration and handlers for backend APIs.
package httpapi

import (
	"log/slog"
	"net/http"

	openapicontract "barnlog/backend/internal/contracts/openapi"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

// RouteDeps contains dependencies required to build HTTP routes.
type RouteDeps struct {
	Logger       *slog.Logger
	FileStoreDir string
}

// Routes builds the public HTTP router for backend endpoints.
func Routes(deps RouteDeps) http.Handler {
	r := chi.NewRouter()
	h := newHandlers(deps.Logger)
	store := newFileStore(deps.FileStoreDir)
	if store == nil {
		deps.Logger.Error("invalid file store dir", slog.String("file_store_dir", deps.FileStoreDir))
	}
	upload := newUploadHandlers(deps.Logger, store)
	server := oapiServerAdapter{
		system: h,
		upload: upload,
	}

	openapicontract.HandlerFromMux(server, r)
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/openapi.json"),
	))
	r.NotFound(h.notFound)

	return r
}
