package prompt

import (
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
)

// Ask for input file if not provided
func AskInputFile() string {
	var path string
	prompt := &survey.Input{
		Message: "Enter the input file path:",
	}
	survey.AskOne(prompt, &path)
	return path
}

// Ask for output file if not provided
func AskOutputFile() string {
	var path string
	prompt := &survey.Input{
		Message: "Enter the output file path:",
	}
	survey.AskOne(prompt, &path)
	return path
}

// Confirm overwrite if file exists
func ConfirmOverwrite(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return true
	}

	var confirm bool
	message := fmt.Sprintf("File %s already exists. Overwrite?", file)
	survey.AskOne(&survey.Confirm{Message: message}, &confirm)
	return confirm
}
