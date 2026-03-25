// Command openapi-export writes derived OpenAPI artifacts from the canonical spec.
package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"sigs.k8s.io/yaml"
)

func main() {
	inPath := flag.String("in", "backend/openapi/openapi.yaml", "path to canonical OpenAPI YAML")
	jsonPath := flag.String("json", "backend/docs/swagger.json", "path to generated OpenAPI JSON")
	yamlPath := flag.String("yaml", "backend/docs/swagger.yaml", "path to generated OpenAPI YAML")
	flag.Parse()

	if err := run(*inPath, *jsonPath, *yamlPath); err != nil {
		slog.Error("openapi export failed", "error", err)
		os.Exit(1)
	}
}

func run(inPath, jsonPath, yamlPath string) error {
	validatedInPath, err := validateCanonicalSpecPath(inPath)
	if err != nil {
		return fmt.Errorf("validate canonical spec path: %w", err)
	}
	validatedJSONPath, err := validateOutputPath(jsonPath, ".json")
	if err != nil {
		return fmt.Errorf("validate JSON output path: %w", err)
	}
	validatedYAMLPath, err := validateOutputPath(yamlPath, ".yaml")
	if err != nil {
		return fmt.Errorf("validate YAML output path: %w", err)
	}

	// #nosec G304 -- validated canonical path allowlist.
	yamlInput, err := os.ReadFile(validatedInPath)
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

	if err := ensureParentDir(validatedJSONPath); err != nil {
		return err
	}
	if err := ensureParentDir(validatedYAMLPath); err != nil {
		return err
	}
	if err := os.WriteFile(validatedJSONPath, jsonOutput, 0o600); err != nil {
		return fmt.Errorf("write OpenAPI JSON: %w", err)
	}
	if err := os.WriteFile(validatedYAMLPath, yamlInput, 0o600); err != nil {
		return fmt.Errorf("write OpenAPI YAML: %w", err)
	}

	return nil
}

func validateCanonicalSpecPath(path string) (string, error) {
	resolvedPath, err := resolvePath(path)
	if err != nil {
		return "", err
	}
	expectedPath, err := resolvePath("backend/openapi/openapi.yaml")
	if err != nil {
		return "", err
	}
	if resolvedPath != expectedPath {
		return "", fmt.Errorf("path %q must resolve to %q", path, "backend/openapi/openapi.yaml")
	}
	if err := ensureNoSymlinkInPath(resolvedPath); err != nil {
		return "", err
	}
	return resolvedPath, nil
}

func validateOutputPath(path, expectedExt string) (string, error) {
	resolvedPath, err := resolvePath(path)
	if err != nil {
		return "", err
	}
	docsRoot, err := resolvePath("backend/docs")
	if err != nil {
		return "", err
	}
	if err := ensurePathWithinRoot(resolvedPath, docsRoot); err != nil {
		return "", err
	}
	if ext := strings.ToLower(filepath.Ext(resolvedPath)); ext != expectedExt {
		return "", fmt.Errorf("path %q must use %s extension", path, expectedExt)
	}
	if err := ensureNoSymlinkInPath(resolvedPath); err != nil {
		return "", err
	}
	return resolvedPath, nil
}

func resolvePath(path string) (string, error) {
	clean := filepath.Clean(path)
	abs, err := filepath.Abs(clean)
	if err != nil {
		return "", fmt.Errorf("resolve absolute path for %q: %w", path, err)
	}
	return abs, nil
}

func ensurePathWithinRoot(path, root string) error {
	relativePath, err := filepath.Rel(root, path)
	if err != nil {
		return fmt.Errorf("resolve path %q under root %q: %w", path, root, err)
	}
	if relativePath == ".." || strings.HasPrefix(relativePath, ".."+string(os.PathSeparator)) {
		return fmt.Errorf("path %q must be under %q", path, root)
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

func ensureNoSymlinkInPath(path string) error {
	current := filepath.Clean(path)
	components := []string{current}

	for {
		parent := filepath.Dir(current)
		if parent == current {
			break
		}
		components = append(components, parent)
		current = parent
	}

	for i := len(components) - 1; i >= 0; i-- {
		component := components[i]
		info, err := os.Lstat(component)
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				continue
			}
			return fmt.Errorf("inspect path component %q: %w", component, err)
		}
		if info.Mode()&os.ModeSymlink != 0 {
			return fmt.Errorf("path %q contains symlink component %q", path, component)
		}
	}

	return nil
}
