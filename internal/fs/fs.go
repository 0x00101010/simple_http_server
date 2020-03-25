package fs

import (
	"os"
)

// Exists checks if file exists on specified paths
func Exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		return false
	}
	return true
}

// AllowRead checks if file has the correct permission to be read
func AllowRead(path string) bool {
	if _, err := os.OpenFile(path, os.O_RDONLY, 0666); err != nil && os.IsPermission(err) {
		return false
	}
	return true
}
