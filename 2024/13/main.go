package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"strconv"

	"github.com/erka/aoc/pkg/log"
	"github.com/samber/lo"
)

//go:embed input.txt
var input []byte

/*
* part 1: 31761
* part 2: 90798500745591
 */
func main() {
	log.Infof("part1: %s", solvePart1(input))
	log.Infof("part2: %s", solvePart2(input))
}

// solve
func solvePart1(input []byte) string {
	return solve(input, 0)
}

// solve
func solvePart2(input []byte) string {
	return solve(input, 10000000000000)
}

/*
* math solution
* 94*a + 22*b = 8400
* 34a + 67*b = 5400
*
* a = (5400 - 67*b)/ 34
* 94*5400 - 94*67*b + 22*34*b = 8400 * 34
* 222000 = 5550*b
 */
var template = `Button A: X+%d, Y+%d
Button B: X+%d, Y+%d
Prize: X=%d, Y=%d`

// solve
func solve(input []byte, resultCorrection int) string {
	prizes := lo.Map(bytes.Split(input, []byte("\n\n")), func(item []byte, _ int) int {
		var a1, b1, a2, b2, ra, rb int
		_, err := fmt.Sscanf(string(item), template, &a1, &b1, &a2, &b2, &ra, &rb)
		if err != nil {
			panic(err)
		}
		// adjust results
		ra += resultCorrection
		rb += resultCorrection

		// solve equations
		y := (ra*b1 - a1*rb) / (b1*a2 - a1*b2)
		x := (rb - b2*y) / b1

		// check if solution is correct
		if a1*x+a2*y == ra && b1*x+b2*y == rb {
			return x*3 + y
		}

		return 0
	})

	output := lo.Sum(prizes)
	return strconv.Itoa(output)
}
