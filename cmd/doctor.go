package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Check Codewise CLI environment",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Codewise CLI Doctor")
		fmt.Println("-------------------")
		fmt.Println("Go version:", runtime.Version())
		fmt.Println("OS/Arch:", runtime.GOOS, runtime.GOARCH)
		fmt.Println("Codewise version:", rootCmd.Version)
	},
}

func init() {
	rootCmd.AddCommand(doctorCmd)
}
