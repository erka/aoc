package main

import (
	_ "embed"
	"strconv"

	"github.com/erka/aoc/pkg/log"
)

//go:embed input.txt
var input []byte

/*
* part1: 138
* part2: 1771
 */
func main() {
	log.Infof("part1: %s", solvePart1(input))
	log.Infof("part2: %s", solvePart2(input))
}

func solvePart1(input []byte) string {
	i := 0
	for _, c := range input {
		switch c {
		case '(':
			i++
		case ')':
			i--
		}
	}
	return strconv.Itoa(i)
}

func solvePart2(input []byte) string {
	i := 0
	for j, c := range input {
		switch c {
		case '(':
			i++
		case ')':
			i--
		}
		if i == -1 {
			return strconv.Itoa(j + 1)
		}
	}
	return ""
}
