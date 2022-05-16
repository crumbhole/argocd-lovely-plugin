This will deploy two Helm charts as one application. An nginx helm chart and a Hello World chart.

1. Install argocd with the argoce-lovely-plugin
```
kubectl apply -k examples/installation/argocd
```

2. Apply the mutliple helm charts application:
```
kubectl apply -n argocd -f examples/multiple-helm-charts/argocd-application/hello-nginx-world.yml
```