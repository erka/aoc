package main

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:embed example.txt
var example []byte

func TestSolvePart1(t *testing.T) {
	result := solve(example, 36, 2)
	require.Equal(t, "4", result)
}

func TestSolvePart2(t *testing.T) {
	result := solve(example, 68, 20)
	require.Equal(t, "55", result)
}

func BenchmarkSolvePart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solve(input, 100, 2)
	}
}

func BenchmarkSolvePart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solve(input, 100, 20)
	}
}
