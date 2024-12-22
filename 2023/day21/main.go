package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"image"
	"strconv"

	"github.com/erka/aoc/pkg/log"
	"github.com/samber/lo"
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

var directions = []image.Point{
	image.Pt(0, -1), // go north
	image.Pt(1, 0),  // go east
	image.Pt(0, 1),  // go south
	image.Pt(-1, 0), // go west
}

// solve
func solve(input []byte, steps int, infinite bool) int {
	scanner := bufio.NewScanner(bytes.NewBuffer(input))
	grid := map[image.Point]rune{}
	start := image.Pt(0, 0)
	rows, cols := 0, 0

	for y := 0; scanner.Scan(); y += 1 {
		rows += 1
		line := scanner.Text()
		cols = len(line)
		for x, w := range line {
			grid[image.Pt(x, y)] = w
			if w == 'S' {
				start = image.Pt(x, y)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Errorf("failed to read input: %v", err)
		return 0
	}
	log.Debugf("cols: %v rows: %v", cols, rows)
	points := []image.Point{
		start,
	}
	cache := map[image.Point][]image.Point{}

	for i := 0; i < steps; i++ {
		if len(points) == 0 {
			log.Error("no points... this should not happen.")
			break
		}
		next := []image.Point{}
		for _, p := range points {
			if cached, ok := cache[p]; ok {
				next = append(next, cached...)
				continue
			}
			values := []image.Point{}
			for _, d := range directions {
				n := p.Add(d)
				x := n
				if infinite {
					x = image.Pt((n.X%cols+cols)%cols, (n.Y%rows+rows)%rows)
				}
				if w, ok := grid[x]; ok {
					if w == '.' || w == 'S' {
						values = append(values, n)
					}
				}
			}
			values = lo.Uniq(values)
			cache[p] = values
			next = append(next, values...)
		}
		points = lo.Uniq(next)
	}

	log.Debug(points)

	return len(points)
}

func solvePart1(input []byte) string {
	return strconv.Itoa(solve(input, 64, false))
}

// solve
func solvePart2(input []byte) string {
	return strconv.Itoa(0)
}
