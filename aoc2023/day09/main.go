package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input []byte

/*
* part 1:
* part 2:
 */
func main() {
	fmt.Printf("part1: %s\n", solve(input))
	fmt.Printf("part2: %s\n", solvePart2(input))
}

func prediction(values []int) int {
	news := make([]int, len(values)-1)
	zeros := true
	for i := 0; i < len(values)-1; i++ {
		news[i] = values[i+1] - values[i]
		zeros = zeros && news[i] == 0
	}
	if zeros {
		return values[len(values)-1]
	}
	return prediction(news) + values[len(values)-1]
}

// solve
func solve(input []byte) string {
	scanner := bufio.NewScanner(bytes.NewBuffer(input))
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		values := make([]int, len(fields))
		for i := range fields {
			v, err := strconv.Atoi(fields[i])
			if err != nil {
				panic(err)
			}
			values[i] = v
		}

		sum += prediction(values)
	}
	if err := scanner.Err(); err != nil {
		return err.Error()
	}
	return strconv.Itoa(sum)
}

// solve
func solvePart2(input []byte) string {
	scanner := bufio.NewScanner(bytes.NewBuffer(input))
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		values := make([]int, len(fields))
		for i := range fields {
			v, err := strconv.Atoi(fields[i])
			if err != nil {
				panic(err)
			}
			values[i] = v
		}
		slices.Reverse(values)
		sum += prediction(values)
	}
	if err := scanner.Err(); err != nil {
		return err.Error()
	}
	return strconv.Itoa(sum)
}
