package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

type yamlProcessor struct {
	output string
}

func (_ yamlProcessor) enabled(_ string) bool {
	// Always enabled, to pick up if nothing else worked
	return true
}

func (y yamlProcessor) init(path string) error {
	if !y.enabled(path) {
		return DisabledProcessorError
	}
	// No preprocessing needed
	return nil
}

func (v *yamlProcessor) scanFile(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if info.IsDir() {
		return nil
	}
	yamlRegexp := regexp.MustCompile(`\.ya?ml$`)
	if yamlRegexp.MatchString(info.Name()) {
		log.Printf("YAML processing %s\n", info.Name())
		yamlcontent, err := ioutil.ReadFile(info.Name())
		if err != nil {
			return err
		}
		v.output += "---\n"
		v.output += string(yamlcontent)
	}
	return nil
}

func (v yamlProcessor) process(input *string, path string) (*string, error) {
	if !v.enabled(path) {
		return input, DisabledProcessorError
	}
	if input == nil {
		v.output = ""
		err := filepath.Walk(path, v.scanFile)
		if err != nil {
			return nil, err
		}
		return &v.output, nil
	}
	return input, nil
}
