package features

// The control of this is via environment variables, as that
// is the way argocd allows you to control plugins
import (
	"strings"
)

type featureId int

//type CollectionType int

// All the features/parameters this module supports
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
	FirstFeature = Plugins
	LastFeature  = HelmfilePatch
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

// Feature is an individual configurable element of the plugin
type Feature struct {
	// Name is the name of the parameter or environment variable, in lower_snake_case
	Name string
	// Title is plain text user readable name for the parameter
	Title string
	// DefaultVal is the default value for this parameter under most circumstances
	DefaultVal string
	// Description is in markdown, but will be rendered as text to plugin.yaml
	Description string
	//	CollectionType CollectionType
}

// EnvName gives the environment variable name to use for a feature
func (f Feature) EnvName() string {
	return strings.ToUpper(f.Name)
}

// Features are all the actual parameters supported by the plugin
var Features = map[featureId]Feature{
	Plugins: {
		Title:       `Plugins`,
		Name:        `lovely_plugins`,
		Description: "Set to a comma separated list of binaries to run as plugins. Read [the documentation](plugins.md) for more on plugins. Will not be used if `Plugins YAML` is set.",
	},
	PluginsYaml: {
		Title:       `Plugins YAML`,
		Name:        `lovely_plugins_yaml`,
		Description: "Set to some YAML describing the binaries to run as plugins. Read [the documentation](plugins.md) for more on plugins. Will override `Plugins` if set.",
	},
	Preprocessors: {
		Title:       `PreProcessors`,
		Name:        `lovely_preprocessors`,
		Description: "Set to a comma separated list of binaries to run as preprocessors. Read [the documentation](plugins.md) for more on plugins. Will not be used if `Preproecessors YAML` is set.",
	},
	PreprocessorsYaml: {
		Title:       `PreProcessors YAML`,
		Name:        `lovely_preprocessors_yaml`,
		Description: "Set to some YAML describing the binaries to run as preprocessors. Read [the documentation](plugins.md) for more on plugins. Will override `PreProcessors` if set.",
	},
	DetectRx: {
		Title:       `Detection Regular Expression`,
		Name:        `lovely_detection_regex`,
		DefaultVal:  `\.ya?ml$`,
		Description: "Allow applications to be detected using a different regex so that a PREPROCESSOR that works on non-yaml files can run on this application. The default is `\\.ya?ml$`. (Note: currently `helmfile.d/` will always trigger an application being detected, raise an issue if this needs configuring too). This is pointless to change unless you have a PREPROCESSOR defined.",
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
		Title:       `Helm Merge`,
		Name:        `lovely_helm_merge`,
		Description: "Set to some yaml you'd like [strategic merged](https://kubectl.docs.kubernetes.io/references/kustomize/kustomization/patchesstrategicmerge/) into the values.yaml (or first `Helm Values` specified file) used by Helm."},
	HelmPatch: {
		Title:       `Helm Patch`,
		Name:        `lovely_helm_patch`,
		Description: "Set to some yaml you'd like [json6902](https://kubectl.docs.kubernetes.io/references/kustomize/kustomization/patchesjson6902/) patched into the values.yaml (or first `Helm Values` specified file) used by Helm.",
	},
	HelmTemplateParams: {
		Title:       `Helm Template Parameters`,
		Name:        `lovely_helm_template_params`,
		Description: "Space separated extra parameters to `Helm template` as you might use on the command line. You're on your own here if you pass rubbish parameters.",
	},
	HelmRepoAddParams: {
		Title:       `Helm Repo Add Parameters`,
		Name:        `lovely_helm_repo_add_params`,
		Description: "Space separated extra parameters to `Helm repo add` as you might use on the command line. You're on your own here if you pass rubbish parameters. `--insecure-skip-tls-verify` if your helm chart is on an insecure HTTPS server.",
	},
	KustomizeMerge: {
		Title:       `Kustomize Merge`,
		Name:        `lovely_kustomize_merge`,
		Description: "Set to some yaml you'd like [strategic merged](https://kubectl.docs.kubernetes.io/references/kustomize/kustomization/patchesstrategicmerge/) into any kustomization.yaml found.",
	},
	KustomizePatch: {
		Title:       `Kustomize Patch`,
		Name:        `lovely_kustomize_patch`,
		Description: "Set to some yaml or json you'd like [json6902](https://kubectl.docs.kubernetes.io/references/kustomize/kustomization/patchesjson6902/) patched into any kustomization.yaml found.",
	},
	AllowGitCheckout: {
		Title:       `Allow Git Checkout`,
		Name:        `lovely_allow_gitcheckout`,
		DefaultVal:  `false`,
		Description: `This is not necessary when using the plugin as a sidecar. Allows kustomize base paths to work. Do **not** just set this without reading [the documentation](allow_git.md)`,
	},
	HelmName: {
		Title:       `Helm Name`,
		Name:        `lovely_helm_name`,
		Description: "This can be used to set the Helm 'name' in the same way as releaseName works in Argo CD's standard Helm processing. (`ARGOCD_APP_NAME` used to be overridable in old versions of ArgoCD, but is no longer). Will default to ARGOCD_APP_NAME from the application.",
	},
	HelmNamespace: {
		Title:       `Helm Namespace`,
		Name:        `lovely_helm_namespace`,
		Description: " This can be used to set the Helm 'namespace' it will apply. Will default to ARGOCD_APP_NAMESPACE from the application.",
	},
	HelmValues: {
		Title:       `Helm Values`,
		Name:        `lovely_helm_values`,
		Description: "This is a space separated list values files you'd like to use when rendering the helm chart. Defaults to `values.yaml` if that exists, but its fine if it doesn't. If you override this the file *must* exist. MERGE and PATCH will be applied to the first file in this list.",
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
		Description: "Set to some yaml you'd like [strategic merged](https://kubectl.docs.kubernetes.io/references/kustomize/kustomization/patchesstrategicmerge/) into any helmfile.yaml used by helmfile.",
	},
	HelmfilePatch: {
		Title:       `Helmfile Patch`,
		Name:        `lovely_helmfile_patch`,
		Description: "to some yaml or json you'd like [json6902](https://kubectl.docs.kubernetes.io/references/kustomize/kustomization/patchesjson6902/) patched into any helmfile.yaml used by Helmfile.",
	},
}
