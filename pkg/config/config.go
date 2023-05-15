package config

import (
	"os"
	"strings"
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

// GetStringListParam returns a string array from the first available configuration source it can find
// or the default value listified if that fails. The separator is used
// to separate list items
func GetStringListParam(name string, defaultVal string, separator string) []string {
	paramsStr := GetStringParam(name, defaultVal)
	if paramsStr == `` {
		return []string{}
	}
	params := strings.Split(paramsStr, separator)
	for i, param := range params {
		params[i] = strings.TrimSpace(param)
	}
	return params
}
