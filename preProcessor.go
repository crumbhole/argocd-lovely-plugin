package main

// The control of this is via environment variables, as that
// is the way argocd allows you to control plugins
import (
	"fmt"
	"os/exec"
	"path/filepath"
)

type preProcessor struct{}

func (preProcessor) name() string {
	return "preprocessor"
}

func getRelPreprocessors(basePath string, path string) ([]string, error) {
	relPath, err := filepath.Rel(basePath, path)
	if err != nil {
		return make([]string, 0), fmt.Errorf("Internal relative path error %s", err)
	}
	return Preprocessors(relPath)
}

func (preProcessor) enabled(basePath string, path string) bool {
	plugins, err := getRelPreprocessors(basePath, path)
	if err != nil {
		return true // Enable for error case so errors get reported
	}
	return len(plugins) > 0
}

func (v preProcessor) generate(basePath string, path string) error {
	if !v.enabled(basePath, path) {
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
