package main

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"testing"
)

const (
	normalPath = "test/"
	errorsPath = "test_errors/"
)

func setupEnv(path string) (map[string]string, error) {
	var envValues map[string]string
	envFile := path + "/env.txt"
	_, err := os.Stat(envFile)
	if !errors.Is(err, os.ErrNotExist) {
		envText, err := os.ReadFile(envFile)
		if err != nil {
			return envValues, err
		}
		if err := yaml.Unmarshal(envText, &envValues); err != nil {
			return envValues, err
		}
		for k, v := range envValues {
			os.Setenv(k, v)
		}
	}
	return envValues, nil
}

func cleanupEnv(env map[string]string) {
	for k := range env {
		os.Unsetenv(k)
	}
}

func checkDir(path string, errorsExpected bool) error {
	env, err := setupEnv(path)
	defer cleanupEnv(env)
	if err != nil {
		return err
	}
	c := Collection{}
	expected, err := os.ReadFile(path + "/expected.txt")
	if err != nil {
		return err
	}

	out, err := c.doAllDirs(path)
	if errorsExpected {
		// We expect an error and the error
		// should match expected.txt
		if err == nil {
			return fmt.Errorf("Expected an error but didn't get one")
		}
		if err.Error() != string(expected) {
			os.WriteFile(path+"/got.txt", []byte(err.Error()), 0444)
			return fmt.Errorf("Expected error >\n%s\n< and got >\n%s\n<", expected, err.Error())
		}
	} else {
		// We don't expect and error and
		// expected.txt should be the output
		if err != nil {
			return err
		}
		if out != string(expected) {
			os.WriteFile(path+"/got.txt", []byte(out), 0444)
			return fmt.Errorf("Expected >\n%s\n< and got >\n%s\n<", expected, out)
		}
	}
	return nil
}

// Finds directories under ./test and evaluates all the .yaml/.ymls
func testDirs(t *testing.T, path string, errorsExpected bool) {
	os.Setenv(`ARGOCD_APP_NAME`, `test`)
	os.Setenv(`ARGOCD_APP_NAMESPACE`, `testnamespace`)
	dirs, err := os.ReadDir(path)
	if err != nil {
		t.Error(err)
	}

	for _, d := range dirs {
		if d.IsDir() {
			t.Run(d.Name(), func(t *testing.T) {
				err := checkDir(path+d.Name(), errorsExpected)
				if err != nil {
					t.Error(err)
				}
			})
		}
	}
}

// Tests with copy
func TestDirectoriesCopy(t *testing.T) {
	testDirs(t, normalPath, false)
}

// Tests with git checkout/clean
func TestDirectoriesGitCheckout(t *testing.T) {
	os.Setenv(`ARGOCD_ENV_LOVELY_ALLOW_GITCHECKOUT`, `true`)
	testDirs(t, normalPath, false)
	os.Unsetenv(`ARGOCD_ENV_LOVELY_ALLOW_GITCHECKOUT`)
}

// Error Tests with copy
func TestDirectoriesError(t *testing.T) {
	testDirs(t, errorsPath, true)
}
