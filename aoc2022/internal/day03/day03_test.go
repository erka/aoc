package day03

import (
	_ "embed"
	"testing"
)

//go:embed input_test1.txt
var input1 []byte

//go:embed input_test2.txt
var input2 []byte

func TestSolutionPart1(t *testing.T) {
	expected := 157
	result := solutionPart1(input1)
	if result != expected {
		t.Fatalf("experted %d, but got: %d", expected, result)
	}

}

func TestSolutionPart2(t *testing.T) {
	expected := 70
	result := solutionPart2(input2)
	if result != expected {
		t.Fatalf("experted %d, but got: %d", expected, result)
	}

}

func TestItemPriority(t *testing.T) {
	tests := []struct {
		input    rune
		expected int
	}{
		{input: 'p', expected: 16},
		{input: 'L', expected: 38},
		{input: 'P', expected: 42},
		{input: 'v', expected: 22},
		{input: 't', expected: 20},
		{input: 's', expected: 19},
		{input: 'r', expected: 18},
		{input: 'Z', expected: 52},
	}
	for _, tt := range tests {
		result := itemPriority(tt.input)
		if result != tt.expected {
			t.Fatalf("experted %d, but got: %d", tt.expected, result)
		}
	}
}
