package encoder_test

import (
	"os/exec"
	"path/filepath"
	"testing"
)

func TestJSONToYAML(t *testing.T) {
	input := filepath.Join("..", "..", "testdata", "sample.json")
	output := filepath.Join("..", "..", "testdata", "JTY_output.yaml")
	binPath := filepath.Join("..", "..", "codewise")

	cmd := exec.Command(binPath, "encode", "--input", input, "--output", output)
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("❌ Command failed: %v\nOutput: %s", err, string(out))
	}
}
