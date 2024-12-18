package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"image"
	"strconv"

	"github.com/erka/aoc/pkg/log"
	"gonum.org/v1/gonum/graph/path"
	"gonum.org/v1/gonum/graph/simple"
)

//go:embed input.txt
var input []byte

/*
* part 1: 302
* part 2: 24,32
 */
func main() {
	log.Infof("part1: %s", solvePart1(input, 70, 1024))
	log.Infof("part2: %s", solvePart2(input, 70))
}

var (
	north      = image.Pt(0, -1)
	east       = image.Pt(1, 0)
	south      = image.Pt(0, 1)
	west       = image.Pt(-1, 0)
	directions = []image.Point{north, east, south, west}
)

func nodeId(p image.Point) int64 {
	return int64(p.Y*1000 + p.X)
}

// solve
func solvePart1(input []byte, maxXY int, limit int) string {
	lines := bytes.Split(bytes.Trim(input, "\n"), []byte{'\n'})
	corrupted := []image.Point{}
	for _, line := range lines {
		var p image.Point
		fmt.Sscanf(string(line), "%d,%d", &p.X, &p.Y)
		corrupted = append(corrupted, p)
	}
	area := image.Rect(0, 0, maxXY+1, maxXY+1)
	corrupted = corrupted[:limit]
	log.Debug(corrupted)
	g := simple.NewUndirectedGraph()
	for x := 0; x <= maxXY; x++ {
		for y := 0; y <= maxXY; y++ {
			p := image.Pt(x, y)
			for _, d := range directions {
				n := p.Add(d)
				if !n.In(area) {
					continue
				}
				g.SetEdge(g.NewEdge(simple.Node(nodeId(p)), simple.Node(nodeId(n))))
			}
		}
	}
	for i := 0; i < limit; i++ {
		p := corrupted[i]
		g.RemoveNode(nodeId(p))
		pt, _ := path.DijkstraFrom(simple.Node(0), g).To(nodeId(image.Pt(maxXY, maxXY)))
		if len(pt) == 0 {
			return fmt.Sprintf("%d,%d", p.X, p.Y)
		}
	}
	pt, _ := path.DijkstraFrom(simple.Node(0), g).To(nodeId(image.Pt(maxXY, maxXY)))
	return strconv.Itoa(len(pt) - 1)
}

// solve
func solvePart2(input []byte, maxXY int) string {
	lines := bytes.Split(bytes.Trim(input, "\n"), []byte{'\n'})
	corrupted := []image.Point{}
	for _, line := range lines {
		var p image.Point
		fmt.Sscanf(string(line), "%d,%d", &p.X, &p.Y)
		corrupted = append(corrupted, p)
	}
	area := image.Rect(0, 0, maxXY+1, maxXY+1)
	g := simple.NewUndirectedGraph()
	for x := 0; x <= maxXY; x++ {
		for y := 0; y <= maxXY; y++ {
			p := image.Pt(x, y)
			for _, d := range directions {
				n := p.Add(d)
				if !n.In(area) {
					continue
				}
				g.SetEdge(g.NewEdge(simple.Node(nodeId(p)), simple.Node(nodeId(n))))
			}
		}
	}
	for i := 0; i < len(corrupted); i++ {
		p := corrupted[i]
		g.RemoveNode(nodeId(p))
		pt, _ := path.DijkstraFrom(simple.Node(0), g).To(nodeId(image.Pt(maxXY, maxXY)))
		if len(pt) == 0 {
			return fmt.Sprintf("%d,%d", p.X, p.Y)
		}
	}
	return ""
}
