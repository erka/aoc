package main

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:embed example.txt
var example []byte

func TestSolvePart1(t *testing.T) {
	result := solvePart1(example, 6, 12)
	require.Equal(t, "22", result)
}

func TestSolvePart2(t *testing.T) {
	result := solvePart2(example, 6)
	require.Equal(t, "6,1", result)
}

func BenchmarkSolvePart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solvePart1(input, 70, 1024)
	}
}

func BenchmarkSolvePart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solvePart2(input, 70)
	}
}
