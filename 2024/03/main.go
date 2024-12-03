package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"

	"github.com/erka/aoc/pkg/log"
)

//go:embed input.txt
var input []byte

/*
* part 1: 178794710
* part 2: 76729637
 */
func main() {
	log.Infof("part1: %s", solvePart1(input))
	log.Infof("part2: %s", solvePart2(input))
}

// solve
func solvePart1(input []byte) string {
	re := regexp.MustCompile(`mul\(\d+,\d+\)`)
	// Find all matches
	matches := re.FindAllString(string(input), -1)
	var a, b, output int
	for _, m := range matches {
		_, err := fmt.Sscanf(m, "mul(%d,%d)", &a, &b)
		if err != nil {
			return err.Error()
		}
		output += a * b
	}

	return strconv.Itoa(output)
}

// solve
func solvePart2(input []byte) string {
	re := regexp.MustCompile(`mul\(\d+,\d+\)|don't\(\)|do\(\)`)
	// Find all matches
	matches := re.FindAllString(string(input), -1)
	var a, b, output int
	instructionApplied := true

	for _, m := range matches {
		switch m {
		case "don't()":
			instructionApplied = false
		case "do()":
			instructionApplied = true
		default:
			if !instructionApplied {
				continue
			}
			_, err := fmt.Sscanf(m, "mul(%d,%d)", &a, &b)
			if err != nil {
				return err.Error()
			}
			output += a * b
		}
	}

	return strconv.Itoa(output)
}
