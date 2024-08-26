package docker

import (
	"os/exec"
	"strings"
	"testing"
)

// TestDockerCmd tests the docker command
func TestDockerCmd(t *testing.T) {

	expectedOutput := "Docker related commands."
	cmd := exec.Command("Codewise-CLI", "docker")

	// Capture the output for redention
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Errorf("expected no error, but got: %v", err)
	}

	// Validate the cli output
	got := strings.TrimSpace(string(output)[:24])
	if got != expectedOutput {
		t.Errorf("expected %v, but got: %v", expectedOutput, got)
	}

}
