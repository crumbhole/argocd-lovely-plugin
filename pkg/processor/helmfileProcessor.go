package processor

import (
	"fmt"
	"github.com/crumbhole/argocd-lovely-plugin/pkg/features"
	"regexp"
)

// HelmfileProcessor handles Chart,yaml files via helm
type HelmfileProcessor struct{}

// Name returns a string for the plugin's name
func (HelmfileProcessor) Name() string {
	return "helmfile"
}

// Enabled returns true only if this proessor can do work
func (HelmfileProcessor) Enabled(_ string, path string) bool {
	return reEntryInDir(path, regexp.MustCompile(`^helmfile\.ya?ml$`)) ||
		reEntryInDir(path, regexp.MustCompile(`^helmfile\.d$`))
}

func (h HelmfileProcessor) helmfileDo(path string, params ...string) (string, error) {
	return execute(path, features.GetHelmfilePath(), params...)
}

// Generate create the text stream for this plugin
func (h HelmfileProcessor) Generate(input *string, basePath string, path string) (*string, error) {
	if !h.Enabled(basePath, path) {
		return input, ErrDisabledProcessor
	}
	if reEntryInDir(path, regexp.MustCompile(`^helmfile\.ya?ml$`)) {
		err := MergeYaml(path+"/helmfile.yaml", features.GetHelmfileMerge(), features.GetHelmfilePatch())
		if err != nil {
			return nil, err
		}
	}
	params := []string{`template`}
	out, err := h.helmfileDo(path, params...)
	if err != nil {
		return nil, fmt.Errorf("error running helmfile: %v", err)
	}
	return &out, err
}
