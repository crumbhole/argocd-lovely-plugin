package main

// The control of this is via environment variables, as that
// is the way argocd allows you to control plugins
import (
	"os"
	"strings"
)

func Plugins() []string {
	// Set LOVELY_PLUGINS to a comma separated list of plugins to run
	pluginsText, got := os.LookupEnv(`LOVELY_PLUGINS`)
	if got {
		plugins := strings.Split(pluginsText, `,`)
		for i, plugin := range plugins {
			plugins[i] = strings.TrimSpace(plugin)
		}
		return plugins
	}
	return make([]string, 0)
}

func KustomizeBinary() string {
	// Set LOVELY_KUSTOMIZE_PATH to the path to the kustomize binary
	kustomize, got := os.LookupEnv(`LOVELY_KUSTOMIZE_PATH`)
	if !got {
		return `kustomize`
	}
	return kustomize
}

func HelmBinary() string {
	// Set LOVELY_HELM_PATH to the path to the helm binary
	helm, got := os.LookupEnv(`LOVELY_HELM_PATH`)
	if !got {
		return `helm`
	}
	return helm
}

func HelmValues() string {
	// Set LOVELY_HELM_VALUES to some yaml you'd like overlayed on any values.yaml files used by helm
	helm, got := os.LookupEnv(`LOVELY_HELM_VALUES`)
	if !got {
		return ``
	}
	return helm
}

func KustomizeExtra() string {
	// Set LOVELY_KUSTOMIZE_EXTRA to some yaml you'd like overlayed on any kustomization.yaml files used
	kustomize, got := os.LookupEnv(`LOVELY_KUSTOMIZE_EXTRA`)
	if !got {
		return ``
	}
	return kustomize
}
