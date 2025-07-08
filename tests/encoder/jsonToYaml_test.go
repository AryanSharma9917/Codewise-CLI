package encoder_test

import (
	"os"
	"os/exec"
	"testing"
)

func TestJSONToYAML(t *testing.T) {
	cmd := exec.Command("./codewise", "JTY", "--file=testdata/JTY.json", "--output=testdata/JTY_output.yaml")

	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("❌ Command failed: %v\nOutput: %s", err, output)
	}

	if _, err := os.Stat("testdata/JTY_output.yaml"); os.IsNotExist(err) {
		t.Fatal("❌ Output file was not created")
	}

	defer os.Remove("testdata/JTY_output.yaml")
}
