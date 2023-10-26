package processor

import (
	"fmt"
	"os"
	"regexp"

	"github.com/crumbhole/argocd-lovely-plugin/pkg/features"
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

func (k KustomizeProcessor) getPath(path string) (string, error) {
	kustYamlPath := path + "/kustomization.yaml"
	if _, err := os.Stat(kustYamlPath); os.IsNotExist(err) {
		kustYamlPath = path + "/kustomization.yml"
		if _, err := os.Stat(kustYamlPath); os.IsNotExist(err) {
			return "", err
		}
	}
	return kustYamlPath, nil
}

// Generate create the text stream for this plugin
func (k KustomizeProcessor) Generate(input *string, basePath string, path string) (*string, error) {
	if !k.Enabled(basePath, path) {
		return input, ErrDisabledProcessor
	}
	kustYamlPath, err := k.getPath(path)
	if err != nil {
		return nil, err
	}
	err = MergeYaml(kustYamlPath, features.GetKustomizeMerge(), features.GetKustomizePatch())
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
			return nil, fmt.Errorf("error reading kustomization.yaml: %w", err)
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
	extraParams, err := features.GetKustomizeParams()
	if err != nil {
		return nil, err
	}
	params = append(params, extraParams...)
	params = append(params, path)
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	out, err := execute(wd, features.GetKustomizePath(), params...)
	if err != nil {
		return nil, fmt.Errorf("error running kustomize: %w", err)
	}
	outstr := "---\n" + out
	return &outstr, nil
}
