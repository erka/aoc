package main

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:embed example1a.txt
var example1a []byte

//go:embed example1b.txt
var example1b []byte

//go:embed example1c.txt
var example1c []byte

//go:embed example2a.txt
var example2a []byte

//go:embed example2b.txt
var example2b []byte

func TestSolvePart1(t *testing.T) {
	tests := []struct {
		name     string
		expected string
		input    []byte
	}{
		{name: "example1a", input: example1a, expected: "140"},
		{name: "example1b", input: example1b, expected: "772"},
		{name: "example1c", input: example1c, expected: "1930"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := solvePart1(tt.input)
			require.Equal(t, tt.expected, result)
		})
	}
}

func TestSolvePart2(t *testing.T) {
	tests := []struct {
		name     string
		expected string
		input    []byte
	}{
		{name: "example1a", input: example1a, expected: "80"},
		{name: "example1b", input: example1b, expected: "436"},
		{name: "example1c", input: example1c, expected: "1206"},
		{name: "example2a", input: example2a, expected: "236"},
		{name: "example2b", input: example2b, expected: "368"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := solvePart2(tt.input)
			require.Equal(t, tt.expected, result)
		})
	}
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
