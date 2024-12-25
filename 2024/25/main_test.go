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
	require.Equal(t, "3", result)
}

func BenchmarkSolvePart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solvePart1(input)
	}
}
