package main

import (
	"bytes"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
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

type helmProcessor struct{}

func (helmProcessor) name() string {
	return "helm"
}

func (helmProcessor) enabled(path string) bool {
	return reFileInDir(path, regexp.MustCompile(`^Chart\.ya?ml$`))
}

func (h helmProcessor) helmDo(path string, params ...string) (string, error) {
	baseParams := [6]string{`--registry-config`, `/tmp/.helm/registry.json`, `--repository-cache`, `/tmp/.helm/cache/repository`, `--repository-config`, `/tmp/.helm/repositories.json`}
	cmdArray := append(baseParams[:], params[:]...)
	cmd := exec.Command(HelmBinary(), cmdArray...)
	cmd.Dir = path
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	out, err := cmd.Output()
	// fmt.Printf("Output from %v in %s is %s\n", params, path, out)
	if err != nil {
		return string(out), fmt.Errorf("%s: %v", err, stderr.String())
	}
	return string(out), nil
}

func (h helmProcessor) repoEnsure(path string, name string, url string) error {
	_, err := h.helmDo(path, `repo`, `add`, name, url)
	return err
}

func (h helmProcessor) values() []string {
	return HelmValues()
}

var requirementsFiles = [...]string{
	`requirements.yaml`,
	`requirements.yml`,
	`Chart.yaml`,
	`Chart.yml`,
}

func (h helmProcessor) reposEnsure(path string) error {
	for _, reqsFile := range requirementsFiles {
		yamlcontent, err := ioutil.ReadFile(path + "/" + reqsFile)
		if err != nil {
			continue
		}
		var deps Dependencies
		err = yaml.Unmarshal(yamlcontent, &deps)
		if err != nil {
			return err
		}
		for _, dep := range deps.Dependencies {
			h.repoEnsure(path, dep.Name, dep.Repository)
		}
	}
	// Add won't cause an update, so we do an update as well.
	// This is a sledgehammer update all as per-repo update isn't in until helm 3.7
	// and argo ships with 3.6
	_, err := h.helmDo(path, `repo`, `update`)
	return err
}

func (h helmProcessor) generate(input *string, path string) (*string, error) {
	if !h.enabled(path) {
		return input, ErrDisabledProcessor
	}
	h.reposEnsure(path)
	_, err := h.helmDo(path, `dependency`, `build`)
	if err != nil {
		return nil, err
	}
	err = MergeYaml(path+"/values.yaml", HelmMerge(), HelmPatch())
	if err != nil {
		return nil, err
	}
	params := []string{`template`}
	for _, valuefile := range h.values() {
		params = append(params, `-f`, valuefile)
	}
	params = append(params,
		`-n`,
		os.Getenv(`ARGOCD_APP_NAMESPACE`),
		os.Getenv(`ARGOCD_APP_NAME`),
		`.`)
	fmt.Printf("%v\n", params)
	out, err := h.helmDo(path, params...)
	return &out, err
}
