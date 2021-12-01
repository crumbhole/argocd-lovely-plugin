package main

import (
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"sigs.k8s.io/kustomize/api/types"
)

const kustomizeIntermediateFilename = `_lovely_resource.yaml`

type kustomizeProcessor struct{}

func (kustomizeProcessor) name() string {
	return "kustomize"
}

func (kustomizeProcessor) enabled(path string) bool {
	return reFileInDir(path, regexp.MustCompile(`^kustomization\.ya?ml$`))
}

func (k kustomizeProcessor) init(path string) error {
	if !k.enabled(path) {
		return ErrDisabledProcessor
	}
	// No preprocessing needed
	return nil
}

func (k kustomizeProcessor) process(input *string, path string) (*string, error) {
	if !k.enabled(path) {
		return input, ErrDisabledProcessor
	}
	kustYamlPath := path + "/kustomization.yaml"
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

		kustomizationFile, err := os.OpenFile(path+"/kustomization.yaml", os.O_WRONLY, 0644)
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
	out, err := exec.Command(KustomizeBinary(), `build`, path).CombinedOutput()
	if err != nil {
		return nil, errors.New(string(out))
	}
	outstr := "---\n" + string(out)
	return &outstr, nil
}
