package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"strconv"

	"github.com/erka/aoc/pkg/mathx"
)

//go:embed input.txt
var input []byte

type node [2]string

/*
 * part1: 20659
 * part2: 15690466351717
 */
func main() {
	fmt.Printf("part1: %s\n", solvePart1(input))
	fmt.Printf("part2: %s\n", solvePart2(input))
}

// solve
func solvePart1(input []byte) string {
	scanner := bufio.NewScanner(bytes.NewBuffer(input))
	scanner.Scan()
	navigation := scanner.Text()
	// empty line
	scanner.Scan()
	network := map[string]node{}
	for scanner.Scan() {
		line := scanner.Text()
		network[line[:3]] = node{line[7:10], line[12:15]}
	}
	if err := scanner.Err(); err != nil {
		return err.Error()
	}
	idx := "AAA"
	steps := 0
	for i := 0; ; i = (i + 1) % len(navigation) {
		node := network[idx]
		switch navigation[i] {
		case 'L':
			idx = node[0]
		case 'R':
			idx = node[1]
		}

		steps += 1
		if idx == "ZZZ" {
			break
		}
	}
	return strconv.Itoa(steps)
}

// solve
func solvePart2(input []byte) string {
	scanner := bufio.NewScanner(bytes.NewBuffer(input))
	scanner.Scan()
	navigation := scanner.Text()
	// empty line
	scanner.Scan()
	network := map[string]node{}
	idx := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		label := line[:3]
		network[label] = node{line[7:10], line[12:15]}
		if label[2] == 'A' {
			idx = append(idx, label)
		}
	}
	if err := scanner.Err(); err != nil {
		return err.Error()
	}

	cnt := len(idx)
	steps := make([]int, cnt)
	for j := 0; j < cnt; j++ {
		steps[j] = find(navigation, network, idx[j])
	}
	return strconv.Itoa(mathx.LCM(steps[0], steps[1], steps[2:]...))
}

func find(navigation string, network map[string]node, idx string) int {
	steps := 0
	for i := 0; ; i = (i + 1) % len(navigation) {
		node := network[idx]
		switch navigation[i] {
		case 'L':
			idx = node[0]
		case 'R':
			idx = node[1]
		}

		steps += 1
		if idx[2] == 'Z' {
			break
		}
	}
	return steps
}
