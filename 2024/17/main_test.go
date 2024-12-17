package main

import (
	_ "embed"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:embed example.txt
var example []byte

func TestRun(t *testing.T) {
	tests := []struct {
		regs [3]int
		prog []int

		output   string
		expected [3]int
	}{
		{
			[3]int{0, 0, 9},
			[]int{2, 6},
			"",
			[3]int{0, 1, 9},
		},
		{
			[3]int{10, 0, 0},
			[]int{5, 0, 5, 1, 5, 4},
			"0,1,2",
			[3]int{10, 0, 0},
		},
		{
			[3]int{2024, 0, 0},
			[]int{0, 1, 5, 4, 3, 0},
			"4,2,5,6,7,7,7,7,3,1,0",
			[3]int{0, 0, 0},
		},
		{
			[3]int{0, 29, 0},
			[]int{1, 7},
			"",
			[3]int{0, 26, 0},
		},
		{
			[3]int{0, 2024, 43690},
			[]int{4, 0},
			"",
			[3]int{0, 44354, 43690},
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			output, regs := solve(tt.prog, tt.regs)
			require.Equal(t, tt.output, output)
			require.Equal(t, tt.expected, regs)
		})
	}
}

func TestSolvePart1(t *testing.T) {
	result := solvePart1(example)
	require.Equal(t, "4,6,3,5,6,3,5,2,1,0", result)
}

func TestSolvePart2(t *testing.T) {
	input := []byte(`Register A: 2024
Register B: 0
Register C: 0

Program: 0,3,5,4,3,0`)
	result := solvePart2(input)
	require.Equal(t, "117440", result)
}

func BenchmarkSolvePart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solvePart1(input)
	}
}

func BenchmarkSolvePart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solvePart2(input)
	}
}
