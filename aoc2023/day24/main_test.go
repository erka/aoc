package main

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:embed example.txt
var example []byte

func TestSolvePart1(t *testing.T) {
	result := solvePart1(example, 7, 27)
	require.Equal(t, "2", result)
}

func TestSolvePart2(t *testing.T) {
	result := solvePart2(example)
	require.Equal(t, "0", result)
}
