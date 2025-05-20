package processor

import (
	"reflect"
	"testing"
)

func TestFilterEnvironment(t *testing.T) {
	tests := []struct {
		name     string
		env      []string
		expected []string
	}{
		{
			name:     "empty environment",
			env:      []string{},
			expected: []string{},
		},
		{
			name:     "no argocd variables",
			env:      []string{"PATH=/usr/bin", "HOME=/home/user"},
			expected: []string{"PATH=/usr/bin", "HOME=/home/user"},
		},
		{
			name:     "argocd variables only",
			env:      []string{"ARGOCD_ENV_FOO=bar", "ARGOCD_ENV_BAZ=qux"},
			expected: []string{"FOO=bar", "BAZ=qux"},
		},
		{
			name:     "mixed environment",
			env:      []string{"PATH=/usr/bin", "ARGOCD_ENV_FOO=bar", "HOME=/home/user", "ARGOCD_ENV_BAZ=qux"},
			expected: []string{"PATH=/usr/bin", "FOO=bar", "HOME=/home/user", "BAZ=qux"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := filterEnvironment(tt.env)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("filterEnvironment() = %v, want %v", got, tt.expected)
			}
		})
	}
}
