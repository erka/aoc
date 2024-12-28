package main

import (
	"bytes"
	_ "embed"
	"slices"

	"github.com/erka/aoc/pkg/log"
)

//go:embed input.txt
var input []byte

/*
* part 1: hepxxyzz
* part 2:
 */
func main() {
	log.Infof("part1: %s", solvePart1(input))
	log.Infof("part2: %s", solvePart2(input))
}

// solve
func solvePart1(input []byte) string {
	passwords := solve(slices.Clone(input))
	return passwords[0]
}

// solve
func solvePart2(input []byte) string {
	passwords := solve(slices.Clone(input))
	return passwords[1]
}

func solve(input []byte) []string {
	line := bytes.Trim(input, "\n")
	l := len(line)
	passwd := []string{}
	for {
		if line[l-1]%'z' == 0 {
			for k := 1; line[l-k]%'z' == 0; k++ {
				line[l-k] = 'a'
				if line[l-k-1] != 'z' {
					line[l-k-1]++
					break
				}
			}
		} else {
			line[l-1]++
		}

		if validPassword(line) {
			passwd = append(passwd, string(line))
			if len(passwd) == 2 {
				break
			}
		}
	}
	return passwd
}

func validPassword(input []byte) bool {
	if bytes.ContainsAny(input, "iol") {
		return false
	}
	var pairs int
	var straight bool

	var last byte
	var beforelast byte
	for i, c := range input {
		if i > 0 {
			if c == last && c != beforelast {
				pairs++
			}
			if c == last+1 && c == beforelast+2 {
				straight = true
			}
		}
		beforelast = last
		last = c
	}

	return pairs >= 2 && straight
}
