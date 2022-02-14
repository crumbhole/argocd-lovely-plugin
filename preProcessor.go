package main

// The control of this is via environment variables, as that
// is the way argocd allows you to control plugins
import (
	"fmt"
	"os/exec"
)

type preProcessor struct{}

func (preProcessor) name() string {
	return "preprocessor"
}

func (preProcessor) enabled(_ string) bool {
	return len(Preprocessors()) > 0
}

func (v preProcessor) generate(path string) error {
	if !v.enabled(path) {
		return ErrDisabledProcessor
	}
	for _, plugin := range Preprocessors() {
		cmd := exec.Command(`bash`, `-c`, plugin)
		cmd.Dir = path
		out, err := cmd.CombinedOutput()
		if err != nil {
			return fmt.Errorf("%s: %s", err, out)
		}
	}
	return nil
}
