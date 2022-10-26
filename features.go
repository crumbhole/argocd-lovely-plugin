package main

// The control of this is via environment variables, as that
// is the way argocd allows you to control plugins
import (
	"fmt"
	"os"
	yaml "sigs.k8s.io/yaml"
	"strconv"
	"strings"
)

const argoPrefix = `ARGOCD_ENV_`

func getArgoEnv(name string, defaultVal string) string {
	result, got := os.LookupEnv(argoPrefix + name)
	if !got {
		result, got = os.LookupEnv(name)
		if !got {
			return defaultVal
		}
	}
	return result
}

func getPlugins(envname string) []string {
	pluginsText := getArgoEnv(envname, ``)
	if pluginsText == `` {
		return make([]string, 0)
	}
	plugins := strings.Split(pluginsText, `,`)
	for i, plugin := range plugins {
		plugins[i] = strings.TrimSpace(plugin)
	}
	return plugins
}

type pluginYaml map[string][]string

func getYamlPlugins(envname string) (pluginYaml, error) {
	pluginsText := getArgoEnv(envname, ``)
	if pluginsText == `` {
		return make(pluginYaml), nil
	}
	var plugins pluginYaml
	err := yaml.Unmarshal([]byte(pluginsText), &plugins)
	if err != nil {
		return make(pluginYaml), err
	}
	return plugins, nil
}

// Path here must be a relative path
func pluginsForPath(path string, yamlEnv string, plainEnv string) ([]string, error) {
	plugins, err := getYamlPlugins(yamlEnv)
	if err != nil {
		return make([]string, 0), fmt.Errorf("%s is invalid: %s", yamlEnv, err)
	}
	if list, contains := plugins[path]; contains {
		return list, nil
	}
	return getPlugins(plainEnv), nil
}

// Plugins returns the list of plugins to run during the generate phase after main processing
// Set ARGOCD_ENV_LOVELY_PLUGINS_YAML to a map of a list of plugins to run per directory
// e.g.
// foo/bar:
// - plugin1
// - plugin2
// helm:
// - plugin3
// Or et ARGOCD_ENV_LOVELY_PLUGINS to a comma separated list of plugins to run after other processing.
// for any directories not in the list from the YAML
func Plugins(path string) ([]string, error) {
	return pluginsForPath(path, `LOVELY_PLUGINS_YAML`, `LOVELY_PLUGINS`)
}

// Preprocessors returns the list of plugins to run before we generate yaml.
// Set ARGOCD_ENV_LOVELY_PREPROCESSORS_YAML to a map of a list of plugins to run per directory
// Set ARGOCD_ENV_LOVELY_PREPROCESSORS to a comma separated list of plugins to run
// for any directories not in the list from the YAML
func Preprocessors(path string) ([]string, error) {
	return pluginsForPath(path, `LOVELY_PREPROCESSORS_YAML`, `LOVELY_PREPROCESSORS`)
}

// KustomizeBinary returns the path to kustomize if overridden, otherwise we search the path
// Set ARGOCD_ENV_LOVELY_KUSTOMIZE_PATH to the path to the kustomize binary
func KustomizeBinary() string {
	return getArgoEnv(`LOVELY_KUSTOMIZE_PATH`, `kustomize`)
}

// KustomizeParams returns extra parameters to pass to kustomize
// Set ARGOCD_ENV_LOVELY_KUSTOMIZE_PARAMS to extra parameters to pass to kustomize
func KustomizeParams() []string {
	paramsStr := getArgoEnv(`LOVELY_KUSTOMIZE_PARAMS`, ``)
	if paramsStr == `` {
		return []string{}
	}
	params := strings.Split(paramsStr, ` `)
	for i, param := range params {
		params[i] = strings.TrimSpace(param)
	}
	return params
}

// HelmBinary returns the path to helm if overridden, otherwise we search the path
// Set ARGOCD_ENV_LOVELY_HELM_PATH to the path to the helm binary
func HelmBinary() string {
	return getArgoEnv(`LOVELY_HELM_PATH`, `helm`)
}

// HelmMerge returns the yaml to strategic merge into values.yaml
// Set ARGOCD_ENV_LOVELY_HELM_MERGE to some yaml you'd like strategic merged into any values.yaml files used by helm
func HelmMerge() string {
	return getArgoEnv(`LOVELY_HELM_MERGE`, ``)
}

// HelmPatch returns the yaml to json6902 patch into values.yaml
// Set ARGOCD_ENV_LOVELY_HELM_PATCH to some yaml you'd like json6902 patched into any values.yaml files used by helm
func HelmPatch() string {
	return getArgoEnv(`LOVELY_HELM_PATCH`, ``)
}

// KustomizeMerge returns the yaml to strategic merge into kustomization.yaml
// Set ARGOCD_ENV_LOVELY_KUSTOMIZE_MERGE to some yaml you'd like strategic merged on any kustomization.yaml files used by kustomize
func KustomizeMerge() string {
	return getArgoEnv(`LOVELY_KUSTOMIZE_MERGE`, ``)
}

// KustomizePatch returns the yaml to json6902 patch into kustomization.yaml
// Set ARGOCD_ENV_LOVELY_KUSTOMIZE_PATCH to some yaml you'd like json6902 patched on any kustomization.yaml files used by kustomize
func KustomizePatch() string {
	return getArgoEnv(`LOVELY_KUSTOMIZE_PATCH`, ``)
}

// AllowGitCheckout establishes if git is safe to use
// Set ARGOCD_ENV_ALLOW_GITCHECKOUT to true to say you've told Argo this is safe
func AllowGitCheckout() bool {
	res, err := strconv.ParseBool(getArgoEnv(`LOVELY_ALLOW_GITCHECKOUT`, `false`))
	if err != nil {
		return false
	}
	return res
}
