package main

import (
	"bytes"
	_ "embed"
	"image"
	"strconv"

	"github.com/erka/aoc/pkg/log"
	"github.com/samber/lo"
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/path"
	"gonum.org/v1/gonum/graph/simple"
)

//go:embed input.txt
var input []byte

/*
* part 1: 1422
* part 2: 1009299
 */
func main() {
	log.Infof("part1: %s", solve(input, 100, 2))
	log.Infof("part2: %s", solve(input, 100, 20))
}

var (
	north      = image.Pt(0, -1)
	east       = image.Pt(1, 0)
	south      = image.Pt(0, 1)
	west       = image.Pt(-1, 0)
	directions = []image.Point{north, east, south, west}
)

// solve
func solve(input []byte, minSavedPicoseconds int, cheatAllowed int) string {
	lines := bytes.FieldsFunc(input, func(r rune) bool { return r == '\n' })
	start := image.Point{}
	end := image.Point{}
	places := map[image.Point]byte{}
	for y, line := range lines {
		for x, c := range line {
			if c == '#' {
				continue
			}
			p := image.Pt(x, y)
			places[p] = c
			switch c {
			case 'S':
				start = p
			case 'E':
				end = p
			}
		}
	}
	g := simple.NewUndirectedGraph()
	for p := range places {
		for _, d := range directions {
			n := p.Add(d)
			if _, ok := places[n]; !ok {
				continue
			}
			g.SetEdge(g.NewEdge(simple.Node(nodeId(p)), simple.Node(nodeId(n))))
		}
	}
	pt, _ := path.DijkstraFrom(simple.Node(nodeId(start)), g).To(nodeId(end))
	pathIds := lo.Map(pt, func(n graph.Node, _ int) image.Point { return nodePoint(n.ID()) })

	output := 0
	for i := 0; i < len(pathIds)-minSavedPicoseconds; i++ {
		p := pathIds[i]
		for j := i + minSavedPicoseconds; j < len(pathIds); j++ {
			e := pathIds[j]
			distance := abs(e.X-p.X) + abs(e.Y-p.Y)
			if distance <= cheatAllowed && j-i >= distance+minSavedPicoseconds {
				output++
			}
		}
	}
	return strconv.Itoa(output)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func nodeId(p image.Point) int64 {
	return int64(p.Y*1000 + p.X)
}

func nodePoint(n int64) image.Point {
	return image.Pt(int(n%1000), int(n/1000))
}
