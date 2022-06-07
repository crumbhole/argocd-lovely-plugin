# Lovely plugins

Lovely has two places where it will call external programs. In all cases the list of extrnal programs to call is an ordered comma delimited list of programs set in an environment variable, which will be called in that order. Each plugin is executed with bash -c <plugin and parameters>, so you can pass parameters as hoped for.

In all cases the working directory will be the location of the application source to be processed.

## ARGOCD_ENV_LOVELY_PREPROCESSORS

These plugins will be called before any other processing. They may add or modify any files inside the current directory, and are intended to do so. This is their only way of interacting with the process.

The idea is that you can modify Chart.yaml or kustomization.yaml before the main processing occurs on those files.

A preprocessor fails by returning a non-zero exit code and error information on stderr.

## ARGOCD_ENV_LOVELY_PLUGINS

These plugins will be called after any other processing happens. These will get the current yaml piped into them on stdin and should output all of the resulting yaml on stdout. If your plugin makes no changes stdin should be echoed to stdout.

A plugin fails by returning a non-zero exit code and error information on stderr.


