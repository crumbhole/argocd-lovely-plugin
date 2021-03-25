package main

import (
	"log"
	"os"
	"os/exec"
	"regexp"
)

const kustomizeIntermediateFilename = `_lovely_resource.yaml`

type kustomizeProcessor struct{}

func (_ kustomizeProcessor) enabled(path string) bool {
	return reFileInDir(path, regexp.MustCompile(`^kustomization\.ya?ml$`))
}

func (k kustomizeProcessor) init(path string) error {
	if !k.enabled(path) {
		return DisabledProcessorError
	}
	// No preprocessing needed
	return nil
}

func (k kustomizeProcessor) process(input *string, path string) (*string, error) {
	if !k.enabled(path) {
		return input, DisabledProcessorError
	}
	if input == nil {
		// Reading from 'stdin' - so put this on disk for kustomize
		intermediateFile, err := os.Create(path + "/" + kustomizeIntermediateFilename)
		if err != nil {
			return nil, err
		}
		defer intermediateFile.Close()
		if _, err := intermediateFile.WriteString(*input); err != nil {
			return nil, err
		}
		kustomizationFile, err := os.OpenFile(path+"/kustomization.yaml", os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			return nil, err
		}
		defer kustomizationFile.Close()
		if _, err := kustomizationFile.WriteString("resources:"); err != nil {
			return nil, err
		}
		if _, err := kustomizationFile.WriteString("- " + kustomizeIntermediateFilename); err != nil {
			return nil, err
		}
	}
	log.Printf("Kustomize processing %s\n", path)
	out, err := exec.Command(`kustomize`, `build`, path).Output()
	if err != nil {
		return nil, err
	}
	outstr := string(out)
	return &outstr, nil
}
