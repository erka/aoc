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

//go:embed example3.txt
var example3 []byte

func TestDay131(t *testing.T) {
	result := solvePart1(example1)
	require.Equal(t, "408", result)
}

func TestDay132(t *testing.T) {
	result := solvePart2(example2)
	require.Equal(t, "400", result)
}

func TestDay1321(t *testing.T) {
	result := solvePart2(example3)
	require.Equal(t, "8", result)
}
