package encoder_test

import (
	"os"
	"os/exec"
	"testing"
)

func TestYAMLToJSON(t *testing.T) {
	cmd := exec.Command("./codewise", "encode", "--input=testdata/sample.yaml", "--output=testdata/YTJ_output.json")

	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("❌ Command failed: %v\nOutput: %s", err, output)
	}

	if _, err := os.Stat("testdata/YTJ_output.json"); os.IsNotExist(err) {
		t.Fatal("❌ Output file was not created")
	}

	defer os.Remove("testdata/YTJ_output.json")
}
