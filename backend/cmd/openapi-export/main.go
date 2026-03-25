// Command openapi-export writes derived OpenAPI artifacts from the canonical spec.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"sigs.k8s.io/yaml"
)

func main() {
	inPath := flag.String("in", "backend/openapi/openapi.yaml", "path to canonical OpenAPI YAML")
	jsonPath := flag.String("json", "backend/docs/swagger.json", "path to generated OpenAPI JSON")
	yamlPath := flag.String("yaml", "backend/docs/swagger.yaml", "path to generated OpenAPI YAML")
	flag.Parse()

	if err := run(*inPath, *jsonPath, *yamlPath); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "openapi export failed: %v\n", err)
		os.Exit(1)
	}
}

func run(inPath, jsonPath, yamlPath string) error {
	yamlInput, err := os.ReadFile(inPath)
	if err != nil {
		return fmt.Errorf("read canonical spec: %w", err)
	}

	var payload map[string]any
	if err := yaml.Unmarshal(yamlInput, &payload); err != nil {
		return fmt.Errorf("decode canonical spec: %w", err)
	}

	jsonOutput, err := json.MarshalIndent(payload, "", "    ")
	if err != nil {
		return fmt.Errorf("encode OpenAPI JSON: %w", err)
	}
	jsonOutput = append(jsonOutput, '\n')

	if err := ensureParentDir(jsonPath); err != nil {
		return err
	}
	if err := ensureParentDir(yamlPath); err != nil {
		return err
	}
	if err := os.WriteFile(jsonPath, jsonOutput, 0o600); err != nil {
		return fmt.Errorf("write OpenAPI JSON: %w", err)
	}
	if err := os.WriteFile(yamlPath, yamlInput, 0o600); err != nil {
		return fmt.Errorf("write OpenAPI YAML: %w", err)
	}

	return nil
}

func ensureParentDir(path string) error {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0o750); err != nil {
		return fmt.Errorf("create parent directory %q: %w", dir, err)
	}
	return nil
}
