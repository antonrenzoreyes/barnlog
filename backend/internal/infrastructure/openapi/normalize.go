// Package openapi contains OpenAPI payload normalization helpers used by backend tooling.
package openapi

import "strings"

// Normalize mutates an OpenAPI payload to match Barnlog contract needs.
func Normalize(payload map[string]any) {
	normalizeVersion(payload)
	normalizePublicSecurity(payload)
	normalizeExternalDocs(payload)
	normalizeUploadAnimalPhotoRequestBody(payload)
	normalizeUploadAnimalPhotoResponses(payload)
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

func normalizeUploadAnimalPhotoRequestBody(payload map[string]any) {
	paths, ok := payload["paths"].(map[string]any)
	if !ok {
		return
	}
	uploads, ok := paths["/uploads/animal-photos"].(map[string]any)
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

	multipart := ensureMap(content, "multipart/form-data")
	schema := ensureMap(multipart, "schema")
	schema["type"] = "object"
	schema["required"] = ensureContainsString(schema["required"], "file")

	properties := ensureMap(schema, "properties")
	properties["file"] = map[string]any{
		"type":        "string",
		"format":      "binary",
		"description": "Animal photo file to upload",
	}

	example := ensureMap(multipart, "example")
	example["file"] = "(binary file)"
	delete(content, "application/x-www-form-urlencoded")
}

func normalizeUploadAnimalPhotoResponses(payload map[string]any) {
	paths, ok := payload["paths"].(map[string]any)
	if !ok {
		return
	}

	uploads, ok := paths["/uploads/animal-photos"].(map[string]any)
	if !ok {
		return
	}
	post, ok := uploads["post"].(map[string]any)
	if !ok {
		return
	}

	responses := ensureMap(post, "responses")
	ensureMap(responses, "400")["description"] = "Bad Request (invalid_multipart | file_required | multiple_files_not_allowed | invalid_file | unsupported_file_type)"
	ensureMap(responses, "413")["description"] = "Request Entity Too Large (file_too_large)"
	ensureMap(responses, "500")["description"] = "Internal Server Error (internal_error)"

	components := ensureMap(payload, "components")
	schemas := ensureMap(components, "schemas")
	uploadFileResponse, ok := schemas["httpapi.uploadFileResponse"].(map[string]any)
	if !ok {
		return
	}

	uploadFileResponse["required"] = ensureContainsString(uploadFileResponse["required"], "file_id")
	uploadFileResponse["required"] = ensureContainsString(uploadFileResponse["required"], "file_name")
	uploadFileResponse["required"] = ensureContainsString(uploadFileResponse["required"], "content_type")
	uploadFileResponse["required"] = ensureContainsString(uploadFileResponse["required"], "size_bytes")
}

func normalizeExternalDocs(payload map[string]any) {
	externalDocs, ok := payload["externalDocs"].(map[string]any)
	if !ok {
		return
	}

	urlValue, hasURL := externalDocs["url"]
	url, isString := urlValue.(string)
	if !hasURL || !isString || strings.TrimSpace(url) == "" {
		delete(payload, "externalDocs")
	}
}

func normalizePublicSecurity(payload map[string]any) {
	payload["security"] = []any{}
	components := ensureMap(payload, "components")
	if _, ok := components["securitySchemes"]; !ok {
		components["securitySchemes"] = map[string]any{}
	}
}

func ensureMap(parent map[string]any, key string) map[string]any {
	value, ok := parent[key].(map[string]any)
	if !ok || value == nil {
		value = map[string]any{}
		parent[key] = value
	}
	return value
}

func ensureContainsString(value any, target string) []any {
	items, ok := value.([]any)
	if !ok {
		items = []any{}
	}
	for _, item := range items {
		if s, ok := item.(string); ok && s == target {
			return items
		}
	}
	return append(items, target)
}
