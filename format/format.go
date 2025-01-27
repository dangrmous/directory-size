package format

// Package for formatting integers into human-readable strings

import "fmt"

func ToHumanReadable(intVal int64, baseTen bool) string {
	KB := int64(1024)
	if baseTen {
		KB = 1000
	}
	MB := KB * KB
	GB := MB * KB
	TB := GB * KB
	PB := TB * KB

	switch {
	case intVal >= PB:
		return fmt.Sprintf("%.2fPB", float64(intVal)/float64(PB))
	case intVal >= TB:
		return fmt.Sprintf("%.2fTB", float64(intVal)/float64(TB))
	case intVal >= GB:
		return fmt.Sprintf("%.2fGB", float64(intVal)/float64(GB))
	case intVal >= MB:
		return fmt.Sprintf("%.2fMB", float64(intVal)/float64(MB))
	case intVal >= KB:
		return fmt.Sprintf("%.2fKB", float64(intVal)/float64(KB))
	default:
		return fmt.Sprintf("%dB", intVal)
	}
}
