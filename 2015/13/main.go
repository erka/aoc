package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/erka/aoc/pkg/log"
	"github.com/samber/lo"
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/encoding"
	"gonum.org/v1/gonum/graph/simple"
)

//go:embed input.txt
var input []byte

/*
* part 1: 733
* part 2: 725
 */
func main() {
	log.Infof("part1: %s", solvePart1(input))
	log.Infof("part2: %s", solvePart2(input))
}

// solve
func solvePart1(input []byte) string {
	g := buildConnections(input)

	output := 0
	guests := graph.NodesOf(g.Nodes())
	for _, guest := range guests {
		w := optimalHappiness(g, guest, []graph.Node{guest})
		output = max(w, output)
	}

	return strconv.Itoa(output)
}

// solve
func solvePart2(input []byte) string {
	g := buildConnections(input)
	guests := graph.NodesOf(g.Nodes())
	me := newNode("me")
	for _, guest := range guests {
		g.SetWeightedEdge(simple.WeightedEdge{F: me, T: guest, W: 0})
		g.SetWeightedEdge(simple.WeightedEdge{F: guest, T: me, W: 0})
	}

	output := 0
	for _, guest := range guests {
		w := optimalHappiness(g, guest, []graph.Node{guest})
		output = max(w, output)
	}

	return strconv.Itoa(output)
}

func buildConnections(input []byte) *simple.WeightedDirectedGraph {
	lines := bytes.Split(bytes.Trim(input, "\n"), []byte("\n"))
	var from, to, reaction string
	var happiness int
	g := simple.NewWeightedDirectedGraph(0, 0)
	for _, line := range lines {
		fmt.Sscanf(string(line), "%s would %s %d happiness units by sitting next to %s.",
			&from, &reaction, &happiness, &to)
		if reaction == "lose" {
			happiness = -happiness
		}
		to = strings.TrimSuffix(to, ".")
		g.SetWeightedEdge(simple.WeightedEdge{F: newNode(from), T: newNode(to), W: float64(happiness)})
	}
	return g
}

func optimalHappiness(g graph.Weighted, start graph.Node, seen []graph.Node) int {
	if len(seen) == g.Nodes().Len() {
		w := 0
		for i := len(seen) - 2; i >= 0; i-- {
			w += int(g.WeightedEdge(seen[i+1].ID(), seen[i].ID()).Weight())
			w += int(g.WeightedEdge(seen[i].ID(), seen[i+1].ID()).Weight())
		}
		w += int(g.WeightedEdge(seen[0].ID(), seen[len(seen)-1].ID()).Weight())
		w += int(g.WeightedEdge(seen[len(seen)-1].ID(), seen[0].ID()).Weight())
		return w
	}
	neighbors := g.From(start.ID())
	output := 0
	for neighbors.Next() {
		n := neighbors.Node()
		if slices.Contains(seen, n) {
			continue
		}
		w := optimalHappiness(g, n, append(seen, n))
		output = max(output, w)
	}
	return output
}

type LabeledNode struct {
	graph.Node
	Label string
}

func (e LabeledNode) Attributes() []encoding.Attribute {
	return []encoding.Attribute{
		{Key: "label", Value: e.Label}, // Add label as a DOT attribute
	}
}

func newNode(label string) LabeledNode {
	id := lo.Sum([]byte(label))
	return LabeledNode{Node: simple.Node(id), Label: label}
}
