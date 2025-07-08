package encoder_test

import (
	"os/exec"
	"path/filepath"
	"testing"
)

func TestBase64Encode(t *testing.T) {
	input := filepath.Join("..", "..", "testdata", "sample.txt")
	output := filepath.Join("..", "..", "testdata", "out.b64")
	binPath := filepath.Join("..", "..", "codewise")

	cmd := exec.Command(binPath, "encode", "--input", input, "--output", output)
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("❌ Command failed: %v\nOutput: %s", err, string(out))
	}
}
