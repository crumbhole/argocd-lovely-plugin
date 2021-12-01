package main

// The control of this is via environment variables, as that
// is the way argocd allows you to control plugins
import (
	"bytes"
	"fmt"
	"io"
	"os/exec"
)

type pluginProcessor struct{}

func (pluginProcessor) name() string {
	return "plugin"
}

func (pluginProcessor) enabled(_ string) bool {
	return len(Plugins()) > 0
}

func (v pluginProcessor) init(path string) error {
	if !v.enabled(path) {
		return ErrDisabledProcessor
	}
	for _, plugin := range PluginsInit() {
		cmd := exec.Command(`bash`, `-c`, plugin)
		cmd.Dir = path
		out, err := cmd.CombinedOutput()
		if err != nil {
			return fmt.Errorf("%s: %s", err, out)
		}
	}
	return nil
}

func (v pluginProcessor) process(input *string, path string) (*string, error) {
	if !v.enabled(path) {
		return input, ErrDisabledProcessor
	}
	currentText := *input
	for _, plugin := range Plugins() {
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
