// Package config contains some helpers for configuration of lovely
package config

import (
	"os"
	"strconv"

	"github.com/go-andiamo/splitter"
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

// GetBoolParam returns a bool, coerced from GetStringParam's value
// or the defaultVal if that fails
func GetBoolParam(name string, defaultVal string) bool {
	boolVal, err := strconv.ParseBool(GetStringParam(name, defaultVal))
	if err != nil {
		defaultBool, _ := strconv.ParseBool(defaultVal)
		return defaultBool
	}
	return boolVal
}

// GetStringListParam returns a string array from the first available configuration source it can find
// or the default value listified if that fails. The separator is used
// to separate list items
func GetStringListParam(name string, defaultVal string, separator rune) ([]string, error) {
	paramsStr := GetStringParam(name, defaultVal)
	if paramsStr == `` {
		return []string{}, nil
	}
	encs := []*splitter.Enclosure{
		splitter.DoubleQuotesBackSlashEscaped, splitter.SingleQuotesBackSlashEscaped,
	}
	s := splitter.MustCreateSplitter(separator, encs...).
		AddDefaultOptions(splitter.TrimSpaces, splitter.NoEmpties)
	return s.Split(paramsStr)
}
