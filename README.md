# argocd-lovely-plugin
An [ArgoCD](https://argoproj.github.io/argo-cd/) plugin that behaves in a way I wish ArgoCD behaved. This is only aimed at using ArgoCD for gitops - I do not use the UI for creating or modifying applications.

## Headline features
- Composite multiple things together to form a single app from multiple directories. For example - two or more helm charts together as a single app. Or a helm chart with a bit of plain yaml (a secret) to supplement it.
- Trivially allows helm + kustomize to work together, just works as you'd hope. Put a helm
- When used with [application sets](https://argocd-applicationset.readthedocs.io/en/stable/) you can apply kustomization and modify helm's values.yaml per application to apply minor differences to your applications trivially.
- Chain several plugins together. Lovely acts as a master plugin runner (acting as the only plugin to ArgoCD), and then runs other ArgoCD compatbile plugins in a chain. This acts a bit like a unix pipe, so you can helm | kustomize | argocd-vault-replacer.

## Why?
- Allows for better GitOps with one argo application per real application.
- Keep complex applications structured with subdirectories
- DRY (Don't repeat yourself) more
  - Allows ArgoCD to kustomize per application.
  - Combines particularly well with [application sets](https://argocd-applicationset.readthedocs.io/en/stable/) to allow broadly similar things to be partially modified by the application.

## Supports
- Helm
- Kustomize
- Plain YAML

Lovely-plugin does not support jsonnet as I do not use jsonnet, and haven't seen the need to use it. I don't know how it would best fit into the structure.

# Installing as an ArgoCD Plugin
You can use [our Kustomization example](https://github.com/crumbhole/argocd-lovely-plugin/tree/main/examples/kustomize/argocd) to install ArgoCD and to bootstrap the installation of the plugin at the same time. However the steps below will detail what is required should you wish to do things more manually.

## General configuration
Lovely is designed for minimal configuration and to do the right thing. The following environment variables can be used to change some behaviour:
- LOVELY_PLUGINS: Set to a comma separated list of binaries to run, in the same way as argocd expects plugins. stdin->plugin->stdout processing yaml. Each plugin is executed with bash -c <plugin and parameters>, so you can pass parameters as hoped for.
- LOVELY_KUSTOMIZE_PATH: Set to a path or binary name to use for kustomize.
- LOVELY_HELM_PATH: Set to a path or binary name to use for helm.

## Helm variation
- LOVELY_HELM_MERGE: to some yaml you'd like [strategic merged](https://kubectl.docs.kubernetes.io/references/kustomize/kustomization/patchesstrategicmerge/) merged into any values.yaml used by helm.
- LOVELY_HELM_PATCH: to some yaml or json you'd like [json6902](https://kubectl.docs.kubernetes.io/references/kustomize/kustomization/patchesjson6902/) patched into any values.yaml used by helm.
- ARGOCD_APP_NAME: This can be used to set the helm 'name' in the same way as releaseName works in ArgoCD's standard helm processing

There is no way to modify any other helm files at this time.

## Kustomize
- LOVELY_KUSTOMIZE_MERGE: to some yaml you'd like [strategic merged](https://kubectl.docs.kubernetes.io/references/kustomize/kustomization/patchesstrategicmerge/) merged into any kustomization.yaml found.
- LOVELY_KUSTOMIZE_PATCH: to some yaml or json you'd like [json6902](https://kubectl.docs.kubernetes.io/references/kustomize/kustomization/patchesjson6902/) patched into any kustomization.yaml found.

There is no way to modify any other files, that's what kustomize itself is for.

## What can I do with it?
Have a look at the [examples directory](examples/README.md) for a list of examples of how you can use this to make nice git repos for your applications. This also refers to the [test directory](test/README.md), which contains a number of examples that also serve as CI/CD tests for this plugin.

## What doesn't it do?
This is not a templating tool, there are plenty of choices out there to that stuff. It just brings together external tools.

# How does it work?

Firstly we scan the working application directory for yaml files. At each shallowest point where one is found then it is marked as a separate sub-application for working on.

For each sub-application
-  If there is a Chart.yaml then helm will run the templating engine.
- If there is a kustomization.yaml then kustomize will be run. If helm was previously run the output of helm will be added to the kustomization.yaml as a resource automatically.
- Only if neither a Chart.yaml nor a kustomization.yaml were found then all the yaml in the directory tree will be concatentated together.
- All plugins will be run, in order given to process the data.

All the sub-application yamls will be concatenated and the result will be fed to ArgoCD (printed to stdout).
