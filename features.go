package main

// The control of this is via environment variables, as that
// is the way argocd allows you to control plugins
import (
	"os"
	"strconv"
)

func vaultEnabled() bool {
	// Set DISABLE_VAULT to anything to disable vault
	disabled, err := strconv.ParseBool(os.Getenv(`LOVELY_DISABLE_VAULT`))
	if err != nil {
		return true
	}
	return !disabled
}
