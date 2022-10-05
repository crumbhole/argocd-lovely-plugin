package main

// The control of this is via environment variables, as that
// is the way argocd allows you to control plugins
import (
	"bytes"
	"fmt"
	"io"
	"os/exec"
	"path/filepath"
)

type pluginProcessor struct{}

func (pluginProcessor) name() string {
	return "plugin"
}

func getRelPlugins(basePath string, path string) ([]string, error) {
	relPath, err := filepath.Rel(basePath, path)
	if err != nil {
		return nil, fmt.Errorf("Internal relative path error %s", err)
	}
	return Plugins(relPath)
}

func (pluginProcessor) enabled(basePath string, path string) bool {
	plugins, err := getRelPlugins(basePath, path)
	if err != nil {
		return true // Enable for error case so errors get reported
	}
	return len(plugins) > 0
}

func (v pluginProcessor) generate(input *string, basePath string, path string) (*string, error) {
	if !v.enabled(basePath, path) {
		return input, ErrDisabledProcessor
	}
	currentText := *input
	plugins, err := getRelPlugins(basePath, path)
	if err != nil {
		return nil, err
	}
	for _, plugin := range plugins {
		cmd := exec.Command(`bash`, `-c`, plugin)
		cmd.Dir = path
		stdin, err := cmd.StdinPipe()
		var stderr bytes.Buffer
		cmd.Stderr = &stderr
		if err != nil {
			return nil, err
		}
		go func() {
			defer stdin.Close()
			io.WriteString(stdin, currentText)
		}()

		out, err := cmd.Output()
		if err != nil {
			return nil, fmt.Errorf("%s: %s", err, stderr.String())
		}
		currentText = string(out)
	}
	return &currentText, nil
}
