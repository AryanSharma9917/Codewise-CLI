package encoder

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func EnvToJSON(inputFile, outputFile string) error {
	file, err := os.Open(inputFile)
	if err != nil {
		return fmt.Errorf("cannot open .env file: %w", err)
	}
	defer file.Close()

	envMap := make(map[string]string)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		// Skip empty lines or comments
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue // ignore malformed lines
		}

		key := strings.TrimSpace(parts[0])
		val := strings.TrimSpace(parts[1])
		envMap[key] = val
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading .env: %w", err)
	}

	jsonData, err := json.MarshalIndent(envMap, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshalling JSON: %w", err)
	}

	return os.WriteFile(outputFile, jsonData, 0644)
}
