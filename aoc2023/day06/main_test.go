package main

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:embed example.txt
var example []byte

func TestDay061(t *testing.T) {
	result := solve([]BestRace{
		{7, 9},
		{15, 40},
		{30, 200},
	})
	require.Equal(t, "288", result)
}

func TestDay062(t *testing.T) {
	result := solve([]BestRace{{71530, 940200}})
	require.Equal(t, "71503", result)
}
