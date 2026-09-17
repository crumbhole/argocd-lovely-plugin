# Helm versions

The lovely plugin containers ship two versions of Helm:

- Helm v4 at `/usr/local/bin/helm`. This is the system version, and is what lovely uses by default.
- Helm v3 at `/usr/local/bin/helm3`, for charts that do not yet work with Helm v4.

## Using Helm v3 for every application

Set `LOVELY_HELM_PATH` in the environment of the lovely sidecar container to make Helm v3 the default for all applications rendered by that plugin:

```yaml
containers:
  - name: lovely-plugin
    image: ghcr.io/crumbhole/argocd-lovely-plugin-cmp:<version>
    env:
      - name: LOVELY_HELM_PATH
        value: /usr/local/bin/helm3
```

## Using Helm v3 for a single application

Set `LOVELY_HELM_PATH` on the application itself, either as a plugin environment variable:

```yaml
apiVersion: argoproj.io/v1alpha1
kind: Application
spec:
  source:
    plugin:
      name: argocd-lovely-plugin-v1.3
      env:
        - name: LOVELY_HELM_PATH
          value: /usr/local/bin/helm3
```

or as a plugin parameter:

```yaml
apiVersion: argoproj.io/v1alpha1
kind: Application
spec:
  source:
    plugin:
      name: argocd-lovely-plugin-v1.3
      parameters:
        - name: lovely_helm_path
          string: /usr/local/bin/helm3
```

You don't need to prefix the environment variable with `ARGOCD_ENV_` in the application; Argo CD does that for you. If you are running lovely outside these containers, `LOVELY_HELM_PATH` can point at any Helm binary you have installed.
