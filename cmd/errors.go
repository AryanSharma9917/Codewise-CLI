package cmd

import (
	"errors"
	"fmt"
	"os"
)

// ExitError prints an error message and exits with status code 1.
func ExitError(message string) error {
	fmt.Fprintf(os.Stderr, "error: %s\n", message)
	return errors.New(message)
}

// LogError prints an info/warning message but does not exit.
func LogError(message string) error {
	fmt.Fprintf(os.Stderr, "info: %s\n", message)
	return errors.New(message)
}

// LogErrorf prints a formatted info/warning message but does not exit.
func LogErrorf(format string, args ...interface{}) error {
	msg := fmt.Sprintf(format, args...)
	fmt.Fprintf(os.Stderr, "info: %s\n", msg)
	return errors.New(msg)
}

// LogInfo prints an info message to stderr.
func LogInfo(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "info: %s\n", fmt.Sprintf(format, args...))
}

// LogSuccess prints a success message to stdout.
func LogSuccess(format string, args ...interface{}) {
	fmt.Printf("✅ %s\n", fmt.Sprintf(format, args...))
}
