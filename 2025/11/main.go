package main

import (
	_ "embed"
	"hash/crc32"
	"log/slog"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/erka/aoc/pkg/iox"
	"github.com/erka/aoc/pkg/log"
	_ "github.com/erka/aoc/pkg/xslog"
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/encoding"
	"gonum.org/v1/gonum/graph/encoding/dot"
	"gonum.org/v1/gonum/graph/simple"
)

//go:embed input.txt
var input []byte

/*
* part 1: 494
* part 2: 296006754704850
 */
func main() {
	slog.Info("part1", slog.String("value", solvePart1(input)))
	slog.Info("part2", slog.String("value", solvePart2(input)))
}

// solve
func solvePart1(input []byte) string {
	g := buildGraph(input)
	n := countPaths(g, newNode("you").ID(), newNode("out").ID())
	return strconv.FormatInt(n, 10)
}

// solve
func solvePart2(input []byte) string {
	g := buildGraph(input)
	ways := [][]string{{"svr", "dac", "fft", "out"}, {"svr", "fft", "dac", "out"}}
	var n int64 = math.MaxInt64
	for _, way := range ways {
		var steps int64 = 1
		for i := range way[1:] {
			c := countPaths(g, newNode(way[i]).id, newNode(way[i+1]).id)
			steps = steps * c
		}
		if steps != 0 {
			n = min(n, steps)
		}
	}

	dotData, err := dot.Marshal(g, "Graph", "", "")
	if err != nil {
		log.Errorf("failed to marshal graph: %v", err)
	}

	_ = os.WriteFile("graph.dot", dotData, 0o644)
	return strconv.FormatInt(n, 10)
}

func buildGraph(input []byte) *simple.DirectedGraph {
	g := simple.NewDirectedGraph()
	for line := range iox.Lines(input) {
		nodes := strings.Fields(line)
		in := nodes[0]
		in = in[:len(in)-1]
		f := newNode(in)
		for _, o := range nodes[1:] {
			t := newNode(o)
			e := g.NewEdge(f, t)
			g.SetEdge(e)
		}
	}
	return g
}

func newNode(s string) node {
	id := crc32.ChecksumIEEE([]byte(s))
	return node{
		label: s,
		id:    int64(id),
	}
}

type node struct {
	label string
	id    int64
}

func (n node) ID() int64 {
	return n.id
}

func (n node) Attributes() []encoding.Attribute {
	a := []encoding.Attribute{
		{Key: "label", Value: n.label}, // Add label as a DOT attribute
	}
	if n.label == "dac" || n.label == "fft" {
		a = append(a, encoding.Attribute{Key: "color", Value: "orange"})
	}
	return a
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
