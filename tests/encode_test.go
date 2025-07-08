package test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/aryansharma9917/Codewise-CLI/pkg/encoder"
)

func createFile(t *testing.T, path string, content string) {
	t.Helper()
	err := os.WriteFile(path, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file %s: %v", path, err)
	}
}

func deleteFile(path string) {
	_ = os.Remove(path)
}

func TestYAMLToJSON(t *testing.T) {
	input := "testdata/sample.yaml"
	output := "testdata/output.json"

	createFile(t, input, `
name: Codewise
version: 1.0
`)

	defer deleteFile(input)
	defer deleteFile(output)

	err := encoder.YAMLToJSON(input, output)
	if err != nil {
		t.Fatalf("YAML to JSON conversion failed: %v", err)
	}

	data, err := os.ReadFile(output)
	if err != nil {
		t.Fatalf("Failed to read output file: %v", err)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		t.Fatalf("Output is not valid JSON: %v", err)
	}

	if result["name"] != "Codewise" {
		t.Errorf("Expected name to be 'Codewise', got '%v'", result["name"])
	}
}

func TestEnvToJSON(t *testing.T) {
	input := "testdata/sample.env"
	output := "testdata/env_output.json"

	createFile(t, input, `
APP_NAME=Codewise
VERSION=1.0
`)

	defer deleteFile(input)
	defer deleteFile(output)

	err := encoder.EnvToJSON(input, output)
	if err != nil {
		t.Fatalf(".env to JSON conversion failed: %v", err)
	}

	data, err := os.ReadFile(output)
	if err != nil {
		t.Fatalf("Failed to read env output file: %v", err)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		t.Fatalf("Output is not valid JSON: %v", err)
	}

	if result["APP_NAME"] != "Codewise" {
		t.Errorf("Expected APP_NAME to be 'Codewise', got '%v'", result["APP_NAME"])
	}
}
