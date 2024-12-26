package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"math"
	"slices"
	"strconv"

	"github.com/erka/aoc/pkg/log"
	"github.com/samber/lo"
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/simple"
)

//go:embed input.txt
var input []byte

/*
* part 1: 141
* part 2: 736
 */
func main() {
	log.Infof("part1: %s", solvePart1(input))
	log.Infof("part2: %s", solvePart2(input))
}

type LabeledNode struct {
	graph.Node
	Label string
}

func newNode(label string) LabeledNode {
	id := lo.Sum([]byte(label))
	return LabeledNode{Node: simple.Node(id), Label: label}
}

// solve
func solvePart1(input []byte) string {
	lines := bytes.Split(bytes.Trim(input, "\n"), []byte("\n"))
	var from, to string
	var distance int
	g := simple.NewWeightedUndirectedGraph(0, 0)
	for _, line := range lines {
		fmt.Sscanf(string(line), "%s to %s = %d", &from, &to, &distance)
		g.SetWeightedEdge(simple.WeightedEdge{F: newNode(from), T: newNode(to), W: float64(distance)})
	}
	output := math.MaxInt
	cities := graph.NodesOf(g.Nodes())
	for _, city := range cities {
		w := shortestPath(g, city, []graph.Node{city})
		output = min(w, output)
	}

	return strconv.Itoa(output)
}

func shortestPath(g graph.Weighted, start graph.Node, seen []graph.Node) int {
	if len(seen) == g.Nodes().Len() {
		return 0
	}
	neighbors := g.From(start.ID())
	output := math.MaxInt
	for neighbors.Next() {
		n := neighbors.Node()
		if slices.Contains(seen, n) {
			continue
		}
		w := shortestPath(g, n, append(seen, n))
		w += int(g.WeightedEdge(start.ID(), n.ID()).Weight())
		output = min(output, w)
	}
	return output
}

// solve
func solvePart2(input []byte) string {
	lines := bytes.Split(bytes.Trim(input, "\n"), []byte("\n"))
	var from, to string
	var distance int
	g := simple.NewWeightedUndirectedGraph(0, 0)
	for _, line := range lines {
		fmt.Sscanf(string(line), "%s to %s = %d", &from, &to, &distance)
		g.SetWeightedEdge(simple.WeightedEdge{F: newNode(from), T: newNode(to), W: float64(distance)})
	}
	output := 0
	cities := graph.NodesOf(g.Nodes())
	for _, city := range cities {
		w := longestPath(g, city, []graph.Node{city})
		output = max(w, output)
	}

	return strconv.Itoa(output)
}

func longestPath(g graph.Weighted, start graph.Node, seen []graph.Node) int {
	if len(seen) == g.Nodes().Len() {
		return 0
	}
	neighbors := g.From(start.ID())
	output := 0
	for neighbors.Next() {
		n := neighbors.Node()
		if slices.Contains(seen, n) {
			continue
		}
		w := longestPath(g, n, append(seen, n))
		w += int(g.WeightedEdge(start.ID(), n.ID()).Weight())
		output = max(output, w)
	}
	return output
}
