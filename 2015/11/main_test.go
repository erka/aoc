package main

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//go:embed example.txt
var example []byte

func TestValidPassword(t *testing.T) {
	tests := []string{
		"hijklmmn",
		"abbceffg",
		"abbcegjk",
	}
	for _, tt := range tests {
		t.Run(tt, func(t *testing.T) {
			assert.False(t, validPassword([]byte(tt)))
		})
	}
}

func TestSolvePart1(t *testing.T) {
	result := solvePart1(example)
	require.Equal(t, "abcdffaa", result)
	result = solvePart1([]byte("ghijklmn"))
	require.Equal(t, "ghjaabcc", result)
}

func TestSolvePart2(t *testing.T) {
	result := solvePart2(example)
	require.Equal(t, "abcdffbb", result)
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
