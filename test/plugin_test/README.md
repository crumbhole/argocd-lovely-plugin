This example:
- Performs an installation of all the manifests defined in the directory.
- Runs the `ARGOCD_ENV_LOVELY_PLUGIN` as defined in `env.txt`. In this example, the plugin trims some of the resulting manifests from the first step.

This is example is somewhat contrived, but demonstrates argocd-lovely-plugin's ability to make use of plugins to perform additional manipulations on manifests. A more real-world example can be found in [examples/argocd-vault-replacer](../../examples/argocd-vault-replacer) where argocd-vault-replacer is used to inject secrets from Hashicorp Vault into manifests that argocd-lovely-plugin has generated.
