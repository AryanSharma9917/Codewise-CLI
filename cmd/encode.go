package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/aryansharma9917/Codewise-CLI/pkg/encoder"
)

var (
	inputFile     string
	outputFile    string
	jsonToYAML    bool
	base64Mode    bool
	base64Decode  bool
	envToJSON     bool

)

var encodeCmd = &cobra.Command{
	Use:   "encode",
	Short: "Convert between formats (YAML ⇄ JSON, base64, .env)",
	Run: func(cmd *cobra.Command, args []string) {
		if inputFile == "" || outputFile == "" {
			fmt.Println("❌ Please provide both --input and --output flags")
			os.Exit(1)
		}

		var err error

		switch {
			case envToJSON:
				err = encoder.EnvToJSON(inputFile, outputFile)
			case base64Mode && base64Decode:
				err = encoder.Base64Decode(inputFile, outputFile)
			case base64Mode:
				err = encoder.Base64Encode(inputFile, outputFile)
			case jsonToYAML:
				err = encoder.JSONToYAML(inputFile, outputFile)
			default:
				err = encoder.YAMLToJSON(inputFile, outputFile)
		}


		if err != nil {
			fmt.Println("❌ Error:", err)
			os.Exit(1)
		}

		switch {
			case envToJSON:
				fmt.Println("✅ .env file converted to JSON successfully.")
			case base64Mode && base64Decode:
				fmt.Println("✅ Base64 decoded successfully.")
			case base64Mode:
				fmt.Println("✅ Base64 encoded successfully.")
			case jsonToYAML:
				fmt.Println("✅ JSON converted to YAML successfully.")
			default:
				fmt.Println("✅ YAML converted to JSON successfully.")
		}
	},
}

func init() {
	rootCmd.AddCommand(encodeCmd)

	encodeCmd.Flags().StringVarP(&inputFile, "input", "i", "", "Input file path")
	encodeCmd.Flags().StringVarP(&outputFile, "output", "o", "", "Output file path")
	encodeCmd.Flags().BoolVar(&jsonToYAML, "json-to-yaml", false, "Convert JSON to YAML")
	encodeCmd.Flags().BoolVar(&base64Mode, "base64", false, "Perform base64 encode/decode")
	encodeCmd.Flags().BoolVar(&base64Decode, "decode", false, "Decode base64 instead of encode (use with --base64)")
	encodeCmd.Flags().BoolVar(&envToJSON, "env-to-json", false, "Convert .env file to JSON")
}
