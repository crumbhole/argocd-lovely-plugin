# How lovely does its thing

Lovely processes your application directory in two main steps.

1. Separate the application into sub-applications
1. Process the files found in that sub-application

## Sub-applications

Separation into sub-applications allows lovely to process each of the sub-applications independently. They can be totally different, and processed using different pipelines if you use the YAML configuration of `plugins` and `preprocessors`.

At a simple level they can just allow you to have a helm chart in one directory and some yaml in another, which is easier to read.

```mermaid
flowchart LR
    start[Recurse Directories] --> files{"Files match
    LOVELY_DETECTION_REGEX"}
    files -->|No| dir{"Has helmfile.d
    subdirectory"}
    dir -->|No| recurse[Keep recursing]
    recurse --> start
    files -->|Yes| mark["Mark as sub-application
    stop recursing deeper"]
    dir -->|Yes| mark
```

## Ensuring a clean copy

Lovely's processing of a sub-application can modify files. As modifications made may not be idempotent, we need to ensure we are working on an unmodified copy of the files. Lovely has 3 strategies for dealing with this:
* Sidecar: When running as a sidecar, we get a fresh copy, so lovely doesn't do anything special. The sidecar images set `LOVELY_SIDECAR=true`.
* Configmap plugin: Normally lovely will copy the sub-application folder only to a temporary folder and process them there. This means that if you refer to files outside the sub-application folder then it will not work.
* Configmap plugin with `LOVELY_ALLOW_GITCHECKOUT=true`. In this case we will perform a git checkout after processing to undo any changes made by lovely. Check [this documentation](doc/allow_git.md).

```mermaid
flowchart LR
    start{LOVELY_SIDECAR} -->|true| nothing["Do no special process"]
	start -->|not true| check{LOVELY_ALLOW_GITCHECKOUT}
	check -->|not true| copy[Copy files to temporary dir]
	check -->|true| git[Git checkout at end]
```

## Processing

Each sub-application is processed like this

```mermaid
flowchart TD
    preprocessors[Each preprocessor] -->|Next preprocessor| preprocess["Call `bash -c` with
    the preprocessor.
    Working directory: sub-application
    Preprocessor should change disk files
    Exit code should be zero to continue"]
    preprocess -->|Exit 0| preprocessors
    preprocessors -->|All preprocessors run| helmcheck{Helm and helmfile present}
    helmcheck -->|Yes| error[Not allowed]
    helmcheck -->|No| helmfile{"helmfile.yaml or
    helmfile.d present"}
    helmfile -->|Yes| helmfileprocess["Run helmfile"]
    helmfile -->|No| helm{"Chart.yaml present"}
    helm -->|Yes| helmprocess["Run helm"]
    helm -->|No| kustomize{"kustomization.yaml present"}
    helmprocess --> kustomize
    helmfileprocess --> kustomize
    kustomize -->|Yes| kustomizeprocess["Run kustomize
    with any helm/helmfile
    output magically included
    as a resource"]
    kustomize -->|No| yaml{"Has helm or helmfile been run"}
    yaml -->|No| yamlprocess["Read yaml from disk"]
    yaml -->|Yes| plugins[Each plugin]
    kustomizeprocess --> plugins
    yamlprocess --> plugins
    plugins -->|Next plugin| plugin["Call `bash -c` with
    the plugin.
    stdin: The yaml to preprocess
    stdout: The processed yaml
    Exit code should be zero to continue"]
    plugin -->|Exit 0| plugins
    plugins -->|All plugins run| Done
```
