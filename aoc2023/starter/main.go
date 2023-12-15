package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"strconv"
)

//go:embed input.txt
var input []byte

/*
* part 1:
* part 2:
 */
func main() {
	fmt.Printf("part1: %s\n", solvePart1(input))
	fmt.Printf("part2: %s\n", solvePart2(input))
}

// solve
func solvePart1(input []byte) string {
	scanner := bufio.NewScanner(bytes.NewBuffer(input))

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
	if err := scanner.Err(); err != nil {
		return err.Error()
	}
	return strconv.Itoa(0)
}

// solve
func solvePart2(input []byte) string {
	scanner := bufio.NewScanner(bytes.NewBuffer(input))
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
	if err := scanner.Err(); err != nil {
		return err.Error()
	}
	return strconv.Itoa(0)
}
