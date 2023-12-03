package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"strconv"
)

//go:embed input.txt
var input []byte

type point struct {
	x, y int
}

/*
* part 1: 532445
* part 2: 79842967
 */
func main() {
	part1, part2 := solve(input)
	fmt.Println("part1: ", part1)
	fmt.Println("part2: ", part2)
}

func solveGearRatios(ratios map[point][]int) string {
	sum := 0
	for _, v := range ratios {
		if len(v) == 1 {
			continue
		}
		m := 1
		for _, i := range v {
			m *= i
		}
		sum += m
	}
	return strconv.Itoa(sum)
}

func isAdjacent(lines [][]byte, l, s, e int) (point, bool) {
	for x := max(0, l-1); x < min(l+2, len(lines)); x++ {
		for y := max(0, s-1); y < min(e+2, len(lines[l])); y++ {
			c := lines[x][y]
			if c != '.' && (c < '0' || c > '9') {
				return point{x: x, y: y}, true
			}
		}
	}
	return point{}, false
}

// solveEngineSchematic
func solve(input []byte) (string, string) {
	var ratios = map[point][]int{}

	lines := bytes.Split(input, []byte{'\n'})
	sum := 0
	s := -1
	e := -1
	n := 0
	for row, line := range lines {
		for col, c := range line {
			switch c {
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				if s == -1 {
					s = col
				}
				e = col
				n = n*10 + int(c-'0')
				if col == len(line)-1 {
					// end of a row
					if p, ok := isAdjacent(lines, row, s, e); ok {
						sum += n
						if lines[p.x][p.y] == '*' {
							ratios[p] = append(ratios[p], n)
						}

					}
					s = -1
					e = -1
					n = 0
				}
			default:
				if n == 0 {
					continue
				}
				if p, ok := isAdjacent(lines, row, s, e); ok {
					sum += n
					if lines[p.x][p.y] == '*' {
						ratios[p] = append(ratios[p], n)
					}
				}
				s = -1
				e = -1
				n = 0

			}
		}
	}
	return strconv.Itoa(sum), solveGearRatios(ratios)
}
