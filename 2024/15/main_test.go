package main

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:embed example.txt
var example []byte

//go:embed example1.txt
var example1 []byte

//go:embed example2.txt
var example2 []byte

func TestSolvePart1a(t *testing.T) {
	result := solvePart1(example)
	require.Equal(t, "10092", result)
}

func TestSolvePart1b(t *testing.T) {
	result := solvePart1(example1)
	require.Equal(t, "2028", result)
}

func TestSolvePart2a(t *testing.T) {
	result := solvePart2(example2)
	require.Equal(t, "618", result)
}

func TestSolvePart2b(t *testing.T) {
	result := solvePart2(example)
	require.Equal(t, "9021", result)
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
