package main

import (
	"errors"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"

	"gopkg.in/yaml.v2"
	"sigs.k8s.io/kustomize/api/types"
)

const kustomizeIntermediateFilename = `_lovely_resource.yaml`

type kustomizeProcessor struct{}

func (kustomizeProcessor) name() string {
	return "kustomize"
}

func (kustomizeProcessor) enabled(_ string, path string) bool {
	return reFileInDir(path, regexp.MustCompile(`^kustomization\.ya?ml$`))
}

func (k kustomizeProcessor) generate(input *string, basePath string, path string) (*string, error) {
	if !k.enabled(basePath, path) {
		return input, ErrDisabledProcessor
	}
	kustYamlPath := path + "/kustomization.yaml"
	if _, err := os.Stat(kustYamlPath); os.IsNotExist(err) {
		kustYamlPath = path + "/kustomization.yml"
		if _, err := os.Stat(kustYamlPath); os.IsNotExist(err) {
			return nil, err
		}
	}
	err := MergeYaml(kustYamlPath, KustomizeMerge(), KustomizePatch())
	if err != nil {
		return nil, err
	}
	if input != nil {
		// Reading from 'stdin' - so put this on disk for kustomize
		intermediateFile, err := os.Create(path + "/" + kustomizeIntermediateFilename)
		if err != nil {
			return nil, err
		}
		if _, err := intermediateFile.WriteString(*input); err != nil {
			intermediateFile.Close()
			return nil, err
		}
		intermediateFile.Close()
		kustContents, err := ioutil.ReadFile(kustYamlPath)
		if err != nil {
			return nil, err
		}
		var kust types.Kustomization
		err = yaml.Unmarshal(kustContents, &kust)
		if err != nil {
			return nil, err
		}
		kust.Resources = append(kust.Resources, kustomizeIntermediateFilename)

		kustomizationFile, err := os.OpenFile(kustYamlPath, os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			return nil, err
		}
		defer kustomizationFile.Close()
		kustYaml, err := yaml.Marshal(&kust)
		if err != nil {
			return nil, err
		}
		if _, err := kustomizationFile.Write(kustYaml); err != nil {
			return nil, err
		}
	}
	params := []string{`build`, `--enable-helm`}
	params = append(params[:], KustomizeParams()[:]...)
	params = append(params, path)
	out, err := exec.Command(KustomizeBinary(), params...).CombinedOutput()
	if err != nil {
		return nil, errors.New(string(out))
	}
	outstr := "---\n" + string(out)
	return &outstr, nil
}
