package main

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:embed example.txt
var example []byte

func TestDay091(t *testing.T) {
	result := solve(example)
	require.Equal(t, "114", result)
}

func TestDay092(t *testing.T) {
	result := solvePart2(example)
	require.Equal(t, "2", result)
}
