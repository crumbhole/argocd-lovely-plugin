# Configuration options

## Configuring argocd-lovely-plugin

There are a number of configuration parameters for lovely. Each of them can be configured via various means.

This is a list of those means, in priority order. If a value is found in one of these locations it will be used and further options for finding it will be ignored.

1. As a plugin parameter, when used as a sidecar in ArgoCD. These are in lower case, instead of upper case.
2. As an environment variable prefixed with `ARGOCD_ENV_`. This is what will happen to `env:` variables entered into the application.
3. As an environment variable not prefixed with anything. If you are configuring using a kubernetes manifest, this is probably what you'd expect.

## Plugin parameters

Environment variables, are conventionally ALL_CAPS, and for our purposes

## Available parameters
|Name | Environment variable | Description | Default |
| ---- | -------------------- | ----------- | ------- |
| Plugins | LOVELY_PLUGINS | Set to a comma separated list of binaries to run as plugins. Read [the documentation](plugins.md) for more on plugins. Will not be used if `Plugins YAML` is set. |  |
| Plugins YAML | LOVELY_PLUGINS_YAML | Set to some YAML describing the binaries to run as plugins. Read [the documentation](plugins.md) for more on plugins. Will override `Plugins` if set. |  |
| PreProcessors | LOVELY_PREPROCESSORS | Set to a comma separated list of binaries to run as preprocessors. Read [the documentation](plugins.md) for more on plugins. Will not be used if `Preproecessors YAML` is set. |  |
| PreProcessors YAML | LOVELY_PREPROCESSORS_YAML | Set to some YAML describing the binaries to run as preprocessors. Read [the documentation](plugins.md) for more on plugins. Will override `PreProcessors` if set. |  |
| Detection Regular Expression | LOVELY_DETECTION_REGEX | Allow applications to be detected using a different regex so that a PREPROCESSOR that works on non-yaml files can run on this application. The default is `\.ya?ml$`. (Note: currently `helmfile.d/` will always trigger an application being detected, raise an issue if this needs configuring too). This is pointless to change unless you have a PREPROCESSOR defined. | \.ya?ml(\.gotmpl)?$ |
| Kustomize Path | LOVELY_KUSTOMIZE_PATH | Path to the kustomize binary used for this application | kustomize |
| Kustomize parameters | LOVELY_KUSTOMIZE_PARAMS | Space separated extra parameters to `kustomize build` as you might use on the command line. `--enable-helm` is already passed always. You're on your own here if you pass rubbish parameters. |  |
| Helm Path | LOVELY_HELM_PATH | Path to the helm binary used for this application | helm |
| Helm Merge | LOVELY_HELM_MERGE | Set to some yaml you'd like [strategic merged](https://kubectl.docs.kubernetes.io/references/kustomize/kustomization/patchesstrategicmerge/) into the values.yaml (or first `Helm Values` specified file) used by Helm. |  |
| Helm Patch | LOVELY_HELM_PATCH | Set to some yaml you'd like [json6902](https://kubectl.docs.kubernetes.io/references/kustomize/kustomization/patchesjson6902/) patched into the values.yaml (or first `Helm Values` specified file) used by Helm. |  |
| Helm Template Parameters | LOVELY_HELM_TEMPLATE_PARAMS | Space separated extra parameters to `Helm template` as you might use on the command line. You're on your own here if you pass rubbish parameters. |  |
| Helm Repo Add Parameters | LOVELY_HELM_REPO_ADD_PARAMS | Space separated extra parameters to `Helm repo add` as you might use on the command line. You're on your own here if you pass rubbish parameters. `--insecure-skip-tls-verify` if your helm chart is on an insecure HTTPS server. |  |
| Helm CRDs | LOVELY_HELM_CRDS | Whether to include CRDs from a helm chart or skip them | true |
| Kustomize Merge | LOVELY_KUSTOMIZE_MERGE | Set to some yaml you'd like [strategic merged](https://kubectl.docs.kubernetes.io/references/kustomize/kustomization/patchesstrategicmerge/) into any kustomization.yaml found. |  |
| Kustomize Patch | LOVELY_KUSTOMIZE_PATCH | Set to some yaml or json you'd like [json6902](https://kubectl.docs.kubernetes.io/references/kustomize/kustomization/patchesjson6902/) patched into any kustomization.yaml found. |  |
|  |  |  |  |
| Helm Name | LOVELY_HELM_NAME | This can be used to set the Helm 'name' in the same way as releaseName works in Argo CD's standard Helm processing. (`ARGOCD_APP_NAME` used to be overridable in old versions of ArgoCD, but is no longer). Will default to ARGOCD_APP_NAME from the application. |  |
| Helm Namespace | LOVELY_HELM_NAMESPACE |  This can be used to set the Helm 'namespace' it will apply. Will default to ARGOCD_APP_NAMESPACE from the application. |  |
| Helm Values | LOVELY_HELM_VALUES | This is a space separated list values files you'd like to use when rendering the helm chart. Defaults to `values.yaml` if that exists, but its fine if it doesn't. If you override this the file *must* exist. MERGE and PATCH will be applied to the first file in this list. |  |
| Helmfile Path | LOVELY_HELMFILE_PATH | Path to the helmfile binary used for this application | helmfile |
| Helmfile CRDs | LOVELY_HELMFILE_CRDS | Whether to include CRDs from helmfile or skip them | true |
| Helmfile Merge | LOVELY_HELMFILE_MERGE | Set to some yaml you'd like [strategic merged](https://kubectl.docs.kubernetes.io/references/kustomize/kustomization/patchesstrategicmerge/) into any helmfile.yaml used by helmfile. |  |
| Helmfile Patch | LOVELY_HELMFILE_PATCH | to some yaml or json you'd like [json6902](https://kubectl.docs.kubernetes.io/references/kustomize/kustomization/patchesjson6902/) patched into any helmfile.yaml used by Helmfile. |  |
| Helmfile Template Parameters | LOVELY_HELMFILE_TEMPLATE_PARAMS | Space separated extra parameters to `Helmfile template` as you might use on the command line. You're on your own here if you pass rubbish parameters. |  |
| Environment variables propagation | LOVELY_ENV_PROPAGATION | Whether to propagate and map ARGOCD_ENV_{VARIABLE} environment variables as {VARIABLE} to the downstream processor (Helm, Kustomize or Helmfile). | false |
## Plugin Name

You can set `PLUGIN_NAME` in the environment of the sidecar to override the default name of the plugin. This allows you to supply multiple pre-configured plugins (with different environment, but the same variation).
