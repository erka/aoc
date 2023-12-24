package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"image"
	"strconv"
	"strings"

	"github.com/erka/aoc/aoc2023/day23/depth"
	"github.com/erka/aoc/pkg/log"
	"gonum.org/v1/gonum/graph/simple"
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

	g := simple.NewDirectedGraph()
	for y, l := range grid {
		for x, c := range l {
			if c == '#' {
				continue
			}
			g.AddNode(simple.Node(vertexID(image.Pt(x, y))))
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
					from, _ := g.NodeWithID(int64(vertexID(p)))
					to, _ := g.NodeWithID(int64(vertexID(q)))
					g.SetEdge(g.NewEdge(from, to))
				}
			}

		}
	}
	// end
	result := 0
	s, _ := g.NodeWithID(int64(vertexID(start)))
	e, _ := g.NodeWithID(int64(vertexID(end)))
	log.Debug(s.ID(), e.ID())

	bfs := &depth.DepthLast{}
	bfs.WalkAll(g, s, e, func(depth int) {
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
