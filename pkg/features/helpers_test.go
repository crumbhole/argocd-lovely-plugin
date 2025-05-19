package features

import (
	"os"
	"testing"
)

func TestGetYamlPlugins(t *testing.T) {
	tests := []struct {
		name     string
		envValue string
		want     pluginYaml
		wantErr  bool
	}{
		{
			name:     "empty environment variable",
			envValue: "",
			want:     make(pluginYaml),
			wantErr:  false,
		},
		{
			name:     "valid yaml with single path",
			envValue: "path/to/app:\n  - plugin1\n  - plugin2",
			want: pluginYaml{
				"path/to/app": {"plugin1", "plugin2"},
			},
			wantErr: false,
		},
		{
			name:     "valid yaml with multiple paths",
			envValue: "path1:\n  - plugin1\npath2:\n  - plugin2\n  - plugin3",
			want: pluginYaml{
				"path1": {"plugin1"},
				"path2": {"plugin2", "plugin3"},
			},
			wantErr: false,
		},
		{
			name:     "invalid yaml",
			envValue: "this is not yaml",
			want:     make(pluginYaml),
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			envName := "TEST_YAML_PLUGINS"
			os.Setenv(envName, tt.envValue)
			defer os.Unsetenv(envName)

			got, err := getYamlPlugins(envName)
			if (err != nil) != tt.wantErr {
				t.Errorf("getYamlPlugins() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				if len(got) != len(tt.want) {
					t.Errorf("getYamlPlugins() got length = %v, want length %v", len(got), len(tt.want))
					return
				}

				for path, plugins := range tt.want {
					gotPlugins, exists := got[path]
					if !exists {
						t.Errorf("getYamlPlugins() missing path %v", path)
						continue
					}

					if len(gotPlugins) != len(plugins) {
						t.Errorf("getYamlPlugins() got %v plugins for path %v, want %v", len(gotPlugins), path, len(plugins))
						continue
					}

					for i, plugin := range plugins {
						if gotPlugins[i] != plugin {
							t.Errorf("getYamlPlugins() got plugin[%d] = %v, want %v", i, gotPlugins[i], plugin)
						}
					}
				}
			}
		})
	}
}

func TestPluginsForPath(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		yamlEnv  string
		plainEnv string
		want     []string
		wantErr  bool
	}{
		{
			name:     "path found in yaml",
			path:     "path/to/app",
			yamlEnv:  "path/to/app:\n  - plugin1\n  - plugin2",
			plainEnv: "plugin3,plugin4",
			want:     []string{"plugin1", "plugin2"},
			wantErr:  false,
		},
		{
			name:     "path not in yaml, using plain env",
			path:     "other/path",
			yamlEnv:  "path/to/app:\n  - plugin1\n  - plugin2",
			plainEnv: "plugin3,plugin4",
			want:     []string{"plugin3", "plugin4"},
			wantErr:  false,
		},
		{
			name:     "empty yaml, using plain env",
			path:     "any/path",
			yamlEnv:  "",
			plainEnv: "plugin1,plugin2",
			want:     []string{"plugin1", "plugin2"},
			wantErr:  false,
		},
		{
			name:     "invalid yaml",
			path:     "any/path",
			yamlEnv:  "this is not yaml",
			plainEnv: "plugin1,plugin2",
			want:     nil,
			wantErr:  true,
		},
		{
			name:     "both envs empty",
			path:     "any/path",
			yamlEnv:  "",
			plainEnv: "",
			want:     []string{},
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			yamlEnvName := "TEST_YAML_PLUGINS"
			plainEnvName := "TEST_PLAIN_PLUGINS"

			os.Setenv(yamlEnvName, tt.yamlEnv)
			os.Setenv(plainEnvName, tt.plainEnv)
			defer os.Unsetenv(yamlEnvName)
			defer os.Unsetenv(plainEnvName)

			got, err := pluginsForPath(tt.path, yamlEnvName, plainEnvName)
			if (err != nil) != tt.wantErr {
				t.Errorf("pluginsForPath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				if len(got) != len(tt.want) {
					t.Errorf("pluginsForPath() got length = %v, want length %v", len(got), len(tt.want))
					return
				}

				for i, plugin := range tt.want {
					if got[i] != plugin {
						t.Errorf("pluginsForPath() got plugin[%d] = %v, want %v", i, got[i], plugin)
					}
				}
			}
		})
	}
}
