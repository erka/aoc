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

//go:embed example21.txt
var example21 []byte

//go:embed example22.txt
var example22 []byte

func TestDay1011(t *testing.T) {
	result := solve(example11)
	require.Equal(t, "4", result)
}

func TestDay1012(t *testing.T) {
	result := solve(example12)
	require.Equal(t, "8", result)
}

func TestDay1021(t *testing.T) {
	result := solvePart2(example21)
	require.Equal(t, "4", result)
}

func TestDay1022(t *testing.T) {
	result := solvePart2(example22)
	require.Equal(t, "10", result)
}
