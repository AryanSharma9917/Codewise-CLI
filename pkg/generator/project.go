package generator

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func ScaffoldProject(projectName string, withDocker, withDeployment bool) {
	if projectName == "" {
		fmt.Println("❌ Please provide a project name using --project")
		os.Exit(1)
	}

	basePath := filepath.Join(".", projectName)

	// Folder structure
	dirs := []string{
		"cmd", "pkg", "internal", "configs", "scripts", "templates",
	}

	for _, dir := range dirs {
		fullPath := filepath.Join(basePath, dir)
		if err := os.MkdirAll(fullPath, 0755); err != nil {
			fmt.Printf("❌ Failed to create %s: %v\n", fullPath, err)
			os.Exit(1)
		}
	}

	// Dockerfile
	if withDocker {
		dockerfile := `FROM golang:1.21
WORKDIR /app
COPY . .
RUN go build -o main .
CMD ["./main"]`

		err := os.WriteFile(filepath.Join(basePath, "Dockerfile"), []byte(dockerfile), 0644)
		if err != nil {
			fmt.Println("❌ Failed to write Dockerfile:", err)
			os.Exit(1)
		}
		fmt.Println("📦 Dockerfile created.")
	}

	// Kubernetes deployment.yaml
	if withDeployment {
		k8sPath := filepath.Join(basePath, "k8s")
		_ = os.MkdirAll(k8sPath, 0755)

		deployment := `apiVersion: apps/v1
kind: Deployment
metadata:
  name: ` + projectName + `
spec:
  replicas: 1
  selector:
    matchLabels:
      app: ` + projectName + `
  template:
    metadata:
      labels:
        app: ` + projectName + `
    spec:
      containers:
      - name: ` + projectName + `
        image: ` + projectName + `:latest
        ports:
        - containerPort: 8080`

		err := os.WriteFile(filepath.Join(k8sPath, "deployment.yaml"), []byte(deployment), 0644)
		if err != nil {
			fmt.Println("❌ Failed to write deployment.yaml:", err)
			os.Exit(1)
		}
		fmt.Println("📄 k8s/deployment.yaml created.")
	}

	setupGitRepo(basePath)

	fmt.Println("✅ Project scaffolded successfully.")
}

func setupGitRepo(basePath string) {
	cmd := exec.Command("git", "init")
	cmd.Dir = basePath
	if err := cmd.Run(); err != nil {
		fmt.Println("❌ Failed to initialize Git:", err)
		return
	}

	gitignore := `# Binaries
*.exe
*.dll
*.so
*.dylib
*.test
*.out

# Vendor
/vendor/

# Logs
*.log

# IDEs and editors
.vscode/
.idea/
*.swp

# Build
/build/
bin/
`

	err := os.WriteFile(filepath.Join(basePath, ".gitignore"), []byte(gitignore), 0644)
	if err != nil {
		fmt.Println("❌ Failed to write .gitignore:", err)
		return
	}

	fmt.Println("🔧 Git repo initialized with .gitignore")
}
