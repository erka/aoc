package main

import (
	"bytes"
	_ "embed"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/erka/aoc/pkg/log"
	"github.com/samber/lo"
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/encoding"
	"gonum.org/v1/gonum/graph/encoding/dot"
	"gonum.org/v1/gonum/graph/path"
	"gonum.org/v1/gonum/graph/simple"
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
	data := bytes.Split(bytes.Trim(input, "\n"), []byte("\n\n"))

	gates := lo.Associate(bytes.Split(data[0], []byte("\n")), func(v []byte) (string, int) {
		// dXX: n
		data := bytes.Split(v, []byte(": "))
		value, _ := strconv.Atoi(string(data[1]))
		return string(data[0]), value
	})

	wires := lo.Map(bytes.Split(data[1], []byte("\n")), func(v []byte, _ int) wire {
		// dXX WWW dYY -> dZZ
		buf := bytes.Split(v, []byte(" -> "))
		w := wire{}
		w.out = string(buf[1])

		buf = bytes.Split(buf[0], []byte(" "))

		w.g1 = string(buf[0])
		w.op = string(buf[1])
		w.g2 = string(buf[2])

		return w
	})

	ops := slices.Clone(wires)
	c := []string{}
	for len(ops) > 0 {
		op := ops[0]
		ops = ops[1:]
		gv1, ok1 := gates[op.g1]
		gv2, ok2 := gates[op.g2]
		if !ok1 || !ok2 {
			// not ready yet
			ops = append(ops, op)
			continue
		}
		c = append(c, op.out)
		var o int
		switch op.op {
		case "AND":
			o = gv1 & gv2
		case "OR":
			o = gv1 | gv2
		case "XOR":
			o = gv1 ^ gv2
		}

		gates[op.out] = o

	}
	bitKeys := lo.Filter(lo.Keys(gates), func(v string, _ int) bool {
		return v[0] == 'z'
	})

	out := lo.Reduce(bitKeys, func(agg int, key string, i int) int {
		if gates[key] == 1 {
			agg += 1 << i
		}
		return agg
	}, 0)

	return strconv.Itoa(out)
}

// solve
func solvePart2(input []byte) string {
	data := bytes.Split(bytes.Trim(input, "\n"), []byte("\n\n"))
	wires := lo.Map(bytes.Split(data[1], []byte("\n")), func(v []byte, _ int) wire {
		// dXX WWW dYY -> dZZ
		buf := bytes.Split(v, []byte(" -> "))
		w := wire{}
		w.out = string(buf[1])

		buf = bytes.Split(buf[0], []byte(" "))

		w.g1 = string(buf[0])
		w.op = string(buf[1])
		w.g2 = string(buf[2])

		return w
	})

	inputs := lo.Filter(wires, func(item wire, _ int) bool {
		return item.g1[1:] == item.g2[1:] && ((item.g1[0] == 'x' && item.g2[0] == 'y') || (item.g1[0] == 'y' && item.g2[0] == 'x'))
	})

	g := simple.NewDirectedGraph()

	for _, wire := range wires {
		out := node(wire.out)
		edgea := LabeledEdge{simple.Edge{F: node(wire.g1), T: out}, wire.op}
		g.SetEdge(edgea)
		edgeb := LabeledEdge{simple.Edge{F: node(wire.g2), T: out}, wire.op}
		g.SetEdge(edgeb)
	}

	anomales := lo.Reduce(inputs, func(acc map[string][]string, in wire, _ int) map[string][]string {
		out := "z" + in.g1[1:]
		pt := path.DijkstraFrom(node(in.g1), g)
		nodes, _ := pt.To(node(out).ID())
		innerops := []string{}
		for i := 0; i < len(nodes)-1; i++ {
			innerops = append(innerops, (g.Edge(nodes[i].ID(), nodes[i+1].ID()).(LabeledEdge)).Label)
		}
		acc[strings.Join(innerops, ",")] = lo.Uniq(append(acc[strings.Join(innerops, ",")], out))
		return acc
	}, map[string][]string{})

	log.Debug(anomales)

	dotData, err := dot.Marshal(g, "Graph", "", "")
	if err != nil {
		log.Errorf("failed to marshal graph: %v", err)
	}

	os.WriteFile("graph.dot", dotData, 0644)

	b := []string{"z37", "z12", "z29", "vvm", "dgr", "dtv", "fgc", "mtj"}
	slices.Sort(b)
	return strings.Join(b, ",")
}

func node(s string) LabeledNode {
	return LabeledNode{simple.Node(nodeid(s)), s, "circle"}
}

func nodeid(s string) int64 {
	return 100000*int64(s[0]) + 1000*int64(s[1]) + int64(s[2])
}

type LabeledNode struct {
	graph.Node
	Label string
	Shape string
}

func (e LabeledNode) Attributes() []encoding.Attribute {
	return []encoding.Attribute{
		{Key: "label", Value: e.Label}, // Add label as a DOT attribute
		{Key: "shape", Value: e.Shape}, // Add label as a DOT attribute
	}
}

type LabeledEdge struct {
	graph.Edge
	Label string
}

func (e LabeledEdge) Attributes() []encoding.Attribute {
	colors := map[string]string{
		"AND": "blue",
		"OR":  "orange",
		"XOR": "black",
	}
	return []encoding.Attribute{
		{Key: "label", Value: e.Label},         // Add label as a DOT attribute
		{Key: "color", Value: colors[e.Label]}, // Add label as a DOT attribute
	}
}

type wire struct {
	g1  string
	op  string
	g2  string
	out string
}
