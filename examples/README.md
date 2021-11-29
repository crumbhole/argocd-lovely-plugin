This directory contains some useful argocd-lovely-plugin examples.

Additional examples also serve as CI/CD tests and can be found in the `test` directory.

|Name|Description|
|--|--|
|[helm_kustomize](https://github.com/crumbhole/argocd-lovely-plugin/tree/main/test/kustomize)|Deploy a Helm chart and then run Kustomize over the deployment to add additional values.|
|[helm_merge](https://github.com/crumbhole/argocd-lovely-plugin/tree/main/test/helm_merge)|Deploy a Helm chart and merge in a text file containing overwriting config.|
|[helm_plus_additions](https://github.com/crumbhole/argocd-lovely-plugin/tree/main/test/helm_plus_additions)|Deploy a Helm chart, a secret and a Kustomized configmap all under one Argo CD application.|