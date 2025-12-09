package main

import (
	"cmp"
	_ "embed"
	"fmt"
	"image"
	"log/slog"
	"slices"
	"strconv"

	"github.com/erka/aoc/pkg/iox"
	_ "github.com/erka/aoc/pkg/xslog"
)

//go:embed input.txt
var input []byte

/*
* part 1: 4754955192
* part 2: 1568849600
 */
func main() {
	slog.Info("part1", slog.String("value", solvePart1(input)))
	slog.Info("part2", slog.String("value", solvePart2(input)))
}

// solve
func solvePart1(input []byte) string {
	_, rects := parse(input)
	return strconv.Itoa(rects[0].square())
}

func parse(input []byte) ([]image.Point, []Rectangle) {
	tiles := []image.Point{}
	for line := range iox.Lines(input) {
		tile := image.Point{}
		_, err := fmt.Sscanf(line, "%d,%d", &tile.X, &tile.Y)
		if err != nil {
			panic(err)
		}
		tiles = append(tiles, tile)
	}
	rects := []Rectangle{}
	for i, a := range tiles[:len(tiles)-1] {
		for _, b := range tiles[i+1:] {
			rects = append(rects, rect(a, b))
		}
	}
	slices.SortFunc(rects, func(a Rectangle, b Rectangle) int {
		return -cmp.Compare(a.square(), b.square())
	})
	return tiles, rects
}

// solve
func solvePart2(input []byte) string {
	tiles, rects := parse(input)

	lines := [][2]image.Point{}
	for i := range tiles[:len(tiles)-1] {
		lines = append(lines, [2]image.Point{tiles[i], tiles[i+1]})
	}
	lines = append(lines, [2]image.Point{tiles[len(tiles)-1], tiles[0]})

	m := 0
	for _, r := range rects {
		if isValidRectangle(r, lines) {
			m = max(m, r.square())
		}
	}

	return strconv.Itoa(m)
}

func rect(a, b image.Point) Rectangle {
	minx, miny, maxx, maxy := min(a.X, b.X), min(a.Y, b.Y), max(a.X, b.X), max(a.Y, b.Y)
	return Rectangle{
		Rectangle: image.Rect(minx, miny, maxx, maxy),
		sq:        (maxy - miny + 1) * (maxx - minx + 1),
	}
}

type Rectangle struct {
	image.Rectangle
	sq int
}

func (r Rectangle) square() int {
	return r.sq
}

func (r Rectangle) PointIn(p image.Point) bool {
	return (r.Min.X < p.X && p.X < r.Max.X &&
		r.Min.Y < p.Y && p.Y < r.Max.Y)
}

func isValidRectangle(r Rectangle, lines [][2]image.Point) bool {
	for _, line := range lines {
		from, to := line[0], line[1]

		var delta image.Point

		switch {
		case from.X < to.X:
			delta = image.Pt(1, 0)
		case from.X > to.X:
			delta = image.Pt(-1, 0)
		case from.Y < to.Y:
			delta = image.Pt(0, 1)
		case from.Y > to.Y:
			delta = image.Pt(0, -1)
		}

		tile := from
		for tile.X != to.X || tile.Y != to.Y {
			tile = tile.Add(delta)
			if r.PointIn(tile) {
				return false
			}
		}

	}
	return true
}
