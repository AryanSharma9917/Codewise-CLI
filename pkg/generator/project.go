package generator

import (
	"fmt"
	"os"
)

func ScaffoldProject() {
	fmt.Println("🚀 Initializing project...")

	files := map[string]string{
		"Dockerfile":      "FROM alpine\nCMD echo Hello Codewise",
		"deployment.yaml": "apiVersion: apps/v1\nkind: Deployment\nmetadata:\n  name: app",
		".env.example":    "PORT=8080\nENV=dev",
		".gitignore":      "node_modules\n*.log\ndist/",
	}

	for name, content := range files {
		if err := os.WriteFile(name, []byte(content), 0644); err != nil {
			fmt.Printf("❌ Could not write %s: %v\n", name, err)
			continue
		}
		fmt.Printf("✅ Created %s\n", name)
	}
}
