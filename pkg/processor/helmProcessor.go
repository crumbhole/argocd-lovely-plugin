package processor

import (
	"fmt"
	"github.com/crumbhole/argocd-lovely-plugin/pkg/features"
	"gopkg.in/yaml.v3"
	"net/url"
	"os"
	"regexp"
	"strings"
)

// Dependency is one repository that this chart is dependent upon
type Dependency struct {
	Name       string `yaml:"name"`
	Repository string `yaml:"repository"`
}

// Dependencies is a list of repositories that this chart is dependent upon
type Dependencies struct {
	Dependencies []Dependency `yaml:"dependencies"`
}

// HelmProcessor handles Chart,yaml files via helm
type HelmProcessor struct{}

// Name returns a string for the plugin's name
func (HelmProcessor) Name() string {
	return "helm"
}

// Enabled returns true only if this proessor can do work
func (HelmProcessor) Enabled(_ string, path string) bool {
	return reEntryInDir(path, regexp.MustCompile(`^Chart\.ya?ml$`))
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func (h HelmProcessor) helmDo(path string, params ...string) (string, error) {
	baseParams := [6]string{`--registry-config`, `/tmp/.helm/registry.json`, `--repository-cache`, `/tmp/.helm/cache/repository`, `--repository-config`, `/tmp/.helm/repositories.json`}
	cmdArray := append(baseParams[:], params[:]...)
	return execute(path, features.HelmBinary(), cmdArray...)
}

func downloadableRepo(repourl string) bool {
	parsedURL, err := url.Parse(repourl)
	if err != nil {
		return true // Bad default possibly
	}
	if parsedURL.Scheme == "oci" ||
		parsedURL.Scheme == "file" ||
		parsedURL.Scheme == "" {
		return false
	}
	return true
}

func (h HelmProcessor) repoEnsure(path string, name string, repourl string) error {
	params := []string{`repo`, `add`, `--force-update`}
	params = append(params[:], features.HelmRepoAddParams()[:]...)
	params = append(params[:], []string{name, repourl}...)
	_, err := h.helmDo(path, params...)
	return err
}

var requirementsFiles = [...]string{
	`requirements.yaml`,
	`requirements.yml`,
	`Chart.yaml`,
	`Chart.yml`,
}

func (h HelmProcessor) reposEnsure(path string) error {
	for _, reqsFile := range requirementsFiles {
		yamlcontent, err := os.ReadFile(path + "/" + reqsFile)
		if err != nil {
			continue
		}
		var deps Dependencies
		err = yaml.Unmarshal(yamlcontent, &deps)
		if err != nil {
			return err
		}
		updateRepos := 0
		for _, dep := range deps.Dependencies {
			if downloadableRepo(dep.Repository) {
				err := h.repoEnsure(path, dep.Name, dep.Repository)
				if err != nil {
					return err
				}
				updateRepos++
			}
		}
		if updateRepos > 0 {
			// Add won't cause an update, so we do an update as well.
			// This is a sledgehammer update all as per-repo update isn't in until helm 3.7
			// and argo ships with 3.6
			_, err := h.helmDo(path, `repo`, `update`)
			return err
		}
	}
	return nil
}

// Generate create the text stream for this plugin
func (h HelmProcessor) Generate(input *string, basePath string, path string) (*string, error) {
	if !h.Enabled(basePath, path) {
		return input, ErrDisabledProcessor
	}
	err := h.reposEnsure(path)
	if err != nil {
		return nil, err
	}
	_, err = h.helmDo(path, `dependency`, `build`)
	if err != nil {
		return nil, err
	}
	err = MergeYaml(path+"/"+features.HelmValues()[0], features.HelmMerge(), features.HelmPatch())
	if err != nil {
		return nil, err
	}
	params := []string{`template`, `--include-crds`}
	params = append(params[:], features.HelmTemplateParams()[:]...)
	if kubeVersion := os.Getenv(`KUBE_VERSION`); kubeVersion != "" && !contains(params, `--kube-version`) {
		params = append(params[:], []string{`--kube-version`, kubeVersion}...)
	}
	// each API version needs to be added with --api-versions (https://github.com/helm/helm/issues/11485)
	if apiVersions := strings.Split(os.Getenv(`KUBE_API_VERSIONS`), `,`); apiVersions != nil {
		for _, apiVersion := range apiVersions {
			params = append(params[:], []string{"--api-versions", apiVersion}...)
		}
	}
	if features.HelmValuesSet() {
		for _, valueFile := range features.HelmValues() {
			params = append(params[:], []string{`-f`, valueFile}...)
		}
	}
	params = append(params[:], []string{`-n`, features.HelmNamespace(), features.HelmName(), `.`}...)
	out, err := h.helmDo(path, params...)
	if err != nil {
		return nil, fmt.Errorf("error running helm: %v", err)
	}
	return &out, err
}
