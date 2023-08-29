# [ytt](https://github.com/carvel-dev/ytt)

`ytt` is a YAML templating tool that works on YAML structure instead of text.

It can be used to achieve complex YAML manipulation (templating, merging,
patching/overlaying, modularizing, etc). Check out [ytt
website](https://carvel.dev/ytt/) for an interactive playground and the
documentation

## ytt lovely plugin

The `lovely-ytt` plugin doesn't set any default configuration: in other words, using the `lovely-ytt` plugin is exactly the same as using the `lovely` plugin, except that the `ytt` binary is built-in.

To actually make use of `ytt`, you need to configure either of the `plugins` or
`preprocessors` parameter (which are [described
here](https://github.com/crumbhole/argocd-lovely-plugin/blob/main/doc/parameter.md#available-parameters)).

For example, to use it as a preprocessor, you need to configure your application as follows:

```yaml
project: default
source:
  plugin:
    name: lovely-ytt
    env:
      - name: your_env_var_name
        value: eureka
    parameters:
      - name: lovely_preprocessors
        string: ytt --data-values-env=ARGOCD_ENV --file . --output-files .
```

The above would run `ytt` on a sub-application folder, and the environment
variables that you specify in `spec.source.plugin.env` will be available for
templating inside your manifest. For completeness, here is what your manifest
could typically look like:

```yaml
#@ load("@ytt:data", "data")
#@yaml/text-templated-strings
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: test-ytt
spec:
  selector:
    matchLabels:
      app: test-ytt
  template:
    metadata:
      labels:
        app: test-ytt
    spec:
      containers:
      - name: test-pod-(@= data.values.your_env_var_name @)
        image: docker.io/nicolaka/netshoot
        command: ["sleep"]
        args: ["infinity"]
```

Finally, it's worth noting that if you want to use `ytt` to template a Helm
Chart, it's probably best to adapt the `preprocessor` parameter a bit, as you
`ytt` doesn't like Helm Charts Jinja2 templates. Therefore, it's probably a
good idea to only template the `values.yaml` file when working with Helm
Charts, and to that end, you should configure the plugin with:

```yaml
source:
  plugin:
    parameters:
      - name: lovely_preprocessors
        string: ytt --data-values-env=ARGOCD_ENV --file values.yaml --output-files .
```

## Notes

`ytt` could also be used as a plugin rather than as a preprocessor as described
above; to that end, do not set the `--output-files .` argument when configuring
the plugin, so that ytt actually writes the rendered manifest to `stdout`.

Concretely, the config would be:

```yaml
source:
  plugin:
    parameters:
      - name: lovely_plugins
        string: ytt --data-values-env=ARGOCD_ENV --file .
```
