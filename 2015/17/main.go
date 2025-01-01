package main

import (
	"bytes"
	_ "embed"
	"strconv"

	"github.com/erka/aoc/pkg/log"
	"github.com/samber/lo"
)

//go:embed input.txt
var input []byte

/*
* part 1: 654
* part 2: 57
 */
func main() {
	log.Infof("part1: %s", solvePart1(input, 150))
	log.Infof("part2: %s", solvePart2(input, 150))
}

// solve
func solvePart1(input []byte, liters int) string {
	lines := bytes.Split(bytes.Trim(input, "\n"), []byte("\n"))
	containers := lo.Map(lines, func(line []byte, _ int) int {
		n, err := strconv.Atoi(string(line))
		if err != nil {
			panic(err)
		}
		return n
	})
	output, _ := find(containers, liters, []int{})
	return strconv.Itoa(output)
}

func find(containers []int, liters int, ways []int) (int, [][]int) {
	if liters == 0 {
		return 1, [][]int{ways}
	}
	if liters < 0 || len(containers) == 0 {
		return 0, [][]int{}
	}
	v1, w1 := find(containers[1:], liters-containers[0], append(ways, containers[0]))
	v2, w2 := find(containers[1:], liters, ways)

	return v1 + v2, append(w1, w2...)
}

// solve
func solvePart2(input []byte, liters int) string {
	lines := bytes.Split(bytes.Trim(input, "\n"), []byte("\n"))
	containers := lo.Map(lines, func(line []byte, _ int) int {
		n, err := strconv.Atoi(string(line))
		if err != nil {
			panic(err)
		}
		return n
	})
	_, ways := find(containers, liters, []int{})
	minimal := lo.MinBy(ways, func(a, b []int) bool {
		return len(a) < len(b)
	})
	output := lo.CountBy(ways, func(item []int) bool {
		return len(item) == len(minimal)
	})
	return strconv.Itoa(output)
}
