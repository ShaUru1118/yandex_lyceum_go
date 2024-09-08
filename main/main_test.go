package main

import (
	"testing"
)

func TestGetUTFLength(t *testing.T) {
	testCases := []struct {
		input    []byte
		expected int
		err      error
	}{
		{[]byte("Hello, 世界"), 9, nil},
		{[]byte("Привет, мир"), 11, nil},
		{[]byte("Invalid \xc3"), 0, ErrInvalidUTF8},
	}

	for _, tc := range testCases {
		actual, err := GetUTFLength(tc.input)
		if actual != tc.expected {
			t.Errorf("Expected length %d, but got %d", tc.expected, actual)
		}

		if tc.err != err {
			t.Errorf("Expected error %v, but got %v", tc.err, err)
		}
	}
}
