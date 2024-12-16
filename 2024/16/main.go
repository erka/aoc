package main

import (
	"bytes"
	_ "embed"
	"image"
	"math"
	"slices"
	"strconv"

	"github.com/erka/aoc/pkg/log"
	"github.com/samber/lo"
)

//go:embed input.txt
var input []byte

/*
* part 1: 122492
* part 2: 520
 */
func main() {
	log.Infof("part1: %s", solvePart1(input))
	log.Infof("part2: %s", solvePart2(input))
}

var (
	north = image.Pt(0, -1)
	east  = image.Pt(1, 0)
	south = image.Pt(0, 1)
	west  = image.Pt(-1, 0)
	dx    = map[image.Point][]image.Point{
		north: {north, east, west},
		east:  {east, north, south},
		south: {south, east, west},
		west:  {west, north, south},
	}
)

type Node struct {
	pos, dir image.Point
}

type Item struct {
	Node
	weight int
	nodes  []image.Point
}

func solvePart1(input []byte) string {
	paths := solve(input)
	log.Debug(display(input, paths[0].nodes))
	return strconv.Itoa(paths[0].weight)
}

// solve
func solvePart2(input []byte) string {
	paths := solve(input)
	nodes := []image.Point{}
	for _, path := range paths {
		nodes = append(nodes, path.nodes...)
	}
	nodes = lo.Uniq(nodes)
	output := len(nodes)
	return strconv.Itoa(output)
}

// solve
func solve(input []byte) []Item {
	lines := bytes.FieldsFunc(input, func(r rune) bool { return r == '\n' })
	start := image.Point{}
	end := image.Point{}
	places := map[image.Point]byte{}
	for y, line := range lines {
		for x, c := range line {
			if c == '#' {
				continue
			}
			p := image.Pt(x, y)
			places[p] = c
			switch c {
			case 'S':
				start = p
			case 'E':
				end = p
			}
		}
	}

	output := math.MaxInt
	queue := []Item{{Node{start, east}, 0, []image.Point{start}}}

	paths := []Item{}
	cache := map[Node]int{}
	for len(queue) > 0 {
		i := queue[0]
		queue = queue[1:]
		if c, ok := cache[i.Node]; ok && c < i.weight {
			continue
		}
		cache[Node{i.pos, i.dir}] = i.weight
		if i.pos == end {
			output = min(output, i.weight)
			paths = append(paths, i)
			paths = lo.Filter(paths, func(i Item, _ int) bool { return i.weight == output })
			continue
		}
		if i.weight > output {
			continue
		}
		for _, d := range dx[i.dir] {
			nd := i.pos.Add(d)
			if _, ok := places[nd]; !ok {
				continue
			}
			w := i.weight + 1
			if d != i.dir {
				w += 1000
			}
			queue = append(queue, Item{Node{nd, d}, w, append(slices.Clone(i.nodes), nd)})
		}
	}

	return paths
}

func display(input []byte, nodes []image.Point) string {
	var b bytes.Buffer
	lines := bytes.FieldsFunc(input, func(r rune) bool { return r == '\n' })
	for y, line := range lines {
		for x, c := range line {
			if c == '.' && slices.Contains(nodes, image.Pt(x, y)) {
				b.WriteByte('O')
				continue
			}
			b.WriteByte(c)
		}
		b.WriteByte('\n')
	}
	return b.String()
}
