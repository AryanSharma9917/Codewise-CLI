apiVersion: argoproj.io/v1alpha1
kind: Application
metadata:
  name: my-app
spec:
  destination:
    namespace: default
    server: https://kubernetes.default.svc
  project: default
  source:
    repoURL: https://github.com/your/repo
    targetRevision: HEAD
    path: .
  syncPolicy:
    automated:
      prune: true
      selfHeal: true
