package features

// The control of this is via environment variables, as that
// is the way argocd allows you to control plugins
import (
	"strings"
)

type featureId int

//type CollectionType int

const (
	// The order of these is the order they will appear in the docs
	Plugins featureId = iota
	PluginsYaml
	Preprocessors
	PreprocessorsYaml
	DetectRx
	KustomizePath
	KustomizeParams
	HelmPath
	HelmMerge
	HelmPatch
	HelmTemplateParams
	HelmRepoAddParams
	KustomizeMerge
	KustomizePatch
	AllowGitCheckout
	HelmName
	HelmNamespace
	HelmValues
	HelmfilePath
	HelmfileMerge
	HelmfilePatch
)

// const (
// 	Single CollectionType = iota
// 	Array
// 	Map
// )

// func (ct CollectionType) String() string {
// 	switch ct {
// 	case Single:
// 		return "string"
// 	case Array:
// 		return "array"
// 	case Map:
// 		return "map"
// 	}
// 	return "unknown"
// }

type Feature struct {
	// Name is the name of the parameter or environment variable, in lower_snake_case
	Name        string
	// Title is plain text user readable name for the parameter
	Title       string
	DefaultVal  string
	// Description is in markdown, but will be rendered as text to plugin.yaml
	Description string
	//	CollectionType CollectionType
}

func (f Feature) EnvName() string {
	return strings.ToUpper(f.Name)
}

var Features = map[featureId]Feature{
	Plugins: {
		Title: `Plugins`,
		Name:  `lovely_plugins`,
	},
	PluginsYaml: {
		Title: `Plugins YAML`,
		Name:  `lovely_plugins_yaml`,
	},
	Preprocessors: {
		Title: `PreProcessors`,
		Name:  `lovely_preprocessors`,
	},
	PreprocessorsYaml: {
		Title: `PreProcessors YAML`,
		Name:  `lovely_preprocessors_yaml`,
	},
	DetectRx: {
		Title:       `Detection Regular Expression`,
		Name:        `lovely_detection_regex`,
		DefaultVal:  `\.ya?ml$`,
		Description: `Regular expression used for detecting filenames that denote applications.`,
		// CollectionType: Single,
	},
	KustomizePath: {
		Title:       `Kustomize Path`,
		Name:        `lovely_kustomize_path`,
		DefaultVal:  `kustomize`,
		Description: "Path to the kustomize binary used for this application",
		// CollectionType: Single,
	},
	KustomizeParams: {
		Title:       `Kustomize parameters`,
		Name:        `lovely_kustomize_params`,
		Description: "Space separated extra parameters to `kustomize build` as you might use on the command line. `--enable-helm` is already passed always. You're on your own here if you pass rubbish parameters.",
		// CollectionType: Array,
	},
	HelmPath: {
		Title:       `Helm Path`,
		Name:        `lovely_helm_path`,
		DefaultVal:  `helm`,
		Description: "Path to the helm binary used for this application",
		// CollectionType: Single,
	},
	HelmMerge: {
		Title: `Helm Merge`,
		Name:  `lovely_helm_merge`,
	},
	HelmPatch: {
		Title: `Helm Patch`,
		Name:  `lovely_helm_patch`,
	},
	HelmTemplateParams: {
		Title: `Helm Template Parameters`,
		Name:  `lovely_helm_template_params`,
	},
	HelmRepoAddParams: {
		Title: `Helm Repo Add Parameters`,
		Name:  `lovely_helm_repo_add_params`,
	},
	KustomizeMerge: {
		Title: `Kustomize Merge`,
		Name:  `lovely_kustomize_merge`,
	},
	KustomizePatch: {
		Title: `Kustomize Patch`,
		Name:  `lovely_kustomize_patch`,
	},
	AllowGitCheckout: {
		Title:      `Allow Git Checkout`,
		Name:       `lovely_allow_gitcheckout`,
		DefaultVal: `false`,
	},
	HelmName: {
		Title: `Helm Name`,
		Name:  `lovely_helm_name`,
	},
	HelmNamespace: {
		Title: `Helm Namespace`,
		Name:  `lovely_helm_namespace`,
	},
	HelmValues: {
		Title: `Helm Values`,
		Name:  `lovely_helm_values`,
	},
	HelmfilePath: {
		Title:       `Helmfile Path`,
		Name:        `lovely_helmfile_path`,
		DefaultVal:  `helmfile`,
		Description: "Path to the helmfile binary used for this application",
		// CollectionType: Single,
	},
	HelmfileMerge: {
		Title:       `Helmfile Merge`,
		Name:        `lovely_helmfile_merge`,
		Description: `Set to some yaml you'd like [strategic merged](https://kubectl.docs.kubernetes.io/references/kustomize/kustomization/patchesstrategicmerge/) into any helmfile.yaml used by helmfile.`,
	},
	HelmfilePatch: {
		Title: `Helmfile Patch`,
		Name:  `lovely_helmfile_patch`,
	},
}

