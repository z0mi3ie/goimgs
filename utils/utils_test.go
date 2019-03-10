package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormalizePath(t *testing.T) {
	testCases := []struct {
		tc       string
		path     string
		expected string
	}{
		{
			tc:       "Path is returned the same as its given",
			path:     "localhost:8080/test",
			expected: "localhost:8080/test",
		},
		{
			tc:       "Path is returned without trailing /",
			path:     "localhost:8080/test/",
			expected: "localhost:8080/test",
		},
	}

	for _, test := range testCases {
		p := NormalizePath(test.path)
		assert.Equal(t, test.expected, p)
	}
}
