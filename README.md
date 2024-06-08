# Codewise

<div align="center">


[![Releases](https://github.com/aryansharma9917/Codewise-CLI/actions/workflows/releases.yml/badge.svg)](https://github.com/aryansharma9917/Codewise-CLI/actions/workflows/releases.yml) [![Go Reference](https://pkg.go.dev/github.com/aryansharma9917/Codewise-CLI)]

</div>

**Codewise** is a CLI tool that provides a basic set of commands to perform tedious tasks such as converting **YAML to JSON** or **JSON to YAML** directly from your terminal. It's built with [Go](https://github.com/golang/go), [Cobra](https://github.com/spf13/cobra), [Viper](https://github.com/spf13/viper), etc.
 
To install the Codewise-CLI, use the command `go install github.com/AryanSharma9917/Codewise-CLI@latest`.
Go will automatically install it in your `$GOPATH/bin` directory, which should be in your `$PATH`.

Once installed, you can use the `codewise-CLI` CLI command. To confirm installation, type `codewise-cli` at the command line.

> **Note** If you are getting an error like `command not found: Codewise-CLI`, then you need to add `$GOPATH/bin` to your `$PATH` environment variable.

## 📚 Documentation

**Check out detailed documentation for getting started and using Codewise CLI** [**here**](https://pkg.go.dev/github.com/aryansharma9917/Codewise-CLI).

## ⭐️ Features

- Convert JSON file to YAML
- Convert YAML file to JSON
- Convert Key-Value to JSON
- Generate Dockerfile for different languages/framework
- Generate Kubernetes manifests for different objects
- Encode and Decode a string to base64
- More coming soon...

## 📝 Usage

```
Usage:
  codewise-cli [command] [flags]
  codewise-cli [command]

Available Commands:
  JTY         Converts a JSON into YAML.
  KVTJ        Converts Key-Value (text) to JSON.
  YTJ         Converts a YAML into JSON.
  docker      Docker related commands. Like generating a Dockerfile for a language.
  encode      It encodes and decodes a string to base64 and vice versa.
  k8s         Kubernetes related commands. Like generating manifest files for kubernetes objects.
```

For detailed usage of each command, visit [here](https://pkg.go.dev/github.com/aryansharma9917/Codewise-CLI)

Eg `codewise-cli JTY --file test.json` with convert JSON into YAML and create a `output.yaml` in your current directory.

Eg `codewise-cli docker dockerfile --lang go` to generate a `Dockerfile` template for go.

eg `codewise-cli k8s manifest --obj deployment` to generate a `deployment.yaml` file with deployment template.

## 📜 License

This project is licensed under the Apache-2.0 license - see the [LICENSE](LICENSE) file for details.

## 🛡 Security

If you discover a security vulnerability within this project, please check the [SECURITY](SECURITY.md) for more information.
