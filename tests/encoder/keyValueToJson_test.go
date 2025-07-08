package encoder_test

import (
	"os"
	"os/exec"
	"testing"
)

func TestKeyValueToJSON(t *testing.T) {
	cmd := exec.Command("./codewise", "KVTJ", "--env=testdata/sample.env", "--output=testdata/env_output.json")

	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("❌ Command failed: %v\nOutput: %s", err, output)
	}

	if _, err := os.Stat("testdata/env_output.json"); os.IsNotExist(err) {
		t.Fatal("❌ Output file was not created")
	}

	defer os.Remove("testdata/env_output.json")
}
