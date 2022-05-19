# argocd-lovely-plugin
An [Argo CD](https://argoproj.github.io/argo-cd/) plugin that behaves in a way we wish Argo CD behaved. This is only aimed at using Argo CD for GitOps - we do not use the UI for creating or modifying applications.

## Headline features
- Composite multiple things together to form a single app from multiple directories. For example - two or more Helm charts together as a single app. Or a Helm chart with a bit of plain yaml (a secret) to supplement it.
- Trivially allows Helm + Kustomize to work together, just works as you'd hope. Put a helm Chart.yaml+values.yaml in a folder, alongside a kustomization.yaml and you can kustomize your helm output or add more objects with kustomize
- When used with [application sets](https://argocd-applicationset.readthedocs.io/en/stable/) you can apply Kustomization and modify Helm's values.yaml per application to apply minor differences to your applications trivially.
- Chain several plugins together. argocd-lovely-plugin acts as a master plugin runner (acting as the only plugin to Argo CD), and then runs other Argo CD compatible plugins in a chain. This acts a bit like a unix pipe, so you can helm | kustomize | argocd-vault-replacer.

## Why?
- Allows for better GitOps with one argo application per real application.
- Keep complex applications structured with subdirectories
- DRY (Don't repeat yourself) more
  - Allows Argo CD to Kustomize per application.
  - Combines particularly well with [application sets](https://argocd-applicationset.readthedocs.io/en/stable/) to allow broadly similar things to be partially modified by the application.

## Supports
- Helm
- Kustomize
- Plain YAML

argocd-lovely-plugin does not support jsonnet as we do not use jsonnet, and haven't seen the need to use it. We don't know how it would best fit into the structure.

# Installing as an Argo CD Plugin
You can use [our Kustomization example](examples/installation/argocd) to install Argo CD and to bootstrap the installation of the plugin at the same time. However the steps below will detail what is required should you wish to do things more manually.

## General configuration
argocd-lovely-plugin is designed for minimal configuration and to do the right thing. The following environment variables can be used to change some behaviour:
- LOVELY_PREPROCESSORS and LOVELY_PLUGINS: Set to a comma separated list of binaries to run during preprocessing and as plugins. Read [this](doc/plugins.md) for more on plugins.
- LOVELY_KUSTOMIZE_PATH: Set to a path or binary name to use for Kustomize.
- LOVELY_HELM_PATH: Set to a path or binary name to use for Helm.

## Helm variation
You can use these environment variables for modifying helm's behaviour, and the values.yaml file. More generic manipulation of any file is available through preprocessing.
- LOVELY_HELM_MERGE: to some yaml you'd like [strategic merged](https://kubectl.docs.kubernetes.io/references/kustomize/kustomization/patchesstrategicmerge/) merged into any values.yaml used by Helm.
- LOVELY_HELM_PATCH: to some yaml or json you'd like [json6902](https://kubectl.docs.kubernetes.io/references/kustomize/kustomization/patchesjson6902/) patched into any values.yaml used by Helm.
- ARGOCD_APP_NAME: This can be used to set the Helm 'name' in the same way as releaseName works in Argo CD's standard Helm processing

There is no way to modify any other Helm files at this time.

## Kustomize
You can use these environment variables for modifying kustomize's behaviour, and the kustomization.yaml file. More generic manipulation of any file is available through preprocessing.
- LOVELY_KUSTOMIZE_MERGE: to some yaml you'd like [strategic merged](https://kubectl.docs.kubernetes.io/references/kustomize/kustomization/patchesstrategicmerge/) merged into any kustomization.yaml found.
- LOVELY_KUSTOMIZE_PATCH: to some yaml or json you'd like [json6902](https://kubectl.docs.kubernetes.io/references/kustomize/kustomization/patchesjson6902/) patched into any kustomization.yaml found.

There is no way to modify any other files, that's what Kustomize itself is for.

## What can I do with it?
Have a look at the [examples directory](examples/README.md) for a list of examples of how you can use this to make nice git repos for your applications. This also refers to the [test directory](test/README.md), which contains a number of examples that also serve as CI/CD tests for this plugin.

## What doesn't it do?
This is not a templating tool, there are plenty of choices out there to that stuff. It just brings together external tools.

# How does it work?
Firstly we scan the working application directory for yaml files. At each shallowest point where one is found then it is marked as a separate sub-application for working on.

For each sub-application
- If there is a Chart.yaml then Helm will run the templating engine.
- If there is a kustomization.yaml then Kustomize will be run. If Helm was previously run the output of Helm will be added to the kustomization.yaml as a resource automatically.
- Only if neither a Chart.yaml nor a kustomization.yaml were found then all the yaml in the directory tree will be concatentated together.
- All plugins will be run, in order given to process the data.

All the sub-application yamls will be concatenated and the result will be fed to Argo CD (printed to stdout).

# Videos
If you prefer to watch videos of things rather than read words, we have compiled some demos.

- [Using Argo CD to Kustomize a Helm chart, and deploy additional manifests](https://youtu.be/OMae_8DHELA)
