package main

import (
	"bytes"
	_ "embed"
	"math"
	"strconv"

	"github.com/erka/aoc/pkg/log"
)

//go:embed input.txt
var input []byte

/*
* part 1: 776160
* part 2: 786240
 */
func main() {
	log.Infof("part1: %s", solvePart1(input))
	log.Infof("part2: %s", solvePart2(input))
}

// solve
func solvePart1(input []byte) string {
	n, _ := strconv.Atoi(string(bytes.Trim(input, "\n")))
	n /= 10
	houses := make([]int, n)
	house := math.MaxInt
	for i := 1; i < n; i++ {
		for elf := i; elf < n; elf += i {
			houses[elf] += i
			if houses[elf] >= n {
				house = min(house, elf)
			}
		}
	}
	return strconv.Itoa(house)
}

// solve
func solvePart2(input []byte) string {
	n, _ := strconv.Atoi(string(bytes.Trim(input, "\n")))
	n /= 10
	houses := make([]int, n)
	house := math.MaxInt
	for i := 1; i < n; i++ {
		c := 0
		for elf := i; elf < n; elf += i {
			houses[elf] += i * 11
			if houses[elf] >= n*10 {
				house = min(house, elf)
			}
			c++
			if c > 50 {
				break
			}
		}
	}
	return strconv.Itoa(house)
}
