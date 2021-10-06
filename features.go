package main

// The control of this is via environment variables, as that
// is the way argocd allows you to control plugins
import (
	"os"
	"strconv"
)

func VaultEnabled() bool {
	// Set LOVELY_DISABLE_VAULT to anything to disable vault
	disabled, err := strconv.ParseBool(os.Getenv(`LOVELY_DISABLE_VAULT`))
	if err != nil {
		return true
	}
	return !disabled
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
