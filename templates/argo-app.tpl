apiVersion: argoproj.io/v1alpha1
kind: Application
metadata:
  name: {{.AppName}}
spec:
  source:
    repoURL: {{.Repo}}
    path: .
    targetRevision: HEAD
  destination:
    server: https://kubernetes.default.svc
    namespace: default
  project: default
