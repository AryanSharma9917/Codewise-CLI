package cmd

import (
	"fmt"
	"os"
)

// ExitError prints an error message and exits with status code 1.
func ExitError(format string, args ...interface{}) error {
	fmt.Fprintf(os.Stderr, "error: %s\n", fmt.Sprintf(format, args...))
	return fmt.Errorf(format, args...)
}

// LogError prints an info/warning message but does not exit.
func LogError(format string, args ...interface{}) error {
	msg := fmt.Sprintf(format, args...)
	fmt.Fprintf(os.Stderr, "info: %s\n", msg)
	return fmt.Errorf(msg)
}

// LogInfo prints an info message to stderr.
func LogInfo(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "info: %s\n", fmt.Sprintf(format, args...))
}

// LogSuccess prints a success message to stdout.
func LogSuccess(format string, args ...interface{}) {
	fmt.Printf("✅ %s\n", fmt.Sprintf(format, args...))
}