// // Plugins returns the list of plugins to run during the generate phase after main processing
// // Set LOVELY_PLUGINS_YAML to a map of a list of plugins to run per directory
// // e.g.
// // foo/bar:
// // - plugin1
// // - plugin2
// // helm:
// // - plugin3
// // Or set LOVELY_PLUGINS to a comma separated list of plugins to run after other processing.
// // for any directories not in the list from the YAML
// func Plugins(path string) ([]string, error) {
// 	return pluginsForPath(path, `LOVELY_PLUGINS_YAML`, `LOVELY_PLUGINS`)
// }

// // Preprocessors returns the list of plugins to run before we generate yaml.
// // Set LOVELY_PREPROCESSORS_YAML to a map of a list of plugins to run per directory
// // Set LOVELY_PREPROCESSORS to a comma separated list of plugins to run
// // for any directories not in the list from the YAML
// func Preprocessors(path string) ([]string, error) {
// 	return pluginsForPath(path, `LOVELY_PREPROCESSORS_YAML`, `LOVELY_PREPROCESSORS`)
// }

// // DetectionRegex returns the regex used to detect applications
// // Set LOVELY_DETECTION_REGEX to the regex to detect applications
// func DetectionRegex() string {
// 	return config.GetStringParam(`LOVELY_DETECTION_REGEX`, `\.ya?ml$`)
// }

// // KustomizeBinary returns the path to kustomize if overridden, otherwise we search the path
// // Set LOVELY_KUSTOMIZE_PATH to the path to the kustomize binary
// func KustomizeBinary() string {
// 	return config.GetStringParam(`LOVELY_KUSTOMIZE_PATH`, `kustomize`)
// }

// // KustomizeParams returns extra parameters to pass to kustomize
// // Set LOVELY_KUSTOMIZE_PARAMS to extra parameters to pass to kustomize
// func KustomizeParams() []string {
// 	paramsStr := config.GetStringParam(`LOVELY_KUSTOMIZE_PARAMS`, ``)
// 	if paramsStr == `` {
// 		return []string{}
// 	}
// 	params := strings.Split(paramsStr, ` `)
// 	for i, param := range params {
// 		params[i] = strings.TrimSpace(param)
// 	}
// 	return params
// }

// KustomizeParams returns extra parameters to pass to kustomize
// Set LOVELY_KUSTOMIZE_PARAMS to extra parameters to pass to kustomize
// func KustomizeParams() []string {
// 	return config.GetStringListParam(`LOVELY_KUSTOMIZE_PARAMS`, ``, ` `)
// }

// // HelmBinary returns the path to helm if overridden, otherwise we search the path
// // Set LOVELY_HELM_PATH to the path to the helm binary
// func HelmBinary() string {
// 	return config.GetStringParam(`LOVELY_HELM_PATH`, `helm`)
// }

// // HelmMerge returns the yaml to strategic merge into values.yaml
// // Set LOVELY_HELM_MERGE to some yaml you'd like strategic merged into any values.yaml files used by helm
// func HelmMerge() string {
// 	return config.GetStringParam(`LOVELY_HELM_MERGE`, ``)
// }

// // HelmPatch returns the yaml to json6902 patch into values.yaml
// // Set LOVELY_HELM_PATCH to some yaml you'd like json6902 patched into any values.yaml files used by helm
// func HelmPatch() string {
// 	return config.GetStringParam(`LOVELY_HELM_PATCH`, ``)
// }

// // HelmTemplateParams returns extra parameters to pass to helm template
// // Set LOVELY_HELM_TEMPLATE_PARAMS to extra parameters to pass to helm template
// func HelmTemplateParams() []string {
// 	paramsStr := config.GetStringParam(`LOVELY_HELM_TEMPLATE_PARAMS`, ``)
// 	if paramsStr == `` {
// 		return []string{}
// 	}
// 	params := strings.Split(paramsStr, ` `)
// 	for i, param := range params {
// 		params[i] = strings.TrimSpace(param)
// 	}
// 	return params
// }

// HelmTemplateParams returns extra parameters to pass to helm template
// Set LOVELY_HELM_TEMPLATE_PARAMS to extra parameters to pass to helm template
// func HelmTemplateParams() []string {
// 	return config.GetStringListParam(`LOVELY_HELM_TEMPLATE_PARAMS`, ``, ` `)
// }

