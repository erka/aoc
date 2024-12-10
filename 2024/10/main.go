package main

import (
	"bytes"
	_ "embed"
	"image"
	"strconv"

	"github.com/erka/aoc/pkg/log"
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/path"
	"gonum.org/v1/gonum/graph/simple"
)

//go:embed input.txt
var input []byte

/*
* part 1: 688
* part 2: 1459
 */
func main() {
	log.Infof("part1: %s", solvePart1(input))
	log.Infof("part2: %s", solvePart2(input))
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
func solvePart1(input []byte) string {
	topo, heads, peaks := makeTopo(input)
	trails := 0
	for _, head := range heads {
		headNode, _ := topo.NodeWithID(nodeId(head))
		trail, _ := path.BellmanFordFrom(headNode, topo)
		for _, peak := range peaks {
			places, _ := trail.To(nodeId(peak))
			if len(places) > 0 {
				trails += 1
			}
		}
	}
	return strconv.Itoa(trails)
}

func makeTopo(input []byte) (*simple.DirectedGraph, []image.Point, []image.Point) {
	area := bytes.FieldsFunc(input, func(r rune) bool { return r == '\n' })
	graph := simple.NewDirectedGraph()
	boundary := image.Rect(0, 0, len(area), len(area[0]))
	trailheads := []image.Point{}
	peaks := []image.Point{}
	for y, line := range area {
		for x, h := range line {
			p := image.Pt(y, x)
			switch h {
			case '0':
				trailheads = append(trailheads, p)
			case '9':
				peaks = append(peaks, p)
				continue
			}
			for _, d := range directions {
				neighbor := p.Add(d)
				if neighbor.In(boundary) {
					nh := area[neighbor.X][neighbor.Y]
					if h < nh && nh-h == 1 {
						graph.SetEdge(simple.Edge{F: simple.Node(nodeId(p)), T: simple.Node(nodeId(neighbor))})
					}
				}
			}
		}
	}

	log.Debug("trailheads:", trailheads)
	log.Debug("peaks:", peaks)
	return graph, trailheads, peaks
}

// solve
func solvePart2(input []byte) string {
	topo, heads, peaks := makeTopo(input)
	ratings := 0
	for _, th := range heads {
		start, _ := topo.NodeWithID(nodeId(th))
		for _, pk := range peaks {
			end, _ := topo.NodeWithID(nodeId(pk))
			var trails [][]graph.Node
			FindAllPaths(topo, start, end, map[int64]bool{}, []graph.Node{}, &trails)
			ratings += len(trails)
		}
	}
	return strconv.Itoa(ratings)
}

func FindAllPaths(g graph.Graph, start, end graph.Node, visited map[int64]bool, path []graph.Node, paths *[][]graph.Node) {
	visited[start.ID()] = true
	path = append(path, start)

	if start.ID() == end.ID() {
		tempPath := make([]graph.Node, len(path))
		copy(tempPath, path)
		*paths = append(*paths, tempPath)
	} else {
		for _, neighbor := range graph.NodesOf(g.From(start.ID())) {
			if !visited[neighbor.ID()] {
				FindAllPaths(g, neighbor, end, visited, path, paths)
			}
		}
	}
	path = path[:len(path)-1]
	visited[start.ID()] = false
}
