package encoder_test

import (
	"os/exec"
	"testing"
)

func TestKeyValueToJSON(t *testing.T) {
	input := "../../testdata/sample.env"
	output := "../../testdata/out.json"

	cmd := exec.Command("codewise", "encode", "--input", input, "--output", output)
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("❌ Command failed: %v\nOutput: %s", err, string(out))
	}
}
