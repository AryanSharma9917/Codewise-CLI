---
layout: home

hero:
  name: Codewise CLI
  text: One command surface for DevOps
  tagline: Scaffold, package, configure, and deploy containerized applications without memorizing disconnected toolchains.
  image:
    src: /logo.png
    alt: Codewise CLI
  actions:
    - theme: brand
      text: Get started
      link: /guide/quick-start
    - theme: alt
      text: Command reference
      link: /commands/
    - theme: alt
      text: View on GitHub
      link: https://github.com/aryansharma9917/codewise-cli

features:
  - icon: 🐳
    title: Container workflow
    details: Generate and validate Dockerfiles, then build consistently tagged images.
  - icon: ☸️
    title: Kubernetes and Helm
    details: Scaffold manifests and charts, preview changes safely, and operate deployments.
  - icon: 🔁
    title: GitOps ready
    details: Describe environments in files and deploy through generated Argo CD Applications.
  - icon: 🧰
    title: Everyday utilities
    details: Convert YAML, JSON, TOML, XML, ENV, and Base64 without reaching for separate tools.
---

## Start in under a minute

```bash
git clone https://github.com/aryansharma9917/codewise-cli.git
cd codewise-cli
make build
./codewise-cli --help
```

Then scaffold the deployment assets for your application:

```bash
./codewise-cli docker init
./codewise-cli k8s init --app myapp --image myapp:latest
./codewise-cli helm init --app myapp --image myapp:latest
```

::: tip Safe exploration
Use `k8s apply --dry-run`, `k8s delete --dry-run`, and `deploy run --dry-run` before connecting Codewise to a cluster.
:::
