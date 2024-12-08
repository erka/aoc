package main

import (
	"bytes"
	_ "embed"
	"image"
	"maps"
	"strconv"

	"github.com/erka/aoc/pkg/log"
)

//go:embed input.txt
var input []byte

/*
* part 1: 244
* part 2: 912
 */
func main() {
	log.Infof("part1: %s", solvePart1(input))
	log.Infof("part2: %s", solvePart2(input))
}

func isAnthena(c byte) bool {
	switch {
	case c >= '0' && c <= '9':
		return true
	case c >= 'A' && c <= 'Z':
		return true
	case c >= 'a' && c <= 'z':
		return true
	default:
		return false
	}
}

// solve
func solvePart1(input []byte) string {
	anthenas, area := parse(input)

	antinodes := make(map[image.Point]struct{})
	for locations := range maps.Values(anthenas) {
		for i := 0; i < len(locations)-1; i++ {
			for j := i + 1; j < len(locations); j++ {
				diff := locations[j].Sub(locations[i])
				loc := locations[i].Sub(diff)
				if loc.In(area) {
					antinodes[loc] = struct{}{}
				}
				loc = locations[j].Add(diff)
				if loc.In(area) {
					antinodes[loc] = struct{}{}
				}
			}
		}
	}

	return strconv.Itoa(len(antinodes))
}

func parse(input []byte) (map[byte][]image.Point, image.Rectangle) {
	lines := bytes.FieldsFunc(input, func(r rune) bool { return r == '\n' })
	anthenas := make(map[byte][]image.Point)
	for y, line := range lines {
		for x, c := range line {
			if isAnthena(c) {
				if _, ok := anthenas[c]; !ok {
					anthenas[c] = []image.Point{}
				}
				anthenas[c] = append(anthenas[c], image.Pt(y, x))
			}
		}
	}

	area := image.Rect(0, 0, len(lines[0]), len(lines))
	return anthenas, area
}

// solve
func solvePart2(input []byte) string {
	anthenas, area := parse(input)

	antinodes := make(map[image.Point]struct{})
	for locations := range maps.Values(anthenas) {
		for i := 0; i < len(locations)-1; i++ {
			for j := i + 1; j < len(locations); j++ {
				antinodes[locations[i]] = struct{}{}
				antinodes[locations[j]] = struct{}{}
				diff := locations[j].Sub(locations[i])
				loc := locations[i].Sub(diff)
				for loc.In(area) {
					antinodes[loc] = struct{}{}
					loc = loc.Sub(diff)
				}
				loc = locations[j].Add(diff)
				for loc.In(area) {
					antinodes[loc] = struct{}{}
					loc = loc.Add(diff)
				}
			}
		}
	}

	return strconv.Itoa(len(antinodes))
}
