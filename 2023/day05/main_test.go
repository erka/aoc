package main

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:embed example.txt
var example []byte

func TestDay051(t *testing.T) {
	seed := newSeed("79")
	seed.Transform(52, 50, 48)
	require.Equal(t, 81, seed.location)
	seed = newSeed("98")
	seed.Transform(50, 98, 2)
	require.Equal(t, 50, seed.location)
	result := solve(example)
	require.Equal(t, "35", result)
}

func TestDay052(t *testing.T) {
	result := solvePart2(example)
	require.Equal(t, "46", result)
}
