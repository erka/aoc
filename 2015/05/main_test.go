package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolvePart1(t *testing.T) {
	input := "aaa\nugknbfddgicrmopn\njchzalrnumimnmhp\ndvszwmarrgswjxmb\nhaegwjzuvuyypxyu"
	output := solvePart1([]byte(input))
	assert.Equal(t, "2", output)
}

func TestContainsVowels(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"ab", false},
		{"cdasd", false},
		{"axupqxo", true},
		{"asdufxyio", true},
		{"avidgew", true},
		{"aaa", true},
	}

	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			output := containsVowels([]byte(test.input))
			assert.Equal(t, test.expected, output)
		})
	}
}

func TestContainsReserved(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"ab", true},
		{"cdasd", true},
		{"axpqxa", true},
		{"asdfxy", true},
		{"avdgew", false},
	}

	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			output := containsReserved([]byte(test.input))
			assert.Equal(t, test.expected, output)
		})
	}
}

func TestContainsDoubles(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"abcdef", false},
		{"cddef", true},
		{"ggew", true},
		{"gewww", true},
	}

	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			output := containsDouble([]byte(test.input))
			assert.Equal(t, test.expected, output)
		})
	}
}

func TestSolutionPart2(t *testing.T) {
	input := "qjhvhtzxzqqjkmpb\nxxyxx\nuurcxstgmygtbstg\nieodomkazucvgmuy\n"
	output := solvePart2([]byte(input))
	assert.Equal(t, "2", output)
}

func TestDoublePair(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"abcdef", false},
		{"cddefdd", true},
		{"ggew", false},
		{"gewww", false},
		{"gewwww", true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			output := containsDoublePair([]byte(tt.input))
			assert.Equal(t, tt.expected, output)
		})
	}
}

func TestContainsMirrorsWithoutOverlap(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"aaa", true},
		{"aabb", false},
		{"aabaa", true},
		{"bbbb", true},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			output := containsMirrors([]byte(tt.input))
			assert.Equal(t, tt.expected, output)
		})
	}
}
