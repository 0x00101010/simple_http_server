package fs

import (
	"os"
	"path/filepath"
)

// Exists checks if file exists on specified paths
func Exists(folder, filename string) bool {
	entirePath := filepath.Join(folder, filename)
	if _, err := os.Stat(entirePath); err != nil {
		return false
	}
	return true
}

// AllowRead checks if file has the correct permission to be read
func AllowRead(folder, filename string) bool {
	entirePath := filepath.Join(folder, filename)

	if _, err := os.OpenFile(entirePath, os.O_RDONLY, 0666); err != nil && os.IsPermission(err) {
		return false
	}
	return true
}

// ReadFile reads the specified file and return byte stream
func ReadFile(folder, filename string) {

}
