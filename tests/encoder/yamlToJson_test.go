package encoder_test

import (
	"os/exec"
	"testing"
)

func TestYAMLToJSON(t *testing.T) {
	input := "../../testdata/sample.yaml"
	output := "../../testdata/out.json"

	cmd := exec.Command("codewise", "encode", "--input", input, "--output", output)
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("❌ Command failed: %v\nOutput: %s", err, string(out))
	}
}
