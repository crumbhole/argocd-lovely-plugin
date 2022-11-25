package main

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/otiai10/copy"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

var processors = []Processor{
	helmProcessor{},
	kustomizeProcessor{},
	yamlProcessor{},
	pluginProcessor{},
}

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
		if c.dirs.AddDirectoryIfYaml(path) {
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

func (c *Collection) processOneDir(path string) (string, error) {
	var result *string
	pre := preProcessor{}
	if pre.enabled(c.baseDir, path) {
		err := pre.generate(c.baseDir, path)
		if err != nil {
			return "", err
		}
	}
	for _, processor := range processors {
		if processor.enabled(c.baseDir, path) {
			out, err := processor.generate(result, c.baseDir, path)
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
		return fmt.Errorf("%s: %v", err, stderr.String())
	}
	clean := exec.Command("git", "clean", "-fdx", ".")
	clean.Dir = path
	clean.Stderr = &stderr
	_, err = clean.Output()
	if err != nil {
		return fmt.Errorf("%s: %v", err, stderr.String())
	}
	return nil
}

// Ensure we have a clean working copy
// ArgoCD doesn't guarantee us an unpatched copy when we run
func (c *Collection) ensureClean(path string) (string, func(string) error, error) {
	if AllowGitCheckout() {
		return path, c.gitClean, c.gitClean(path)
	}
	newPath, err := c.makeTmpCopy(path)
	return newPath, os.RemoveAll, err
}

func (c *Collection) doAllDirs(path string) (string, error) {
	workingPath, cleanup, err := c.ensureClean(path)
	if err != nil {
		log.Fatal(err)
	}
	c.baseDir = workingPath
	defer cleanup(workingPath)
	err = c.scanDir(workingPath)
	if err != nil {
		log.Fatal(err)
	}
	output, err := c.processAllDirs()
	if err != nil {
		log.Fatal(err)
	}
	return output, err
}

func parseArgs() (bool, error) {
	if len(os.Args[1:]) == 0 {
		return false, nil
	}
	if len(os.Args[1:]) > 1 {
		return false, errors.New("Too many arguments. Only one optional argument allowed of 'init'")
	}
	if os.Args[1] == `init` {
		return true, nil
	}
	return false, errors.New("Invalid argument. Only one optional argument allowed of 'init'")
}

func main() {
	initMode, err := parseArgs()
	if err != nil {
		log.Fatal(err)
	}
	if initMode {
		return
	}
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
