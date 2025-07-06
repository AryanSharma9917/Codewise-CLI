package encoder

import (
	"encoding/base64"
	"fmt"
	"os"
)

func Base64Encode(inputFile, outputFile string) error {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		return fmt.Errorf("cannot read file: %w", err)
	}

	encoded := base64.StdEncoding.EncodeToString(data)
	return os.WriteFile(outputFile, []byte(encoded), 0644)
}

func Base64Decode(inputFile, outputFile string) error {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		return fmt.Errorf("cannot read file: %w", err)
	}

	decoded, err := base64.StdEncoding.DecodeString(string(data))
	if err != nil {
		return fmt.Errorf("base64 decode error: %w", err)
	}

	return os.WriteFile(outputFile, decoded, 0644)
}

// echo "hello codewise" > input.txt
// go run main.go encode --base64 --input=input.txt --output=encoded.txt


// go run main.go encode --base64 --decode --input=encoded.txt --output=output.txt
// cat output.txt  # => hello codewise
