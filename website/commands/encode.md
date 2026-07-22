# Encode command

```bash
codewise encode --input FILE --output FILE [flags]
```

| Flag | Short | Description |
| --- | --- | --- |
| `--input` | `-i` | Input path |
| `--output` | `-o` | Output path |
| `--json-to-yaml` | | Convert JSON to YAML |
| `--env-to-json` | | Convert an ENV file to JSON |
| `--base64` | | Use Base64 mode |
| `--decode` | | Decode rather than encode in Base64 mode |
| `--force` | `-f` | Overwrite an existing output file |

## Examples

```bash
# YAML to JSON (default structured conversion)
codewise encode -i app.yaml -o app.json

# JSON to YAML
codewise encode -i app.json -o app.yaml --json-to-yaml

# ENV to JSON
codewise encode -i .env -o env.json --env-to-json

# Base64 round trip
codewise encode -i input.txt -o encoded.txt --base64
codewise encode -i encoded.txt -o decoded.txt --base64 --decode
```

TOML and XML conversions are inferred from input and output file extensions.
