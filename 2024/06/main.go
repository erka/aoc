package main

import (
	"bytes"
	_ "embed"
	"image"
	"strconv"

	"github.com/erka/aoc/pkg/log"
	"gonum.org/v1/gonum/graph/simple"
)

//go:embed input.txt
var input []byte

/*
* part 1: 5409
* part 2: 2022
 */
func main() {
	log.Infof("part1: %s", solvePart1(input))
	log.Infof("part2: %s", solvePart2(input))
}

var (
	north = image.Pt(0, -1)
	east  = image.Pt(1, 0)
	south = image.Pt(0, 1)
	west  = image.Pt(-1, 0)
)

func rightDir(dir image.Point) image.Point {
	switch dir {
	case north:
		return east
	case east:
		return south
	case south:
		return west
	case west:
		return north
	}
	return north
}

type Guard struct {
	dir image.Point
	pos image.Point
}

func nodeId(p image.Point) int64 {
	return int64(p.Y*1000 + p.X)
}

func (g *Guard) set(pos image.Point, dir image.Point) {
	g.pos.X = pos.X
	g.pos.Y = pos.Y
	g.dir = dir // go north
}

func (g *Guard) next() image.Point {
	return g.pos.Add(g.dir)
}

func (g *Guard) move(area [][]byte) bool {
	next := g.next()
	if area[next.Y][next.X] == '#' {
		g.turnRight()
		return false
	}
	g.pos = next
	return true
}

func (g *Guard) turnRight() {
	g.dir = rightDir(g.dir)
}

// solve
func solvePart1(input []byte) string {
	area := bytes.FieldsFunc(input, func(r rune) bool { return r == '\n' })

	boundary := image.Rect(0, 0, len(area), len(area[0]))

	start := image.Pt(0, 0)
	for y, line := range area {
		for x, s := range line {
			if s == '^' {
				start = image.Pt(x, y)
			}
		}
	}

	graph := simple.NewDirectedGraph()
	guard := &Guard{}
	guard.set(start, north)
	nextPos := guard.next()
	for nextPos.In(boundary) {
		prevPos := guard.pos
		if guard.move(area) {
			graph.SetEdge(graph.NewEdge(simple.Node(nodeId(prevPos)), simple.Node(nodeId(guard.pos))))
			nextPos = guard.next()
		}
	}
	return strconv.Itoa(graph.Nodes().Len())
}

// solve
func solvePart2(input []byte) string {
	area := bytes.FieldsFunc(input, func(r rune) bool { return r == '\n' })

	boundary := image.Rect(0, 0, len(area), len(area[0]))

	start := image.Pt(0, 0)

	for y, line := range area {
		for x, s := range line {
			if s == '^' {
				start = image.Pt(x, y)
			}
		}
	}

	loops := map[image.Point]struct{}{}
	graph := simple.NewDirectedGraph()
	guard := &Guard{}
	guard.set(start, north)
	nextPos := guard.next()
	for nextPos.In(boundary) {
		prevPos := guard.pos
		if guard.move(area) {
			graph.SetEdge(graph.NewEdge(simple.Node(nodeId(prevPos)), simple.Node(nodeId(guard.pos))))
		}
		nextPos = guard.next()
		if !nextPos.In(boundary) {
			break
		}

		if area[nextPos.Y][nextPos.X] == '#' || area[nextPos.Y][nextPos.X] == '^' {
			// skip if there is an obstacle or original start
			continue
		}
		area[nextPos.Y][nextPos.X] = '#'
		if detectLoop(start, boundary, area) {
			loops[nextPos] = struct{}{}
		}
		area[nextPos.Y][nextPos.X] = '.'
	}
	return strconv.Itoa(len(loops))
}

func detectLoop(start image.Point, boundary image.Rectangle, area [][]byte) bool {
	graph := simple.NewDirectedGraph()
	guard := &Guard{}
	guard.set(start, north)
	nextPos := guard.next()
	for nextPos.In(boundary) {
		prevPos := guard.pos
		if guard.move(area) {
			graph.SetEdge(graph.NewEdge(simple.Node(nodeId(prevPos)), simple.Node(nodeId(guard.pos))))
		}
		nextPos = guard.next()
		if graph.HasEdgeFromTo(nodeId(guard.pos), nodeId(nextPos)) {
			return true
		}
	}
	return false
}
