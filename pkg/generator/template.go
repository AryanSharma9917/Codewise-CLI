package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

// TemplateData holds dynamic values used in templates
type TemplateData struct {
	AppName string
	Repo    string
}

// RenderTemplate renders a Go template from the templates/ folder into an output file.
func RenderTemplate(name, output string, data TemplateData) error {
	// Build full path to the template file
	tmplPath := filepath.Join("templates", name+".tpl")

	// Parse the template file
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		return fmt.Errorf("❌ Failed to parse template: %w", err)
	}

	// Create the output file
	out, err := os.Create(output)
	if err != nil {
		return fmt.Errorf("❌ Failed to create output file: %w", err)
	}
	defer out.Close()

	// Execute the template with the provided data
	if err := tmpl.Execute(out, data); err != nil {
		return fmt.Errorf("❌ Failed to render template: %w", err)
	}

	return nil
}
