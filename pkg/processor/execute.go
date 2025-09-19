package processor

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"regexp"

	"github.com/crumbhole/argocd-lovely-plugin/pkg/features"
)

const (
	envPrefixArgoCD = "ARGOCD_ENV_"
)

func execute(ctx context.Context, path string, command string, params ...string) (string, error) {
	cmd := exec.CommandContext(ctx, command, params...)
	cmd.Dir = path

	if features.GetEnvPropagation() {
		cmd.Env = filterEnvironment(os.Environ())
	}

	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	out, err := cmd.Output()
	//	fmt.Printf("Output from %s %v in %s is %s and err is %s / %s\n", command, params, path, out, err, cmd.Stderr)
	if err != nil {
		return string(out), fmt.Errorf("%w: %v", err, stderr.String())
	}
	return string(out), nil
}

func filterEnvironment(env []string) []string {
	filtered := make([]string, 0, len(env))
	argoRegex := regexp.MustCompile(`^` + regexp.QuoteMeta(envPrefixArgoCD))
	for _, e := range env {
		if argoRegex.MatchString(e) {
			e = argoRegex.ReplaceAllString(e, "")
		}
		filtered = append(filtered, e)
	}
	return filtered
}
