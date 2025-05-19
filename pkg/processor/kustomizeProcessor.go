package processor

import (
	"fmt"
	"os"
	"path/filepath"
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
	if err := MergeYaml(kustYamlPath, features.GetKustomizeMerge(), features.GetKustomizePatch()); err != nil {
		return nil, err
	}
	if input != nil {
		if err := k.handleIntermediateFile(input, path, kustYamlPath); err != nil {
			return nil, err
		}
	}
	out, err := k.runKustomize(path)
	if err != nil {
		return nil, err
	}
	outstr := "---\n" + out
	return &outstr, nil
}

// handleIntermediateFile writes the input to an intermediate file and updates kustomization.yaml
func (k KustomizeProcessor) handleIntermediateFile(input *string, path, kustYamlPath string) error {
	// #nosec - G304 we've chosen both parts of this
	intermediateFile, err := os.Create(filepath.Join(path, kustomizeIntermediateFilename))
	if err != nil {
		return err
	}
	if _, err := intermediateFile.WriteString(*input); err != nil {
		_ = intermediateFile.Close()
		return err
	}
	if err := intermediateFile.Close(); err != nil {
		return err
	}
	return k.addIntermediateToKustomization(kustYamlPath)
}

// addIntermediateToKustomization adds the intermediate file to the kustomization.yaml resources
func (k KustomizeProcessor) addIntermediateToKustomization(kustYamlPath string) error {
	// #nosec - G304 we've created this path
	kustContents, err := os.ReadFile(kustYamlPath)
	if err != nil {
		return err
	}
	var kust types.Kustomization
	if err := yaml.Unmarshal(kustContents, &kust); err != nil {
		return fmt.Errorf("error reading kustomization.yaml: %w", err)
	}
	kust.Resources = append(kust.Resources, kustomizeIntermediateFilename)
	// #nosec - G304 we've created this path
	kustomizationFile, err := os.OpenFile(kustYamlPath, os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer func() { _ = kustomizationFile.Close() }()
	kustYaml, err := yaml.Marshal(&kust)
	if err != nil {
		return err
	}
	if _, err := kustomizationFile.Write(kustYaml); err != nil {
		return err
	}
	return nil
}

// runKustomize builds the kustomize output for the given path
func (k KustomizeProcessor) runKustomize(path string) (string, error) {
	params := []string{"build", "--enable-helm"}
	extraParams, err := features.GetKustomizeParams()
	if err != nil {
		return "", err
	}
	params = append(params, extraParams...)
	params = append(params, path)
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	out, err := execute(wd, features.GetKustomizePath(), params...)
	if err != nil {
		return "", fmt.Errorf("error running kustomize: %w", err)
	}
	return out, nil
}
