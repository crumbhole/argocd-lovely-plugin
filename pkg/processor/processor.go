// Package processor contains an interface for all processing and implementations
// of the processors which mangle the yaml
package processor

import (
	"errors"
	"os"
	"regexp"
)

// Processor interface is the interface for all processing engines
// name - name of processor for display and selection purposes
// enabled - does this processor believe it has work to do
// generate - call this to invoke main phase on this processor
// Historically we had an init phase, which has been removed
type Processor interface {
	Name() string
	Enabled(basePath string, path string) bool
	Generate(input *string, basePath string, path string) (*string, error)
}

func reEntryInDir(path string, re *regexp.Regexp) string {
	files, err := os.ReadDir(path)
	if err != nil {
		return ""
	}

	for _, file := range files {
		if re.MatchString(file.Name()) {
			return file.Name()
		}
	}
	return ""
}

// ErrDisabledProcessor is the error to return when attempting to use a disabled processor
var ErrDisabledProcessor = errors.New("internal Error: attempt to use disabled processor")
