package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"strconv"

	"github.com/erka/aoc/pkg/log"
	"github.com/samber/lo"
)

//go:embed input.txt
var input []byte

/*
* part 1: 631
* part 2: 665
 */
func main() {
	log.Infof("part1: %s", solvePart1(input))
	log.Infof("part2: %s", solvePart2(input))
}

func validAdjacent(x, y int) bool {
	d := max(x, y) - min(x, y)
	return d > 0 && d < 4
}

// ugly but works
func safe(in string, tolerateLevel int) bool {
	report := lo.Map(lo.Words(in), func(item string, index int) int {
		i, _ := strconv.Atoi(item)
		return i
	})

	return safeReport(report, tolerateLevel) || safeReport(lo.Reverse(report), tolerateLevel)
}

func safeReport(report []int, tolerateLevel int) bool {
	return checkSafety(report, tolerateLevel, asc) || checkSafety(report, tolerateLevel, desc)
}

func checkSafety(ints []int, tolerateLevel int, fn func(x, y int) bool) bool {
	currentLevel := 0
	l := ints[0]
	for i := 1; i < len(ints); i++ {
		c := ints[i]
		if !validAdjacent(l, c) || fn(l, c) {
			currentLevel += 1
			continue
		}
		l = c
	}
	return currentLevel <= tolerateLevel
}

func asc(x, y int) bool {
	return x < y
}

func desc(x, y int) bool {
	return x > y
}

// solve
func solvePart1(input []byte) string {
	scanner := bufio.NewScanner(bytes.NewBuffer(input))
	output := 0
	for scanner.Scan() {
		line := scanner.Text()
		log.Debug(line)
		if safe(line, 0) {
			output += 1
		}
	}
	if err := scanner.Err(); err != nil {
		log.Errorf("failed to read input: %v", err)
		return ""
	}
	return strconv.Itoa(output)
}

// solve
func solvePart2(input []byte) string {
	scanner := bufio.NewScanner(bytes.NewBuffer(input))
	output := 0
	for scanner.Scan() {
		line := scanner.Text()
		log.Debug(line)
		if safe(line, 1) {
			output += 1
		}
	}
	if err := scanner.Err(); err != nil {
		log.Errorf("failed to read input: %v", err)
		return ""
	}
	return strconv.Itoa(output)
}
