# Configuring argocd-lovely-plugin

There are a number of configuration parameters for lovely. Each of them can be configured via various means.

This is a list of those means, in priority order. If a value is found in one of these locations it will be used and further options for finding it will be ignored.

1. As a plugin parameter, when used as a sidecar in ArgoCD. These are in lower case, instead of upper case.
2. As an environment variable prefixed with `ARGOCD_ENV_`. This is what will happen to `env:` variables entered into the application.
3. As an environment variable not prefixed with anything. If you are configuring using a kubernetes manifest, this is probably what you'd expect.

## Plugin parameters

Environment variables, are conventionally ALL_CAPS, and for our purposes
