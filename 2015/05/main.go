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
* part1: 236
* part2: 51
 */
func main() {
	log.Infof("part1: %s", solvePart1(input))
	log.Infof("part2: %s", solvePart2(input))
}

func solvePart1(input []byte) string {
	lines := bytes.Split(bytes.Trim(input, "\n"), []byte("\n"))
	output := lo.Filter(lines, func(line []byte, _ int) bool {
		if containsReserved(line) {
			return false
		}
		if !containsDouble(line) {
			return false
		}
		return containsVowels(line)
	})
	return strconv.Itoa(len(output))
}

func containsVowels(input []byte) bool {
	vowels := []byte{'a', 'e', 'i', 'o', 'u'}
	count := lo.Sum(lo.Map(vowels, func(v byte, _ int) int {
		return bytes.Count(input, []byte{v})
	}))
	return count >= 3
}

func containsDouble(line []byte) bool {
	doubles := false
	for i := 1; i < len(line); i++ {
		if line[i-1] == line[i] {
			doubles = true
			break
		}
	}
	return doubles
}

func containsReserved(line []byte) bool {
	for _, exc := range []string{"ab", "cd", "pq", "xy"} {
		if bytes.Contains(line, []byte(exc)) {
			return true
		}
	}
	return false
}

func solvePart2(input []byte) string {
	lines := bytes.Split(bytes.Trim(input, "\n"), []byte("\n"))
	output := lo.Filter(lines, func(line []byte, _ int) bool {
		if !containsDoublePair(line) {
			return false
		}
		return containsMirrors(line)
	})
	return strconv.Itoa(len(output))
}

func containsDoublePair(line []byte) bool {
	for i := 0; i < len(line)-1; i += 1 {
		pair := line[i : i+2]
		if bytes.Count(line, pair) > 1 {
			return true
		}
	}
	return false
}

func containsMirrors(line []byte) bool {
	for i := 2; i < len(line); i++ {
		if line[i-2] == line[i] {
			return true
		}
	}
	return false
}
