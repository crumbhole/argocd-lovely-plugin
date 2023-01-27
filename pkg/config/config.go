package config

import (
	"os"
)

const paramPrefix = `PARAM_`
const argoPrefix = `ARGOCD_ENV_`

// GetStringParam returns a string from the first available configuration source it can find
// or the defaultVal if that fails
func GetStringParam(name string, defaultVal string) string {
	result, got := os.LookupEnv(paramPrefix + name)
	if got {
		return result
	}
	result, got = os.LookupEnv(argoPrefix + name)
	if got {
		return result
	}
	result, got = os.LookupEnv(name)
	if got {
		return result
	}
	return defaultVal
}
