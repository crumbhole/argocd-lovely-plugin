package processor

// The control of this is via environment variables, as that
// is the way argocd allows you to control plugins
import (
	"fmt"
	"os/exec"
	"path/filepath"
)

// PreProcessor is not a strict processor, as it handles files on disk
type PreProcessor struct{}

// Name returns a string for the plugin's name
func (PreProcessor) Name() string {
	return "preprocessor"
}

func getRelPreprocessors(basePath string, path string) ([]string, error) {
	relPath, err := filepath.Rel(basePath, path)
	if err != nil {
		return make([]string, 0), fmt.Errorf("internal relative path error %s", err)
	}
	return Preprocessors(relPath)
}

// Enabled returns true only if this proessor can do work
func (PreProcessor) Enabled(basePath string, path string) bool {
	plugins, err := getRelPreprocessors(basePath, path)
	if err != nil {
		return true // Enable for error case so errors get reported
	}
	return len(plugins) > 0
}

// Generate create the text stream for this plugin
func (v PreProcessor) Generate(basePath string, path string) error {
	if !v.Enabled(basePath, path) {
		return ErrDisabledProcessor
	}
	plugins, err := getRelPreprocessors(basePath, path)
	if err != nil {
		return err
	}
	for _, plugin := range plugins {
		cmd := exec.Command(`bash`, `-c`, plugin)
		cmd.Dir = path
		out, err := cmd.CombinedOutput()
		if err != nil {
			return fmt.Errorf("%s: %s", err, out)
		}
	}
	return nil
}
