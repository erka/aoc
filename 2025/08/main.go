package main

import (
	"cmp"
	_ "embed"
	"log/slog"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/erka/aoc/pkg/iox"
	"github.com/erka/aoc/pkg/log"
	_ "github.com/erka/aoc/pkg/xslog"
	"github.com/samber/lo"
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/encoding/dot"
	"gonum.org/v1/gonum/graph/path"
	"gonum.org/v1/gonum/graph/simple"
	"gonum.org/v1/gonum/graph/topo"
)

//go:embed input.txt
var input []byte

/*
* part 1: 50568
* part 2: 36045012
 */
func main() {
	slog.Info("part1", slog.String("value", solvePart1(input, 1000)))
	slog.Info("part2", slog.String("value", solvePart2(input)))
}

// solve
func solvePart1(input []byte, limit int) string {
	g, _, edges := buildGraph(input)

	k := 0
	for _, e := range edges {
		shortest := path.DijkstraFrom(e.F, g)
		nodes, _ := shortest.To(e.T.ID())
		k += 1
		if len(nodes) == 0 {
			g.SetWeightedEdge(e)
			if k == limit {
				break
			}
		}
	}

	components := topo.ConnectedComponents(g)

	slices.SortFunc(components, func(a []graph.Node, b []graph.Node) int {
		return -cmp.Compare(len(a), len(b))
	})
	values := (lo.Map(components, func(item []graph.Node, index int) int {
		return len(item)
	}))

	dotData, err := dot.Marshal(g, "Graph", "", "")
	if err != nil {
		log.Errorf("failed to marshal graph: %v", err)
	}

	_ = os.WriteFile("graph.dot", dotData, 0o644)
	return strconv.Itoa(values[0] * values[1] * values[2])
}

// solve
func solvePart2(input []byte) string {
	g, points, edges := buildGraph(input)

	for _, e := range edges {
		shortest := path.DijkstraFrom(e.F, g)
		nodes, _ := shortest.To(e.T.ID())
		if len(nodes) == 0 {
			g.SetWeightedEdge(e)
			components := topo.ConnectedComponents(g)
			if len(components) == 1 && len(components[0]) == len(points) {
				return strconv.Itoa((e.F.(point)).x * (e.T.(point)).x)
			}
		}
	}

	return strconv.Itoa(0)
}

func buildGraph(input []byte) (*simple.WeightedUndirectedGraph, []point, []simple.WeightedEdge) {
	points := []point{}
	for line := range iox.Lines(input) {
		xyz := lo.Map(strings.SplitN(line, ",", 3), func(s string, _ int) int {
			n, _ := strconv.Atoi(s)
			return n
		})
		points = append(points, newPt(xyz[0], xyz[1], xyz[2]))
	}

	g := simple.NewWeightedUndirectedGraph(0, 0)
	edges := []simple.WeightedEdge{}
	for i, p := range points[:len(points)-1] {
		for k := i; k < len(points); k++ {
			if k == i {
				continue
			}
			d := p.distanceTo(points[k])
			edges = append(edges, simple.WeightedEdge{F: p, T: points[k], W: d})
		}
	}

	slices.SortFunc(edges, func(a, b simple.WeightedEdge) int {
		return cmp.Compare(a.W, b.W)
	})
	return g, points, edges
}

type point struct {
	x, y, z, id int
}

func (p point) distanceTo(o point) float64 {
	return math.Sqrt(float64((o.x-p.x)*(o.x-p.x) + (o.y-p.y)*(o.y-p.y) + (o.z-p.z)*(o.z-p.z)))
}

func (p point) ID() int64 {
	return int64(p.x*10000000 + p.y*1000 + p.z)
}

func newPt(x, y, z int) point {
	return point{
		x:  x,
		y:  y,
		z:  z,
		id: x*10000000 + y*1000 + z,
	}
}

func getEdges(g *simple.WeightedDirectedGraph, n point) []point {
	it := g.From(n.ID())
	to := []point{}
	for it.Next() {
		p := it.Node().(point)
		if g.HasEdgeBetween(n.ID(), p.ID()) {
			to = append(to, p)
		}
	}
	return to
}

func Intersect(a, b []point) []point {
	set := make(map[point]struct{}, len(a))
	for _, v := range a {
		set[v] = struct{}{}
	}

	var out []point
	for _, v := range b {
		if _, ok := set[v]; ok {
			out = append(out, v)
		}
	}
	return out
}
