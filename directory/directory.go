package directory

import (
	"fmt"
	"os"
	"path/filepath"
)

type ByteCount struct {
	IntegerVal int64
}

func (t ByteCount) ToHumanReadable() string {
	const (
		KB = 1000
		MB = 1000 * KB
		GB = 1000 * MB
		TB = 1000 * GB
		PB = 1000 * TB
	)

	switch {
	case t.IntegerVal >= PB:
		return fmt.Sprintf("%.2fPB", float64(t.IntegerVal)/PB)
	case t.IntegerVal >= TB:
		return fmt.Sprintf("%.2fTB", float64(t.IntegerVal)/TB)
	case t.IntegerVal >= GB:
		return fmt.Sprintf("%.2fGB", float64(t.IntegerVal)/GB)
	case t.IntegerVal >= MB:
		return fmt.Sprintf("%.2fMB", float64(t.IntegerVal)/MB)
	case t.IntegerVal >= KB:
		return fmt.Sprintf("%.2fKB", float64(t.IntegerVal)/KB)
	default:
		return fmt.Sprintf("%dB", t.IntegerVal)
	}
}

// GetDirectorySize calculates the total IntegerVal of a directory in bytes.
// If recursive is false, it only calculates the IntegerVal of the files in the top-level directory.
func GetDirectorySize(dir string, recursive bool) (total ByteCount, err error) {
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
			return ByteCount{0}, err
		}
	} else {
		// Only calculate IntegerVal of top-level files (non-recursive)
		entries, err := os.ReadDir(dir)
		if err != nil {
			return ByteCount{0}, err
		}

		for _, entry := range entries {
			// Skip directories
			if entry.IsDir() {
				continue
			}

			// Get file info and add the IntegerVal
			fileInfo, err := entry.Info()
			if err != nil {
				return ByteCount{0}, err
			}

			size += fileInfo.Size()
		}
	}

	return ByteCount{size}, nil
}
