// Package httpapi provides HTTP route registration and handlers for backend APIs.
package httpapi

import "log/slog"

type handlers struct {
	logger *slog.Logger
}

func newHandlers(logger *slog.Logger) handlers {
	return handlers{logger: logger}
}
