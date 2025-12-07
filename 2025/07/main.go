package main

import (
	_ "embed"
	"image"
	"log/slog"
	"slices"
	"strconv"
	"strings"

	"github.com/erka/aoc/pkg/iox"
	_ "github.com/erka/aoc/pkg/xslog"
	"github.com/samber/lo"
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/simple"
)

//go:embed input.txt
var input []byte

/*
* part 1: 1570
* part 2: 15118009521693
 */
func main() {
	slog.Info("part1", slog.String("value", solvePart1(input)))
	slog.Info("part2", slog.String("value", solvePart2(input)))
}

// solve
func solvePart1(input []byte) string {
	_, _, split := buildPaths(input)
	return strconv.Itoa(split)
}

// solve
func solvePart2(input []byte) string {
	g, beams, _ := buildPaths(input)

	from := nodeID(beams[0][0])
	ends := lo.Map(beams[len(beams)-1], func(p image.Point, index int) int64 {
		return nodeID(p)
	})
	var timelines int64
	for _, to := range ends {
		timelines += countPaths(g, from, to)
	}
	return strconv.FormatInt(timelines, 10)
}

func buildPaths(input []byte) (*simple.DirectedGraph, map[int][]image.Point, int) {
	var y, split int

	g := simple.NewDirectedGraph()
	beams := map[int][]image.Point{}
	for line := range iox.Lines(input) {
		if y == 0 {
			x := strings.IndexByte(line, 'S')
			p := image.Pt(x, y)
			beams[y] = append(beams[y], p)
			n, ok := g.NodeWithID(nodeID(p))
			if ok {
				g.AddNode(n)
			}
			y += 1
			continue
		}
		for x, ch := range line {
			if slices.Contains(beams[y-1], image.Pt(x, y-1)) {
				switch ch {
				case '.':
					p := image.Pt(x, y)
					beams[y] = append(beams[y], p)
					n, ok := g.NodeWithID(nodeID(p))
					if ok {
						g.AddNode(n)
					}
					f, _ := g.NodeWithID(nodeID(image.Pt(x, y-1)))
					e := g.NewEdge(f, n)
					g.SetEdge(e)
				case '^':
					p1, p2 := image.Pt(x-1, y), image.Pt(x+1, y)
					beams[y] = slices.Compact(append(beams[y], p1, p2))
					f, _ := g.NodeWithID(nodeID(image.Pt(x, y-1)))
					for _, p := range []image.Point{p1, p2} {
						n, ok := g.NodeWithID(nodeID(p))
						if ok {
							g.AddNode(n)
						}
						e := g.NewEdge(f, n)
						g.SetEdge(e)
					}
					split += 1
				}
			}
		}

		y += 1
	}
	return g, beams, split
}

func nodeID(p image.Point) int64 {
	return int64(p.Y*1000 + p.X)
}

// Topological sort (Kahn's algorithm)
func countPaths(g graph.Directed, fromID, toID int64) int64 {
	inDeg := make(map[int64]int)
	nodes := g.Nodes()
	for nodes.Next() {
		u := nodes.Node().ID()
		inDeg[u] = 0
	}
	nodes.Reset()
	for nodes.Next() {
		u := nodes.Node().ID()
		it := g.From(u)
		for it.Next() {
			v := it.Node().ID()
			inDeg[v]++
		}
	}

	queue := make([]int64, 0)
	for id, d := range inDeg {
		if d == 0 {
			queue = append(queue, id)
		}
	}

	topo := make([]int64, 0, len(inDeg))
	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		topo = append(topo, u)

		it := g.From(u)
		for it.Next() {
			v := it.Node().ID()
			inDeg[v]--
			if inDeg[v] == 0 {
				queue = append(queue, v)
			}
		}
	}

	// DP counting
	ways := make(map[int64]int64, len(inDeg))
	ways[fromID] = 1

	for _, u := range topo {
		it := g.From(u)
		for it.Next() {
			v := it.Node().ID()
			ways[v] += ways[u]
		}
	}

	return ways[toID]
}
