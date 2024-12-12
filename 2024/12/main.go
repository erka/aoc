package main

import (
	"bytes"
	_ "embed"
	"image"
	"slices"
	"strconv"

	"github.com/erka/aoc/pkg/log"
)

//go:embed input.txt
var input []byte

/*
* part 1: 1465112
* part 2: 893790
 */
func main() {
	log.Infof("part1: %s", solvePart1(input))
	log.Infof("part2: %s", solvePart2(input))
}

var (
	north      = image.Pt(0, -1)
	east       = image.Pt(1, 0)
	south      = image.Pt(0, 1)
	west       = image.Pt(-1, 0)
	directions = []image.Point{north, east, south, west}
)

// solve
func solvePart1(input []byte) string {
	farm := bytes.FieldsFunc(input, func(r rune) bool { return r == '\n' })
	plots := make(map[image.Point]byte)
	for y, line := range farm {
		for x, c := range line {
			p := image.Pt(y, x)
			plots[p] = c
		}
	}
	output := 0
	seen := make(map[image.Point]struct{})
	for p, c := range plots {
		if _, ok := seen[p]; ok {
			continue
		}
		log.Debugf("plot: %v %s", p, string(c))
		queue := []image.Point{p}
		area := 0
		perimeter := 0
		for len(queue) > 0 {
			p := queue[0]
			queue = queue[1:]
			seen[p] = struct{}{}

			area += 1
			perimeter += 4

			for _, d := range directions {
				q := p.Add(d)
				if plots[q] == c {
					perimeter -= 1
					if _, ok := seen[q]; !ok && !slices.Contains(queue, q) {
						queue = append(queue, q)
					}
				}
			}
		}
		log.Debugf("%s area: %d, perimeter: %d", string(c), area, perimeter)
		output += area * perimeter
	}

	return strconv.Itoa(output)
}

// solve
func solvePart2(input []byte) string {
	farm := bytes.FieldsFunc(input, func(r rune) bool { return r == '\n' })
	plots := make(map[image.Point]byte)
	for y, line := range farm {
		for x, c := range line {
			p := image.Pt(y, x)
			plots[p] = c
		}
	}
	output := 0
	seen := make(map[image.Point]struct{})
	for p, c := range plots {
		if _, ok := seen[p]; ok {
			continue
		}
		log.Debugf("plot: %v %s", p, string(c))
		queue := []image.Point{p}
		area := 0
		sides := 0
		for len(queue) > 0 {
			p := queue[0]
			queue = queue[1:]
			seen[p] = struct{}{}

			area += 1

			for _, d := range directions {
				q := p.Add(d)
				if plots[q] != c {
					r := p.Add(image.Point{-d.Y, d.X})
					if plots[r] != c || plots[r.Add(d)] == c {
						sides++
					}
				}
				if plots[q] == c {
					if _, ok := seen[q]; !ok && !slices.Contains(queue, q) {
						queue = append(queue, q)
					}
				}
			}
		}
		log.Debugf("%s area: %d, sides: %d", string(c), area, sides)
		output += area * sides
	}

	return strconv.Itoa(output)
}
