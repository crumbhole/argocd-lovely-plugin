package processor

import (
	jsonpatch "github.com/evanphx/json-patch"
	"os"
	kyamlyaml "sigs.k8s.io/kustomize/kyaml/yaml"
	"sigs.k8s.io/kustomize/kyaml/yaml/merge2"
	yaml "sigs.k8s.io/yaml"
)

// MergeYaml performs a json patch and a strategic merge, both if needed to some yaml on disk
// Performs the strategic merge FIRST, followed by the json patch.
// This is supposed to be Kubernetes strategic merge and RFC6902 patching
// Accepts these patches as yaml. Will write a file to disk from these if
// the file does not already exist, strategic merge patching an on disk
// file will effectively create it.
func MergeYaml(path string, mergetext string, patchtext string) error {
	if mergetext == `` && patchtext == `` {
		return nil
	}

	basetext, err := os.ReadFile(path)
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
		patchjson, err := yaml.YAMLToJSON([]byte(patchtext))
		if err != nil {
			// May not be yaml. Let decode patch whine.
			patchjson = []byte(patchtext)
		}
		patch, err := jsonpatch.DecodePatch(patchjson)
		if err != nil {
			return err
		}
		jsonmergedtext, err := yaml.YAMLToJSON([]byte(mergedtext))
		if err != nil {
			return err
		}

		modifiedjson, err := patch.Apply(jsonmergedtext)
		if err != nil {
			return err
		}

		patchedtext, err := yaml.JSONToYAML(modifiedjson)
		if err != nil {
			return err
		}
		mergedtext = string(patchedtext)
	}

	// #nosec - G306 - this needs to be readable by argocd
	err = os.WriteFile(path, []byte(mergedtext), 0644)
	return err
}
