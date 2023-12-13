package main

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:embed example.txt
var example []byte

func TestDay121(t *testing.T) {

	tt := map[string]string{
		"???????????#????#?? 2,1,11,1": "3",
		"???.### 1,1,3":                "1",
		"??.?? 1,1":                    "4",
		"???.????#?? 1,3":              "15",
		"??#.?.?#.#?##?###? 2,1,8":     "1",
		"?#?#??????#?? 4,1,1":          "9",
	}

	for tk, tv := range tt {
		result := solvePart1([]byte(tk))
		require.Equal(t, tv, result)
	}
	result := solvePart1(example)
	require.Equal(t, "34", result)
}

func TestDay122(t *testing.T) {
	tt := map[string]string{
		"???.### 1,1,3":             "1",
		".??..??...?##. 1,1,3":      "16384",
		"?#?#?#?#?#?#?#? 1,3,1,6":   "1",
		"????.#...#... 4,1,1":       "16",
		"????.######..#####. 1,6,5": "2500",
		"?###???????? 3,2,1":        "506250",
	}

	for tk, tv := range tt {
		result := solvePart2([]byte(tk))
		require.Equal(t, tv, result)
	}

	result := solvePart2(example)
	require.Equal(t, "998928", result)
}
