package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/spf13/cobra"
)

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Check Codewise CLI environment",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Codewise CLI Doctor")
		fmt.Println("-------------------")

		// Go info
		fmt.Println("Go version:", runtime.Version())
		fmt.Println("OS/Arch:", runtime.GOOS, runtime.GOARCH)

		// Codewise version
		fmt.Println("Codewise version:", rootCmd.Version)

		// Working directory
		if wd, err := os.Getwd(); err == nil {
			fmt.Println("Working directory:", wd)
		}

		// Git version
		if out, err := exec.Command("git", "--version").Output(); err == nil {
			fmt.Println("Git:", string(out))
		} else {
			fmt.Println("Git: not found")
		}

		// Config path (future use)
		configPath := os.ExpandEnv("$HOME/.codewise/config.yaml")
		if _, err := os.Stat(configPath); err == nil {
			fmt.Println("Config file:", configPath)
		} else {
			fmt.Println("Config file:", configPath, "(not found)")
		}
	},
}

func init() {
	rootCmd.AddCommand(doctorCmd)
}
