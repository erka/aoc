package main

import (
	_ "embed"
	"testing"
)

//go:embed example1.txt
var example1 []byte

//go:embed example2.txt
var example2 []byte

func TestDay011(t *testing.T) {
	result := solve(example1, digits)
	if result != "142" {
		t.Fail()
	}
}

func TestDay012(t *testing.T) {
	result := solve(example2, wordsAndDigits)
	if result != "281" {
		println("expected 281, got: ", result)
		t.Fail()
	}
	result = solve([]byte("78twoninepdghsneightone"), wordsAndDigits)
	if result != "71" {
		println("expected 71, got: ", result)
		t.Fail()
	}
	result = solve([]byte("eight3fiveninefivemtxm9eightwot"), wordsAndDigits)
	if result != "82" {
		println("expected 82, got: ", result)
		t.Fail()
	}
}
