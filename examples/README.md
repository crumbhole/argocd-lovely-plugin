This directory contains some useful argocd-lovely-plugin examples.

Additional examples also serve as CI/CD tests and can be found in the `test` directory.

## Noteworthy Usage Examples

|Name|Description|
|--|--|
|[helm_kustomize](https://github.com/crumbhole/argocd-lovely-plugin/tree/main/test/kustomize)|Deploy a Helm chart and then run Kustomize over the deployment to add additional values.|
|[helm_merge](https://github.com/crumbhole/argocd-lovely-plugin/tree/main/test/helm_merge)|Deploy a Helm chart and merge in a text file containing overwriting config.|
|[helm_plus_additions](https://github.com/crumbhole/argocd-lovely-plugin/tree/main/test/helm_plus_additions)|Deploy a Helm chart, a secret and a Kustomized configmap all under one Argo CD application.|
|[argocd-vault-replacer](https://github.com/crumbhole/argocd-lovely-plugin/tree/main/examples/argocd-vault-replacer)|Demonstrates how to use argocd-lovely-plugin with other Argo CD plugins, in this case argocd-vault-replacer to inject secrets from Hashicorp Vault.|
|[kustomize_patch](https://github.com/crumbhole/argocd-lovely-plugin/tree/main/examples/kustomize_patch)|Demonstrates defining a Kustomize patch outside of a kustomization.yaml manifest.|

## Installation Examples

|Name|Description|
|--|--|
|[Installing argocd-lovely-plugin into Argo CD using Kustomize](https://github.com/crumbhole/argocd-lovely-plugin/tree/main/examples/installation/argocd)|An example of how to install and configure Argo CD and argo-cd-lovely-plugin using Kustomize|
|[Installing argocd-lovely-plugin and argocd-vault-replacer into Argo CD using Kustomize](https://github.com/crumbhole/argocd-lovely-plugin/tree/main/examples/installation/argocd-with-argocd-vault-replacer)|An example of how to install and configure Argo CD, argo-cd-lovely-plugin and argocd-vault-replacer using Kustomize|