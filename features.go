package main

// The control of this is via environment variables, as that
// is the way argocd allows you to control plugins
import (
	"os"
	"strings"
)

func getPlugins(envname string) []string {
	pluginsText, got := os.LookupEnv(envname)
	if got {
		plugins := strings.Split(pluginsText, `,`)
		for i, plugin := range plugins {
			plugins[i] = strings.TrimSpace(plugin)
		}
		return plugins
	}
	return make([]string, 0)
}

// Plugins returns the list of plugins to run during the generate phase after main processing
// Set LOVELY_PLUGINS to a comma separated list of plugins to run after other processing.
func Plugins() []string {
	return getPlugins(`LOVELY_PLUGINS`)
}

// Preprocessors returns the list of plugins to run before we generate yaml.
// Set LOVELY_PREPROCESSORS to a comma separated list of plugins to run on the directory before any other processing.
func Preprocessors() []string {
	return getPlugins(`LOVELY_PREPROCESSORS`)
}

// KustomizeBinary returns the path to kustomize if overridden, otherwise we search the path
// Set LOVELY_KUSTOMIZE_PATH to the path to the kustomize binary
func KustomizeBinary() string {
	kustomize, got := os.LookupEnv(`LOVELY_KUSTOMIZE_PATH`)
	if !got {
		return `kustomize`
	}
	return kustomize
}

// HelmBinary returns the path to helm if overridden, otherwise we search the path
// Set LOVELY_HELM_PATH to the path to the helm binary
func HelmBinary() string {
	helm, got := os.LookupEnv(`LOVELY_HELM_PATH`)
	if !got {
		return `helm`
	}
	return helm
}

// HelmMerge returns the yaml to strategic merge into values.yaml
// Set LOVELY_HELM_MERGE to some yaml you'd like strategic merged into any values.yaml files used by helm
func HelmMerge() string {
	helm, got := os.LookupEnv(`LOVELY_HELM_MERGE`)
	if !got {
		return ``
	}
	return helm
}

// HelmPatch returns the yaml to json6902 patch into values.yaml
// Set LOVELY_HELM_PATCH to some yaml you'd like json6902 patched into any values.yaml files used by helm
func HelmPatch() string {
	helm, got := os.LookupEnv(`LOVELY_HELM_PATCH`)
	if !got {
		return ``
	}
	return helm
}

// KustomizeMerge returns the yaml to strategic merge into kustomization.yaml
// Set LOVELY_KUSTOMIZE_MERGE to some yaml you'd like strategic merged on any kustomization.yaml files used by kustomize
func KustomizeMerge() string {
	kustomize, got := os.LookupEnv(`LOVELY_KUSTOMIZE_MERGE`)
	if !got {
		return ``
	}
	return kustomize
}

// KustomizePatch returns the yaml to json6902 patch into kustomization.yaml
// Set LOVELY_KUSTOMIZE_PATCH to some yaml you'd like json6902 patched on any kustomization.yaml files used by kustomize
func KustomizePatch() string {
	kustomize, got := os.LookupEnv(`LOVELY_KUSTOMIZE_PATCH`)
	if !got {
		return ``
	}
	return kustomize
}
