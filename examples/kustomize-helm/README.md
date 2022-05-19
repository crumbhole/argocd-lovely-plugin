On occasion, you may find that a Helm chart you are consuming lacks the necessary values for your needs. The argocd-lovely-plugin allows you to add additional resources to your Helm chart configuration and also allows you to manipluate the output of Helm Template to add or remove additional changes.

This example:
- Installs the helm chart defined in `chart`
- Deploys a separate secret defined in `secret`
- Kustomizes the output of the chart to tweak the liveness and readiness probe config (something we aren’t able to do via this particular `values.yaml`). The patch is in `liveness-patch.yml` in `chart`.

1. Install argoCD with the argocd-lovely-plugin. Optional: If you want to use the ArgoCD UI, port-forward the server pod and grab the admin password:
```
cd examples/installation/argocd
kubectl apply -k .
kubectl port-forward svc/argocd-server -n argocd 8080:443
kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d; echo
```

2. Apply the kustomize-helm application:
```
kubectl apply -n argocd -f https://raw.githubusercontent.com/crumbhole/argocd-lovely-plugin/main/examples/kustomize-helm/application/kustomize-helm.yml
kubectl -n helm-plus-additions wait deployment helm-plus-additions-hello-world --for condition=Available=True
```

In your cluster, you should see a namespace called 'helm-plus-additions' containing your pod. If you look at the manifest, you should see the kustomized liveness and readiness probes.
In addition, you should find your secret (`mysecret`) in this namespace:
```
# Get the secret
kubectl -n helm-plus-additions get secret mysecret

# output the yaml for the deployment and look for failurethreshold to be '3', not '5'
kubectl -n helm-plus-additions get deployment helm-plus-additions-hello-world -o yaml | grep "failureThreshold:"
```


When finished, you can delete the argoCD Application and the multiple-helm namespace:
```
kubectl delete -n argocd -f https://raw.githubusercontent.com/crumbhole/argocd-lovely-plugin/main/examples/kustomize-helm/application/kustomize-helm.yml && kubectl delete namespace helm-plus-additions
```


This is a somewhat contrived example as it would make logical sense in the real world to include the secret.yml contents within the kustomization.
However, the aim of this example is to demonstrate that if you *just* needed to add some manifests to the output of a Helm chart, you can simply place the manifest alongside and the argocd-lovely-plugin will do the rest. For more complicated things, you will probably want to use Kustomize with the argocd-lovely-plugin.
