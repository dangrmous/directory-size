package osfilesystem

import (
	"os"
	"path/filepath"
)

// FileSystem defines the interface implemented by OSFileSystem to allow mocking.
type FileSystem interface {
	ReadDir(dirname string) ([]os.DirEntry, error)
	Walk(root string, fn filepath.WalkFunc) error
}

// OSFileSystem provides a real implementation of FileSystem using os and filepath packages.
type OSFileSystem struct{}

// ReadDir calls os.ReadDir to read the contents of a directory.
func (OSFileSystem) ReadDir(dirname string) ([]os.DirEntry, error) {
	return os.ReadDir(dirname)
}

// Walk calls filepath.Walk to traverse the directory tree.
func (OSFileSystem) Walk(root string, fn filepath.WalkFunc) error {
	return filepath.Walk(root, fn)
}
