package main

import (
	"testing"
)

func TestValid(t *testing.T) {
	tests := []struct {
		name           string
		number         string
		expectedResult bool
	}{
		{
			name:           "just return default value (false)",
			number:         "",
			expectedResult: false,
		},

		{
			name:           "less than 2 chars",
			number:         "1",
			expectedResult: false,
		},
		{
			name:           "valid credit card",
			number:         "4539148803436467",
			expectedResult: true,
		},
		{
			name:           "invalid credit card",
			number:         "8273123273520569",
			expectedResult: false,
		},

		{
			name:           "valid credit card with spaces",
			number:         "4539 1488 0343 6467",
			expectedResult: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsValid(test.number)

			if result != test.expectedResult {
				t.Errorf("expectec result for %s is %v, got %v", test.number, test.expectedResult, result)
			}
		})
	}
}

func BenchmarkIsValid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsValid("4539 1488 0343 6467")
	}
}
