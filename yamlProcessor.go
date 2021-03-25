package main

import (
	"io/ioutil"
	"log"
	"regexp"
)

type yamlProcessor struct{}

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

func (v yamlProcessor) process(input *string, path string) (*string, error) {
	if !v.enabled(path) {
		return input, DisabledProcessorError
	}
	if input == nil {
		yamlRegexp := regexp.MustCompile(`\.ya?ml$`)
		files, err := ioutil.ReadDir(path) // Sorted already, so stable
		if err != nil {
			return nil, err
		}
		output := ""
		for _, f := range files {
			if !f.IsDir() && yamlRegexp.MatchString(f.Name()) {
				log.Printf("YAML processing %s\n", f.Name())
				yamlcontent, err := ioutil.ReadFile(f.Name())
				if err != nil {
					return nil, err
				}
				output += "---\n"
				output += string(yamlcontent)
			}
		}
		return &output, nil
	}
	return input, nil
}
