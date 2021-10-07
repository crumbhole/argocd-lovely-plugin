package main

import (
	jsonpatch "github.com/evanphx/json-patch"
	"io/ioutil"
	kyamlyaml "sigs.k8s.io/kustomize/kyaml/yaml"
	"sigs.k8s.io/kustomize/kyaml/yaml/merge2"
	yaml "sigs.k8s.io/yaml"
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

	basetext, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	var mergedtext string

	if mergetext != `` {
		// Could allow access to infer and merge options. Need use cases.
		mergedtext, err = merge2.MergeStrings(mergetext, string(basetext), false, kyamlyaml.MergeOptions{
			ListIncreaseDirection: kyamlyaml.MergeOptionsListAppend,
		})
		if err != nil {
			return err
		}
	} else {
		mergedtext = string(basetext)
	}

	if patchtext != `` {
		patch, err := jsonpatch.DecodePatch([]byte(patchtext))
		if err != nil {
			return err
		}
		jsontext, err := yaml.YAMLToJSON([]byte(mergedtext))
		if err != nil {
			return err
		}

		modifiedjson, err := patch.Apply(jsontext)
		if err != nil {
			return err
		}

		patchedtext, err := yaml.JSONToYAML(modifiedjson)
		if err != nil {
			return err
		}
		mergedtext = string(patchedtext)
	}

	if err := ioutil.WriteFile(path, []byte(mergedtext), 0644); err != nil {
		return err
	}
	return nil
}
