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

func normalizeJSONFile(path string) error {
	content, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read %s: %w", path, err)
	}

	var payload map[string]any
	if err := json.Unmarshal(content, &payload); err != nil {
		return fmt.Errorf("unmarshal %s: %w", path, err)
	}

	openapi.Normalize(payload)

	normalized, err := json.MarshalIndent(payload, "", "    ")
	if err != nil {
		return fmt.Errorf("marshal %s: %w", path, err)
	}
	normalized = append(normalized, '\n')

	if err := os.WriteFile(path, normalized, 0o644); err != nil {
		return fmt.Errorf("write %s: %w", path, err)
	}
	return nil
}

func normalizeYAMLFile(path string) error {
	content, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read %s: %w", path, err)
	}

	var payload map[string]any
	if err := yaml.Unmarshal(content, &payload); err != nil {
		return fmt.Errorf("unmarshal %s: %w", path, err)
	}

	openapi.Normalize(payload)

	normalized, err := yaml.Marshal(payload)
	if err != nil {
		return fmt.Errorf("marshal %s: %w", path, err)
	}
	if err := os.WriteFile(path, normalized, 0o644); err != nil {
		return fmt.Errorf("write %s: %w", path, err)
	}
	return nil
}
