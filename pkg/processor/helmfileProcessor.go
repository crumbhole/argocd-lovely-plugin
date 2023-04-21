package processor

import (
	"fmt"
	// "os"
	"regexp"
	// "strings"
)

// HelmfileProcessor handles Chart,yaml files via helm
type HelmfileProcessor struct{}

// Name returns a string for the plugin's name
func (HelmfileProcessor) Name() string {
	return "helmfile"
}

// Enabled returns true only if this proessor can do work
func (HelmfileProcessor) Enabled(_ string, path string) bool {
	return reFileInDir(path, regexp.MustCompile(`^helmfile\.ya?ml$`))
}

func (h HelmfileProcessor) helmfileDo(path string, params ...string) (string, error) {
	// baseParams := [6]string{`--registry-config`, `/tmp/.helm/registry.json`, `--repository-cache`, `/tmp/.helm/cache/repository`, `--repository-config`, `/tmp/.helm/repositories.json`}
	// cmdArray := append(baseParams[:], params[:]...)
	return execute(path, HelmfileBinary(), params...)
}

// Generate create the text stream for this plugin
func (h HelmfileProcessor) Generate(input *string, basePath string, path string) (*string, error) {
	if !h.Enabled(basePath, path) {
		return input, ErrDisabledProcessor
	}
	// err := h.reposEnsure(path)
	// if err != nil {
	// 	return nil, err
	// }
	// _, err = h.helmDo(path, `dependency`, `build`)
	// if err != nil {
	// 	return nil, err
	// }
	err := MergeYaml(path+"/helmfile.yaml", HelmfileMerge(), HelmfilePatch())
	if err != nil {
		return nil, err
	}
	params := []string{`template`, `--include-crds`}
	// params = append(params[:], HelmTemplateParams()[:]...)
	// if kubeVersion := os.Getenv(`KUBE_VERSION`); kubeVersion != "" && !contains(params, `--kube-version`) {
	// 	params = append(params[:], []string{`--kube-version`, kubeVersion}...)
	// }
	// // each API version needs to be added with --api-versions (https://github.com/helm/helm/issues/11485)
	// if apiVersions := strings.Split(os.Getenv(`KUBE_API_VERSIONS`), `,`); apiVersions != nil {
	// 	for _, apiVersion := range apiVersions {
	// 		params = append(params[:], []string{"--api-versions", apiVersion}...)
	// 	}
	// }
	// params = append(params[:], []string{`-n`, HelmNamespace(), HelmName(), `.`}...)
	out, err := h.helmfileDo(path, params...)
	if err != nil {
		return nil, fmt.Errorf("Error running helm: %v", err)
	}
	return &out, err
}
