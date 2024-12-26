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
* part 1: 1333
* part 2: 2046
 */
func main() {
	log.Infof("part1: %s", solvePart1(input))
	log.Infof("part2: %s", solvePart2(input))
}

// solve
func solvePart1(input []byte) string {
	lines := bytes.Split(bytes.Trim(input, "\n"), []byte("\n"))
	output := lo.Sum(lo.Map(lines, func(line []byte, _ int) int {
		l := 0
		for i := 0; i < len(line); i++ {
			switch line[i] {
			case '"':
				continue
			case '\\':
				if line[i+1] == 'x' {
					i += 3
				} else {
					i += 1
				}
				l += 1
			default:
				l += 1
			}
		}
		return len(line) - l
	}))
	return strconv.Itoa(output)
}

// solve
func solvePart2(input []byte) string {
	lines := bytes.Split(bytes.Trim(input, "\n"), []byte("\n"))
	output := lo.Sum(lo.Map(lines, func(line []byte, _ int) int {
		l := 2
		for i := 0; i < len(line); i++ {
			switch line[i] {
			case '"', '\\':
				l += 2
			default:
				l += 1
			}
		}
		return l - len(line)
	}))
	return strconv.Itoa(output)
}
