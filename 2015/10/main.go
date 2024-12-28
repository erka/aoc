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
* part 1: 360154
* part 2: 5103798
 */
func main() {
	log.Infof("part1: %s", solvePart1(input))
	log.Infof("part2: %s", solvePart2(input))
}

// solve
func solvePart1(input []byte) string {
	line := bytes.Trim(input, "\n")
	for i := 0; i < 40; i++ {
		queue := make([]byte, 0, 2*len(line))
		for k := 0; k < len(line); {
			j := k + 1
			for ; j < len(line) && line[j] == line[k]; j++ {
			}
			queue = append(queue, '0'+byte(j-k), line[k])
			k = j
		}
		line = queue
	}
	return strconv.Itoa(len(line))
}

// solve
func solvePart2(input []byte) string {
	line := bytes.Trim(input, "\n")
	for i := 0; i < 50; i++ {
		queue := make([]byte, 0, 2*len(line))
		for k := 0; k < len(line); {
			j := k + 1
			for ; j < len(line) && line[j] == line[k]; j++ {
			}
			queue = append(queue, '0'+byte(j-k), line[k])
			k = j
		}
		line = queue
	}
	return strconv.Itoa(len(line))
}
