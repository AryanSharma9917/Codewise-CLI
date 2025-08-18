package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/aryansharma9917/Codewise-CLI/pkg/encoder"
)

var (
	inputFile    string
	outputFile   string
	jsonToYAML   bool
	base64Mode   bool
	base64Decode bool
	envToJSON    bool
)

var encodeCmd = &cobra.Command{
	Use:   "encode",
	Short: "Convert between formats (YAML ⇄ JSON, base64, .env, TOML ⇄ JSON, XML ⇄ JSON)",
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

		// Auto-detect based on extensions
		case strings.HasSuffix(inputFile, ".toml") && strings.HasSuffix(outputFile, ".json"):
			err = encoder.TOMLToJSON(inputFile, outputFile)

		case strings.HasSuffix(inputFile, ".json") && strings.HasSuffix(outputFile, ".toml"):
			err = encoder.JSONToTOML(inputFile, outputFile)

		case strings.HasSuffix(inputFile, ".json") && strings.HasSuffix(outputFile, ".xml"):
			err = encoder.JSONToXML(inputFile, outputFile)

		case strings.HasSuffix(inputFile, ".xml") && strings.HasSuffix(outputFile, ".json"):
			err = encoder.XMLToJSON(inputFile, outputFile)

		default:
			err = encoder.YAMLToJSON(inputFile, outputFile)
		}

		if err != nil {
			fmt.Println("❌ Error:", err)
			os.Exit(1)
		}

		// Success messages
		switch {
		case envToJSON:
			fmt.Println("✅ .env file converted to JSON successfully.")
		case base64Mode && base64Decode:
			fmt.Println("✅ Base64 decoded successfully.")
		case base64Mode:
			fmt.Println("✅ Base64 encoded successfully.")
		case jsonToYAML:
			fmt.Println("✅ JSON converted to YAML successfully.")
		case strings.HasSuffix(inputFile, ".toml") && strings.HasSuffix(outputFile, ".json"):
			fmt.Println("✅ TOML converted to JSON successfully.")
		case strings.HasSuffix(inputFile, ".json") && strings.HasSuffix(outputFile, ".toml"):
			fmt.Println("✅ JSON converted to TOML successfully.")
		case strings.HasSuffix(inputFile, ".json") && strings.HasSuffix(outputFile, ".xml"):
			fmt.Println("✅ JSON converted to XML successfully.")
		case strings.HasSuffix(inputFile, ".xml") && strings.HasSuffix(outputFile, ".json"):
			fmt.Println("✅ XML converted to JSON successfully.")
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
// JSON → TOML