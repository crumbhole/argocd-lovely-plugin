This will deploy two Helm charts as one application. An nginx helm chart and a Hello World chart.
Both charts have some arbitrary changes made to their Values.yaml just to prove that that's a thing we can do.

1. Install argoCD with the argocd-lovely-plugin. Optional: If you want to use the ArgoCD UI, port-forward the server pod and grab the admin password:
```
cd examples/installation/argocd
kubectl apply -k .
kubectl port-forward svc/argocd-server -n argocd 8080:443
kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d; echo
```

2. Apply the mutliple helm charts application:
```
kubectl apply -n argocd -f https://raw.githubusercontent.com/crumbhole/argocd-lovely-plugin/main/examples/multiple-helm-charts/argocd-application/hello-nginx-world.yml
```

In your cluster, you should see a namespace called 'multiple-helm' containing two pods. One using the nginx image (chart 1) and one using the docker/whalesay image (from Chart 2).
In this example, chart 2's deployment ends up in a `CrashLoopBackOff` because it hasn't been fully configured.


When finished, you can delete the argoCD Application and the multiple-helm namespace:
```
kubectl delete -n argocd -f https://raw.githubusercontent.com/crumbhole/argocd-lovely-plugin/main/examples/multiple-helm-charts/argocd-application/hello-nginx-world.yml && kubectl delete namespace multiple-helm
```