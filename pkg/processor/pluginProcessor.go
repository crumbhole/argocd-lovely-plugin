package processor

// The control of this is via environment variables, as that
// is the way argocd allows you to control plugins
import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os/exec"
	"path/filepath"

	"github.com/crumbhole/argocd-lovely-plugin/pkg/features"
)

// PluginProcessor runs post processing plugins on a stream of yaml text
type PluginProcessor struct{}

// Name returns a string for the plugin's name
func (PluginProcessor) Name() string {
	return "plugin"
}

func getRelPlugins(basePath string, path string) ([]string, error) {
	relPath, err := filepath.Rel(basePath, path)
	if err != nil {
		return nil, fmt.Errorf("internal relative path error %w", err)
	}
	return features.GetPlugins(relPath)
}

// Enabled returns true only if this proessor can do work
func (PluginProcessor) Enabled(basePath string, path string) bool {
	plugins, err := getRelPlugins(basePath, path)
	if err != nil {
		return true // Enable for error case so errors get reported
	}
	return len(plugins) > 0
}

// Generate create the text stream for this plugin
func (v PluginProcessor) Generate(ctx context.Context, input *string, basePath string, path string) (*string, error) {
	if !v.Enabled(basePath, path) {
		return input, ErrDisabledProcessor
	}
	currentText := *input
	plugins, err := getRelPlugins(basePath, path)
	if err != nil {
		return nil, err
	}
	for _, plugin := range plugins {
		// #nosec - G204 the whole point is to run a user specified binary here
		cmd := exec.CommandContext(ctx, `bash`, `-c`, plugin)
		cmd.Dir = path
		stdin, err := cmd.StdinPipe()
		var stderr bytes.Buffer
		cmd.Stderr = &stderr
		if err != nil {
			return nil, err
		}
		go func() {
			defer func() {
				_ = stdin.Close()
			}()
			_, _ = io.WriteString(stdin, currentText)
		}()

		out, err := cmd.Output()
		if err != nil {
			return nil, fmt.Errorf("%w: %s", err, stderr.String())
		}
		currentText = string(out)
	}
	return &currentText, nil
}
