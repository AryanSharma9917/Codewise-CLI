package generator_test

import (
	"os"
	"testing"

	"github.com/aryansharma9917/Codewise-CLI/pkg/generator"
)

func TestRenderTemplate(t *testing.T) {
	t.Helper()

	templateName := "github-action"
	output := "tests/testdata/test-output.yml"

	data := generator.TemplateData{
		AppName: "testapp",
		Repo:    "https://github.com/test/repo",
	}

	err := generator.RenderTemplate(templateName, output, data)
	if err != nil {
		t.Fatalf("❌ RenderTemplate failed: %v", err)
	}

	// Confirm that output file was created
	if _, err := os.Stat(output); os.IsNotExist(err) {
		t.Fatalf("❌ Output file not created: %s", output)
	}

	// Clean up
	defer func() {
		if err := os.Remove(output); err != nil {
			t.Logf("⚠️ Failed to clean up output file: %v", err)
		}
	}()
}
