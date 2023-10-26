package processor

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStripQuotes(t *testing.T) {
	assert.Equal(t, `Hello World`, stripQuotes(`Hello World`))
	assert.Equal(t, `Hello World`, stripQuotes(`"Hello World"`))
	assert.Equal(t, `Hello World`, stripQuotes(`"Hello World`))
	assert.Equal(t, `Hello World`, stripQuotes(`Hello World"`))
	assert.Equal(t, `Hello World`, stripQuotes(`'Hello World'`))
	assert.Equal(t, `Hello World`, stripQuotes(`'Hello World"`))
	assert.Equal(t, `Hello World`, stripQuotes(`"Hello World'`))
}
