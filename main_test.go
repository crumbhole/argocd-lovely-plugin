package main

import (
	"fmt"
	"github.com/otiai10/copy"
	"io/ioutil"
	"os"
	"testing"
)

const (
	testsPath     = "test/"
	testsPathCopy = "test_copy/"
)

func checkDir(c Collection, path string) error {
	out, err := c.doAllDirs(true, path)
	if err != nil {
		return err
	}
	out, err = c.doAllDirs(false, path)
	if err != nil {
		return err
	}
	expected, err := ioutil.ReadFile(path + "/expected.txt")
	if err != nil {
		return err
	}
	if out != string(expected) {
		return fmt.Errorf("Expected %s and got %s", expected, out)
	}
	return nil
}

// Finds directories under ./test and evaluates all the .yaml/.ymls
func TestDirectories(t *testing.T) {
	os.Setenv(`ARGOCD_APP_NAME`, `test`)
	os.Setenv(`ARGOCD_APP_NAMESPACE`, `testnamespace`)
	opt := copy.Options{
		OnDirExists: func(_ string, _ string) copy.DirExistsAction {
			return copy.Replace
		},
	}
	err := copy.Copy(testsPath, testsPathCopy, opt)
	if err != nil {
		t.Error(err)
	}
	dirs, err := ioutil.ReadDir(testsPathCopy)
	if err != nil {
		t.Error(err)
	}
	c := Collection{}

	for _, d := range dirs {
		if d.IsDir() {
			t.Logf("Testing dir %s", testsPathCopy+d.Name())
			err := checkDir(c, testsPathCopy+d.Name())
			if err != nil {
				t.Error(err)
			}
		}
	}
}
