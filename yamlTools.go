package main

import (
	"errors"
	//yaml2 "gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"sigs.k8s.io/kustomize/kyaml/yaml"
	"sigs.k8s.io/kustomize/kyaml/yaml/merge2"
)

// Performs a json patch and a strategic merge, both if needed to some yaml on disk
// Performs the strategic merge FIRST, followed by the json patch.
// This is supposed to be Kubernetes strategic merge and RFC6902 patching
// Accepts these patches as yaml. Will write a file to disk from these if
// the file does not already exist, strategic merge patching an on disk
// file will effectively create it.

func MergeYaml(path string, mergetext string, patchtext string) error {
	if mergetext == `` && patchtext == `` {
		return nil
	}

	//	var base map[string]interface{}
	var basetext []byte = []byte(``)
	var newtext string

	_, err := os.Stat(path)
	if !errors.Is(err, os.ErrNotExist) {
		basetext, err = ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		// Could allow access to infer and merge options. Need use cases.
		newtext, err = merge2.MergeStrings(mergetext, string(basetext), false, yaml.MergeOptions{
			ListIncreaseDirection: yaml.MergeOptionsListAppend,
		})
		if err != nil {
			return err
		}
	} else {
		newtext = mergetext
	}

	if err := ioutil.WriteFile(path, []byte(newtext), 0644); err != nil {
		return err
	}
	return nil
}
