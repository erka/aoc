package main

import (
	"strconv"
	"testing"
)

func TestSolvePart1(t *testing.T) {
	tests := []struct {
		input    []byte
		expected string
	}{
		{[]byte(`2x3x4`), "58"},
		{[]byte(`1x1x10`), "43"},
		{[]byte("1x1x10\n2x3x4\n"), "101"},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := solvePart1(tt.input); got != tt.expected {
				t.Errorf("test %d: expected %s, got %s", i, tt.expected, got)
			}
		})
	}
}

func TestSolvePart2(t *testing.T) {
	tests := []struct {
		input    []byte
		expected string
	}{
		{[]byte(`2x3x4`), "34"},
		{[]byte(`1x1x10`), "14"},
		{[]byte("1x1x10\n2x3x4\n"), "48"},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := solvePart2(tt.input); got != tt.expected {
				t.Errorf("test %d: expected %s, got %s", i, tt.expected, got)
			}
		})
	}
}
