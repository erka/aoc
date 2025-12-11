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

func TestSolvePart1(t *testing.T) {
	result := solvePart1(example1)
	require.Equal(t, "5", result)
}

func TestSolvePart2(t *testing.T) {
	result := solvePart2(example2)
	require.Equal(t, "2", result)
}

func BenchmarkSolvePart1(b *testing.B) {
	for b.Loop() {
		solvePart1(input)
	}
}

func BenchmarkSolvePart2(b *testing.B) {
	for b.Loop() {
		solvePart2(input)
	}
}
