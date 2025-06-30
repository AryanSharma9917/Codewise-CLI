package encoder

import (
	"encoding/json"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func YAMLToJSON(inputFile, outputFile string) error {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		return fmt.Errorf("cannot read file: %w", err)
	}

	var raw interface{}
	if err := yaml.Unmarshal(data, &raw); err != nil {
		return fmt.Errorf("YAML unmarshal failed: %w", err)
	}

	jsonData, err := json.MarshalIndent(raw, "", "  ")
	if err != nil {
		return fmt.Errorf("JSON marshal failed: %w", err)
	}

	if err := os.WriteFile(outputFile, jsonData, 0644); err != nil {
		return fmt.Errorf("cannot write file: %w", err)
	}

	return nil
}


// Example usage:
// apiVersion: v1
// kind: Pod
// metadata:
//   name: example-pod
