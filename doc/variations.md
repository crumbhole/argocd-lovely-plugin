# Lovely variations

All lovely plugin containers have helm, kustomize, helmfile, bash, git, and yq built in to them.

| Container name | Versioned | Contains |
|----------------|-----------|----------|
| ghcr.io/crumbhole/lovely | No | Plain lovely, but unversioned, `called argocd-lovely-plugin` |
| ghcr.io/crumbhole/lovely-vault-ver | Yes | [argocd-vault-replacer](https://github.com/crumbhole/argocd-vault-replacer) |
| ghcr.io/crumbhole/lovely-vault | No | [argocd-vault-replacer](https://github.com/crumbhole/argocd-vault-replacer) |
| ghcr.io/crumbhole/lovely-hera-ver | Yes | [hera](hera.md) |
| ghcr.io/crumbhole/lovely-hera | No | [hera](hera.md) |
| ghcr.io/crumbhole/lovely-hera-vault-ver | Yes | [argocd-vault-replacer](https://github.com/crumbhole/argocd-vault-replacer) + [hera](hera.md) |
| ghcr.io/crumbhole/lovely-hera-vault | No | [argocd-vault-replacer](https://github.com/crumbhole/argocd-vault-replacer) + [hera](hera.md) |

Each of these comes ready setup to perform the processing using these plugins.

## Versioned vs unversioned

The versioned containers require you to specify the version of the plugin as in `pluginname-version`. When you upgrade the plugin you will need to update all your applications, but this approach allows you to run multiple versions of the same plugin. The version will match the container tag, as in `plugin-1.2` will be usableplugin

The unversioned do not have the version, so all versions of the unversioned plugin have the same identifier. This means you cannot run two versions of the same plugin, but you also do not have to do any work when upgrading the plugin, your applications will automatically use the new plugin.

## Lovely does not do [discovery](https://argo-cd.readthedocs.io/en/stable/operator-manual/config-management-plugins/#write-discovery-rules-for-your-plugin)

You must specify that you'd like your application processed with lovely explitly by putting the name of the plugin in the application.
