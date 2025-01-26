package directory

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
	"time"
)

// Mock File System Interfaces
type mockFileSystem struct {
	readDirFunc func(string) ([]os.DirEntry, error)
	walkFunc    func(string, filepath.WalkFunc) error
}

func (m mockFileSystem) ReadDir(dirname string) ([]os.DirEntry, error) {
	return m.readDirFunc(dirname)
}

func (m mockFileSystem) Walk(root string, fn filepath.WalkFunc) error {
	return m.walkFunc(root, fn)
}

// Mock Directory Entry Implementation
type mockDirEntry struct {
	name  string
	isDir bool
	size  int64
}

func (f mockDirEntry) Name() string               { return f.name }
func (f mockDirEntry) IsDir() bool                { return f.isDir }
func (f mockDirEntry) Type() os.FileMode          { return 0 }
func (f mockDirEntry) Info() (os.FileInfo, error) { return mockFileInfo{f.name, f.isDir, f.size}, nil }

// Mock File Info Implementation
type mockFileInfo struct {
	name  string
	isDir bool
	size  int64
}

func (f mockFileInfo) Name() string       { return f.name }
func (f mockFileInfo) Size() int64        { return f.size }
func (f mockFileInfo) Mode() os.FileMode  { return 0 }
func (f mockFileInfo) ModTime() time.Time { return time.Time{} }
func (f mockFileInfo) IsDir() bool        { return f.isDir }
func (f mockFileInfo) Sys() interface{}   { return nil }

// Test Scenarios
func TestGetDirectorySize_NonRecursive(t *testing.T) {
	// Mock ReadDir to return files and directories
	mockFS := mockFileSystem{
		readDirFunc: func(dirname string) ([]os.DirEntry, error) {
			return []os.DirEntry{
				mockDirEntry{"file1.txt", false, 50},
				mockDirEntry{"file2.txt", false, 100},
				mockDirEntry{"subdir", true, 0}, // A directory
			}, nil
		},
	}

	size, err := GetDirectorySize(mockFS, "/mock/path", false)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expectedSize := int64(150) // Only file sizes should be included
	if size != expectedSize {
		t.Errorf("Expected size %d, got %d", expectedSize, size)
	}
}

func TestGetDirectorySize_Recursive(t *testing.T) {
	// Mock Walk to traverse files recursively
	mockFS := mockFileSystem{
		walkFunc: func(root string, fn filepath.WalkFunc) error {
			fn("/mock/path/file1.txt", mockFileInfo{"file1.txt", false, 50}, nil)
			fn("/mock/path/file2.txt", mockFileInfo{"file2.txt", false, 100}, nil)
			fn("/mock/path/subdir/file3.txt", mockFileInfo{"file3.txt", false, 200}, nil)
			return nil
		},
	}

	size, err := GetDirectorySize(mockFS, "/mock/path", true)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expectedSize := int64(350) // Sum of all file sizes
	if size != expectedSize {
		t.Errorf("Expected size %d, got %d", expectedSize, size)
	}
}

func TestGetDirectorySize_ErrorHandling(t *testing.T) {
	// Mock ReadDir to return an error
	mockFS := mockFileSystem{
		readDirFunc: func(dirname string) ([]os.DirEntry, error) {
			return nil, errors.New("mock error")
		},
	}

	_, err := GetDirectorySize(mockFS, "/mock/path", false)
	if err == nil {
		t.Fatal("Expected an error, but got nil")
	}
}
