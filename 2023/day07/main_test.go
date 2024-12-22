package main

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:embed example.txt
var example []byte

func TestDay071(t *testing.T) {
	result := solve(example, part1, handType)
	require.Equal(t, "6440", result)
}

func TestDay072(t *testing.T) {
	result := solve(example, part2, handTypeJoker)
	require.Equal(t, "5905", result)
}

func TestJoker(t *testing.T) {
	c := handTypeJoker(hits("248JJ"))
	require.Equal(t, ThreeOfKind, c)
}
