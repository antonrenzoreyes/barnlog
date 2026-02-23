package httpapi

import (
	"encoding/json"
	"net/http"
	"strings"

	"barnlog/backend/docs"
)

// OpenAPIDoc serves the OpenAPI document used by Swagger UI.
func OpenAPIDoc(w http.ResponseWriter, _ *http.Request) {
	doc := docs.SwaggerInfo.ReadDoc()

	var payload map[string]any
	if err := json.Unmarshal([]byte(doc), &payload); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	if version, ok := payload["openapi"].(string); ok && strings.HasPrefix(version, "3.1.") {
		payload["openapi"] = "3.0.3"
	}
	normalizeAnimalsRequestBody(payload)
	normalizeUploadPhotoRequestBody(payload)
	body, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_, _ = w.Write(body)
}

func normalizeUploadPhotoRequestBody(payload map[string]any) {
	paths, ok := payload["paths"].(map[string]any)
	if !ok {
		return
	}
	uploads, ok := paths["/uploads/photos"].(map[string]any)
	if !ok {
		return
	}
	post, ok := uploads["post"].(map[string]any)
	if !ok {
		return
	}
	requestBody, ok := post["requestBody"].(map[string]any)
	if !ok {
		return
	}
	content, ok := requestBody["content"].(map[string]any)
	if !ok {
		return
	}
	content["multipart/form-data"] = map[string]any{
		"schema": map[string]any{
			"type":     "object",
			"required": []string{"photo"},
			"properties": map[string]any{
				"photo": map[string]any{
					"type":        "string",
					"format":      "binary",
					"description": "Photo file to upload",
				},
			},
		},
		"example": map[string]any{
			"photo": "(binary file)",
		},
	}
	delete(content, "application/x-www-form-urlencoded")
}

func normalizeAnimalsRequestBody(payload map[string]any) {
	paths, ok := payload["paths"].(map[string]any)
	if !ok {
		return
	}
	animals, ok := paths["/animals"].(map[string]any)
	if !ok {
		return
	}
	post, ok := animals["post"].(map[string]any)
	if !ok {
		return
	}
	requestBody, ok := post["requestBody"].(map[string]any)
	if !ok {
		return
	}
	content, ok := requestBody["content"].(map[string]any)
	if !ok {
		return
	}
	applicationJSON, ok := content["application/json"].(map[string]any)
	if !ok {
		return
	}
	applicationJSON["schema"] = map[string]any{
		"$ref": "#/components/schemas/httpapi.createAnimalRequest",
	}
	applicationJSON["example"] = map[string]any{
		"name":      "Nanny",
		"species":   "goat",
		"tag":       "G-7",
		"birthdate": "2021-03-04",
		"photo_id":  "photo_1",
	}
}
