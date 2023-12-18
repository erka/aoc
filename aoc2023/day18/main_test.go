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
	require.Equal(t, "62", result)
}

func TestSolvePart2(t *testing.T) {
	result := solvePart2(example)
	require.Equal(t, "952408144115", result)
}
