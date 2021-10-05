package main

// The control of this is via environment variables, as that
// is the way argocd allows you to control plugins
import (
	"log"
	"os/exec"
)

const vaultBin = `argo-vault-replacer`

type vaultProcessor struct{}

func (_ vaultProcessor) enabled(_ string) bool {
	// Set DISABLE_VAULT to anything to disable vault
	if _, err := exec.LookPath(vaultBin); err == nil {
		return VaultEnabled()
	}
	return false
}

func (v vaultProcessor) init(path string) error {
	if !v.enabled(path) {
		return DisabledProcessorError
	}
	// No preprocessing needed
	return nil
}

func (v vaultProcessor) process(input *string, path string) (*string, error) {
	if !v.enabled(path) {
		return input, DisabledProcessorError
	}
	log.Printf("Vault processing %s\n", path)
	cmd := exec.Command(vaultBin)
	cmd.Dir = path
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	outstr := string(out)
	return &outstr, nil
}
