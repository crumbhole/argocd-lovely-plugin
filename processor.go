package main

import (
	"errors"
	"os"
	"regexp"
)

// Processor interface is the interface for all processing engines
// name - name of processor for display and selection purposes
// enabled - does this processor believe it has work to do
// generate - call this to invoke main phase on this processor
// Historically we had an init phase, which has been removed
type Processor interface {
	name() string
	enabled(basePath string, path string) bool
	generate(input *string, basePath string, path string) (*string, error)
}

func reFileInDir(path string, re *regexp.Regexp) bool {
	files, err := os.ReadDir(path)
	if err != nil {
		return false
	}

	for _, file := range files {
		if re.Match([]byte(file.Name())) {
			return true
		}
	}
	return false
}

// ErrDisabledProcessor is the error to return when attempting to use a disabled processor
var ErrDisabledProcessor = errors.New("Internal Error: attempt to use disabled processor")
