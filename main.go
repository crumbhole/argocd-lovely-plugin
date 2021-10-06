package main

import (
	"errors"
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
			log.Printf("Processing %s with %s\n", path, processor.name())
			out, err := processor.process(result, path)
			if err != nil {
				return "", err
			}
			//			print(*out)
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
			log.Printf("Init %s with %s\n", path, processor.name())
			err := processor.init(path)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (c *Collection) doAllDirs(init bool, path string) (string, error) {
	err := c.scanDir(path)
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
	print(output)
}
