package features

// The control of this is via environment variables, as that
// is the way argocd allows you to control plugins
import (
	"github.com/crumbhole/argocd-lovely-plugin/pkg/config"
	"os"
	"strconv"
)

// GetPlugins returns the list of plugins to run during the generate phase after main processing
// Set LOVELY_PLUGINS_YAML to a map of a list of plugins to run per directory
// e.g.
// foo/bar:
// - plugin1
// - plugin2
// helm:
// - plugin3
// Or et LOVELY_PLUGINS to a comma separated list of plugins to run after other processing.
// for any directories not in the list from the YAML
func GetPlugins(path string) ([]string, error) {
	f := Features[Plugins]
	fy := Features[PluginsYaml]
	return pluginsForPath(path, fy.EnvName(), f.EnvName())
}

// GetPreprocessors returns the list of plugins to run before we generate yaml.
// Set LOVELY_PREPROCESSORS_YAML to a map of a list of plugins to run per directory
// Set LOVELY_PREPROCESSORS to a comma separated list of plugins to run
// for any directories not in the list from the YAML
func GetPreprocessors(path string) ([]string, error) {
	f := Features[Preprocessors]
	fy := Features[PreprocessorsYaml]
	return pluginsForPath(path, fy.EnvName(), f.EnvName())
}

// GetDetectionRegex returns the regex used to detect applications
// Set LOVELY_DETECTION_REGEX to the regex to detect applications
func GetDetectionRegex() string {
	f := Features[DetectRx]
	return config.GetStringParam(f.EnvName(), f.DefaultVal)
}

// GetKustomizePath returns the path to kustomize if overridden, otherwise we search the path
// Set LOVELY_KUSTOMIZE_PATH to the path to the kustomize binary
func GetKustomizePath() string {
	f := Features[KustomizePath]
	return config.GetStringParam(f.EnvName(), f.DefaultVal)
}

// GetKustomizeParams returns extra parameters to pass to kustomize
// Set LOVELY_KUSTOMIZE_PARAMS to extra parameters to pass to kustomize
func GetKustomizeParams() []string {
	f := Features[KustomizeParams]
	return config.GetStringListParam(f.EnvName(), f.DefaultVal, ` `)
}

// GetHelmPath returns the path to helm if overridden, otherwise we search the path
// Set LOVELY_HELM_PATH to the path to the helm binary
func GetHelmPath() string {
	f := Features[HelmPath]
	return config.GetStringParam(f.EnvName(), f.DefaultVal)
}

// GetHelmMerge returns the yaml to strategic merge into values.yaml
// Set LOVELY_HELM_MERGE to some yaml you'd like strategic merged into any values.yaml files used by helm
func GetHelmMerge() string {
	f := Features[HelmMerge]
	return config.GetStringParam(f.EnvName(), f.DefaultVal)
}

// GetHelmPatch returns the yaml to json6902 patch into values.yaml
// Set LOVELY_HELM_PATCH to some yaml you'd like json6902 patched into any values.yaml files used by helm
func GetHelmPatch() string {
	f := Features[HelmPatch]
	return config.GetStringParam(f.EnvName(), f.DefaultVal)
}

// GetHelmTemplateParams returns extra parameters to pass to helm template
// Set LOVELY_HELM_TEMPLATE_PARAMS to extra parameters to pass to helm template
func GetHelmTemplateParams() []string {
	f := Features[HelmTemplateParams]
	return config.GetStringListParam(f.EnvName(), f.DefaultVal, ` `)
}

// GetHelmRepoAddParams returns extra parameters to pass to helm repo add
// Set LOVELY_HELM_REPO_ADD_PARAMS to extra parameters to pass to helm template
func GetHelmRepoAddParams() []string {
	f := Features[HelmRepoAddParams]
	return config.GetStringListParam(f.EnvName(), f.DefaultVal, ` `)
}

// GetKustomizeMerge returns the yaml to strategic merge into kustomization.yaml
// Set LOVELY_KUSTOMIZE_MERGE to some yaml you'd like strategic merged on any kustomization.yaml files used by kustomize
func GetKustomizeMerge() string {
	f := Features[KustomizeMerge]
	return config.GetStringParam(f.EnvName(), f.DefaultVal)
}

// GetKustomizePatch returns the yaml to json6902 patch into kustomization.yaml
// Set LOVELY_KUSTOMIZE_PATCH to some yaml you'd like json6902 patched on any kustomization.yaml files used by kustomize
func GetKustomizePatch() string {
	f := Features[KustomizePatch]
	return config.GetStringParam(f.EnvName(), f.DefaultVal)
}

// GetAllowGitCheckout establishes if git is safe to use
// Set ALLOW_GITCHECKOUT to true to say you've told Argo this is safe
func GetAllowGitCheckout() bool {
	f := Features[AllowGitCheckout]
	res, err := strconv.ParseBool(config.GetStringParam(f.EnvName(), f.DefaultVal))
	if err != nil {
		return false
	}
	return res
}

// GetHelmName gives us the application name for helm
// Set LOVELY_HELM_NAME to override the default of ARGOCD_APP_NAME
func GetHelmName() string {
	f := Features[HelmName]
	nameOverride := config.GetStringParam(f.EnvName(), f.DefaultVal)
	if nameOverride == `` {
		return os.Getenv(`ARGOCD_APP_NAME`)
	}
	return nameOverride
}

// GetHelmNamespace gives us the namespace for helm
func GetHelmNamespace() string {
	f := Features[HelmNamespace]
	namespaceOverride := config.GetStringParam(f.EnvName(), f.DefaultVal)
	if namespaceOverride == `` {
		return os.Getenv(`ARGOCD_APP_NAMESPACE`)
	}
	return namespaceOverride11
}

// GetHelmValues gives us the values file we're going to use for helm
// Set LOVELY_HELM_VALUES to the path to use for the values file
func GetHelmValues() []string {
	f := Features[HelmValues]
	return config.GetStringListParam(f.EnvName(), `values.yaml`, ` `)
}

// GetHelmValuesSet returns true if HelmValues() is explicily set
func GetHelmValuesSet() bool {
	f := Features[HelmValues]
	return len(config.GetStringListParam(f.EnvName(), f.DefaultVal, ` `)) != 0
}

// GetHelmfilePath returns the path to helm if overridden, otherwise we search the path
// Set LOVELY_HELMFILE_PATH to the path to the helm binary
func GetHelmfilePath() string {
	f := Features[HelmfilePath]
	return config.GetStringParam(f.EnvName(), f.DefaultVal)
}

// GetHelmfileMerge returns the yaml to strategic merge into values.yaml
// Set LOVELY_HELMFILE_MERGE to some yaml you'd like strategic merged into any helmfile.yaml files used by helmfile
func GetHelmfileMerge() string {
	f := Features[HelmfileMerge]
	return config.GetStringParam(f.EnvName(), f.DefaultVal)
}

// GetHelmfilePatch returns the yaml to json6902 patch into values.yaml
// Set LOVELY_HELMFILE_PATCH to some yaml you'd like json6902 patched into any helmfile.yaml files used by helmfile
func GetHelmfilePatch() string {
	f := Features[HelmfilePatch]
	return config.GetStringParam(f.EnvName(), f.DefaultVal)
}
