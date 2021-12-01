package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
)

type yamlProcessor struct {
	output string
}

func (yamlProcessor) name() string {
	return "yaml"
}

func (yamlProcessor) enabled(_ string) bool {
	// Always enabled, to pick up if nothing else worked
	return true
}

func (y yamlProcessor) init(path string) error {
	if !y.enabled(path) {
		return ErrDisabledProcessor
	}
	// No preprocessing needed
	return nil
}

func (y *yamlProcessor) scanFile(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if info.IsDir() {
		return nil
	}
	yamlRegexp := regexp.MustCompile(`\.ya?ml$`)
	if yamlRegexp.MatchString(info.Name()) {
		yamlcontent, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		y.output += "---\n"
		y.output += string(yamlcontent)
	}
	return nil
}

func (y yamlProcessor) process(input *string, path string) (*string, error) {
	if !y.enabled(path) {
		return input, ErrDisabledProcessor
	}
	if input == nil {
		y.output = ""
		err := filepath.Walk(path, y.scanFile)
		if err != nil {
			return nil, err
		}
		return &y.output, nil
	}
	return input, nil
}
