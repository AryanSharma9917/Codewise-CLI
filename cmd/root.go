package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var version string = "v1.1.0" 

var rootCmd = &cobra.Command{
	Use:   "codewise",
	Short: "CLI that helps you scaffold, encode, validate, and automate DevOps workflows easily.",
	Long: `Codewise is a CLI tool designed for DevOps and SREs to:

- Convert between YAML, JSON, .env and base64
- Generate GitHub Actions or ArgoCD templates
- Scaffold new GitOps-ready projects with best practices`,
	Run: func(cmd *cobra.Command, args []string) {
		showVersion, _ := cmd.Flags().GetBool("version")
		if showVersion {
			fmt.Printf("Codewise CLI version: %s\n", version)
			return
		}
		PrintBanner()
		cmd.Help()
	},
}

// PrintBanner displays an ASCII banner
func PrintBanner() {
	fmt.Println(`
   ____          _       _           
  / ___|___   __| | ___ | |__  _   _ 
 | |   / _ \ / _  |/ _ \| '_ \| | | |
 | |__| (_) | (_| | (_) | |_) | |_| |
  \____\___/ \__,_|\___/|_.__/ \__, |
                               |___/ 
Codewise CLI - Simplify your DevOps
`)
}

func init() {
	rootCmd.Flags().BoolP("version", "v", false, "Print the version")
}

// Execute is the entry point
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
