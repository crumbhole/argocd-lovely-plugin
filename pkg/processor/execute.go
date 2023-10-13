package processor

import (
	"bytes"
	"fmt"
	"os/exec"
)

func execute(path string, command string, params ...string) (string, error) {
	cmd := exec.Command(command, params...)
	cmd.Dir = path
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	out, err := cmd.Output()
	//	fmt.Printf("Output from %s %v in %s is %s and err is %s / %s\n", command, params, path, out, err, cmd.Stderr)
	if err != nil {
		return string(out), fmt.Errorf("%w: %v", err, stderr.String())
	}
	return string(out), nil
}
