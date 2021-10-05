package main

import (
	"bytes"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"regexp"
)

type Dependency struct {
	Name       string `yaml:"name"`
	Repository string `yaml:"repository"`
}

type Dependencies struct {
	Dependencies []Dependency `yaml:"dependencies"`
}

type helmProcessor struct{}

func (_ helmProcessor) name() string {
	return "helm"
}

func (_ helmProcessor) enabled(path string) bool {
	return reFileInDir(path, regexp.MustCompile(`^Chart\.ya?ml$`))
}

func (h helmProcessor) addRepo(name string, url string) error {
	log.Printf("Helm repo add %s %s\n", name, url)
	cmd := exec.Command(`helm`, `repo`, `add`, name, url)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	_, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("%s: %v", err, stderr)
	}
	return nil
}

var requirementsFiles = [...]string{
	`requirements.yaml`,
	`requirements.yml`,
	`Chart.yaml`,
	`Chart.yml`,
}

func (h helmProcessor) addRepos(path string) error {
	for _, reqsFile := range requirementsFiles {
		yamlcontent, err := ioutil.ReadFile(path + "/" + reqsFile)
		if err != nil {
			continue
		}
		log.Printf("Requirements %s, %s\n", reqsFile, yamlcontent)
		var deps Dependencies
		err = yaml.Unmarshal(yamlcontent, &deps)
		for _, dep := range deps.Dependencies {
			h.addRepo(dep.Name, dep.Repository)
		}
	}
	return nil
}

func (h helmProcessor) init(path string) error {
	if !h.enabled(path) {
		return DisabledProcessorError
	}
	//	cmd := exec.Command(`helm`, `repo`, `add` name, url)
	// cmd.Dir := path
	h.addRepos(path)
	log.Printf("Helm init %s\n", path)
	cmd := exec.Command(`helm`, `dependency`, `build`)
	cmd.Dir = path
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	_, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("%s: %v", err, stderr)
	}
	return nil
}

func (h helmProcessor) process(input *string, path string) (*string, error) {
	if !h.enabled(path) {
		return input, DisabledProcessorError
	}
	log.Printf("Helm processing %s\n", path)
	cmd := exec.Command(`helm`,
		`template`,
		`-n`,
		os.Getenv(`ARGOCD_APP_NAMESPACE`),
		os.Getenv(`ARGOCD_APP_NAME`),
		`.`)
	cmd.Dir = path
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	out, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("%s: %v", err, stderr)
	}
	outstr := string(out)
	return &outstr, nil
}
