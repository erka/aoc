package day04

import (
	_ "embed"
	"testing"
)

//go:embed input_test.txt
var input []byte

func TestSolution(t *testing.T) {
	tests := []struct {
		input    comparator
		expected int
	}{
		{input: fully, expected: 2},
		{input: overlap, expected: 4},
	}
	for _, tt := range tests {
		result := solution(input, tt.input)
		if result != tt.expected {
			t.Fatalf("experted %d, but got: %d", tt.expected, result)
		}
	}
}
