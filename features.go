package main

// The control of this is via environment variables, as that
// is the way argocd allows you to control plugins
import (
	"os"
	"strings"
)

const argo_prefix = `ARGOCD_ENV_`

func getArgoEnv(name string, defaultVal string) string {
	result, got := os.LookupEnv(argo_prefix + name)
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

// Plugins returns the list of plugins to run during the generate phase after main processing
// Set ARGOCD_ENV_LOVELY_PLUGINS to a comma separated list of plugins to run after other processing.
func Plugins() []string {
	return getPlugins(`LOVELY_PLUGINS`)
}

// Preprocessors returns the list of plugins to run before we generate yaml.
// Set ARGOCD_ENV_LOVELY_PREPROCESSORS to a comma separated list of plugins to run on the directory before any other processing.
func Preprocessors() []string {
	return getPlugins(`LOVELY_PREPROCESSORS`)
}

// KustomizeBinary returns the path to kustomize if overridden, otherwise we search the path
// Set ARGOCD_ENV_LOVELY_KUSTOMIZE_PATH to the path to the kustomize binary
func KustomizeBinary() string {
	return getArgoEnv(`LOVELY_KUSTOMIZE_PATH`, `kustomize`)
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

// HelmValues returns a list of values files to be fed to helm
// Set ARGOCD_ENV_LOVELY_HELM_VALUES to a comma separated list
// of files.
func HelmValues() []string {
	return getPlugins(`LOVELY_HELM_VALUES`)
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
