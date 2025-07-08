package encoder_test

import (
	"os/exec"
	"path/filepath"
	"testing"
)

func TestKeyValueToJSON(t *testing.T) {
	input := filepath.Join("..", "..", "testdata", "sample.env")
	output := filepath.Join("..", "..", "testdata", "env_output.json")
	binPath := filepath.Join("..", "..", "codewise")

	cmd := exec.Command(binPath, "encode", "--input", input, "--output", output)
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("❌ Command failed: %v\nOutput: %s", err, string(out))
	}
}
