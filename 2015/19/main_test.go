package main

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:embed example.txt
var example1 []byte

//go:embed example2.txt
var example2 []byte

func TestSolvePart1(t *testing.T) {
	result := solvePart1(example1)
	require.Equal(t, "4", result)
}

func TestSolvePart2(t *testing.T) {
	result := solvePart2(example2)
	require.Equal(t, "6", result)
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
