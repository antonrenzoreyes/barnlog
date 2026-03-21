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
			"/uploads/animal-photos": map[string]any{
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
	if security, ok := payload["security"].([]any); !ok || len(security) != 0 {
		t.Fatalf("expected root security to be an empty array")
	}

	components := payload["components"].(map[string]any)
	securitySchemes := components["securitySchemes"].(map[string]any)
	if len(securitySchemes) != 0 {
		t.Fatalf("expected empty securitySchemes, got %v entries", len(securitySchemes))
	}

	paths := payload["paths"].(map[string]any)
	uploads := paths["/uploads/animal-photos"].(map[string]any)
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
	file := properties["file"].(map[string]any)
	required := schema["required"].([]any)
	foundFileRequired := false
	for _, v := range required {
		if s, ok := v.(string); ok && s == "file" {
			foundFileRequired = true
			break
		}
	}
	if !foundFileRequired {
		t.Fatalf(`expected schema.required to contain "file"`)
	}

	if file["type"] != "string" {
		t.Fatalf("expected file type string, got %v", file["type"])
	}
	if file["format"] != "binary" {
		t.Fatalf("expected file format binary, got %v", file["format"])
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

func TestNormalizePreservesExistingMultipartMetadata(t *testing.T) {
	payload := map[string]any{
		"paths": map[string]any{
			"/uploads/animal-photos": map[string]any{
				"post": map[string]any{
					"requestBody": map[string]any{
						"content": map[string]any{
							"multipart/form-data": map[string]any{
								"encoding": map[string]any{"file": map[string]any{"contentType": "image/jpeg"}},
							},
						},
					},
				},
			},
		},
	}

	Normalize(payload)

	paths := payload["paths"].(map[string]any)
	uploads := paths["/uploads/animal-photos"].(map[string]any)
	post := uploads["post"].(map[string]any)
	requestBody := post["requestBody"].(map[string]any)
	content := requestBody["content"].(map[string]any)
	multipart := content["multipart/form-data"].(map[string]any)

	if _, ok := multipart["encoding"].(map[string]any); !ok {
		t.Fatalf("expected existing multipart metadata to be preserved")
	}
}

func TestNormalizeMarksUploadResponseFieldsAsRequired(t *testing.T) {
	payload := map[string]any{
		"components": map[string]any{
			"schemas": map[string]any{
				"httpapi.uploadFileResponse": map[string]any{
					"properties": map[string]any{
						"file_id":      map[string]any{"type": "string"},
						"file_name":    map[string]any{"type": "string"},
						"content_type": map[string]any{"type": "string"},
						"size_bytes":   map[string]any{"type": "integer"},
					},
				},
			},
		},
		"paths": map[string]any{
			"/uploads/animal-photos": map[string]any{
				"post": map[string]any{
					"responses": map[string]any{
						"400": map[string]any{"description": "Bad Request"},
						"413": map[string]any{"description": "Request Entity Too Large"},
						"500": map[string]any{"description": "Internal Server Error"},
					},
				},
			},
		},
	}

	Normalize(payload)

	components := payload["components"].(map[string]any)
	schemas := components["schemas"].(map[string]any)
	uploadFile := schemas["httpapi.uploadFileResponse"].(map[string]any)
	required := uploadFile["required"].([]any)

	requiredSet := map[string]bool{}
	for _, field := range required {
		if s, ok := field.(string); ok {
			requiredSet[s] = true
		}
	}

	for _, field := range []string{"file_id", "file_name", "content_type", "size_bytes"} {
		if !requiredSet[field] {
			t.Fatalf("expected required to contain %q, got %#v", field, required)
		}
	}

	paths := payload["paths"].(map[string]any)
	uploads := paths["/uploads/animal-photos"].(map[string]any)
	post := uploads["post"].(map[string]any)
	responses := post["responses"].(map[string]any)

	if responses["400"].(map[string]any)["description"] == "Bad Request" {
		t.Fatalf("expected 400 description to be normalized with upload error codes")
	}
	if responses["413"].(map[string]any)["description"] == "Request Entity Too Large" {
		t.Fatalf("expected 413 description to be normalized with upload error code")
	}
	if responses["500"].(map[string]any)["description"] == "Internal Server Error" {
		t.Fatalf("expected 500 description to be normalized with upload error code")
	}
}
