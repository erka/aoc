package main

import (
	_ "embed"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:embed example1.txt
var example1 []byte

//go:embed example2.txt
var example2 []byte

func TestSolvePart1(t *testing.T) {
	tests := []struct {
		input    []byte
		expected string
	}{
		{example1, "4"},
		{example2, "2024"},
	}

	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			result := solvePart1(tt.input)
			require.Equal(t, tt.expected, result)
		})
	}
}

func TestSolvePart2(t *testing.T) {
	// result := solvePart2(example)
	// require.Equal(t, "0", result)
}

func BenchmarkSolvePart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solvePart1(input)
	}
}

func BenchmarkSolvePart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solvePart2(input)
	}
}
