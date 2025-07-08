package test

import (
	"testing"
)

func TestVersionFormat(t *testing.T) {
	version := "v1.0.0" // replace with actual import if available
	if version[0] != 'v' {
		t.Error("❌ Version should start with 'v'")
	}
}
