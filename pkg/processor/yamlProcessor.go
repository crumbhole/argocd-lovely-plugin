package processor

import (
	"os"
	"path/filepath"
	"regexp"
)

// YamlProcessor is the fallback processor for gathering plain yaml files
type YamlProcessor struct {
	output string
}

// Name returns a string for the plugin's name
func (YamlProcessor) Name() string {
	return "yaml"
}

// Enabled returns true only if this proessor can do work
func (YamlProcessor) Enabled(_ string, _ string) bool {
	// Always enabled, to pick up if nothing else worked
	return true
}

func (y *YamlProcessor) scanFile(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if info.IsDir() {
		return nil
	}
	yamlRegexp := regexp.MustCompile(`^[^\.].*\.ya?ml$`)
	if yamlRegexp.MatchString(info.Name()) {
		yamlcontent, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		if len(y.output) > 0 && y.output[len(y.output)-1] != '\n' {
			y.output += "\n"
		}
		y.output += "---\n"
		y.output += string(yamlcontent)
	}
	return nil
}

// Generate create the text stream for this plugin
func (y YamlProcessor) Generate(input *string, basePath string, path string) (*string, error) {
	if !y.Enabled(basePath, path) {
		return input, ErrDisabledProcessor
	}
	if input == nil {
		y.output = ""
		err := filepath.Walk(path, y.scanFile)
		if err != nil {
			return nil, err
		}
		return &y.output, nil
	}
	return input, nil
}
