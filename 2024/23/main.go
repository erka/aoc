package main

import (
	"bytes"
	"cmp"
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/erka/aoc/pkg/log"
	"github.com/samber/lo"
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/simple"
	"gonum.org/v1/gonum/graph/topo"
)

//go:embed input.txt
var input []byte

/*
* part 1: 1218
* part 2: ah,ap,ek,fj,fr,jt,ka,ln,me,mp,qa,ql,zg
 */
func main() {
	log.Infof("part1: %s", solvePart1(input))
	log.Infof("part2: %s", solvePart2(input))
}

// solve
func solvePart1(input []byte) string {
	edges := lo.Map(
		bytes.Split(bytes.Trim(input, "\n"), []byte("\n")),
		func(item []byte, _ int) []string {
			return strings.Split(string(item), "-")
		},
	)

	g := simple.NewUndirectedGraph()
	for _, c := range edges {
		g.SetEdge(g.NewEdge(node(c[0]), node(c[1])))
	}

	cycles := map[string]struct{}{}
	nodes := graph.NodesOf(g.Nodes())
	for _, node := range nodes {
		neighbours := graph.NodesOf(g.From(node.ID()))
		for i := 0; i < len(neighbours)-1; i++ {
			for j := i + 1; j < len(neighbours); j++ {
				if g.HasEdgeBetween(neighbours[i].ID(), neighbours[j].ID()) {
					l := []string{str(node), str(neighbours[i]), str(neighbours[j])}
					slices.Sort(l)
					cycles[strings.Join(l, ",")] = struct{}{}
				}
			}
		}
	}

	values := lo.Keys(cycles)
	slices.Sort(values)
	log.Debugf("cycles:\n%v", strings.Join(values, "\n"))
	output := len(lo.Filter(values, func(str string, _ int) bool {
		item := strings.Split(str, ",")
		return item[0][0] == 't' || item[1][0] == 't' || item[2][0] == 't'
	}))

	return strconv.Itoa(output)
}

// solve
func solvePart2(input []byte) string {
	edges := lo.Map(
		bytes.Split(bytes.Trim(input, "\n"), []byte("\n")),
		func(item []byte, _ int) []string {
			return strings.Split(string(item), "-")
		},
	)

	g := simple.NewUndirectedGraph()
	for _, c := range edges {
		g.SetEdge(g.NewEdge(node(c[0]), node(c[1])))
	}
	cliques := topo.BronKerbosch(g)

	slices.SortFunc(cliques, func(a []graph.Node, b []graph.Node) int {
		return -cmp.Compare(len(a), len(b))
	})

	pcs := lo.Map(cliques[0], func(item graph.Node, _ int) string {
		return str(item)
	})

	slices.Sort(pcs)
	return strings.Join(pcs, ",")
}

func node(s string) simple.Node {
	id := int64(s[0])*1000 + int64(s[1])
	return simple.Node(id)
}

func str(n graph.Node) string {
	return fmt.Sprintf("%s%s", string(rune(n.ID()/1000)), string(rune(n.ID()%1000)))
}
