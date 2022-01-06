On occasion, you may find that a Helm chart you are consuming lacks the necessary values for your needs. The argocd-lovely-plugin allows you to add additional resources to your helm chart configuration.

This example:
- Installs the helm chart defined in `chart`
- Deploys a secret defined in `secret`
- Deploys a Kustomized configmap defined in `configmap`

This example demonstrates that it is possible to combine a number of different deployment approaches into one argocd application using the argocd-lovely-plugin. You could even take this further and make use of an argocd-lovely-plugin plugin to further manipulate the values. For example, setting [argocd-vault-replacer](https://github.com/crumbhole/argocd-vault-replacer) as an argocd-lovely-plugin plugin will allow you to pull secrets from Hashicorp Vault.

Separating the different deployment types out into directories is not necessary, but is purely there for human readability. Refer to [helm_kustomize](../helm_kustomize) and [helm_merge](../helm_merge) to see examples where a human-friendly directory structure is not used. Ultimately, the end result is the same.