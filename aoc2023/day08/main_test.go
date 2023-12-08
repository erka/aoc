package main

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:embed example11.txt
var example11 []byte

//go:embed example12.txt
var example12 []byte

//go:embed example2.txt
var example2 []byte

func TestDay081(t *testing.T) {
	result := solvePart1(example11)
	require.Equal(t, "2", result)
	result = solvePart1(example12)
	require.Equal(t, "6", result)
}

func TestDay082(t *testing.T) {
	result := solvePart2(example2)
	require.Equal(t, "6", result)
}
