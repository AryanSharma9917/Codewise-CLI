package encoder_test

import (
	"os"
	"testing"

	"github.com/aryansharma9917/Codewise-CLI/pkg/encoder"
)

func TestYAMLToJSON(t *testing.T) {
	input := "tests/testdata/sample.yaml"
	output := "tests/testdata/output.json"

	// Create sample input
	content := `
name: Codewise
version: 1.0
`
	err := os.WriteFile(input, []byte(content), 0644)
	if err != nil {
		t.Fatalf("❌ Failed to write sample YAML: %v", err)
	}

	err = encoder.YAMLToJSON(input, output)
	if err != nil {
		t.Errorf("❌ YAMLToJSON failed: %v", err)
	}

	defer os.Remove(output)
}
