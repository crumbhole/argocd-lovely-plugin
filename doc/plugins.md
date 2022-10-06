# Lovely plugins

Lovely has two places where it will call external programs, which we'll refer to as plugins. Each plugin is executed with `bash -c <plugin and parameters>`, so you can pass parameters as hoped for.

In all cases the working directory will be the location of the application source to be processed. So each subapplication will be processed independently.

## Two ways of configuring plugins/preprocessors

You can specify directory specific plugins for each sub-application using the `YAML` suffixed version of the environment variable. If no appropriate plugin exists and a generic plugin environment variable is set then that will be run instead.

The YAML variants take the subdirectory path to the sub-application without a leading or trailing `/`.

YAML example
```yaml
path:
- plugin1
- plugin2
other/path:
- plugin3
```

As JSON is YAML you can use JSON here if you're so inclined.

The non-YAML environment variable takes an ordered comma delimited list of programs set in an environment variable.

Example:
```plugin1, plugin2```

The order matters in both YAML and non yaml, and the plugins will be called in that order.

## Preprocessors

Configured using `ARGOCD_ENV_LOVELY_PREPROCESSORS` and `ARGOCD_ENV_LOVELY_PREPROCESSORS_YAML`.

These plugins will be called before any other processing. They may add or modify any files inside the current directory, and are intended to do so. This is their only way of interacting with the process.

The idea is that you can modify Chart.yaml or kustomization.yaml before the main processing occurs on those files.

A preprocessor fails by returning a non-zero exit code and error information on stderr.


## Plugins

Configured using `ARGOCD_ENV_LOVELY_PLUGINS` and `ARGOCD_ENV_LOVELY_PLUGINS_YAML`.

These plugins will be called after any other processing happens. These will get the current yaml piped into them on stdin and should output all of the resulting yaml on stdout. If your plugin makes no changes stdin should be echoed to stdout.

A plugin fails by returning a non-zero exit code and error information on stderr.


