package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/hexops/gotextdiff"
	"github.com/hexops/gotextdiff/myers"
	"github.com/hexops/gotextdiff/span"
	"github.com/otiai10/copy"
	"gopkg.in/yaml.v3"
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

func setupEnv(t *testing.T, path string) error {
	t.Helper()
	var envValues map[string]string
	envFile := path + "/env.txt"
	_, err := os.Stat(envFile)
	if !errors.Is(err, os.ErrNotExist) {
		// #nosec - G304 test framework only
		envText, err := os.ReadFile(envFile)
		if err != nil {
			return err
		}
		if err := yaml.Unmarshal(envText, &envValues); err != nil {
			return err
		}
		for k, v := range envValues {
			t.Setenv(k, v)
		}
	}
	return nil
}

func matchREExpected(path string, givenValue string) error {
	// #nosec - G304 test framework only
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
	// #nosec - G304 test framework only
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
		err := os.Remove(got)
		if err != nil {
			return err
		}
		// #nosec - G306 - this is just for test logging/helping
		_ = os.WriteFile(got, []byte(givenValue), 0444)
	}
	return err
}

func checkDir(t *testing.T, path string, errorsExpected bool) error {
	t.Helper()
	err := setupEnv(t, path)
	if err != nil {
		return err
	}
	c := Collection{}
	ctx := context.Background()

	out, fullError := c.doAllDirs(ctx, path)
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
				err := checkDir(t, path+d.Name(), errorsExpected)
				if err != nil {
					t.Error(err)
				}
			})
		}
	}
}

// TestDirectories runs Tests as sidecar only
func TestDirectories(t *testing.T) {
	err := os.RemoveAll(copyPath)
	if err != nil {
		t.Error(err)
	}
	opt := copy.Options{
		OnDirExists: func(_ string, _ string) copy.DirExistsAction {
			return copy.Replace
		},
	}
	err = copy.Copy(normalPath, copyPath, opt)
	if err != nil {
		t.Error(err)
	}
	testDirs(t, copyPath, false)
	err = os.RemoveAll(copyPath)
	if err != nil {
		t.Error(err)
	}
}

// TestDirectoriesError runs Error Tests with copy
func TestDirectoriesError(t *testing.T) {
	testDirs(t, errorsPath, true)
}
