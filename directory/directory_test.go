package directory_test

import (
	"io/fs"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/dangrmous/directory-size/directory"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockFileSystem is a mock implementation of osfilesystem.FileSystem
type MockFileSystem struct {
	mock.Mock
}

func (mfs *MockFileSystem) Walk(root string, fn filepath.WalkFunc) error {
	fn("/test-dir", MockFileInfo{
		size: 100,
	}, nil)
	return nil
}

// Mock implementation of ReadDir
func (mfs *MockFileSystem) ReadDir(path string) ([]os.DirEntry, error) {
	args := mfs.Called(path)
	return args.Get(0).([]os.DirEntry), args.Error(1)
}

// MockDirEntry mocks the os.DirEntry interface
type MockDirEntry struct {
	name  string
	size  int64
	isDir bool
}

func (mde *MockDirEntry) Type() fs.FileMode {
	return fs.FileMode(2147483648)
}

func (mde *MockDirEntry) Name() string { return mde.name }
func (mde *MockDirEntry) IsDir() bool  { return mde.isDir }
func (mde *MockDirEntry) Info() (os.FileInfo, error) {
	return &MockFileInfo{size: mde.size}, nil
}

// MockFileInfo mocks the os.FileInfo interface
type MockFileInfo struct {
	size  int64
	name  string
	isDir bool
}

func (mfi MockFileInfo) Size() int64        { return mfi.size }
func (mfi MockFileInfo) IsDir() bool        { return false }
func (mfi MockFileInfo) Name() string       { return "" }
func (mfi MockFileInfo) Mode() os.FileMode  { return 0 }
func (mfi MockFileInfo) ModTime() time.Time { return time.Time{} }
func (mfi MockFileInfo) Sys() interface{}   { return nil }

func TestGetDirectorySize(t *testing.T) {
	t.Run("Non-recursive, only files in the top-level directory", func(t *testing.T) {
		mockFS := new(MockFileSystem)

		// Mock response for ReadDir
		mockFS.On("ReadDir", "/test-dir").Return([]os.DirEntry{
			&MockDirEntry{name: "file1.txt", size: 100, isDir: false},
			&MockDirEntry{name: "file2.txt", size: 200, isDir: false},
			&MockDirEntry{name: "subdir", isDir: true}, // Subdirectory, should be ignored
		}, nil)

		// Call the function
		size, err := directory.GetDirectorySize(mockFS, "/test-dir", false)

		// Assert results
		assert.NoError(t, err)
		assert.Equal(t, int64(300), size)
		mockFS.AssertExpectations(t)
	})

	t.Run("Recursive", func(t *testing.T) {
		mockFS := new(MockFileSystem)

		mockFS.On("Walk", "/").Return(nil)

		size, err := directory.GetDirectorySize(mockFS, "/test-dir", true)

		// Mock response for Walk

		// Assert results
		assert.Error(t, err)
		assert.Equal(t, int64(0), size)
		mockFS.AssertExpectations(t)
	})

}
