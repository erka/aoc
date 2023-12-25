package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"slices"
	"strconv"
	"strings"

	"github.com/erka/aoc/pkg/log"
	"github.com/samber/lo"
	"gonum.org/v1/gonum/graph/network"
	"gonum.org/v1/gonum/graph/simple"
	"gonum.org/v1/gonum/graph/topo"
)

//go:embed input.txt
var input []byte

/*
* part 1:
* part 2:
 */
func main() {
	log.Infof("part1: %s", solvePart1(input))
	log.Infof("part2: %s", solvePart2(input))
}

// solve
func solvePart1(input []byte) string {
	scanner := bufio.NewScanner(bytes.NewBuffer(input))

	nodeIDs := map[string]int64{}
	IDNodes := map[int64]string{}
	g := simple.NewUndirectedGraph()
	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Split(line, ": ")
		from := values[0]
		to := strings.Fields(values[1])
		for _, n := range append(to, from) {

			if _, ok := nodeIDs[n]; !ok {

				node := g.NewNode()
				nodeIDs[n] = node.ID()
				IDNodes[node.ID()] = n
				g.AddNode(node)
			}
		}
		fromNode, ok := g.NodeWithID(nodeIDs[from])
		if ok {
			panic("no node " + from)
		}
		for _, t := range to {
			toNode, ok := g.NodeWithID(nodeIDs[t])
			if ok {
				panic("no node " + t)
			}
			if !g.HasEdgeBetween(fromNode.ID(), toNode.ID()) {
				g.SetEdge(g.NewEdge(fromNode, toNode))
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Errorf("failed to read input: %v", err)
		return ""
	}

	eb := network.EdgeBetweenness(g)

	values := lo.Values(eb)
	slices.Sort(values)
	topValues := values[len(values)-3:]

	log.Debug("Len: %v", len(eb))
	for k, v := range eb {
		if slices.Contains(topValues, v) {
			g.RemoveEdge(k[0], k[1])
		}
	}
	components := topo.ConnectedComponents(g)
	if len(components) != 2 {
		log.Error("expected two components, but found ", len(components))
	}
	return strconv.Itoa(len(components[0]) * len(components[1]))
}

// solve
func solvePart2(input []byte) string {
	scanner := bufio.NewScanner(bytes.NewBuffer(input))
	for scanner.Scan() {
		line := scanner.Text()
		log.Debug(line)
	}
	if err := scanner.Err(); err != nil {
		return err.Error()
	}
	return strconv.Itoa(0)
}
