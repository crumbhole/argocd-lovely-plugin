package main

import (
	"errors"
	"io/ioutil"
	"regexp"
)

// Processor interface is the interface for all processing engines
// name - name of processor for display and selection purposes
// enabled - does this processor believe it has work to do
// init - call this to invoke init phase on this processor
// process - call this to invoke main phase on this processor
type Processor interface {
	name() string
	enabled(path string) bool
	init(path string) error
	process(input *string, path string) (*string, error)
}

func reFileInDir(path string, re *regexp.Regexp) bool {
	files, err := ioutil.ReadDir(path)
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
