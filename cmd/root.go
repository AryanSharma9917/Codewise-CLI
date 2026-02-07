package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "codewise",
	Short: "CLI that helps you scaffold, encode, validate, and automate DevOps workflows easily.",
	Long: `Codewise CLI

A powerful platform-style CLI for DevOps workflows including:

• Environment-aware deployments
• Kubernetes orchestration
• Helm automation
• Docker tooling
• Template generation
• Encoding utilities
`,
	Run: func(cmd *cobra.Command, args []string) {

		// Banner ONLY when no subcommand is provided
		PrintBanner()

		_ = cmd.Help()
	},
	Version: "v1.1.0",
}

// Execute runs the root command.
func Execute() {

	if err := rootCmd.Execute(); err != nil {
		fmt.Println("error:", err)
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
