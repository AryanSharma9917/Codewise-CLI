package test

import (
	"encoding/json"
	"io/ioutil"
	"os/exec"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// TestYamlToJsonCmd tests the YTJ command
func TestYamlToJsonCmd(t *testing.T) {

	// Execute the yamlToJson command
	cmd := exec.Command("Codewise-CLI", "YTJ", "-f", "testdata/YTJ.yaml")

	// Capture the output
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Errorf("expected no error, but got: %v", err)
	}

	// Validate the cli output
	expectedOutput := "Operation completed successfully. Check the output.json file."
	got := strings.TrimSpace(string(output))
	if got != expectedOutput {
		t.Errorf("expected %v, but got: %v", expectedOutput, got)
	}

	// Read the actual output JSON file
	actualJSON, err := ioutil.ReadFile("output.json")
	if err != nil {
		t.Fatalf("could not read output.json: %v", err)
	}

	// Read the expected JSON file
	expectedJSON, err := ioutil.ReadFile("testdata/YTJ_output.json")
	if err != nil {
		t.Fatalf("could not read testdata/YTJ_output.json: %v", err)
	}

	// Unmarshal JSON files
	var actual, expected map[string]interface{}
	if err := json.Unmarshal(actualJSON, &actual); err != nil {
		t.Fatalf("could not unmarshal actual JSON: %v", err)
	}
	if err := json.Unmarshal(expectedJSON, &expected); err != nil {
		t.Fatalf("could not unmarshal expected JSON: %v", err)
	}

	// Compare JSON files
	if diff := cmp.Diff(expected, actual); diff != "" {
		t.Errorf("Mismatch (-expected +actual):\n%s", diff)
	}
}

// TestYamlToJsonCmdWithOutputFlag tests the YTJ command with the output flag
func TestYamlToJsonCmdWithOutputFlag(t *testing.T) {

	// Execute the yamlToJson command
	cmd := exec.Command("Codewise-CLI", "YTJ", "-f", "testdata/YTJ.yaml", "-o", "YTJ_output.json")

	// Capture the output
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Errorf("expected no error, but got: %v", err)
	}

	// Validate the cli output
	expectedOutput := "Operation completed successfully. Check the YTJ_output.json file."
	got := strings.TrimSpace(string(output))
	if got != expectedOutput {
		t.Errorf("expected %v, but got: %v", expectedOutput, got)
	}

	// Read the actual output JSON file
	actualJSON, err := ioutil.ReadFile("YTJ_output.json")
	if err != nil {
		t.Fatalf("could not read YTJ_output.json: %v", err)
	}

	// Read the expected JSON file
	expectedJSON, err := ioutil.ReadFile("testdata/YTJ_output.json")
	if err != nil {
		t.Fatalf("could not read testdata/YTJ_output.json: %v", err)
	}

	// Unmarshal JSON files
	var actual, expected map[string]interface{}
	if err := json.Unmarshal(actualJSON, &actual); err != nil {
		t.Fatalf("could not unmarshal actual JSON: %v", err)
	}
	if err := json.Unmarshal(expectedJSON, &expected); err != nil {
		t.Fatalf("could not unmarshal expected JSON: %v", err)
	}

	// Compare JSON files
	if diff := cmp.Diff(expected, actual); diff != "" {
		t.Errorf("Mismatch (-expected +actual):\n%s", diff)
	}
}
