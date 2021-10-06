package main

import (
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

func MergeYaml(path string, overlaytext string) error {
	if overlaytext == `` {
		return nil
	}

	var base map[string]interface{}
	var basetext []byte = []byte(``)
	_, err := os.Stat(path)
	if !errors.Is(err, os.ErrNotExist) {
		basetext, err = ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		if err := yaml.Unmarshal(basetext, &base); err != nil {
			return err
		}
	} else {
		base = make(map[string]interface{})
	}
	var overlay map[string]interface{}
	if err := yaml.Unmarshal([]byte(overlaytext), &overlay); err != nil {
		return err
	}

	for k, v := range overlay {
		base[k] = v
	}

	newtext, err := yaml.Marshal(base)
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(path, newtext, 0644); err != nil {
		return err
	}
	return nil
}
