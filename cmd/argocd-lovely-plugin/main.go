package main

import (
	"errors"
	"fmt"
	"github.com/crumbhole/argocd-lovely-plugin/pkg/processor"
	"log"
	"os"
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

func (c *Collection) doAllDirs(path string) (string, error) {
	c.baseDir = path
	err := c.scanDir(path)
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
