package cmd

import (
	"fmt"
	"os"

	"github.com/aryansharma9917/Codewise-CLI/pkg/encoder"
	"github.com/spf13/cobra"
)

var (
	inputFile  string
	outputFile string
	jsonToYAML bool
)

var encodeCmd = &cobra.Command{
	Use:   "encode",
	Short: "Convert files between formats (YAML <-> JSON, base64, .env)",
	Run: func(cmd *cobra.Command, args []string) {
		if inputFile == "" || outputFile == "" {
			fmt.Println("❌ Please provide both --input and --output flags")
			os.Exit(1)
		}

		var err error

		if jsonToYAML {
			err = encoder.JSONToYAML(inputFile, outputFile)
		} else {
			err = encoder.YAMLToJSON(inputFile, outputFile)
		}

		if err != nil {
			fmt.Println("❌ Error:", err)
			os.Exit(1)
		}

		if jsonToYAML {
			fmt.Println("✅ JSON converted to YAML successfully.")
		} else {
			fmt.Println("✅ YAML converted to JSON successfully.")
		}
	},
}

func init() {
	rootCmd.AddCommand(encodeCmd)

	encodeCmd.Flags().StringVarP(&inputFile, "input", "i", "", "Input file path")
	encodeCmd.Flags().StringVarP(&outputFile, "output", "o", "", "Output file path")
	encodeCmd.Flags().BoolVar(&jsonToYAML, "json-to-yaml", false, "Convert JSON to YAML instead")
}
