package httpapi

import (
	"net/http"

	spec "barnlog/backend/openapi"
)

// OpenAPIDoc serves the OpenAPI document used by Swagger UI.
func OpenAPIDoc(w http.ResponseWriter, _ *http.Request) {
	body, err := spec.JSON()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_, _ = w.Write(body)
}
