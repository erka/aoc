package main

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:embed example.txt
var example []byte

func TestDay041(t *testing.T) {
	result := solve(example)
	require.Equal(t, "13", result)
}

func TestDay042(t *testing.T) {
	result := solveCopies(example)
	require.Equal(t, "30", result)
}
