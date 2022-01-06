# Using argocd-vault-replacer as an argocd-lovely-plugin plugin

This allows you to combine the features of argocd-lovely-plugin with other Argo CD plugins. In this example, we:

1. Deploy a helm chart in accordance with its values. One of the values will be taken from Hashicorp Vault.
2. Deploy a separate kubernetes secret, with the value being taken from Hashicorp Vault.
3. Deploy a kubernetes configmap using Kustomize. Some of the data within the configmap will be taken from Hashicorp Vault.

This is an obviously contrived example, but demonstrates how simple it is to combine argocd-lovely-plugin with other Argo CD plugins.

## Prerequisites:
1. You will need a running instance of Vault, with a KV secret at the path `example/foo/bar`, another at `secondexample/foo/bar` and a third at `thirdexample/foo/bar`.
2. You will need to install Argo CD, argocd-vault-replacer and argocd-lovely-plugin. An example of how to do this can be found at [examples/installation/argocd-with-argocd-vault-replacer).](../installation/argocd-with-argocd-vault-replacer). Ensure you modify the `VAULT_ADDR` environment variable to match your Vault installation.
3. You will need to set up the [Vault Kubernetes Authentication](https://github.com/crumbhole/argocd-vault-replacer/blob/main/README.md#vault-kubernetes-authentication).

## Argo CD Application Manifest
An example Argo CD application manifest could look like this:

```yaml
apiVersion: argoproj.io/v1alpha1
kind: Application
metadata:
  name: example
  namespace: argocd
  finalizers:
    - resources-finalizer.argocd.argoproj.io
spec:
  destination:
    server: 'https://kubernetes.default.svc'
    namespace: example
  project: example-project
  syncPolicy:
    automated:
      prune: true
      selfHeal: true
    syncOptions:
      - PrunePropagationPolicy=background
      - PruneLast=true
      - CreateNamespace=true
  source:
    repoURL: 'https://github.com/crumbhole/argocd-lovely-plugin.git'
    path: examples/argocd-vault-replacer
    targetRevision: HEAD
    plugin:
      name: argocd-lovely-plugin
```