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
* part 2:
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

type Guard struct {
	dir  image.Point
	pos  image.Point
	path *simple.DirectedGraph
}

func vertexID(p image.Point) int64 {
	return int64(p.Y*1000 + p.X)
}

func (g *Guard) set(pos image.Point, dir image.Point) {
	g.pos.X = pos.X
	g.pos.Y = pos.Y
	g.path.AddNode(simple.Node(vertexID(g.pos)))
	g.dir = dir // go north
}

func (g *Guard) move(area [][]byte, boundary image.Rectangle) bool {
	next := g.pos.Add(g.dir)
	if !next.In(boundary) {
		return false
	}

	if area[next.Y][next.X] == '#' {
		g.turnRight()
		return true
	}
	b := g.path.Node(vertexID(g.pos))
	n := simple.Node(vertexID(next))
	if g.path.Node(n.ID()) == nil {
		g.path.AddNode(n)
	}
	if g.path.Edge(b.ID(), n.ID()) == nil {
		g.path.NewEdge(b, n)
	}
	g.pos = next
	return true
}

func (g *Guard) turnRight() {
	switch g.dir {
	case north:
		g.dir = east
	case east:
		g.dir = south
	case south:
		g.dir = west
	case west:
		g.dir = north
	}
}

func (g *Guard) walk(area [][]byte, boundary image.Rectangle) {
	for g.move(area, boundary) {
	}
}

// solve
func solvePart1(input []byte) string {
	area := bytes.FieldsFunc(input, func(r rune) bool { return r == '\n' })

	boundary := image.Rect(0, 0, len(area), len(area[0]))

	guard := &Guard{
		path: simple.NewDirectedGraph(),
	}

	for y, line := range area {
		for x, s := range line {
			if s == '^' {
				guard.set(image.Pt(x, y), north)
			}
		}
	}

	guard.walk(area, boundary)

	return strconv.Itoa(guard.path.Nodes().Len())
}

// solve
func solvePart2(input []byte) string {
	area := bytes.FieldsFunc(input, func(r rune) bool { return r == '\n' })

	boundary := image.Rect(0, 0, len(area), len(area[0]))

	guard := &Guard{
		path: simple.NewDirectedGraph(),
	}

	for y, line := range area {
		for x, s := range line {
			if s == '^' {
				guard.set(image.Pt(x, y), north)
			}
		}
	}
	guard.walk(area, boundary)

	nodes := guard.path.Nodes()
	for nodes.Next() {
		node := nodes.Node()
		y, x := node.ID()/1000, node.ID()%1000
		log.Debug(node)
	}

	return strconv.Itoa(0)
}
