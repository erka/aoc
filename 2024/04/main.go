package main

import (
	"bytes"
	_ "embed"
	"image"
	"strconv"

	"github.com/erka/aoc/pkg/log"
)

//go:embed input.txt
var input []byte

/*
* part 1: 2547
* part 2: 1939
 */
func main() {
	log.Infof("part1: %s", solvePart1(input))
	log.Infof("part2: %s", solvePart2(input))
}

var allDirections = []image.Point{
	image.Pt(0, -1), // go north
	image.Pt(1, -1),
	image.Pt(-1, -1),
	image.Pt(1, 0), // go east
	image.Pt(0, 1), // go south
	image.Pt(1, 1),
	image.Pt(-1, 1),
	image.Pt(-1, 0), // go west
}

var xDirections = []image.Point{
	image.Pt(-1, -1),
	image.Pt(1, -1),
	image.Pt(-1, 1),
	image.Pt(1, 1),
}

var word = [3]byte{'M', 'A', 'S'}

func countXmas(puzzle [][]byte, start func(image.Point) image.Point, directions []image.Point) int {
	count := 0
	area := image.Rect(0, 0, len(puzzle), len(puzzle[0]))
	for _, direction := range directions {
		found := true
		p := start(direction)
		for step := 0; step < len(word); step++ {
			if !p.In(area) || puzzle[p.X][p.Y] != word[step] {
				found = false
				break
			}
			p = p.Add(direction)
		}
		if found {
			count += 1
		}
	}
	return count
}

// solve
func solvePart1(input []byte) string {
	output := 0
	puzzle := bytes.Split(input, []byte("\n"))
	puzzle = puzzle[:len(puzzle)-1]
	for y, line := range puzzle {
		for x, sym := range line {
			if sym == 'X' {
				start := image.Pt(y, x)
				log.Debugf("staring point %v", start)
				output += countXmas(puzzle, func(dir image.Point) image.Point {
					return start.Add(dir) // move forward
				}, allDirections)
			}
		}
	}

	return strconv.Itoa(output)
}

// solve
func solvePart2(input []byte) string {
	output := 0
	puzzle := bytes.Split(input, []byte("\n"))
	puzzle = puzzle[:len(puzzle)-1]
	for y, line := range puzzle {
		for x, sym := range line {
			if sym == 'A' {
				p := image.Pt(y, x)
				log.Debugf("staring point %v", p)
				c := countXmas(puzzle, func(dir image.Point) image.Point {
					return p.Sub(dir) // move opposite
				}, xDirections)
				if c == 2 {
					output += 1
				}
			}
		}
	}

	return strconv.Itoa(output)
}
