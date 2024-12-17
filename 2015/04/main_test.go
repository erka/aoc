package main

import (
	"strconv"
	"testing"
)

func TestSolvePart1(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"abcdef", "609043"},
		{"pqrstuv", "1048970"},
	}

	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := solvePart1([]byte(tt.input)); got != tt.expected {
				t.Errorf("test %d: expected %s, got %s", i, tt.expected, got)
			}
		})
	}
}
