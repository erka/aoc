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
* part 1:  303876485655
* part 2: 146111650210682
 */
func main() {
	log.Infof("part1: %s", solvePart1(input))
	log.Infof("part2: %s", solvePart2(input))
}

func f1(result int, values []int) bool {
	if len(values) == 1 {
		return result == values[0]
	}

	copy := slices.Clone(values[1:])
	copy[0] *= values[0]
	if f1(result, copy) {
		return true
	}
	copy = slices.Clone(values[1:])
	copy[0] += values[0]
	return f1(result, copy)
}

func f2(result int, values []int) bool {
	if len(values) == 1 {
		return result == values[0]
	}

	copy := slices.Clone(values[1:])
	copy[0], _ = strconv.Atoi(strconv.Itoa(values[0]) + strconv.Itoa(copy[0]))
	if f2(result, copy) {
		return true
	}

	copy = slices.Clone(values[1:])
	copy[0] *= values[0]
	if f2(result, copy) {
		return true
	}

	copy = slices.Clone(values[1:])
	copy[0] += values[0]
	return f2(result, copy)
}

// solve
func solvePart1(input []byte) string {
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

		if f1(result, values) {
			output += result
		}
	}

	return strconv.Itoa(output)
}

// solve
func solvePart2(input []byte) string {
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

		if f1(result, values) || f2(result, values) {
			output += result
		}
	}

	return strconv.Itoa(output)
}
