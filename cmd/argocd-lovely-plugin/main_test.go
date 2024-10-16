package main

import (
	"errors"
	"fmt"
	"github.com/hexops/gotextdiff"
	"github.com/hexops/gotextdiff/myers"
	"github.com/hexops/gotextdiff/span"
	"github.com/otiai10/copy"
	"gopkg.in/yaml.v3"
	"os"
	"regexp"
	"testing"
)

func prettyDiff(expected, got string) string {
	edits := myers.ComputeEdits(span.URIFromPath("expected.txt"), expected, got)
	return fmt.Sprint(gotextdiff.ToUnified("expected.txt", "got.txt", expected, edits))
}

const (
	normalPath = "test/"
	copyPath   = "test_copy/"
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

func matchREExpected(path string, givenValue string) error {
	expected, err := os.ReadFile(path + "/regexp.txt")
	var pathErr *os.PathError
	if errors.As(err, &pathErr) {
		return fmt.Errorf("Couldn't find expected.txt nor regexp.txt for test")
	}
	if err != nil {
		return err
	}
	expectre := regexp.MustCompile(string(expected))
	if expectre.MatchString(givenValue) {
		return nil
	}
	return fmt.Errorf("Expected regex >\n%s\n< and got >\n%s\n<", expected, givenValue)
}

func matchExpected(path string, givenValue string) error {
	expected, err := os.ReadFile(path + "/expected.txt")
	var pathErr *os.PathError
	if errors.As(err, &pathErr) {
		return matchREExpected(path, givenValue)
	}
	if err != nil {
		return err
	}
	if string(expected) == givenValue {
		return nil
	}
	return fmt.Errorf("%s", prettyDiff(string(expected), givenValue))
}

func matchExpectedWithStore(path string, givenValue string) error {
	err := matchExpected(path, givenValue)
	if err != nil {
		got := path + "/got.txt"
		os.Remove(got)
		// #nosec - G306 - this is just for test logging/helping
		_ = os.WriteFile(got, []byte(givenValue), 0444)
	}
	return err
}

func checkDir(path string, errorsExpected bool) error {
	env, err := setupEnv(path)
	defer cleanupEnv(env)
	if err != nil {
		return err
	}
	c := Collection{}

	out, fullError := c.doAllDirs(path)
	if errorsExpected {
		// We expect an error and the error
		// should match expected.txt
		if fullError == nil {
			return fmt.Errorf("Expected an error but didn't get one")
		}
		return matchExpectedWithStore(path, fullError.Error())
	}
	// We don't expect and error and
	// expected.txt should be the output
	if fullError != nil {
		return fullError
	}
	return matchExpectedWithStore(path, out)
}

// Finds directories under ./test and evaluates all the .yaml/.ymls
func testDirs(t *testing.T, path string, errorsExpected bool) {
	t.Helper()
	t.Setenv(`ARGOCD_APP_NAME`, `test`)
	t.Setenv(`ARGOCD_APP_NAMESPACE`, `testnamespace`)
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

// TestDirectories runs Tests as sidecar only
func TestDirectories(t *testing.T) {
	os.RemoveAll(copyPath)
	opt := copy.Options{
		OnDirExists: func(_ string, _ string) copy.DirExistsAction {
			return copy.Replace
		},
	}
	err := copy.Copy(normalPath, copyPath, opt)
	if err != nil {
		t.Error(err)
	}
	testDirs(t, copyPath, false)
	os.RemoveAll(copyPath)
}

// TestDirectoriesError runs Error Tests with copy
func TestDirectoriesError(t *testing.T) {
	testDirs(t, errorsPath, true)
}
