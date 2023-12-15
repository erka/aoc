package main

import (
	"bufio"
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
	scanner := bufio.NewScanner(bytes.NewBuffer(input))

	for scanner.Scan() {
		line := scanner.Text()
		log.Debug(line)
	}
	if err := scanner.Err(); err != nil {
		log.Errorf("failed to read input: %v", err)
		return ""
	}
	return strconv.Itoa(0)
}

// solve
func solvePart2(input []byte) string {
	scanner := bufio.NewScanner(bytes.NewBuffer(input))
	for scanner.Scan() {
		line := scanner.Text()
		log.Debug(line)
	}
	if err := scanner.Err(); err != nil {
		return err.Error()
	}
	return strconv.Itoa(0)
}
