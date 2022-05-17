You can use argocd-lovely-plugin to patch Argo CD applicationsets.

Consider a situation where you have a Helm chart deployed as an applicationset, and you need to modify a value in the values.yaml depending on the cluster name.

In this example, we will be deploying the Helm chart found in [helm_only](../applicationsets/helm_only). You should note that the values.yaml defines the serviceAccount name as "foo". We want this to be the name of our cluster instead.

1. Install argoCD with the argocd-lovely-plugin. Optional: If you want to use the ArgoCD UI, port-forward the server pod and grab the admin password:
```
cd examples/installation/argocd
kubectl apply -k .
kubectl port-forward svc/argocd-server -n argocd 8080:443
kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d; echo
```

2. Our [applicationset](../applicationsets/applicationset.yaml#L20-L25) contains a `LOVELY_HELM_PATCH` environment variable. This is a JSON string that contains a patch to be applied to the values.yaml.

Apply the applicationset to your cluster:
```
kubectl apply -n argocd -f https://raw.githubusercontent.com/crumbhole/argocd-lovely-plugin/main/examples/applicationsets/applicationset.yaml
```

3. We only have one cluster as part of this demonstration, but you can see that the name of the ServiceAccount for the deployment is called `in-cluster`, which is the what Argo CD calls the default cluster.
```
kubectl -n example get serviceAccounts

NAME         SECRETS   AGE
default      1         1m
in-cluster   1         1m
```

When finished, you can delete the argoCD Applicationset and the example namespace:
```
kubectl delete -n argocd -f https://raw.githubusercontent.com/crumbhole/argocd-lovely-plugin/main/examples/applicationsets/applicationset.yaml && kubectl delete namespace example
```