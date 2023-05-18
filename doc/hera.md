# [Hera](https://github.com/argoproj-labs/hera)

Hera is a project that allows you to write [argo-workflow](https://github.com/argoproj/argo-workflows/) templates (and workflows) in python.

You can use Argo CD and the hera lovely plugin to commit those to git and have them rendered to yaml and injected directly into your cluster; gitops style.

If you commit .py files to git, and ensure your templates are rendered to disk in them with

```
obj.toFile('.')
```

then this plugin will automatically deploy them to your cluster. It doesn't make much sense usually to push plain `Workflows` to through Argo CD.

See [dag.py](../examples/hera/dag.py) for an example.

## Hera lovely plugin

If you'd like to use the hera lovely plugin to process standard yaml, you will need to unset the environment variables
```
LOVELY_DETECTION_REGEX
LOVELY_PREPROCESSORS
```
So that it detects yaml rather than requiring python.
