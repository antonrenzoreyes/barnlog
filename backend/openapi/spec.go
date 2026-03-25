package openapi

import (
	_ "embed"
	"encoding/json"
	"fmt"

	"sigs.k8s.io/yaml"
)

// specYAML is the canonical OpenAPI document used by backend and frontend tooling.
//
//go:embed openapi.yaml
var specYAML []byte

// YAML returns a copy of the canonical OpenAPI YAML document.
func YAML() []byte {
	out := make([]byte, len(specYAML))
	copy(out, specYAML)
	return out
}

// JSON returns the canonical OpenAPI document encoded as JSON.
func JSON() ([]byte, error) {
	var payload map[string]any
	if err := yaml.Unmarshal(specYAML, &payload); err != nil {
		return nil, fmt.Errorf("decode embedded OpenAPI YAML: %w", err)
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("encode OpenAPI JSON: %w", err)
	}
	return body, nil
}
