package utils

import (
	"errors"
	"fmt"
	"os"
)

var (
	errRequiredEnvVarMissing = errors.New("Missing required env variable: %s")
)

// NormalizePath ensures that all paths will be processed at the
// same starting state when the target path is run through this function.
func NormalizePath(path string) string {
	if path[len(path)-1] == '/' {
		path = path[:len(path)-1]
	}
	return path
}

// MustGetEnvironmentVariable returns the environment variable's value or panics if
// it is missing
func MustGetEnvironmentVariable(k string) string {
	v := os.Getenv(k)
	fmt.Println(v)
	if len(v) == 0 {
		panic(errRequiredEnvVarMissing)
	}
	return v
}
