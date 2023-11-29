package test

import (
	"os/exec"
	"strings"
	"testing"
)

// TestRootCmd tests the root command (Codewise-CLI)
func TestRootCmd(t *testing.T) {

	expectedOutput := "Do all your tedious tasks with a single command"
	cmd := exec.Command("Codewise-CLI")

	// Capture the output
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Errorf("expected no error, but got: %v", err)
	}

	// Validate the cli output
	got := strings.TrimSpace(string(output)[:47])
	if got != expectedOutput {
		t.Errorf("expected %v, but got: %v", expectedOutput, got)
	}

}

// TestRootCmdHelpFlag tests the root command (Codewise-CLI) with help flag
func TestRootCmdHelpFlag(t *testing.T) {

	expectedOutput := "Do all your tedious tasks with a single command"
	cmd := exec.Command("Codewise-CLI", "--help")

	// Capture the output
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Errorf("expected no error, but got: %v", err)
	}

	// Validate the cli output
	got := strings.TrimSpace(string(output)[:47])
	if got != expectedOutput {
		t.Errorf("expected %v, but got: %v", expectedOutput, got)
	}

}
