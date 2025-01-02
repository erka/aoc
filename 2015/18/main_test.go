package main

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:embed example.txt
var example []byte

//go:embed example2.txt
var example2 []byte

func TestSolvePart1(t *testing.T) {
	result := solvePart1(example, 4)
	require.Equal(t, "4", result)
}

func TestSolvePart2(t *testing.T) {
	result := solvePart2(example2, 5)
	require.Equal(t, "17", result)
}

func BenchmarkSolvePart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solvePart1(input, 100)
	}
}

func BenchmarkSolvePart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solvePart2(input, 100)
	}
}
