package config

import (
	"os"
	"testing"
)

func setupAll(t *testing.T) {
	t.Helper()
	t.Setenv(`PARAM_TESTING_XXZ`, `p123`)
	t.Setenv(`ARGOCD_ENV_TESTING_XXZ`, `e667`)
	t.Setenv(`TESTING_XXZ`, `zxd`)
	t.Setenv(`PARAM_LIST`, `abc def`)
}

func TestParam(t *testing.T) {
	setupAll(t)
	res := GetStringParam(`TESTING_XXZ`, `abc`)
	if res != `p123` {
		t.Errorf("PARAM param not returned, got %s, expected p123", res)
	}
}

func TestParamList(t *testing.T) {
	setupAll(t)
	res, err := GetStringListParam(`LIST`, ``, ' ')
	if err != nil {
		t.Errorf("Didn't expect an error in GetStringListParam")
	}
	if len(res) != 2 {
		t.Errorf("LIST is not length 2, is length %d", len(res))
	}
	if res[0] != `abc` {
		t.Errorf("LIST[0] = %s not abc", res[0])
	}
	if res[1] != `def` {
		t.Errorf("LIST[0] = %s not def", res[0])
	}
}

func TestArgoCDEnv(t *testing.T) {
	setupAll(t)
	err := os.Unsetenv(`PARAM_TESTING_XXZ`)
	if err != nil {
		t.Error(err)
	}
	res := GetStringParam(`TESTING_XXZ`, `abc`)
	if res != `e667` {
		t.Errorf("ARGOCD_ENV param not returned, got %s, expected e667", res)
	}
}

func TestEnv(t *testing.T) {
	setupAll(t)
	err := os.Unsetenv(`PARAM_TESTING_XXZ`)
	if err != nil {
		t.Error(err)
	}
	err = os.Unsetenv(`ARGOCD_ENV_TESTING_XXZ`)
	if err != nil {
		t.Error(err)
	}
	res := GetStringParam(`TESTING_XXZ`, `abc`)
	if res != `zxd` {
		t.Errorf("Env param not returned, got %s, expected zxd", res)
	}
}

func TestDefault(t *testing.T) {
	setupAll(t)
	err := os.Unsetenv(`PARAM_TESTING_XXZ`)
	if err != nil {
		t.Error(err)
	}
	err = os.Unsetenv(`ARGOCD_ENV_TESTING_XXZ`)
	if err != nil {
		t.Error(err)
	}
	err = os.Unsetenv(`TESTING_XXZ`)
	if err != nil {
		t.Error(err)
	}
	res := GetStringParam(`TESTING_XXZ`, `abc`)
	if res != `abc` {
		t.Errorf("Default param not returned, got %s, expected abc", res)
	}
}

// Regression tests: empty env vars must be skipped and fall through to the
// next priority level, rather than being treated as a set value.

func TestEmptyParamFallsBackToArgoCDEnv(t *testing.T) {
	t.Setenv(`PARAM_TESTING_XXZ`, ``)
	t.Setenv(`ARGOCD_ENV_TESTING_XXZ`, `e667`)
	t.Setenv(`TESTING_XXZ`, `zxd`)
	res := GetStringParam(`TESTING_XXZ`, `abc`)
	if res != `e667` {
		t.Errorf("Empty PARAM_ should fall through to ARGOCD_ENV_, got %s, expected e667", res)
	}
}

func TestEmptyParamFallsBackToEnv(t *testing.T) {
	t.Setenv(`PARAM_TESTING_XXZ`, ``)
	err := os.Unsetenv(`ARGOCD_ENV_TESTING_XXZ`)
	if err != nil {
		t.Error(err)
	}
	t.Setenv(`TESTING_XXZ`, `zxd`)
	res := GetStringParam(`TESTING_XXZ`, `abc`)
	if res != `zxd` {
		t.Errorf("Empty PARAM_ should fall through to bare env, got %s, expected zxd", res)
	}
}

func TestEmptyArgoCDEnvFallsBackToEnv(t *testing.T) {
	err := os.Unsetenv(`PARAM_TESTING_XXZ`)
	if err != nil {
		t.Error(err)
	}
	t.Setenv(`ARGOCD_ENV_TESTING_XXZ`, ``)
	t.Setenv(`TESTING_XXZ`, `zxd`)
	res := GetStringParam(`TESTING_XXZ`, `abc`)
	if res != `zxd` {
		t.Errorf("Empty ARGOCD_ENV_ should fall through to bare env, got %s, expected zxd", res)
	}
}

func TestAllEmptyFallsBackToDefault(t *testing.T) {
	t.Setenv(`PARAM_TESTING_XXZ`, ``)
	t.Setenv(`ARGOCD_ENV_TESTING_XXZ`, ``)
	t.Setenv(`TESTING_XXZ`, ``)
	res := GetStringParam(`TESTING_XXZ`, `abc`)
	if res != `abc` {
		t.Errorf("All empty vars should return default, got %s, expected abc", res)
	}
}

// ApplicationSet use-case: PARAM_ is always present but rendered as empty string
// by Go templates; the real value lives in ARGOCD_ENV_ on the sidecar container.
func TestApplicationSetEmptyParamUsesArgoCDEnv(t *testing.T) {
	t.Setenv(`PARAM_TESTING_XXZ`, ``)
	t.Setenv(`ARGOCD_ENV_TESTING_XXZ`, `sidecar_val`)
	res := GetStringParam(`TESTING_XXZ`, `abc`)
	if res != `sidecar_val` {
		t.Errorf("ApplicationSet scenario: empty PARAM_ should use ARGOCD_ENV_, got %s, expected sidecar_val", res)
	}
}
