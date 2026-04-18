package cmd

import (
	"testing"

	"github.com/spf13/cobra"
)

func findCommandByName(parent *cobra.Command, use string) *cobra.Command {
	for _, c := range parent.Commands() {
		if c.Name() == use {
			return c
		}
	}
	return nil
}

func TestVersionCommandIsRegistered(t *testing.T) {
	if got := findCommandByName(rootCmd, "version"); got == nil {
		t.Fatalf("expected version command to be registered on root")
	}
}

func TestRootVersionMatchesConstant(t *testing.T) {
	if rootCmd.Version != CLI_VERSION {
		t.Fatalf("expected root version %q to match CLI version %q", rootCmd.Version, CLI_VERSION)
	}
}

func TestK8sInitCommandIsRegistered(t *testing.T) {
	k8s := findCommandByName(rootCmd, "k8s")
	if k8s == nil {
		t.Fatalf("expected k8s command to be registered on root")
	}

	if got := findCommandByName(k8s, "init"); got == nil {
		t.Fatalf("expected k8s init subcommand to be registered")
	}
}
