package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRunWritesArtifacts(t *testing.T) {
	chdirToTempRepo(t)

	input := []byte(strings.Join([]string{
		"openapi: 3.0.3",
		"info:",
		"  title: Test API",
		`  version: "1.0"`,
		"paths: {}",
		"",
	}, "\n"))
	if err := os.MkdirAll("backend/openapi", 0o750); err != nil {
		t.Fatalf("create openapi dir: %v", err)
	}
	if err := os.WriteFile("backend/openapi/openapi.yaml", input, 0o600); err != nil {
		t.Fatalf("write canonical spec: %v", err)
	}

	if err := run("backend/openapi/openapi.yaml", "backend/docs/swagger.json", "backend/docs/swagger.yaml"); err != nil {
		t.Fatalf("run() error: %v", err)
	}

	jsonBody, err := os.ReadFile("backend/docs/swagger.json")
	if err != nil {
		t.Fatalf("read swagger.json: %v", err)
	}
	var payload map[string]any
	if err := json.Unmarshal(jsonBody, &payload); err != nil {
		t.Fatalf("unmarshal swagger.json: %v", err)
	}
	if payload["openapi"] != "3.0.3" {
		t.Fatalf("expected openapi version 3.0.3, got %#v", payload["openapi"])
	}

	yamlBody, err := os.ReadFile("backend/docs/swagger.yaml")
	if err != nil {
		t.Fatalf("read swagger.yaml: %v", err)
	}
	if string(yamlBody) != string(input) {
		t.Fatalf("expected swagger.yaml to match input")
	}

	assertFileMode(t, "backend/docs/swagger.json", 0o600)
	assertFileMode(t, "backend/docs/swagger.yaml", 0o600)
}

func TestValidateCanonicalSpecPathRejectsSymlink(t *testing.T) {
	chdirToTempRepo(t)

	if err := os.MkdirAll("backend/openapi", 0o750); err != nil {
		t.Fatalf("create openapi dir: %v", err)
	}
	if err := os.WriteFile("backend/openapi/real.yaml", []byte("openapi: 3.0.3\n"), 0o600); err != nil {
		t.Fatalf("write real spec: %v", err)
	}
	if err := os.Symlink("real.yaml", "backend/openapi/openapi.yaml"); err != nil {
		if os.IsPermission(err) {
			t.Skipf("symlink not permitted: %v", err)
		}
		t.Fatalf("create symlink: %v", err)
	}

	_, err := validateCanonicalSpecPath("backend/openapi/openapi.yaml")
	if err == nil || !strings.Contains(err.Error(), "symlink") {
		t.Fatalf("expected symlink validation error, got %v", err)
	}
}

func TestValidateOutputPathRejectsSymlinkedDocsRoot(t *testing.T) {
	chdirToTempRepo(t)

	if err := os.MkdirAll("backend", 0o750); err != nil {
		t.Fatalf("create backend dir: %v", err)
	}
	if err := os.MkdirAll("docs-target", 0o750); err != nil {
		t.Fatalf("create docs target dir: %v", err)
	}
	if err := os.Symlink(filepath.Join("..", "..", "docs-target"), "backend/docs"); err != nil {
		if os.IsPermission(err) {
			t.Skipf("symlink not permitted: %v", err)
		}
		t.Fatalf("create docs symlink: %v", err)
	}

	_, err := validateOutputPath("backend/docs/swagger.json", ".json")
	if err == nil || !strings.Contains(err.Error(), "symlink") {
		t.Fatalf("expected symlink validation error, got %v", err)
	}
}

func TestValidateOutputPathRejectsWrongExtension(t *testing.T) {
	chdirToTempRepo(t)

	if err := os.MkdirAll("backend/docs", 0o750); err != nil {
		t.Fatalf("create docs dir: %v", err)
	}

	_, err := validateOutputPath("backend/docs/swagger.txt", ".json")
	if err == nil || !strings.Contains(err.Error(), "must use .json extension") {
		t.Fatalf("expected extension validation error, got %v", err)
	}
}

func chdirToTempRepo(t *testing.T) {
	t.Helper()
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("getwd: %v", err)
	}

	tempDir := t.TempDir()
	if err := os.Chdir(tempDir); err != nil {
		t.Fatalf("chdir temp dir: %v", err)
	}
	t.Cleanup(func() {
		if err := os.Chdir(wd); err != nil {
			t.Fatalf("restore cwd: %v", err)
		}
	})
}

func assertFileMode(t *testing.T, path string, expected os.FileMode) {
	t.Helper()
	info, err := os.Stat(path)
	if err != nil {
		t.Fatalf("stat %s: %v", path, err)
	}
	if info.Mode().Perm() != expected {
		t.Fatalf("expected %s mode %o, got %o", path, expected, info.Mode().Perm())
	}
}
