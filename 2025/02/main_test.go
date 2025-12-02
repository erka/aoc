package main

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:embed example.txt
var example []byte

func TestSolvePart1(t *testing.T) {
	result := solvePart1(example)
	require.Equal(t, "1227775554", result)
}

func TestSolvePart2(t *testing.T) {
	result := solvePart2(example)
	require.Equal(t, "4174379265", result)
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
