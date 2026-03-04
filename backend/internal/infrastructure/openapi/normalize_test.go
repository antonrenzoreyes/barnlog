package openapi

import "testing"

func TestNormalize(t *testing.T) {
	payload := map[string]any{
		"openapi": "3.1.0",
		"externalDocs": map[string]any{
			"description": "",
			"url":         "",
		},
		"paths": map[string]any{
			"/uploads/photos": map[string]any{
				"post": map[string]any{
					"requestBody": map[string]any{
						"content": map[string]any{
							"application/x-www-form-urlencoded": map[string]any{
								"schema": map[string]any{"type": "file"},
							},
						},
					},
				},
			},
		},
	}

	Normalize(payload)

	if got := payload["openapi"]; got != "3.0.3" {
		t.Fatalf("expected openapi 3.0.3, got %v", got)
	}
	if _, exists := payload["externalDocs"]; exists {
		t.Fatalf("did not expect externalDocs when url is empty")
	}

	paths := payload["paths"].(map[string]any)
	uploads := paths["/uploads/photos"].(map[string]any)
	post := uploads["post"].(map[string]any)
	requestBody := post["requestBody"].(map[string]any)
	content := requestBody["content"].(map[string]any)

	if _, exists := content["application/x-www-form-urlencoded"]; exists {
		t.Fatalf("did not expect application/x-www-form-urlencoded in content")
	}

	multipart, ok := content["multipart/form-data"].(map[string]any)
	if !ok {
		t.Fatalf("expected multipart/form-data content to exist")
	}

	schema := multipart["schema"].(map[string]any)
	properties := schema["properties"].(map[string]any)
	photo := properties["photo"].(map[string]any)

	if photo["type"] != "string" {
		t.Fatalf("expected photo type string, got %v", photo["type"])
	}
	if photo["format"] != "binary" {
		t.Fatalf("expected photo format binary, got %v", photo["format"])
	}
}

func TestNormalizeKeepsValidExternalDocs(t *testing.T) {
	payload := map[string]any{
		"openapi": "3.0.3",
		"externalDocs": map[string]any{
			"url": "https://example.com/docs",
		},
	}

	Normalize(payload)

	if _, exists := payload["externalDocs"]; !exists {
		t.Fatalf("expected externalDocs to remain when url is valid")
	}
}
