package main

import (
	"testing"
)

func TestHandleInput(t *testing.T) {
	tests := []struct {
		name        string
		inputStr    string
		inputAmount int
		expected    []SlicedMap
		expectedErr string
	}{
		{
			name:        "normal behavior with K less than the number of unique words",
			inputStr:    "aa bb cc aa cc cc cc aa ab ac bb",
			inputAmount: 3,
			expected: []SlicedMap{
				{"cc", 4},
				{"aa", 3},
				{"bb", 2},
			},
			expectedErr: "",
		},
		{
			name:        "empty list of words",
			inputStr:    "",
			inputAmount: 3,
			expectedErr: "ERROR: empty string",
		},
		{
			name:        "list of words where K is greater than the number of unique words",
			inputStr:    "aa bb cc",
			inputAmount: 5,
			expected: []SlicedMap{
				{"aa", 1},
				{"bb", 1},
				{"cc", 1},
			},
			expectedErr: "",
		},
		{
			name:        "amount of words is zero",
			inputStr:    "aa bb cc",
			inputAmount: 0,
			expectedErr: "ERROR: ammount of words less than 0",
		},
		{
			name:        "amount of words is negative",
			inputStr:    "aa bb cc",
			inputAmount: -1,
			expectedErr: "ERROR: ammount of words less than 0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := handleInput(tt.inputStr, tt.inputAmount)

			if tt.expectedErr != "" {
				if err == nil || err.Error() != tt.expectedErr {
					t.Errorf("Expected error: %v, got: %v", tt.expectedErr, err)
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if !compareSlices(result, tt.expected) {
					t.Errorf("Expected: %v, got: %v", tt.expected, result)
				}
			}
		})
	}
}

func compareSlices(a, b []SlicedMap) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
