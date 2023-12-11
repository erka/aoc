package main

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:embed example.txt
var example []byte

func TestDay111(t *testing.T) {
	result := solve(example, 2)
	require.Equal(t, "374", result)
}

func TestDay112(t *testing.T) {
	result := solve(example, 10)
	require.Equal(t, "1030", result)
}

func TestDay113(t *testing.T) {
	result := solve(example, 100)
	require.Equal(t, "8410", result)
}
