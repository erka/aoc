package main

import (
	"bytes"
	_ "embed"
	"strconv"
	"strings"

	"github.com/erka/aoc/pkg/log"
	"github.com/samber/lo"
)

//go:embed input.txt
var input []byte

/*
* part 1: 296
* part 2: 619970556776002
 */
func main() {
	log.Infof("part1: %s", solvePart1(input))
	log.Infof("part2: %s", solvePart2(input))
}

var seen = make(map[string]int)

func possibleDesigns(d string, patterns []string, seen map[string]int) int {
	if n, ok := seen[d]; ok {
		return n
	}
	if len(d) == 0 {
		return 1
	}
	n := 0
	for _, p := range patterns {
		if strings.HasPrefix(d, p) {
			x := possibleDesigns(d[len(p):], patterns, seen)
			n += x
		}
	}
	seen[d] = n
	return n
}

// solve
func solvePart1(input []byte) string {
	data := bytes.Split(bytes.TrimSpace(input), []byte{'\n', '\n'})

	towelPatterns := lo.Map(bytes.Split(data[0], []byte{',', ' '}), func(b []byte, _ int) string { return string(b) })
	desiredDesigns := lo.Map(bytes.Split(data[1], []byte{'\n'}), func(b []byte, _ int) string { return string(b) })
	output := 0
	for _, d := range desiredDesigns {
		n := possibleDesigns(d, towelPatterns, seen)
		if n > 0 {
			output += 1
		}
	}
	return strconv.Itoa(output)
}

// solve
func solvePart2(input []byte) string {
	data := bytes.Split(bytes.TrimSpace(input), []byte{'\n', '\n'})

	towelPatterns := lo.Map(bytes.Split(data[0], []byte{',', ' '}), func(b []byte, _ int) string { return string(b) })
	desiredDesigns := lo.Map(bytes.Split(data[1], []byte{'\n'}), func(b []byte, _ int) string { return string(b) })
	output := 0
	for _, d := range desiredDesigns {
		n := possibleDesigns(d, towelPatterns, seen)
		if n > 0 {
			output += n
		}
	}
	return strconv.Itoa(output)
}