// HelmRepoAddParams returns extra parameters to pass to helm repo add
// Set LOVELY_HELM_REPO_ADD_PARAMS to extra parameters to pass to helm template
// func HelmRepoAddParams() []string {
// 	return config.GetStringListParam(`LOVELY_HELM_REPO_ADD_PARAMS`, ``, ` `)
// }

// // HelmRepoAddParams returns extra parameters to pass to helm repo add
// // Set LOVELY_HELM_REPO_ADD_PARAMS to extra parameters to pass to helm template
// func HelmRepoAddParams() []string {
// 	paramsStr := config.GetStringParam(`LOVELY_HELM_REPO_ADD_PARAMS`, ``)
// 	if paramsStr == `` {
// 		return []string{}
// 	}
// 	params := strings.Split(paramsStr, ` `)
// 	for i, param := range params {
// 		params[i] = strings.TrimSpace(param)
// 	}
// 	return params
// }

// // KustomizeMerge returns the yaml to strategic merge into kustomization.yaml
// // Set LOVELY_KUSTOMIZE_MERGE to some yaml you'd like strategic merged on any kustomization.yaml files used by kustomize
// func KustomizeMerge() string {
// 	return config.GetStringParam(`LOVELY_KUSTOMIZE_MERGE`, ``)
// }

// // KustomizePatch returns the yaml to json6902 patch into kustomization.yaml
// // Set LOVELY_KUSTOMIZE_PATCH to some yaml you'd like json6902 patched on any kustomization.yaml files used by kustomize
// func KustomizePatch() string {
// 	return config.GetStringParam(`LOVELY_KUSTOMIZE_PATCH`, ``)
// }

// // AllowGitCheckout establishes if git is safe to use
// // Set ALLOW_GITCHECKOUT to true to say you've told Argo this is safe
// func AllowGitCheckout() bool {
// 	res, err := strconv.ParseBool(config.GetStringParam(`LOVELY_ALLOW_GITCHECKOUT`, `false`))
// 	if err != nil {
// 		return false
// 	}
// 	return res
// }

// // HelmName gives us the application name for helm
// // Set LOVELY_HELM_NAME to override the default of ARGOCD_APP_NAME
// func HelmName() string {
// 	nameOverride := config.GetStringParam(`LOVELY_HELM_NAME`, ``)
// 	if nameOverride == `` {
// 		return os.Getenv(`ARGOCD_APP_NAME`)
// 	}
// 	return nameOverride
// }

// // HelmNamespace gives us the namespace for helm
// func HelmNamespace() string {
// 	return os.Getenv(`ARGOCD_APP_NAMESPACE`)
// }

// // HelmValues gives us the values file we're going to use for helm
// // Set LOVELY_HELM_VALUES to the path to use for the values file
// func HelmValues() string {
// 	return config.GetStringParam(`LOVELY_HELM_VALUES`, `values.yaml`)
// }

// // HelmValues gives us the values files we're going to use for helm
// // Set LOVELY_HELM_VALUES to the path to use for the values file
// func HelmValues() []string {
// 	return config.GetStringListParam(`LOVELY_HELM_VALUES`, `values.yaml`, ` `)
// }

// // HelmValuesSet returns true if HelmValues() is explicily set
// func HelmValuesSet() bool {
// 	return len(config.GetStringListParam(`LOVELY_HELM_VALUES`, ``, ` `)) != 0
// }

// // HelmValuesSet returns true if HelmValues() is explicily set
// func HelmValuesSet() bool {
// 	return config.GetStringParam(`LOVELY_HELM_VALUES`, ``) != ``
// }

// // HelmfileBinary returns the path to helm if overridden, otherwise we search the path
// // Set LOVELY_HELMFILE_PATH to the path to the helm binary
// func HelmfileBinary() string {
// 	return config.GetStringParam(`LOVELY_HELMFILE_PATH`, `helmfile`)
// }

// // HelmfileMerge returns the yaml to strategic merge into values.yaml
// // Set LOVELY_HELMFILE_MERGE to some yaml you'd like strategic merged into any helmfile.yaml files used by helmfile
// func HelmfileMerge() string {
// 	return config.GetStringParam(`LOVELY_HELMFILE_MERGE`, ``)
// }

// // HelmfilePatch returns the yaml to json6902 patch into values.yaml
// // Set LOVELY_HELMFILE_PATCH to some yaml you'd like json6902 patched into any helmfile.yaml files used by helmfile
// func HelmfilePatch() string {
// 	return config.GetStringParam(`LOVELY_HELMFILE_PATCH`, ``)
// }
