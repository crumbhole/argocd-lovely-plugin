package main

import (
	"errors"
	"io/ioutil"
	"regexp"
)

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

var DisabledProcessorError = errors.New("Internal Error: attempt to use disabled processor")
