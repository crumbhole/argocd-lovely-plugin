On occasion, you may find that a Helm chart you are consuming lacks the necessary values for your needs. The argocd-lovely-plugin allows you to add additional resources to your Helm chart configuration.

This example:
- Installs the helm chart defined in `chart`
- Deploys a separate secret defined in `secret`
- Kustomizes the output of the chart to tweak the liveness and readiness probe config (something we aren’t able to do via this particular `values.yaml`). The patch is in `liveness-patch.yml` in `chart`.

1. Install argoCD with the argoce-lovely-plugin
```
kubectl apply -k examples/installation/argocd
```

2. Apply the kustomize-helm application:
```
kubectl apply -n argocd -f https://raw.githubusercontent.com/crumbhole/argocd-lovely-plugin/main/examples/kustomize-helm/application/kustomize-helm.yml
```

In your cluster, you should see a namespace called 'helm-plus-additions' containing your pod. If you look at the manifest, you should see the kustomized liveness and readiness probes.
In addition, you should find your secret (`mysecret`) in this namespace.


When finished, you can delete the argoCD Application and the multiple-helm namespace:
```
kubectl delete -n argocd -f https://raw.githubusercontent.com/crumbhole/argocd-lovely-plugin/main/examples/kustomize-helm/application/kustomize-helm.yml && kubectl delete namespace helm-plus-additions
```