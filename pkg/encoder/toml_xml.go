package encoder

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"os"

	"github.com/BurntSushi/toml"
)

// JSON → TOML
func JSONToTOML(inputFile, outputFile string) error {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		return err
	}

	var jsonData map[string]interface{}
	if err := json.Unmarshal(data, &jsonData); err != nil {
		return err
	}

	var buf bytes.Buffer
	if err := toml.NewEncoder(&buf).Encode(jsonData); err != nil {
		return err
	}

	return os.WriteFile(outputFile, buf.Bytes(), 0644)
}

// TOML → JSON
func TOMLToJSON(inputFile, outputFile string) error {
	var tomlData map[string]interface{}
	if _, err := toml.DecodeFile(inputFile, &tomlData); err != nil {
		return err
	}

	data, err := json.MarshalIndent(tomlData, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(outputFile, data, 0644)
}

// JSON → XML
func JSONToXML(inputFile, outputFile string) error {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		return err
	}

	var jsonData map[string]interface{}
	if err := json.Unmarshal(data, &jsonData); err != nil {
		return err
	}

	xmlData, err := xml.MarshalIndent(jsonData, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(outputFile, xmlData, 0644)
}

// XML → JSON
func XMLToJSON(inputFile, outputFile string) error {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		return err
	}

	var xmlData map[string]interface{}
	if err := xml.Unmarshal(data, &xmlData); err != nil {
		return err
	}

	jsonData, err := json.MarshalIndent(xmlData, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(outputFile, jsonData, 0644)
}
