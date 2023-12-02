package main

import (
	_ "embed"
	"testing"
)

//go:embed example.txt
var example []byte

func TestDay021(t *testing.T) {
	result, _ := solve(example, cubes)
	if result != "8" {
		println("expected 8, got: ", result)
		t.Fail()
	}
}

func TestDay022(t *testing.T) {
	_, result := solve(example, cubes)
	if result != "2286" {
		println("expected 2286, got: ", result)
		t.Fail()
	}
}
