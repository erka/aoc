package main

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:embed example.txt
var example []byte

func TestDay03(t *testing.T) {
	part1, part2 := solve(example)
	require.Equal(t, "4361", part1)
	require.Equal(t, "467835", part2)
}
