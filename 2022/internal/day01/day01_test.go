package day01

import (
	"bytes"
	"io"
	"testing"
)

const input = `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000
`

func TestSolution(t *testing.T) {
	r := bytes.NewReader([]byte(input))
	tests := []struct {
		input    int
		expected int
	}{
		{input: 1, expected: 24000},
		{input: 3, expected: 45000},
	}
	for _, tt := range tests {
		r.Seek(0, io.SeekStart)
		result := solution(r, tt.input)
		if result != tt.expected {
			t.Fatalf("experted %d but got: %d", tt.expected, result)
		}
	}
}
