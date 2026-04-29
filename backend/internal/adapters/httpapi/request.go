package httpapi

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"mime"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5/middleware"
)

const (
	jsonContentType      = "application/json"
	sourceHeaderName     = "X-Barnlog-Source"
	requestIDHeaderName  = "X-Request-Id"
	defaultRequestSource = "http.api"
	maxJSONBodyBytes     = 1 << 20 // 1 MiB
)

// RequestMeta carries request-scoped metadata used by command handlers.
type RequestMeta struct {
	Source    string
	RequestID string
}

type requestMetaContextKey struct{}

func withRequestMeta(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		source := strings.TrimSpace(r.Header.Get(sourceHeaderName))
		if source == "" {
			source = defaultRequestSource
		}

		requestID := strings.TrimSpace(r.Header.Get(requestIDHeaderName))
		if requestID == "" {
			requestID = strings.TrimSpace(middleware.GetReqID(r.Context()))
		}

		meta := RequestMeta{
			Source:    source,
			RequestID: requestID,
		}
		ctx := context.WithValue(r.Context(), requestMetaContextKey{}, meta)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func requestMeta(ctx context.Context) (RequestMeta, bool) {
	meta, ok := ctx.Value(requestMetaContextKey{}).(RequestMeta)
	return meta, ok
}

func decodeJSONRequest(w http.ResponseWriter, r *http.Request, dst any) (status int, code string, ok bool) {
	contentType := strings.TrimSpace(r.Header.Get("Content-Type"))
	if contentType == "" {
		return http.StatusUnsupportedMediaType, "unsupported_media_type", false
	}

	mediaType, _, err := mime.ParseMediaType(contentType)
	if err != nil || mediaType != jsonContentType {
		return http.StatusUnsupportedMediaType, "unsupported_media_type", false
	}

	r.Body = http.MaxBytesReader(w, r.Body, maxJSONBodyBytes)
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	if err := dec.Decode(dst); err != nil {
		var maxBytesErr *http.MaxBytesError
		if errors.As(err, &maxBytesErr) {
			return http.StatusRequestEntityTooLarge, "request_too_large", false
		}
		return http.StatusBadRequest, "invalid_json", false
	}

	if err := dec.Decode(&struct{}{}); err != io.EOF {
		return http.StatusBadRequest, "invalid_json", false
	}

	return 0, "", true
}
