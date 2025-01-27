package format

// Package for formatting integers into human-readable strings

import "fmt"

func ToHumanReadable(intVal int64) string {
	const (
		KB = 1024
		MB = 1024 * KB
		GB = 1024 * MB
		TB = 1024 * GB
		PB = 1024 * TB
	)

	switch {
	case intVal >= PB:
		return fmt.Sprintf("%.2fPB", float64(intVal)/PB)
	case intVal >= TB:
		return fmt.Sprintf("%.2fTB", float64(intVal)/TB)
	case intVal >= GB:
		return fmt.Sprintf("%.2fGB", float64(intVal)/GB)
	case intVal >= MB:
		return fmt.Sprintf("%.2fMB", float64(intVal)/MB)
	case intVal >= KB:
		return fmt.Sprintf("%.2fKB", float64(intVal)/KB)
	default:
		return fmt.Sprintf("%dB", intVal)
	}
}
