package encoder

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	inputJsonFile  string
	outputYamlFile string
)

// jsonToYamlCmd is the command for converting JSON to YAML
var jsonToYamlCmd = &cobra.Command{
	Use:   "JTY [flags]",
	Short: "Converts a JSON into YAML.",
	Run: func(cmd *cobra.Command, args []string) {
		convertJsonToYaml()
	},
}

// init initializes flags and sets required parameters
func init() {
	// Flags for the JYT command
	jsonToYamlCmd.Flags().StringVarP(&outputYamlFile, "output", "o", "", "Output YAML file name (default is output.yaml)")
	jsonToYamlCmd.Flags().StringVarP(&inputJsonFile, "file", "f", "", "Input JSON file name")
	err := jsonToYamlCmd.MarkFlagRequired("file")
	checkNilErr(err)
}

// convertJsonToYaml handles the JSON to YAML conversion
func convertJsonToYaml() {
	// Initialize viper and read the JSON file
	vp := viper.New()
	vp.SetConfigFile(inputJsonFile)
	err := vp.ReadInConfig()
	checkNilErr(err)

	// Set output file to default if not provided
	if outputYamlFile == "" {
		outputYamlFile = "output.yaml"
	}
	vp.SetConfigFile(outputYamlFile)

	// Write the YAML file
	err = vp.WriteConfig()
	checkNilErr(err)

	// Output completion message
	fmt.Printf("Operation completed successfully. Check the %s file.\n", outputYamlFile)
}

//  Declaration of checkNilErr function to handle errors
// checkNilErr checks if the error is nil and panics if it is not
func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
