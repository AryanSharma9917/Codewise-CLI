package generator

import (
	"bytes"
	"fmt"
	"os"
	"text/template"
)

type TemplateData struct {
	AppName string
	Repo    string
}

// RenderTemplate renders a template by name from the /templates directory (for CLI usage).
func RenderTemplate(templateName, outputPath string, data TemplateData) error {
	templatePath := fmt.Sprintf("templates/%s.tpl", templateName)
	return RenderTemplateFromFile(templatePath, outputPath, data)
}

// RenderTemplateFromFile reads and renders a template file using provided data (for test usage).
func RenderTemplateFromFile(templatePath, outputPath string, data TemplateData) error {
	tplContent, err := os.ReadFile(templatePath)
	if err != nil {
		return fmt.Errorf("❌ Failed to read template: %v", err)
	}

	tpl, err := template.New("tpl").Parse(string(tplContent))
	if err != nil {
		return fmt.Errorf("❌ Failed to parse template: %v", err)
	}

	var buf bytes.Buffer
	if err := tpl.Execute(&buf, data); err != nil {
		return fmt.Errorf("❌ Failed to execute template: %v", err)
	}

	if err := os.WriteFile(outputPath, buf.Bytes(), 0644); err != nil {
		return fmt.Errorf("❌ Failed to write file: %v", err)
	}

	return nil
}
