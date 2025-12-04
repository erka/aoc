package main

import (
	"bytes"
	_ "embed"
	"image"
	"log/slog"
	"slices"
	"strconv"

	_ "github.com/erka/aoc/pkg/xslog"
)

//go:embed input.txt
var input []byte

/*
* part 1: 1518
* part 2: 8665
 */
func main() {
	slog.Info("part1", slog.String("value", solvePart1(input)))
	slog.Info("part2", slog.String("value", solvePart2(input)))
}

var directions = []image.Point{
	image.Pt(1, 0),
	image.Pt(0, 1),
	image.Pt(-1, 0),
	image.Pt(0, -1),
	image.Pt(1, 1),
	image.Pt(-1, 1),
	image.Pt(-1, -1),
	image.Pt(1, -1),
}

// solve
func solvePart1(input []byte) string {
	lines := bytes.Split(bytes.Trim(input, "\n"), []byte("\n"))
	area := image.Rect(0, 0, len(lines[0]), len(lines))
	accessable := findAccessablePaperRolls(area, lines)
	return strconv.Itoa(len(accessable))
}

func findAccessablePaperRolls(area image.Rectangle, lines [][]byte) []image.Point {
	accessable := []image.Point{}
	for x := 0; x < area.Dx(); x++ {
		for y := 0; y < area.Dy(); y++ {
			if lines[y][x] != '@' {
				continue
			}
			neighborsOn := 0
			pos := image.Pt(x, y)
			for _, d := range directions {
				n := pos.Add(d)
				if !n.In(area) {
					continue
				}
				if lines[n.Y][n.X] == '@' {
					neighborsOn += 1
				}
			}
			if neighborsOn < 4 {
				accessable = append(accessable, pos)
			}

		}
	}
	return accessable
}

// solve
func solvePart2(input []byte) string {
	lines := slices.Clone(bytes.Split(bytes.Trim(input, "\n"), []byte("\n")))
	area := image.Rect(0, 0, len(lines[0]), len(lines))
	total := 0
	for {
		accessable := findAccessablePaperRolls(area, lines)
		if len(accessable) == 0 {
			break
		}
		total += len(accessable)
		for _, a := range accessable {
			lines[a.Y][a.X] = '.'
		}
	}
	return strconv.Itoa(total)
}
