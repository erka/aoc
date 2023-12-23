package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"image"
	"strconv"
	"strings"

	"github.com/erka/aoc/pkg/log"
)

//go:embed input.txt
var input []byte

/*
* part 1: 2174
* part 2: 6506
 */
func main() {
	log.Infof("part1: %s", solvePart1(input))
	log.Infof("part2: %s", solvePart2(input))
}

type Grid []string

func (g Grid) String() string {
	b := strings.Builder{}
	b.WriteRune('\n')
	for _, row := range g {
		b.WriteString(row)
		b.WriteRune('\n')
	}
	return b.String()
}

var (
	north   = image.Pt(0, -1)
	east    = image.Pt(1, 0)
	south   = image.Pt(0, 1)
	west    = image.Pt(-1, 0)
	offsets = []image.Point{north, east, south, west}
)

func (g Grid) size() (int, int) {
	return len(g), len(g[0])
}

func solve(input []byte, slopes bool) int {
	scanner := bufio.NewScanner(bytes.NewBuffer(input))
	grid := Grid{}
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, line)
	}
	if err := scanner.Err(); err != nil {
		log.Errorf("failed to read input: %v", err)
		return -1
	}

	rows, cols := grid.size()
	start := image.Pt(1, 0)
	end := image.Pt(cols-2, rows-1)
	log.Debugf("%v %v cols: %v rows: %v", start, end, cols, rows)
	bounds := image.Rect(0, 0, cols, rows)

	graph := NewDirectedGraph()
	for y, l := range grid {
		for x, c := range l {
			if c == '#' {
				continue
			}
			graph.AddVertex(vertexID(image.Pt(x, y)))
		}
	}

	for y, l := range grid {
		for x, c := range l {
			if c == '#' {
				continue
			}
			p := image.Pt(x, y)
			for _, off := range offsets {
				q := p.Add(off)
				if q.In(bounds) {
					switch grid[q.Y][q.X] {
					case '#':
						continue
					case '>':
						if off != east && slopes {
							continue
						}
					case 'v':
						if off != south && slopes {
							continue
						}
					}
					graph.AddEdge(vertexID(p), vertexID(q))
				}
			}

		}
	}
	// end
	result := 0
	DFS(graph, vertexID(start), vertexID(end), func(depth int) {
		if depth > result {
			result = depth
			log.Debug(depth)
		}
	})
	return result
}

// solve
func solvePart1(input []byte) string {
	value := solve(input, true)
	return strconv.Itoa(value)
}

// solve
func solvePart2(input []byte) string {
	value := solve(input, false)
	return strconv.Itoa(value)
}

func vertexID(p image.Point) int {
	return p.Y*1000 + p.X
}

type Vertex struct {
	// Key is the unique identifier of the vertex
	Key int
	// Vertices will describe vertices connected to this one
	// The key will be the Key value of the connected vertice
	// with the value being the pointer to it
	Vertices map[int]*Vertex
}

// We then create a constructor function for the Vertex
func NewVertex(key int) *Vertex {
	return &Vertex{
		Key:      key,
		Vertices: map[int]*Vertex{},
	}
}

func (v *Vertex) String() string {
	s := strconv.Itoa(v.Key) + ":"

	for _, neighbor := range v.Vertices {
		s += " " + strconv.Itoa(neighbor.Key)
	}

	return s
}

type Graph struct {
	// Vertices describes all vertices contained in the graph
	// The key will be the Key value of the connected vertice
	// with the value being the pointer to it
	Vertices map[int]*Vertex
	// This will decide if it's a directed or undirected graph
	directed bool
}

// We defined constructor functions that create
// new directed or undirected graphs respectively

func NewDirectedGraph() *Graph {
	return &Graph{
		Vertices: map[int]*Vertex{},
		directed: true,
	}
}

// AddVertex creates a new vertex with the given
// key and adds it to the graph
func (g *Graph) AddVertex(key int) {
	v := NewVertex(key)
	g.Vertices[key] = v
}

// The AddEdge method adds an edge between two vertices in the graph
func (g *Graph) AddEdge(k1, k2 int) {
	v1 := g.Vertices[k1]
	v2 := g.Vertices[k2]

	// return an error if one of the vertices doesn't exist
	if v1 == nil || v2 == nil {
		panic("not all vertices exist")
	}

	// do nothing if the vertices are already connected
	if _, ok := v1.Vertices[v2.Key]; ok {
		return
	}

	// Add a directed edge between v1 and v2
	// If the graph is undirected, add a corresponding
	// edge back from v2 to v1, effectively making the
	// edge between v1 and v2 bidirectional
	v1.Vertices[v2.Key] = v2
	if !g.directed && v1.Key != v2.Key {
		v2.Vertices[v1.Key] = v1
	}

	// Add the vertices to the graph's vertex map
	g.Vertices[v1.Key] = v1
	g.Vertices[v2.Key] = v2
}

func (g *Graph) String() string {
	s := ""
	i := 0
	for _, v := range g.Vertices {
		if i != 0 {
			s += "\n"
		}
		s += v.String()
		i++
	}
	return s
}

var (
	visited = map[int]struct{}{}
	value   = struct{}{}
)

func DFS(g *Graph, start int, end int, visitCb func(int)) {
	if end == start {
		visitCb(len(visited))
		return
	}
	visited[start] = value
	// for each of the adjacent vertices, call the function recursively
	// if it hasn't yet been visited
	for _, v := range g.Vertices[start].Vertices {
		if _, ok := visited[v.Key]; ok {
			continue
		}
		DFS(g, v.Key, end, visitCb)
		delete(visited, v.Key)
	}
}
