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
		baseTen  bool   // Whether to represent values in baseTen or not
		expected string // Expected human-readable format
	}{
		{0, false, "0B"},                    // Test zero bytes
		{500, false, "500B"},                // Less than 1 KB
		{1024, false, "1.00KB"},             // Exactly 1 KB
		{1536, false, "1.50KB"},             // 1.5 KB
		{1048576, false, "1.00MB"},          // Exactly 1 MB
		{1572864, false, "1.50MB"},          // 1.5 MB
		{1073741824, false, "1.00GB"},       // Exactly 1 GB
		{1610612736, false, "1.50GB"},       // 1.5 GB
		{1099511627776, false, "1.00TB"},    // Exactly 1 TB
		{1649267441664, false, "1.50TB"},    // 1.5 TB
		{1125899906842624, false, "1.00PB"}, // Exactly 1 PB
		{0, true, "0B"},                     // Decimal mode tests
		{500, true, "500B"},                 // Less than 1 KB
		{1000, true, "1.00KB"},              // Exactly 1 KB
		{1500, true, "1.50KB"},              // 1.5 KB
		{1000000, true, "1.00MB"},           // Exactly 1 MB
		{1500000, true, "1.50MB"},           // 1.5 MB
		{1000000000, true, "1.00GB"},        // Exactly 1 GB
		{1500000000, true, "1.50GB"},        // 1.5 GB
		{1000000000000, true, "1.00TB"},     // Exactly 1 TB
		{1500000000000, true, "1.50TB"},     // 1.5 TB
		{1000000000000000, true, "1.00PB"},  // Exactly 1 PB
	}

	// Loop through each test case
	for _, tc := range testCases {
		t.Run(tc.expected, func(t *testing.T) {
			// Call the function
			result := format.ToHumanReadable(tc.input, tc.baseTen)

			// Assert the result is as expected
			assert.Equal(t, tc.expected, result)
		})
	}
}
