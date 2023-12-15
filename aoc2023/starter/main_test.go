package main

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:embed example.txt
var example []byte

func TestDay111(t *testing.T) {
	result := solvePart1(example)
	require.Equal(t, "0", result)
}

func TestDay112(t *testing.T) {
	result := solvePart2(example)
	require.Equal(t, "0", result)
}
