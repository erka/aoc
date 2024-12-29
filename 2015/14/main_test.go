package main

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:embed example.txt
var example []byte

func TestSolvePart1(t *testing.T) {
	result := solvePart1(example, 1000)
	require.Equal(t, "1120", result)
}

func TestSolvePart2(t *testing.T) {
	result := solvePart2(example, 1000)
	require.Equal(t, "689", result)
}

func BenchmarkSolvePart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solvePart1(input, 1000)
	}
}

func BenchmarkSolvePart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solvePart2(input, 1000)
	}
}
