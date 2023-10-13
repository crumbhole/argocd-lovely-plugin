package main

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/crumbhole/argocd-lovely-plugin/pkg/features"
	"github.com/crumbhole/argocd-lovely-plugin/pkg/processor"
	"github.com/otiai10/copy"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

// Collection is a list of sub-applications making up this application
type Collection struct {
	baseDir string
	dirs    PackageDirectories
}

func (c *Collection) scanFile(path string, info os.DirEntry, err error) error {
	if err != nil {
		return err
	}
	if info.IsDir() {
		if c.dirs.AddDirectoryIfWanted(path) {
			return filepath.SkipDir
		}
		return nil
	}
	return nil
}

func (c *Collection) scanDir(path string) error {
	return filepath.WalkDir(path, c.scanFile)
}

func (c *Collection) processAllDirs() (string, error) {
	result := ""
	for _, path := range c.dirs.GetPackages() {
		output, err := c.processOneDir(path)
		if err != nil {
			return "", err
		}
		result += output
	}
	return result, nil
}

// check exclusive checks for processors which must not both simultaneously be used
// Currently this means helm and helmfile, as they cannot cope with both being used.
func (c *Collection) checkExclusive(path string) error {
	helmfileP := processor.HelmfileProcessor{}
	helmP := processor.HelmProcessor{}
	if helmfileP.Enabled(c.baseDir, path) && helmP.Enabled(c.baseDir, path) {
		return errors.New("helmfile.yaml and Chart.yaml should not both exist in the same directory")
	}
	return nil
}

func (c *Collection) processOneDir(path string) (string, error) {
	var result *string
	pre := processor.PreProcessor{}
	if pre.Enabled(c.baseDir, path) {
		err := pre.Generate(c.baseDir, path)
		if err != nil {
			return "", err
		}
	}
	err := c.checkExclusive(path)
	if err != nil {
		return "", err
	}
	for _, processor := range []processor.Processor{
		processor.HelmfileProcessor{},
		processor.HelmProcessor{},
		processor.KustomizeProcessor{},
		processor.YamlProcessor{},
		processor.PluginProcessor{},
	} {
		if processor.Enabled(c.baseDir, path) {
			out, err := processor.Generate(result, c.baseDir, path)
			if err != nil {
				return "", err
			}
			result = out
		}
	}
	return *result, nil
}

// We copy the directory in case we patch some of the files for kustomize or helm
func (c *Collection) makeTmpCopy(path string) (string, error) {
	tmpPath, err := os.MkdirTemp(os.TempDir(), "lovely-plugin-")
	if err != nil {
		return tmpPath, err
	}
	err = os.RemoveAll(tmpPath)
	if err != nil {
		return tmpPath, err
	}
	err = copy.Copy(path, tmpPath)
	return tmpPath, err
}

func (c *Collection) gitClean(path string) error {
	chkout := exec.Command("git", "checkout", "HEAD", "--", ".")
	chkout.Dir = path
	var stderr bytes.Buffer
	chkout.Stderr = &stderr
	_, err := chkout.Output()
	if err != nil {
		return fmt.Errorf("%w: %v", err, stderr.String())
	}
	clean := exec.Command("git", "clean", "-fdx", ".")
	clean.Dir = path
	clean.Stderr = &stderr
	_, err = clean.Output()
	if err != nil {
		return fmt.Errorf("%w: %v", err, stderr.String())
	}
	return nil
}

func donothing(_ string) error {
	return nil
}

// Ensure we have a clean working copy
// ArgoCD doesn't guarantee us an unpatched copy when we run
// as a configmap plugin. It does when as a sidecar.
func (c *Collection) ensureClean(path string) (string, func(string) error, error) {
	sidecar, got := os.LookupEnv(`LOVELY_SIDECAR`)
	if got && sidecar == "true" {
		return path, donothing, nil
	}
	if features.GetAllowGitCheckout() {
		return path, c.gitClean, c.gitClean(path)
	}
	newPath, err := c.makeTmpCopy(path)
	return newPath, os.RemoveAll, err
}

func (c *Collection) doAllDirs(path string) (string, error) {
	workingPath, cleanup, err := c.ensureClean(path)
	if err != nil {
		return "", err
	}
	c.baseDir = workingPath
	defer func() {
		err := cleanup(workingPath)
		if err != nil {
			fmt.Printf("%s", err)
		}
	}()
	err = c.scanDir(workingPath)
	if err != nil {
		return "", err
	}
	output, err := c.processAllDirs()
	if err != nil {
		return "", err
	}
	return output, nil
}

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	c := Collection{}
	output, err := c.doAllDirs(dir)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(output)
}
