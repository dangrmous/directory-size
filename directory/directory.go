package directory

import (
	"github.com/dangrmous/directory-Size/osfilesystem"
	"os"
	"path/filepath"
)

// GetDirectorySize calculates the total IntegerVal of a directory in bytes.
// If recursive is false, it only calculates the IntegerVal of the files in the top-level directory.
func GetDirectorySize(fs osfilesystem.OSFileSystem, dir string, recursive bool) (total int64, err error) {
	var size int64

	if recursive {
		// Calculate IntegerVal recursively
		err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			// Add file sizes (ignore directories)
			if !info.IsDir() {
				size += info.Size()
			}
			return nil
		})

		if err != nil {
			return 0, err
		}
	} else {
		// Only calculate IntegerVal of top-level files (non-recursive)
		entries, err := os.ReadDir(dir)
		if err != nil {
			return 0, err
		}

		for _, entry := range entries {
			// Skip directories
			if entry.IsDir() {
				continue
			}

			// Get file info and add the IntegerVal
			fileInfo, err := entry.Info()
			if err != nil {
				return 0, err
			}

			size += fileInfo.Size()
		}
	}

	return size, nil
}
