package encoder_test

import (
	"os/exec"
	"testing"
)

func TestBase64Encode(t *testing.T) {
	input := "../../testdata/sample.txt"
	output := "../../testdata/out.b64"

	cmd := exec.Command("../../codewise", "encode", "--input", input, "--output", output)
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("❌ Command failed: %v\nOutput: %s", err, string(out))
	}
}
