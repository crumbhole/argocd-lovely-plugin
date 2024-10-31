package processor

import (
	"bytes"
	"fmt"
	"github.com/crumbhole/argocd-lovely-plugin/pkg/features"
	"os"
	"os/exec"
	"regexp"
)

const (
	EnvPrefixArgoCD = "ARGOCD_ENV_"
)

func execute(path string, command string, params ...string) (string, error) {
	cmd := exec.Command(command, params...)
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
	var filtered []string
	argoRegex := regexp.MustCompile(`^` + regexp.QuoteMeta(EnvPrefixArgoCD))
	for _, e := range env {
		if argoRegex.MatchString(e) {
			e = argoRegex.ReplaceAllString(e, "")
		}
		filtered = append(filtered, e)
	}
	return filtered
}
