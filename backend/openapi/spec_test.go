package openapi

import (
	"encoding/json"
	"testing"
)

func TestYAML(t *testing.T) {
	t.Parallel()

	yaml := YAML()
	if len(yaml) == 0 {
		t.Fatal("expected embedded OpenAPI YAML to be non-empty")
	}
}

func TestJSON(t *testing.T) {
	t.Parallel()

	payload := decodeJSONPayload(t)

	openapiVersion, ok := payload["openapi"].(string)
	if !ok || openapiVersion == "" {
		t.Fatalf("expected openapi version in document, got %v", payload["openapi"])
	}
}

func TestJSONUploadAnimalPhotoContract(t *testing.T) {
	t.Parallel()

	payload := decodeJSONPayload(t)
	post := uploadAnimalPhotoPost(t, payload)

	requestBody := mustMap(t, post["requestBody"], "requestBody")
	content := mustMap(t, requestBody["content"], "requestBody.content")
	multipart := mustMap(t, content["multipart/form-data"], "requestBody.content.multipart/form-data")
	schema := mustMap(t, multipart["schema"], "requestBody.content.multipart/form-data.schema")
	required := mustSlice(t, schema["required"], "requestBody.content.multipart/form-data.schema.required")
	if !containsString(required, "file") {
		t.Fatalf("expected multipart schema required fields to contain file, got %v", required)
	}

	properties := mustMap(t, schema["properties"], "requestBody.content.multipart/form-data.schema.properties")
	fileProperty := mustMap(t, properties["file"], "requestBody.content.multipart/form-data.schema.properties.file")
	if got := fileProperty["format"]; got != "binary" {
		t.Fatalf("expected file format=binary, got %v", got)
	}

	responses := mustMap(t, post["responses"], "responses")
	for _, status := range []string{"400", "413", "500"} {
		if _, ok := responses[status]; !ok {
			t.Fatalf("expected upload endpoint to define %s response", status)
		}
	}
}

func TestJSONUploadFileResponseRequiredFields(t *testing.T) {
	t.Parallel()

	payload := decodeJSONPayload(t)
	components := mustMap(t, payload["components"], "components")
	schemas := mustMap(t, components["schemas"], "components.schemas")
	uploadFileResponse := mustMap(t, schemas["httpapi.uploadFileResponse"], "components.schemas.httpapi.uploadFileResponse")
	required := mustSlice(t, uploadFileResponse["required"], "components.schemas.httpapi.uploadFileResponse.required")
	for _, field := range []string{"file_id", "file_name", "content_type", "size_bytes"} {
		if !containsString(required, field) {
			t.Fatalf("expected uploadFileResponse required fields to contain %q, got %v", field, required)
		}
	}
}

func decodeJSONPayload(t *testing.T) map[string]any {
	t.Helper()

	body, err := JSON()
	if err != nil {
		t.Fatalf("JSON() returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(body, &payload); err != nil {
		t.Fatalf("expected valid JSON output, got error: %v", err)
	}
	return payload
}

func uploadAnimalPhotoPost(t *testing.T, payload map[string]any) map[string]any {
	t.Helper()

	paths := mustMap(t, payload["paths"], "paths")
	uploadPath := mustMap(t, paths["/uploads/animal-photos"], "paths./uploads/animal-photos")
	return mustMap(t, uploadPath["post"], "paths./uploads/animal-photos.post")
}

func mustMap(t *testing.T, value any, field string) map[string]any {
	t.Helper()

	result, ok := value.(map[string]any)
	if !ok {
		t.Fatalf("expected %s to be map[string]any, got %T", field, value)
	}
	return result
}

func mustSlice(t *testing.T, value any, field string) []any {
	t.Helper()

	result, ok := value.([]any)
	if !ok {
		t.Fatalf("expected %s to be []any, got %T", field, value)
	}
	return result
}

func containsString(values []any, target string) bool {
	for _, value := range values {
		if asString, ok := value.(string); ok && asString == target {
			return true
		}
	}
	return false
}
