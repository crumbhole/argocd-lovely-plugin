package features

// The control of this is via environment variables, as that
// is the way argocd allows you to control plugins
import (
	"fmt"
	"github.com/crumbhole/argocd-lovely-plugin/pkg/config"
	"os"
	"gopkg.in/yaml.v3"
	"strconv"
)

func getPlugins(envname string) []string {
	return config.GetStringListParam(envname, ``, `,`)
}

type pluginYaml map[string][]string

func getYamlPlugins(envname string) (pluginYaml, error) {
	pluginsText := config.GetStringParam(envname, ``)
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
	return config.GetStringParam(`LOVELY_DETECTION_REGEX`, `\.ya?ml$`)
}

// KustomizeBinary returns the path to kustomize if overridden, otherwise we search the path
// Set LOVELY_KUSTOMIZE_PATH to the path to the kustomize binary
func KustomizeBinary() string {
	return config.GetStringParam(`LOVELY_KUSTOMIZE_PATH`, `kustomize`)
}

// KustomizeParams returns extra parameters to pass to kustomize
// Set LOVELY_KUSTOMIZE_PARAMS to extra parameters to pass to kustomize
func KustomizeParams() []string {
	return config.GetStringListParam(`LOVELY_KUSTOMIZE_PARAMS`, ``, ` `)
}

// HelmBinary returns the path to helm if overridden, otherwise we search the path
// Set LOVELY_HELM_PATH to the path to the helm binary
func HelmBinary() string {
	return config.GetStringParam(`LOVELY_HELM_PATH`, `helm`)
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
	return config.GetStringListParam(`LOVELY_HELM_TEMPLATE_PARAMS`, ``, ` `)
}

// HelmRepoAddParams returns extra parameters to pass to helm repo add
// Set LOVELY_HELM_REPO_ADD_PARAMS to extra parameters to pass to helm template
func HelmRepoAddParams() []string {
	return config.GetStringListParam(`LOVELY_HELM_REPO_ADD_PARAMS`, ``, ` `)
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

// HelmValues gives us the values files we're going to use for helm
// Set LOVELY_HELM_VALUES to the path to use for the values file
func HelmValues() []string {
	return config.GetStringListParam(`LOVELY_HELM_VALUES`, `values.yaml`, ` `)
}

// HelmValuesSet returns true if HelmValues() is explicily set
func HelmValuesSet() bool {
	return len(config.GetStringListParam(`LOVELY_HELM_VALUES`, ``, ` `)) != 0
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
