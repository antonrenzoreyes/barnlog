package openapi

import "strings"

// Normalize mutates an OpenAPI payload to match Barnlog contract needs.
func Normalize(payload map[string]any) {
	normalizeVersion(payload)
	normalizeUploadPhotoRequestBody(payload)
}

func normalizeVersion(payload map[string]any) {
	version, ok := payload["openapi"].(string)
	if !ok {
		return
	}
	if strings.HasPrefix(version, "3.1.") {
		payload["openapi"] = "3.0.3"
	}
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
