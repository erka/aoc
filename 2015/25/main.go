package main

import (
	_ "embed"
	"fmt"
	"strconv"

	"github.com/erka/aoc/pkg/log"
)

//go:embed input.txt
var input []byte

/*
* part 1: 8997277
 */
func main() {
	log.Infof("part1: %s", solvePart1(input))
}

// solve
func solvePart1(input []byte) string {
	var row, col int
	fmt.Sscanf(string(input), "Enter the code at row %d, column %d.", &row, &col)
	index := (row+col-1)*(row+col)/2 - row + 1
	num := 20151125
	for i := 1; i < index; i++ {
		num = (num * 252533) % 33554393
	}

	return strconv.Itoa(num)
}
