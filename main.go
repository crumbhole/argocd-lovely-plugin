package main

import (
	"errors"
	"fmt"
	"github.com/otiai10/copy"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

var processors = []Processor{
	helmProcessor{},
	kustomizeProcessor{},
	yamlProcessor{},
	pluginProcessor{},
}

type Collection struct {
	dirs PackageDirectories
}

func (c *Collection) scanFile(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if info.IsDir() {
		if c.dirs.KnownSubDirectory(path) {
			// We don't allow subdirectories of paths with yaml in
			// to be packages in their own right
			return filepath.SkipDir
		}
		return nil
	}
	yamlRegexp := regexp.MustCompile(`\.ya?ml$`)
	dir := filepath.Dir(path)
	if yamlRegexp.MatchString(path) {
		c.dirs.AddDirectory(dir)
	}
	return nil
}

func (c *Collection) scanDir(path string) error {
	return filepath.Walk(path, c.scanFile)
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
	var result *string = nil
	for _, processor := range processors {
		if processor.enabled(path) {
			out, err := processor.process(result, path)
			if err != nil {
				return "", err
			}
			result = out
		}
	}
	return *result, nil
}

func (c *Collection) initAllDirs() error {
	for _, path := range c.dirs.GetPackages() {
		err := c.initOneDir(path)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Collection) initOneDir(path string) error {
	for _, processor := range processors {
		if processor.enabled(path) {
			err := processor.init(path)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// We copy the directory in case we patch some of the files for kustomize or helm
// ArgoCD doesn't guarantee us an unpatched copy when we run
func (c *Collection) makeTmpCopy(path string) (string, error) {
	tmpPath, err := ioutil.TempDir(os.TempDir(), "lovely-plugin-")
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

func (c *Collection) doAllDirs(init bool, path string) (string, error) {
	workingPath := path
	if !init {
		var err error
		workingPath, err = c.makeTmpCopy(path)
		if err != nil {
			log.Fatal(err)
		}
		defer os.RemoveAll(workingPath)
	}
	err := c.scanDir(workingPath)
	if err != nil {
		log.Fatal(err)
	}
	if init {
		err := c.initAllDirs()
		if err != nil {
			log.Fatal(err)
		}
		return ``, err
	} else {
		output, err := c.processAllDirs()
		if err != nil {
			log.Fatal(err)
		}
		return output, err
	}
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
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	c := Collection{}
	output, err := c.doAllDirs(initMode, dir)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(output)
}
