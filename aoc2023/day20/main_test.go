package main

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:embed example1.txt
var example1 []byte

//go:embed example2.txt
var example2 []byte

func TestSolvePart11(t *testing.T) {
	result := solvePart1(example1)
	require.Equal(t, "32000000", result)
}

func TestSolvePart12(t *testing.T) {
	result := solvePart1(example2)
	require.Equal(t, "11687500", result)
}
