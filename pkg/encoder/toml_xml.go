package encoder

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/clbanning/mxj/v2"
)

// JSON → TOML
func JSONToTOML(inputFile, outputFile string) error {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		return fmt.Errorf("cannot read JSON file: %w", err)
	}

	var jsonData map[string]interface{}
	if err := json.Unmarshal(data, &jsonData); err != nil {
		return fmt.Errorf("failed to parse JSON: %w", err)
	}

	var buf bytes.Buffer
	if err := toml.NewEncoder(&buf).Encode(jsonData); err != nil {
		return fmt.Errorf("failed to encode TOML: %w", err)
	}

	return os.WriteFile(outputFile, buf.Bytes(), 0644)
}

// TOML → JSON
func TOMLToJSON(inputFile, outputFile string) error {
	var tomlData map[string]interface{}
	if _, err := toml.DecodeFile(inputFile, &tomlData); err != nil {
		return fmt.Errorf("failed to decode TOML: %w", err)
	}

	data, err := json.MarshalIndent(tomlData, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to encode JSON: %w", err)
	}

	return os.WriteFile(outputFile, data, 0644)
}

// JSON → XML
func JSONToXML(inputFile, outputFile string) error {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		return fmt.Errorf("cannot read JSON file: %w", err)
	}

	var jsonData map[string]interface{}
	if err := json.Unmarshal(data, &jsonData); err != nil {
		return fmt.Errorf("failed to parse JSON: %w", err)
	}

	mv := mxj.Map(jsonData)
	xmlData, err := mv.XmlIndent("", "  ")
	if err != nil {
		return fmt.Errorf("failed to encode XML: %w", err)
	}

	return os.WriteFile(outputFile, xmlData, 0644)
}

// XML → JSON
func XMLToJSON(inputFile, outputFile string) error {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		return fmt.Errorf("cannot read XML file: %w", err)
	}

	mv, err := mxj.NewMapXml(data)
	if err != nil {
		return fmt.Errorf("failed to parse XML: %w", err)
	}

	jsonData, err := json.MarshalIndent(mv, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to encode JSON: %w", err)
	}

	return os.WriteFile(outputFile, jsonData, 0644)
}
