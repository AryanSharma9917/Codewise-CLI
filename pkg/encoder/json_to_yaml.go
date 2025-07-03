package encoder

import (
	"encoding/json"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func JSONToYAML(inputFile, outputFile string) error {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		return fmt.Errorf("cannot read file: %w", err)
	}

	var raw interface{}
	if err := json.Unmarshal(data, &raw); err != nil {
		return fmt.Errorf("JSON unmarshal failed: %w", err)
	}

	yamlData, err := yaml.Marshal(raw)
	if err != nil {
		return fmt.Errorf("YAML marshal failed: %w", err)
	}

	if err := os.WriteFile(outputFile, yamlData, 0644); err != nil {
		return fmt.Errorf("cannot write file: %w", err)
	}

	return nil
}
