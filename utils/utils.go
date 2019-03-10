package utils

// NormalizePath ensures that all paths will be processed at the
// same starting state when the target path is run through this function.
func NormalizePath(path string) string {
	if path[len(path)-1] == '/' {
		path = path[:len(path)-1]
	}
	return path
}
