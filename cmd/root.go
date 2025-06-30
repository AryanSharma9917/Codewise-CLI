package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "codewise",
	Short: "Codewise-CLI is a modular DevOps CLI tool",
	Long:  "CLI that helps you scaffold, encode, validate, and automate DevOps workflows easily.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
