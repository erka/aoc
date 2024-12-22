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

func TestMix(t *testing.T) {
	if n := mix(42, 15); n != 37 {
		t.Error("mix(42, 15) != 37", n)
	}
}

func TestPrune(t *testing.T) {
	if n := prune(100000000); n != 16113920 {
		t.Error("prune(100000000) != 16113920", n)
	}
}

func TestNext(t *testing.T) {
	s := next(123)
	require.Equal(t, int64(15887950), s)
}

func TestPrice(t *testing.T) {
	tests := []struct {
		input    int64
		expected int
	}{
		{123, 3},
		{15887950, 0},
		{16495136, 6},
	}
	for _, tt := range tests {
		require.Equal(t, tt.expected, price(tt.input))
	}
}

func TestSolvePart1(t *testing.T) {
	result := solvePart1(example1)
	require.Equal(t, "37327623", result)
}

func TestSolvePart2(t *testing.T) {
	result := solvePart2(example2)
	require.Equal(t, "23", result)
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
