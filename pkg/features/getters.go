package features

// The control of this is via environment variables, as that
// is the way argocd allows you to control plugins
import (
	"github.com/crumbhole/argocd-lovely-plugin/pkg/config"
	"os"
	"strconv"
	"strings"
)

// Plugins returns the list of plugins to run during the generate phase after main processing
// Set LOVELY_PLUGINS_YAML to a map of a list of plugins to run per directory
// e.g.
// foo/bar:
// - plugin1
// - plugin2
// helm:
// - plugin3
// Or et LOVELY_PLUGINS to a comma separated list of plugins to run after other processing.
// for any directories not in the list from the YAML
func Plugins(path string) ([]string, error) {
	return pluginsForPath(path, `LOVELY_PLUGINS_YAML`, `LOVELY_PLUGINS`)
}

// Preprocessors returns the list of plugins to run before we generate yaml.
// Set LOVELY_PREPROCESSORS_YAML to a map of a list of plugins to run per directory
// Set LOVELY_PREPROCESSORS to a comma separated list of plugins to run
// for any directories not in the list from the YAML
func Preprocessors(path string) ([]string, error) {
	return pluginsForPath(path, `LOVELY_PREPROCESSORS_YAML`, `LOVELY_PREPROCESSORS`)
}

// DetectionRegex returns the regex used to detect applications
// Set LOVELY_DETECTION_REGEX to the regex to detect applications
func DetectionRegex() string {
	f := Features[DetectRx]
	return config.GetStringParam(f.EnvName(), f.DefaultVal)
}

// KustomizeBinary returns the path to kustomize if overridden, otherwise we search the path
// Set LOVELY_KUSTOMIZE_PATH to the path to the kustomize binary
func KustomizeBinary() string {
	f := Features[KustomizePath]
	return config.GetStringParam(f.EnvName(), f.DefaultVal)
}

// KustomizeParams returns extra parameters to pass to kustomize
// Set LOVELY_KUSTOMIZE_PARAMS to extra parameters to pass to kustomize
func KustomizeParams() []string {
	paramsStr := config.GetStringParam(`LOVELY_KUSTOMIZE_PARAMS`, ``)
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
// Set LOVELY_HELM_PATH to the path to the helm binary
func HelmBinary() string {
	f := Features[HelmPath]
	return config.GetStringParam(f.EnvName(), f.DefaultVal)
}

// HelmMerge returns the yaml to strategic merge into values.yaml
// Set LOVELY_HELM_MERGE to some yaml you'd like strategic merged into any values.yaml files used by helm
func HelmMerge() string {
	return config.GetStringParam(`LOVELY_HELM_MERGE`, ``)
}

// HelmPatch returns the yaml to json6902 patch into values.yaml
// Set LOVELY_HELM_PATCH to some yaml you'd like json6902 patched into any values.yaml files used by helm
func HelmPatch() string {
	return config.GetStringParam(`LOVELY_HELM_PATCH`, ``)
}

// HelmTemplateParams returns extra parameters to pass to helm template
// Set LOVELY_HELM_TEMPLATE_PARAMS to extra parameters to pass to helm template
func HelmTemplateParams() []string {
	paramsStr := config.GetStringParam(`LOVELY_HELM_TEMPLATE_PARAMS`, ``)
	if paramsStr == `` {
		return []string{}
	}
	params := strings.Split(paramsStr, ` `)
	for i, param := range params {
		params[i] = strings.TrimSpace(param)
	}
	return params
}

// HelmRepoAddParams returns extra parameters to pass to helm repo add
// Set LOVELY_HELM_REPO_ADD_PARAMS to extra parameters to pass to helm template
func HelmRepoAddParams() []string {
	paramsStr := config.GetStringParam(`LOVELY_HELM_REPO_ADD_PARAMS`, ``)
	if paramsStr == `` {
		return []string{}
	}
	params := strings.Split(paramsStr, ` `)
	for i, param := range params {
		params[i] = strings.TrimSpace(param)
	}
	return params
}

// KustomizeMerge returns the yaml to strategic merge into kustomization.yaml
// Set LOVELY_KUSTOMIZE_MERGE to some yaml you'd like strategic merged on any kustomization.yaml files used by kustomize
func KustomizeMerge() string {
	return config.GetStringParam(`LOVELY_KUSTOMIZE_MERGE`, ``)
}

// KustomizePatch returns the yaml to json6902 patch into kustomization.yaml
// Set LOVELY_KUSTOMIZE_PATCH to some yaml you'd like json6902 patched on any kustomization.yaml files used by kustomize
func KustomizePatch() string {
	return config.GetStringParam(`LOVELY_KUSTOMIZE_PATCH`, ``)
}

// AllowGitCheckout establishes if git is safe to use
// Set ALLOW_GITCHECKOUT to true to say you've told Argo this is safe
func AllowGitCheckout() bool {
	res, err := strconv.ParseBool(config.GetStringParam(`LOVELY_ALLOW_GITCHECKOUT`, `false`))
	if err != nil {
		return false
	}
	return res
}

// HelmName gives us the application name for helm
// Set LOVELY_HELM_NAME to override the default of ARGOCD_APP_NAME
func HelmName() string {
	nameOverride := config.GetStringParam(`LOVELY_HELM_NAME`, ``)
	if nameOverride == `` {
		return os.Getenv(`ARGOCD_APP_NAME`)
	}
	return nameOverride
}

// HelmNamespace gives us the namespace for helm
func HelmNamespace() string {
	return os.Getenv(`ARGOCD_APP_NAMESPACE`)
}

// HelmValues gives us the values file we're going to use for helm
// Set LOVELY_HELM_VALUES to the path to use for the values file
func HelmValues() string {
	return config.GetStringParam(`LOVELY_HELM_VALUES`, `values.yaml`)
}

// HelmValuesSet returns true if HelmValues() is explicily set
func HelmValuesSet() bool {
	return config.GetStringParam(`LOVELY_HELM_VALUES`, ``) != ``
}

// HelmfileBinary returns the path to helm if overridden, otherwise we search the path
// Set LOVELY_HELMFILE_PATH to the path to the helm binary
func HelmfileBinary() string {
	return config.GetStringParam(`LOVELY_HELMFILE_PATH`, `helmfile`)
}

// HelmfileMerge returns the yaml to strategic merge into values.yaml
// Set LOVELY_HELMFILE_MERGE to some yaml you'd like strategic merged into any helmfile.yaml files used by helmfile
func HelmfileMerge() string {
	return config.GetStringParam(`LOVELY_HELMFILE_MERGE`, ``)
}

// HelmfilePatch returns the yaml to json6902 patch into values.yaml
// Set LOVELY_HELMFILE_PATCH to some yaml you'd like json6902 patched into any helmfile.yaml files used by helmfile
func HelmfilePatch() string {
	return config.GetStringParam(`LOVELY_HELMFILE_PATCH`, ``)
}
