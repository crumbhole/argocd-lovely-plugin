package config

import (
	"os"
	"testing"
)

func setupAll() {
	os.Setenv(`PARAM_TESTING_XXZ`, `p123`)
	os.Setenv(`ARGOCD_ENV_TESTING_XXZ`, `e667`)
	os.Setenv(`TESTING_XXZ`, `zxd`)
	os.Setenv(`PARAM_LIST`, `abc def`)
}

func TestParam(t *testing.T) {
	setupAll()
	res := GetStringParam(`TESTING_XXZ`, `abc`)
	if res != `p123` {
		t.Errorf("PARAM param not returned, got %s, expected p123", res)
	}
}

func TestParamList(t *testing.T) {
	setupAll()
	res := GetStringListParam(`LIST`, ``, ` `)
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
	setupAll()
	os.Unsetenv(`PARAM_TESTING_XXZ`)
	res := GetStringParam(`TESTING_XXZ`, `abc`)
	if res != `e667` {
		t.Errorf("ARGOCD_ENV param not returned, got %s, expected e667", res)
	}
}

func TestEnv(t *testing.T) {
	setupAll()
	os.Unsetenv(`PARAM_TESTING_XXZ`)
	os.Unsetenv(`ARGOCD_ENV_TESTING_XXZ`)
	res := GetStringParam(`TESTING_XXZ`, `abc`)
	if res != `zxd` {
		t.Errorf("Env param not returned, got %s, expected zxd", res)
	}
}

func TestDefault(t *testing.T) {
	setupAll()
	os.Unsetenv(`PARAM_TESTING_XXZ`)
	os.Unsetenv(`ARGOCD_ENV_TESTING_XXZ`)
	os.Unsetenv(`TESTING_XXZ`)
	res := GetStringParam(`TESTING_XXZ`, `abc`)
	if res != `abc` {
		t.Errorf("Default param not returned, got %s, expected abc", res)
	}
}
