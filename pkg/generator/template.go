package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

type TemplateData struct {
	AppName string
	Repo    string
}

func RenderTemplate(name, outputPath string, data TemplateData) {
	templateMap := map[string]string{
		"github-action": "templates/github-action.tpl",
		"argo-app":      "templates/argo-app.tpl",
	}

	tplPath, ok := templateMap[name]
	if !ok {
		fmt.Println("❌ Unknown template:", name)
		return
	}

	tpl, err := template.ParseFiles(tplPath)
	if err != nil {
		fmt.Println("❌ Failed to parse template:", err)
		return
	}

	if err := os.MkdirAll(filepath.Dir(outputPath), 0755); err != nil {
		fmt.Println("❌ Failed to create output dir:", err)
		return
	}

	outFile, err := os.Create(outputPath)
	if err != nil {
		fmt.Println("❌ Failed to create output file:", err)
		return
	}
	defer outFile.Close()

	if err := tpl.Execute(outFile, data); err != nil {
		fmt.Println("❌ Failed to render template:", err)
		return
	}

	fmt.Println("✅ Template rendered at:", outputPath)
}
