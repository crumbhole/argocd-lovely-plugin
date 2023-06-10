This directory contains some useful argocd-lovely-plugin examples.

Additional examples also serve as CI/CD tests and can be found in the [test](../test) directory.

## Noteworthy Usage Examples

|Name|Description|
|--|--|
|[helm_kustomize](../test/helm_kustomize)|Deploy a Helm chart and then run Kustomize over the deployment to add additional values.|
|[helm_merge](../test/helm_merge)|Deploy a Helm chart and merge in a text file containing overwriting config.|
|[helm_plus_additions](../test/helm_plus_additions)|Deploy a Helm chart, a secret and a Kustomized configmap all under one Argo CD application.|
|[argocd-vault-replacer](../examples/argocd-vault-replacer)|Demonstrates how to use argocd-lovely-plugin with other Argo CD plugins, in this case argocd-vault-replacer to inject secrets from Hashicorp Vault.|
|[kustomize_patch](../test/kustomize_patch)|Demonstrates defining a Kustomize patch outside of a kustomization.yaml manifest.|
|[preprocessor](../test/preprocessor)|Demonstrates modifying a helm chart prior to it being downloaded and rendered.|
|[applicationsets](../examples/applicationsets)|Demonstrates performing a Kustomize patch on an applicationset.|

## Installation Examples

|Name|Description|
|--|--|
|[Installing argocd-lovely-plugin with argocd-vault-replacer into Argo CD as a sidecar using Kustomize](../examples/installation/argocd-sidecar)|An example of how to install and configure Argo CD and argocd-lovely-plugin, including argocd-vault-replacer plugin using Kustomize|
|[Installing argocd-lovely-plugin and the official argocd-vault-plugin into Argo CD as a sidecar using Kustomize](../examples/installation/argocd-with-vault-plugin-sidecar)|An example of how to install and configure Argo CD and argo-cd-lovely-plugin, including the official argocd-vault-plugin using Kustomize|
|[Installing argocd-lovely-plugin into Argo CD using Kustomize](../examples/installation/legacy-argocd)|An example of how to install and configure Argo CD and argo-cd-lovely-plugin using Kustomize (deprecated)|
|[Installing argocd-lovely-plugin and argocd-vault-replacer into Argo CD using Kustomize](../examples/installation/legacy-argocd-with-argocd-vault-replacer)|An example of how to install and configure Argo CD, argo-cd-lovely-plugin and argocd-vault-replacer using Kustomize (deprecated)|
