package utils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormalizePath(t *testing.T) {
	testPath := "localhost:8080/test"
	testCases := []struct {
		tc       string
		path     string
		expected string
	}{
		{
			tc:       "Path is returned the same as its given",
			path:     testPath,
			expected: testPath,
		},
		{
			tc:       "Path is returned without trailing /",
			path:     testPath + "/",
			expected: testPath,
		},
	}

	for _, test := range testCases {
		p := NormalizePath(test.path)
		assert.Equal(t, test.expected, p)
	}
}

func TestMustGetEnvironmentVariable(t *testing.T) {
	testEnvKey := "GOIMGSSETENVVARKEY"
	testEnvValue := "OK"

	testCases := []struct {
		tc          string
		expectPanic bool
	}{
		{
			tc:          "successfully get a set environment variable",
			expectPanic: false,
		},
		{
			tc:          "assert that a panic is encountered when no env has been set",
			expectPanic: true,
		},
	}

	for _, test := range testCases {
		if !test.expectPanic {
			os.Setenv(testEnvKey, testEnvValue)
			assert.Equal(t, testEnvValue, MustGetEnvironmentVariable(testEnvKey))
		} else {
			assert.Panics(t, func() { MustGetEnvironmentVariable(testEnvKey) })
		}
		os.Unsetenv(testEnvKey)
	}
}
