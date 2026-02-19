// Package httpapi provides HTTP route registration and handlers for backend APIs.
package httpapi

import (
	"log/slog"
)

type handlers struct {
	logger *slog.Logger
}

type uploadHandlers struct {
	logger     *slog.Logger
	photoStore photoStore
}

func newHandlers(logger *slog.Logger) handlers {
	return handlers{logger: logger}
}

func newUploadHandlers(logger *slog.Logger, photoStore photoStore) uploadHandlers {
	return uploadHandlers{
		logger:     logger,
		photoStore: photoStore,
	}
}
