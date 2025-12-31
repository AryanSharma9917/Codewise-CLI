package cmd

import (
	"bytes"
	"fmt"
	"os/exec"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
)

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Check Codewise CLI environment",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Codewise CLI Doctor")
		fmt.Println("-------------------")

		// Go
		fmt.Println("Go version:", runtime.Version())

		// OS
		fmt.Println("OS/Arch:", runtime.GOOS, runtime.GOARCH)

		// Codewise
		fmt.Println("Codewise version:", rootCmd.Version)

		// Docker
		checkDocker()

		// Git
		checkGit()
	},
}

func init() {
	rootCmd.AddCommand(doctorCmd)
}

// ---- helpers ----

func checkDocker() {
	cmd := exec.Command("docker", "--version")

	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		fmt.Println("Docker: not installed")
		return
	}

	fmt.Println("Docker:", strings.TrimSpace(out.String()))
}

func checkGit() {
	cmd := exec.Command("git", "--version")

	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		fmt.Println("Git: not installed")
		return
	}

	fmt.Println("Git:", strings.TrimSpace(out.String()))
}
