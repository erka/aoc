package main

import (
	"bytes"
	_ "embed"
	"strconv"

	"github.com/erka/aoc/pkg/log"
)

//go:embed input.txt
var input []byte

/*
* part 1:
* part 2:
 */
func main() {
	log.Infof("part1: %s", solvePart1(input))
	log.Infof("part2: %s", solvePart2(input))
}

// solve
func solvePart1(input []byte) string {
	lines := bytes.Split(bytes.Trim(input, "\n"), []byte("\n"))
	log.Debugf("lines: %v", lines)
	return strconv.Itoa(0)
}

// solve
func solvePart2(input []byte) string {
	_ = input
	return strconv.Itoa(0)
}
