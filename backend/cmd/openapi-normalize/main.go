// Command openapi-normalize normalizes generated OpenAPI docs for Barnlog.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"barnlog/backend/internal/infrastructure/openapi"

	"gopkg.in/yaml.v3"
)

func main() {
	jsonPath := flag.String("json", "backend/docs/swagger.json", "path to OpenAPI JSON document")
	yamlPath := flag.String("yaml", "backend/docs/swagger.yaml", "path to OpenAPI YAML document")
	flag.Parse()

	if err := normalizeJSONFile(*jsonPath); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "normalize json: %v\n", err)
		os.Exit(1)
	}
	if err := normalizeYAMLFile(*yamlPath); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "normalize yaml: %v\n", err)
		os.Exit(1)
	}
}

type unmarshalFn func([]byte, any) error
type marshalFn func(any) ([]byte, error)

func normalizeJSONFile(path string) error {
	return normalizeFile(path, json.Unmarshal, func(v any) ([]byte, error) {
		b, err := json.MarshalIndent(v, "", "    ")
		if err != nil {
			return nil, err
		}
		return append(b, '\n'), nil
	})
}

func normalizeYAMLFile(path string) error {
	return normalizeFile(path, yaml.Unmarshal, yaml.Marshal)
}

func normalizeFile(path string, unmarshal unmarshalFn, marshal marshalFn) error {
	//nolint:gosec // path is controlled by local tooling (Makefile/flags) in this repository.
	content, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read %s: %w", path, err)
	}

	var payload map[string]any
	if err := unmarshal(content, &payload); err != nil {
		return fmt.Errorf("unmarshal %s: %w", path, err)
	}

	openapi.Normalize(payload)

	normalized, err := marshal(payload)
	if err != nil {
		return fmt.Errorf("marshal %s: %w", path, err)
	}

	//nolint:gosec // generated API docs are intended repo artifacts, not secrets.
	if err := os.WriteFile(path, normalized, 0o600); err != nil {
		return fmt.Errorf("write %s: %w", path, err)
	}
	return nil
}
