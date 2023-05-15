# argocd-lovely-plugin
An [Argo CD](https://argoproj.github.io/argo-cd/) plugin that behaves in a way we wish Argo CD behaved. This is only aimed at using Argo CD for GitOps - we do not use the UI for creating or modifying applications.

## Headline features
- Composite multiple things together to form a single app from multiple directories. For example - two or more Helm charts together as a single app. Or a Helm chart with a bit of plain yaml (a secret) to supplement it.
- Trivially allows Helm + Kustomize to work together, just works as you'd hope. Put a helm Chart.yaml+values.yaml in a folder, alongside a kustomization.yaml and you can kustomize your helm output or add more objects with kustomize
- When used with [application sets](https://argocd-applicationset.readthedocs.io/en/stable/) you can apply Kustomization and modify Helm's values.yaml per application to apply minor differences to your applications trivially.
- Chain several plugins together. argocd-lovely-plugin acts as a master plugin runner (acting as the only plugin to Argo CD), and then runs other Argo CD compatible plugins in a chain. This acts a bit like a unix pipe, so you can helm | kustomize | argocd-vault-replacer.
- Can also use [helmfiles](https://helmfile.readthedocs.io/en/latest/) and combine them with other things. These can either be a `helmfile.yaml` or some yaml in `helmfile.d/`

## Why?
- Allows for better GitOps with one argo application per real application.
- Keep complex applications structured with subdirectories
- DRY (Don't repeat yourself) more
  - Allows Argo CD to Kustomize per application.
  - Combines particularly well with [application sets](https://argocd-applicationset.readthedocs.io/en/stable/) to allow broadly similar things to be partially modified by the application.

## Supports
- Helm
- Helmfile
- Kustomize
- Plain YAML

argocd-lovely-plugin does not support jsonnet as we do not use jsonnet, and haven't seen the need to use it. We don't know how it would best fit into the structure.

## Plain yaml

All the yaml in the directory and all subdirectories will be used as part of the application only if it is not a kustomize or helm chart. `.hidden` yaml files are not included.

## Supported Argo CD Versions
We aim to match the [Argo CD supported versions](https://argo-cd.readthedocs.io/en/stable/operator-manual/installation/#supported-versions) by testing against the Argo CD N and N -1 versions of Argo CD. You can see the current versions of Argo CD that we test against by looking in the [CI bootstrap directory](.github/workflows/assets/bootstrap) in this repo.

# Installing
We offer three pre-built container options. We only support the use of these containers, the binaries provided are for convenience:

1. `argocd-lovely-plugin-cmp` to install as a CMP sidecar plugin.
2. `argocd-lovely-plugin-cmp-vault` to install as a CMP sidecar plugin with [argocd-vault-replacer](https://github.com/crumbhole/argocd-vault-replacer) already baked in. [The Dockerfile](Dockerfile.cmp-vault) also serves as a blueprint should you wish to package your own plugins to chain with Lovely.
3. The deprecated `argocd-lovely-plugin` if you wish to install as an older-style configMap plugin.

## Installing as an Argo CD Sidecar Plugin
We recommend you install as an Argo CD CMP Sidecar Plugin. [Argo CD's documentation](https://argo-cd.readthedocs.io/en/stable/operator-manual/config-management-plugins/#sidecar-plugin) has steps on how to achieve this, or you can see [our Kustomization example](examples/installation/argocd). You can also observe how we install Lovely for our CI tests in the [CI bootstrap directory](.github/workflows/assets/bootstrap) in this repo.

argocd-lovely-plugin has no discovery rules, so will not run by default. You must reference the plugin by name in your application spec. For example:

```yaml
apiVersion: argoproj.io/v1alpha1
kind: Application
...
spec:
  source:
    plugin:
      name: argocd-lovely-plugin-v1.0
...
```
For more information, please refer to the [Argo CD Documentation on discovery](https://argo-cd.readthedocs.io/en/stable/operator-manual/config-management-plugins/#write-discovery-rules-for-your-plugin).

## Installing as an Argo CD ConfigMap Plugin (Deprecated)
You can use [our Kustomization example](examples/installation/legacy-argocd) to install Argo CD and to bootstrap the installation of the plugin at the same time.

At the moment the helmfile binary is not installed for you if you are running as a configmap plugin, nor is that documented here. You must get the helmfile binary into your repo-server yourself.


## Environment variables

argocd-lovely-plugin is configured through environment variables. These can be set in both the argocd-repo-server and in the application itself.

If you are passing the configuration in as application environment variables in Argo CD 2.4 or higher you must not put the `ARGOCD_ENV_` prefix on them, as Argo CD does that for you.

Otherwise argocd-lovely-plugin will accept either form of all of the variables, with or without `ARGOCD_ENV_`, with the `ARGOCD_ENV_` version taking precedence if you set both of them.

See [this](doc/config.md) for more details on how configuration works.

## General configuration
argocd-lovely-plugin is designed for minimal configuration and to do the right thing. The following environment variables can be used to change some behaviour:
- `LOVELY_PREPROCESSORS` and `LOVELY_PLUGINS`: Set to a comma separated list of binaries to run during preprocessing and as plugins. Read [this](doc/plugins.md) for more on plugins.
- `LOVELY_PREPROCESSORS_YAML` and `LOVELY_PLUGINS_YAML`: Set to some yaml or json for a list of binaries to run during preprocessing and as plugins. Read [this](doc/plugins.md) for more on plugins.
- `LOVELY_KUSTOMIZE_PATH`: Set to a path or binary name to use for Kustomize.
- `LOVELY_HELM_PATH`: Set to a path or binary name to use for Helm.
- `LOVELY_HELMFILE_PATH`: Set to a path or binary name to use for helmfile.
- `LOVELY_ALLOW_GITCHECKOUT`: Allows kustomize base paths to work. Do **not** just set this without reading [this](doc/allow_git.md)
- `LOVELY_DETECTION_REGEX`: Allow applications to be detected using a different regex so that a PREPROCESSOR that works on non-yaml files can run on this application. The default is `\.ya?ml$`. (Note: currently `helmfile.d/` will always trigger an application being detected, raise an issue if this needs configuring too). This is pointless to change unless you have a PREPROCESSOR defined.

## Helm variation
You can use these environment variables for modifying helm's behaviour, and the values.yaml file. More generic manipulation of any file is available through preprocessing.
- `LOVELY_HELM_VALUES`: This is a space separated list values files you'd like to use when rendering the helm chart. Defaults to `values.yaml`. If you override this the file *must* exist. MERGE and PATCH will be applied to the first file in this list.
- `LOVELY_HELM_MERGE`: to some yaml you'd like [strategic merged](https://kubectl.docs.kubernetes.io/references/kustomize/kustomization/patchesstrategicmerge/) merged into the values.yaml (or `LOVELY_HELM_VALUES` specified file)used by Helm.
- `LOVELY_HELM_PATCH`: to some yaml or json you'd like [json6902](https://kubectl.docs.kubernetes.io/references/kustomize/kustomization/patchesjson6902/) patched into the values.yaml (or `LOVELY_HELM_VALUES` specified file) used by Helm.
- `LOVELY_HELM_TEMPLATE_PARAMS`: Space separated extra parameters to `Helm template` as you might use on the command line. You're on your own here if you pass rubbish parameters.
- `LOVELY_HELM_REPO_ADD_PARAMS`: Space separated extra parameters to `Helm repo add` as you might use on the command line. You're on your own here if you pass rubbish parameters. `--insecure-skip-tls-verify` if your helm chart is on an insecure HTTPS server.
- `LOVELY_HELM_NAME`: This can be used to set the Helm 'name' in the same way as releaseName works in Argo CD's standard Helm processing. (`ARGOCD_APP_NAME` used to be overridable in old versions of ArgoCD, but is no longer)

There is no way to modify any other Helm files at this time.

## Helmfile variation
You can use these environment variables for modifying helmfiles's behaviour. More generic manipulation of any file is available through preprocessing. This cannot be used with helmfile.d files.
- `LOVELY_HELMFILE_MERGE`: to some yaml you'd like [strategic merged](https://kubectl.docs.kubernetes.io/references/kustomize/kustomization/patchesstrategicmerge/) merged into any helmfile.yaml used by helmfile.
- `LOVELY_HELMFILE_PATCH`: to some yaml or json you'd like [json6902](https://kubectl.docs.kubernetes.io/references/kustomize/kustomization/patchesjson6902/) patched into any helmfile.yaml used by Helmfile.

## Kustomize
You can use these environment variables for modifying kustomize's behaviour, and the kustomization.yaml file. More generic manipulation of any file is available through preprocessing.
- `LOVELY_KUSTOMIZE_MERGE`: to some yaml you'd like [strategic merged](https://kubectl.docs.kubernetes.io/references/kustomize/kustomization/patchesstrategicmerge/) merged into any kustomization.yaml found.
- `LOVELY_KUSTOMIZE_PATCH`: to some yaml or json you'd like [json6902](https://kubectl.docs.kubernetes.io/references/kustomize/kustomization/patchesjson6902/) patched into any kustomization.yaml found.
- `LOVELY_KUSTOMIZE_PARAMS`: Space separated extra parameters to `kustomize build` as you might use on the command line. `--enable-helm` is already passed always. You're on your own here if you pass rubbish parameters.

There is no way to modify any other files, that's what Kustomize itself is for.

You can use the [helm chart inflation generator](https://kubectl.docs.kubernetes.io/references/kustomize/builtins/#_helmchartinflationgenerator_) of kustomize this way. See [the test](test/helm_only_in_kustomize) for an example of this. If you do this none of the helm environment variables will have any effect as you can set those in your kustomization.yaml instead. There is no way to merge/patch your values.yaml with lovely only (you should run a preprocessor for that). Despite this, that is the recommended way to use helm and kustomize together. `LOVELY_HELM_NAME` will also have no effect here.

## ARGOCD_ENV_ support

All argocd-lovely-plugin environment variables may be prefixed with `ARGOCD_ENV_` for Argo CD 2.4 compatibility. If you are deranged and define both the `ARGOCD_ENV_` version will be used. When you put an environment variable into an application in 2.4 or later it will automatically get prefixed with `ARGOCD_ENV_` so you must use the non prefixed variable name there.

## What can I do with it?
Have a look at the [examples directory](examples/README.md) for a list of examples of how you can use this to make nice git repos for your applications. This also refers to the [test directory](test/README.md), which contains a number of examples that also serve as CI/CD tests for this plugin.

## What doesn't it do?
This is not a templating tool, there are plenty of choices out there to that stuff. It just brings together external tools.

# How does it work?
Firstly we scan the working application directory for yaml files. At each shallowest point where one is found then it is marked as a separate sub-application for working on.

For each sub-application
- Pre-processors will be run if defined
- If there is a `helmfile.yaml` or `helmfile.d` directory then helmfile will run the templating engine.
- If there is a `Chart.yaml` then Helm will run the templating engine.
- If there is a `kustomization.yaml` then Kustomize will be run. If Helm or helmfile was previously run the output of that will be added to the kustomization.yaml as a resource automatically.
- Only if neither a `Chart.yaml` nor a `kustomization.yaml` were found then all the yaml in the directory tree will be concatentated together.
- All plugins will be run, in order given to process the data.

All the sub-application yamls will be concatenated and the result will be fed to Argo CD (printed to stdout).

Helm and helmfile cannot be used in the same subdirectory, this will cause an error. Yaml files in the same directory as the `helmfile.d` are silently ignored (so may be referenced from your `helmfile.d` files.

# Debugging lovely's behaviour locally

You can download `argocd-lovely-plugin` binary and run it in an application directory. Errors will go to stderr, and the rendered yaml will appear on stdout.
- You will need helm, helmfile and kustomize on your path if you use those. You will also need git and bash.
- You should set up the expected environment variables. Remember `ARGOCD_APP_NAME` needs to be set for helm chart rendering, or can be overridden with `LOVELY_HELM_NAME`. One or other must be set.

# Videos
If you prefer to watch videos of things rather than read words, we have compiled some demos.

- [Using Argo CD to Kustomize a Helm chart, and deploy additional manifests](https://youtu.be/OMae_8DHELA)
- [A bit about the general design of lovely plugin](https://youtu.be/5BLHwWlgd1k)
- [Promoted at ArgoCon EU 2023 by Michael Crenshaw from Intuit/Argo CD](https://youtu.be/uYKjSlarlN4?t=1175)
