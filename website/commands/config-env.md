# Config and environment commands

## `config init`

Create `~/.codewise/config.yaml` with defaults:

```bash
codewise config init
```

The command leaves an existing configuration unchanged.

## `config view`

```bash
codewise config view
```

## `env create <name>`

```bash
codewise env create dev
codewise env create staging --interactive
```

| Flag | Short | Description |
| --- | --- | --- |
| `--interactive` | `-i` | Prompt for namespace, context, repository, branch, image, and tag |

## `env list`

```bash
codewise env list
```

## `env delete <name>`

```bash
codewise env delete dev
codewise env delete dev --yes
```

| Flag | Description |
| --- | --- |
| `--yes` | Skip the confirmation prompt |
