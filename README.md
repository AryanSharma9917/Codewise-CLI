# Codewise-CLI

<div align="center">
  <img src="/logo/logo_pic.png" alt="Codewise Logo" width="220"/>
</div>

Codewise is a powerful CLI tool that simplifies common DevOps tasks such as:

  * Encoding/decoding files
  * JSON/YAML conversions
  * Dockerfile generation
  * Kubernetes manifest scaffolding
  * Rendering templates with Go's `text/template` engine

-----

## 🚀 Getting Started

### Clone the Repository

```bash
git clone https://github.com/aryansharma9917/Codewise-CLI.git
cd Codewise-CLI
```

### Build the Binary

```bash
go build -o codewise main.go
```

### (Optional) Install Globally

```bash
sudo mv codewise /usr/local/bin/
```

-----

## Usage

```bash
codewise <command> [flags]
```

-----

## Commands & Examples

### `encode` — Format Conversion & Encoding

| Conversion Type | Description           | Example                                                                     |
| :-------------- | :-------------------- | :-------------------------------------------------------------------------- |
| `JTY`           | JSON to YAML          | `codewise encode --input input.json --output output.yaml --type JTY`      |
| `YTJ`           | YAML to JSON          | `codewise encode --input input.yaml --output output.json --type YTJ`      |
| `KVTJ`          | .env to JSON          | `codewise encode --input .env --output env.json --type KVTJ`              |
| `B64E`          | Base64 Encode         | `codewise encode --input input.txt --output encoded.txt --type B64E`      |
| `B64D`          | Base64 Decode         | `codewise encode --input encoded.txt --output decoded.txt --type B64D`    |

### `generate` — Starter File Generators

Generate a Dockerfile:

```bash
codewise generate dockerfile --output Dockerfile
```

Generate a Kubernetes manifest:

```bash
codewise generate k8s --output deployment.yaml
```

### `template` — Render Templated YAMLs

Use Go `.tpl` template and `.yaml` values file to generate output YAML:

```bash
codewise template --template template.tpl --values values.yaml --output rendered.yaml
```

-----

## 🐳 Docker Usage

### 🔨 Build Docker Image

```bash
docker build -t aryansharma04/codewise-cli:latest .
```

### Run Using Docker

```bash
docker run --rm -v $(pwd):/app aryansharma04/codewise-cli:latest <command>
```

Example:

```bash
docker run --rm -v $(pwd):/app aryansharma04/codewise-cli:latest encode --input /app/input.json --output /app/output.yaml --type JTY
```

-----

## ✅ Running Tests

Make sure Go is installed:

```bash
go test ./... -v
```

Test coverage includes:

  * ✅ JSON to YAML
  * ✅ YAML to JSON
  * ✅ ENV to JSON
  * ✅ Base64 encode/decode
  * ✅ Template rendering

-----

## 📁 Project Structure

```bash
.
├── cmd/             # CLI command handlers
├── pkg/
│   ├── encoder/     # Encoding & conversion logic
│   ├── generator/   # Project scaffolding logic
│   └── validator/   # Future: Schema & config validation
├── templates/       # Go template files
├── tests/           # Unit tests
├── testdata/        # Sample input files for testing
├── Dockerfile
├── Makefile
├── go.mod
├── main.go
└── README.md
```

-----

## 🤝 Contributing

```bash
# Fork and clone the repo
git clone https://github.com/<your-username>/Codewise-CLI.git
cd Codewise-CLI

# Create a feature branch
git checkout -b my-feature

# Make changes and push
git add .
git commit -m "Add awesome feature"
git push origin my-feature
```

Then open a Pull Request 🚀

-----

## 🛡 License

Licensed under the MIT License. See `LICENSE` for more.

-----

## 📬 Author

Aryan Sharma

GitHub: [@aryansharma9917](https://github.com/aryansharma9917)
