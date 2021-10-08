# argocd-lovely-plugin
An [ArgoCD](https://argoproj.github.io/argo-cd/) plugin to perform various manipulations in a sensible order to ultimately output YAML for Argo CD to put into your cluster. If you use it as the reader in a unix pipe, it will instead read from stdin. It will still under some circumstances taint the current directory with files so expects to be run in a directory it can mess up.

## Why?
- Allows for better GitOps with one argo application per real application.
  - Process a helm chart through kustomize more easily.
  - Trivially add some extra resources to a helm chart by popping them in a directory alongside your main one.
  - Chain additional ArgoCD plugins. For example, you can use [argo-vault-replacer](https://github.com/crumbhole/argocd-vault-replacer) to pull secrets from [Hashicorp Vault](https://www.vaultproject.io/) as a plugin to this one to combine secrets and helm/kustomize more easily.
- DRY (Don't repeat yourself) more
  - Allows ArgoCD to kustomize per application.
  - Combines particularly well with [application sets](https://argocd-applicationset.readthedocs.io/en/stable/) to allow broadly similar things to be partially modified by the application.

## Downsides
- Helm is no longer special, and is just a templating tool. This is the case for any use of modified helm charts in ArgoCD. We cannot run hooks or anything any more. You don't get Helm specific support in ArgoCD.

# Installing as an ArgoCD Plugin
You can use [our Kustomization example](https://github.com/crumbhole/argocd-lovely-plugin/tree/main/examples/kustomize/argocd) to install ArgoCD and to bootstrap the installation of the plugin at the same time. However the steps below will detail what is required should you wish to do things more manually.

## General configuration
Lovely is designed for minimal configuration and to do the right thing. The following environment variables can be used to change some behaviour:
- LOVELY_PLUGINS: Set to a comma separated list of binaries to run, in the same way as argocd expects plugins. stdin->plugin->stdout processing yaml. Each plugin is executed with bash -c <plugin and parameters>, so you can pass parameters as hoped for.
- LOVELY_KUSTOMIZE_PATH: Set to a path or binary name to use for kustomize.
- LOVELY_HELM_PATH: Set to a path or binary name to use for helm.

## Helm variation
- LOVELY_HELM_VALUES: to some yaml you'd like overlayed on any values.yaml files used by helm. This will override any settings in values.yaml that already exist. You cannot use this to append to existing yaml lists.

## Kustomize
- LOVELY_KUSTOMIZE_EXTRAS: to some yaml you'd like overlayed on any kustomization.yaml files used by kustomize. This will override any settings in kustomization.yaml that already exist. You cannot use this to append to existing yaml lists.

## What can I do with it?
Have a look at the [test suite](test/README.md) for a lot of examples of how you can use this to make nice git repos for your applications.

## What doesn't it do?
This is not a templating tool, there are plenty of choices out there to that stuff. It just brings together external tools.
