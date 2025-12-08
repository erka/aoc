package main

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:embed example.txt
var example []byte

func TestSolvePart1(t *testing.T) {
	result := solvePart1(example, 10)
	require.Equal(t, "40", result)
}

func TestSolvePart2(t *testing.T) {
	result := solvePart2(example)
	require.Equal(t, "25272", result)
}

func BenchmarkSolvePart1(b *testing.B) {
	for b.Loop() {
		solvePart1(input, 10)
	}
}

func BenchmarkSolvePart2(b *testing.B) {
	for b.Loop() {
		solvePart2(input)
	}
}
