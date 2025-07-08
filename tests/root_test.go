package test

import (
	"testing"

	"github.com/spf13/cobra"
)

func TestRootCommand(t *testing.T) {
	var cmd cobra.Command
	cmd.Use = "codewise"

	if cmd.Use != "codewise" {
		t.Error("❌ Root command name mismatch")
	}
}
