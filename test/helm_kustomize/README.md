On occasion, you may find that a Helm chart you are consuming lacks the necessary values for your needs. The argocd-lovely-plugin allows you to add additional resources to your helm chart configuration.

This example:
- Installs the helm chart defined in `Chart.yaml`
- Deploys a Kustomized configmap defined in `kustomization.yaml`, derived from `index.html`

It is possible to create separate subdirectories for the Helm Chart and the Kustomization manifests. An example of this can be found in [helm_plus_additions](../helm_plus_additions).

Alternatively, if Kustomization is too much for your needs, it is possible to perform some basic merges using flat .yaml files. An example of this can be found at [helm_merge](../helm_merge).