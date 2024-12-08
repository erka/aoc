package main

import (
	"bytes"
	_ "embed"
	"slices"
	"strconv"

	"github.com/erka/aoc/pkg/log"
	"github.com/samber/lo"
)

//go:embed input.txt
var input []byte

/*
* part 1: 303876485655
* part 2: 146111650210682
 */
func main() {
	log.Infof("part1: %s", solvePart1(input))
	log.Infof("part2: %s", solvePart2(input))
}

type fn func(a, b int) int

func sum(a, b int) int {
	return a + b
}

func mul(a, b int) int {
	return a * b
}

func concat(a, b int) int {
	v, _ := strconv.Atoi(strconv.Itoa(a) + strconv.Itoa(b))
	return v
}

func f1(result int, values []int, ops []fn) bool {
	if len(values) == 1 {
		return result == values[0]
	}

	for _, op := range ops {
		copy := slices.Clone(values[1:])
		copy[0] = op(values[0], values[1])
		if f1(result, copy, ops) {
			return true
		}
	}
	return false
}

// solve
func solvePart(input []byte, ops []fn) string {
	output := 0
	lines := bytes.FieldsFunc(input, func(r rune) bool { return r == '\n' })
	for i, line := range lines {
		log.Debugf("line %d", i)
		data := bytes.FieldsFunc(line, func(r rune) bool { return r == ':' })
		if len(data) != 2 {
			return "invalid data at row"
		}

		result, err := strconv.Atoi(string(data[0]))
		if err != nil {
			return err.Error()
		}
		values := lo.Map(bytes.Fields(data[1]), func(item []byte, _ int) int {
			v, _ := strconv.Atoi(string(item))
			return v
		})

		if f1(result, values, ops) {
			output += result
		}
	}

	return strconv.Itoa(output)
}

func solvePart1(input []byte) string {
	return solvePart(input, []fn{mul, sum})
}

// solve
func solvePart2(input []byte) string {
	return solvePart(input, []fn{mul, sum, concat})
}
