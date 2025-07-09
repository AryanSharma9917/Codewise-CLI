package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "codewise",
	Short: "CLI that helps you scaffold, encode, validate, and automate DevOps workflows easily.",
	Long: `Codewise CLI v1.1.0

This CLI helps you with common DevOps tasks like:
- JSON/YAML conversions
- Base64 encoding/decoding
- Dockerfile and Kubernetes manifest generation
- Project scaffolding and GitHub Actions template rendering
`,
	Version: "v1.1.0",
}

// Execute runs the root command.
func Execute() {
	PrintBanner()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("❌", err)
		os.Exit(1)
	}
}

// PrintBanner prints the CLI banner.
func PrintBanner() {
	fmt.Print(`
   ____          _       _           
  / ___|___   __| | ___ | |__  _   _ 
 | |   / _ \ / _  |/ _ \| '_ \| | | |
 | |__| (_) | (_| | (_) | |_) | |_| |
  \____\___/ \__,_|\___/|_.__/ \__, |
                               |___/ 
Codewise CLI - Simplify your DevOps
`)
}
