package generator_test

import (
	"os"
	"testing"

	"github.com/aryansharma9917/Codewise-CLI/pkg/generator"
)

func TestRenderTemplate(t *testing.T) {
	templateName := "github-action"
	output := "tests/testdata/test-output.yml"

	data := generator.TemplateData{
		AppName: "testapp",
		Repo:    "https://github.com/test/repo",
	}

	err := generator.RenderTemplate(templateName, output, data)
	if err != nil {
		t.Errorf("❌ RenderTemplate failed: %v", err)
	}

	// Clean up
	defer os.Remove(output)
}
