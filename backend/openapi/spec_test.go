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

	body, err := JSON()
	if err != nil {
		t.Fatalf("JSON() returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(body, &payload); err != nil {
		t.Fatalf("expected valid JSON output, got error: %v", err)
	}

	openapiVersion, ok := payload["openapi"].(string)
	if !ok || openapiVersion == "" {
		t.Fatalf("expected openapi version in document, got %v", payload["openapi"])
	}
}
