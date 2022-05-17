You can use argocd-lovely-plugin to patch Argo CD applicationsets.

```yaml
apiVersion: argoproj.io/v1alpha1
kind: ApplicationSet
metadata:
  name: example-set
spec:
  generators:
    - clusters:
        selector:
          matchLabels:
            example.biz/appset-common: "true"
  template:
    metadata:
      name: 'example-{{name}}'
    spec:
      destination:
        name: '{{name}}'
        namespace: example
      project: applicationsets
      source:
        path: test/helm_only
        repoURL: 'https://github.com/crumbhole/argocd-lovely-plugin.git'
        targetRevision: HEAD
        plugin:
          name: argocd-lovely-plugin
          env:
            - name: LOVELY_HELM_PATCH
              value: |
                [{ "op": "add", "path": "/spec/containers/0/env", "value": { "action": "add", "name": "cluster", "value": "{{name}}" } }]
      syncPolicy:
        automated:
          prune: true
          selfHeal: true
        syncOptions:
        - CreateNamespace=true
```

This ApplicationSet manifest will deploy our test chart from `test/helm_only` and then inject the cluster name into it as an environment variable.
