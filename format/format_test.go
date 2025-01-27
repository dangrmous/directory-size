package format_test

import (
	"testing"

	"github.com/dangrmous/directory-size/format"
	"github.com/stretchr/testify/assert"
)

func TestToHumanReadable(t *testing.T) {
	// Define test cases
	testCases := []struct {
		input    int64  // Bytes as input
		expected string // Expected human-readable format
	}{
		{0, "0B"},                    // Test zero bytes
		{500, "500B"},                // Less than 1 KB
		{1024, "1.00KB"},             // Exactly 1 KB
		{1536, "1.50KB"},             // 1.5 KB
		{1048576, "1.00MB"},          // Exactly 1 MB
		{1572864, "1.50MB"},          // 1.5 MB
		{1073741824, "1.00GB"},       // Exactly 1 GB
		{1610612736, "1.50GB"},       // 1.5 GB
		{1099511627776, "1.00TB"},    // Exactly 1 TB
		{1649267441664, "1.50TB"},    // 1.5 TB
		{1125899906842624, "1.00PB"}, // Exactly 1 PB
	}

	// Loop through each test case
	for _, tc := range testCases {
		t.Run(tc.expected, func(t *testing.T) {
			// Call the function
			result := format.ToHumanReadable(tc.input)

			// Assert the result is as expected
			assert.Equal(t, tc.expected, result)
		})
	}
}
