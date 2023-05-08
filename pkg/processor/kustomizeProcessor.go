package processor

import (
	"fmt"
	"os"
	"regexp"

	"gopkg.in/yaml.v3"
	"sigs.k8s.io/kustomize/api/types"
)

const kustomizeIntermediateFilename = `_lovely_resource.yaml`

// KustomizeProcessor handles kustomization.yaml via kustomize
type KustomizeProcessor struct{}

// Name returns a string for the plugin's name
func (KustomizeProcessor) Name() string {
	return "kustomize"
}

// Enabled returns true only if this proessor can do work
func (KustomizeProcessor) Enabled(_ string, path string) bool {
	return reEntryInDir(path, regexp.MustCompile(`^kustomization\.ya?ml$`))
}

// Generate create the text stream for this plugin
func (k KustomizeProcessor) Generate(input *string, basePath string, path string) (*string, error) {
	if !k.Enabled(basePath, path) {
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
		kustContents, err := os.ReadFile(kustYamlPath)
		if err != nil {
			return nil, err
		}
		var kust types.Kustomization
		err = yaml.Unmarshal(kustContents, &kust)
		if err != nil {
			return nil, fmt.Errorf("error reading kustomization.yaml: %v", err)
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
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	out, err := execute(wd, KustomizeBinary(), params...)
	if err != nil {
		return nil, fmt.Errorf("error running kustomize: %v", err)
	}
	outstr := "---\n" + string(out)
	return &outstr, nil
}
