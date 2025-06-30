package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/aryansharma9917/Codewise-CLI/pkg/encoder"
)

var encodeCmd = &cobra.Command{
	Use:   "encode",
	Short: "Convert YAML to JSON",
	Run: func(cmd *cobra.Command, args []string) {
		if err := encoder.YAMLToJSON("input.yaml", "output.json"); err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		fmt.Println("✅ YAML converted to JSON successfully.")
	},
}

func init() {
	rootCmd.AddCommand(encodeCmd)
}
