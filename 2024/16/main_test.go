package main

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:embed example1.txt
var example1 []byte

//go:embed example2.txt
var example2 []byte

func TestSolvePart11(t *testing.T) {
	result := solvePart1(example1)
	require.Equal(t, "7036", result)
}

func TestSolvePart12(t *testing.T) {
	result := solvePart1(example2)
	require.Equal(t, "11048", result)
}

func TestSolvePart21(t *testing.T) {
	result := solvePart2(example1)
	require.Equal(t, "45", result)
}

func TestSolvePart22(t *testing.T) {
	result := solvePart2(example2)
	require.Equal(t, "64", result)
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
