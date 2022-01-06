On occasion, you may find that a Helm chart you are consuming lacks the necessary values for your needs. The argocd-lovely-plugin allows you to add additional resources to your helm chart configuration.

This example:
- Installs the helm chart defined in `Chart.yaml`
- Merges configuration into the resulting install by adding the Environment Variable defined in `env.txt`

Alternatively, if this is too basic for your needs, you may wish to Kustomize the Helm chart. An example of this can be found at [helm_kustomize](../helm_kustomize).