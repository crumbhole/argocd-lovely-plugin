# POC for argocd sidecar plugins

## Prereqs
- k3d
- kubectl
- kubens
- docker

## Do the thing
1. Add this to your hosts file:
```
127.0.0.1 k3d-registry.localhost
```

2. run `k3d cluster create --config k3d.conf`. Wait for the cluster to create.

3. `docker build -t k3d-registry.localhost:5000/lovely . && docker push k3d-registry.localhost:5000/lovely`

4. `kubectl create ns argocd && kubens argocd`

5. `kubectl apply -k argocd/`


Argocd should deploy, and then the 3 test apps should also successfully deploy.