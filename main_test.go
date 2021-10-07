package main

import (
	"errors"
	"fmt"
	"github.com/otiai10/copy"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"testing"
)

const (
	testsPath     = "test/"
	testsPathCopy = "test_copy/"
)

func setupEnv(path string) (map[string]string, error) {
	var envValues map[string]string
	envFile := path + "/env.yaml"
	_, err := os.Stat(envFile)
	if !errors.Is(err, os.ErrNotExist) {
		envText, err := ioutil.ReadFile(envFile)
		if err != nil {
			return envValues, err
		}
		if err := yaml.Unmarshal(envText, &envValues); err != nil {
			return envValues, err
		}
		// Unlink the env.yaml so it never gets involved in making the yamls.
		if err := os.Remove(envFile); err != nil {
			return envValues, err
		}
		for k, v := range envValues {
			os.Setenv(k, v)
		}
	}
	return envValues, nil
}

func cleanupEnv(env map[string]string) {
	for k, _ := range env {
		os.Unsetenv(k)
	}
}

func checkDir(c Collection, path string) error {
	env, err := setupEnv(path)
	defer cleanupEnv(env)
	if err != nil {
		return err
	}
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
		ioutil.WriteFile(path+"/got.txt", []byte(out), 0444)
		return fmt.Errorf("Expected >\n%s\n< and got >\n%s\n<", expected, out)
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
	err := os.RemoveAll(testsPathCopy)
	if err != nil {
		t.Error(err)
	}
	err = copy.Copy(testsPath, testsPathCopy, opt)
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
			t.Run(d.Name(), func(t *testing.T) {
				err := checkDir(c, testsPathCopy+d.Name())
				if err != nil {
					t.Error(err)
				}
			})
		}
	}
}
