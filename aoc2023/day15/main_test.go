package main

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:embed example.txt
var example []byte

func TestDay111(t *testing.T) {
	value := hashAlg("HASH")
	require.Equal(t, 52, value)
	result := solvePart1(example)
	require.Equal(t, "1320", result)
}

func TestDay112(t *testing.T) {
	result := solvePart2(example)
	require.Equal(t, "145", result)
}
