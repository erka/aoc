package main

import (
	"bytes"
	_ "embed"
	"errors"
	"fmt"
	"image"
	"io"
	"math"
	"strconv"

	"github.com/erka/aoc/pkg/log"
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/path"
	"gonum.org/v1/gonum/graph/simple"
)

//go:embed input.txt
var input []byte

/*
* part 1: 270084
* part 2:
 */
func main() {
	log.Infof("part1: %s", solvePart1(input))
	log.Infof("part2: %s", solvePart2(input))
}

type keypad map[byte]image.Point

var numericKeypad = keypad{
	'A': image.Pt(0, 0),
	'0': image.Pt(1, 0),
	'1': image.Pt(2, 1),
	'2': image.Pt(1, 1),
	'3': image.Pt(0, 1),
	'4': image.Pt(2, 2),
	'5': image.Pt(1, 2),
	'6': image.Pt(0, 2),
	'7': image.Pt(2, 3),
	'8': image.Pt(1, 3),
	'9': image.Pt(0, 3),
}

var directionalKeypad = keypad{
	'>': image.Pt(0, 0),
	'v': image.Pt(1, 0),
	'<': image.Pt(2, 0),
	'^': image.Pt(1, 1),
	'A': image.Pt(0, 1),
}

var symbols = map[image.Point]byte{
	image.Pt(1, 0):  '<',
	image.Pt(-1, 0): '>',
	image.Pt(0, 1):  '^',
	image.Pt(0, -1): 'v',
	image.Pt(0, 0):  'A',
}

var directionEdges = []simple.Edge{
	{F: simple.Node('A'), T: simple.Node('>')},
	{F: simple.Node('A'), T: simple.Node('^')},
	{F: simple.Node('^'), T: simple.Node('v')},
	{F: simple.Node('>'), T: simple.Node('v')},
	{F: simple.Node('v'), T: simple.Node('<')},
}

var numericEdges = []simple.Edge{
	{F: simple.Node('A'), T: simple.Node('0')},
	{F: simple.Node('A'), T: simple.Node('3')},
	{F: simple.Node('0'), T: simple.Node('2')},
	{F: simple.Node('1'), T: simple.Node('2')},
	{F: simple.Node('1'), T: simple.Node('4')},
	{F: simple.Node('2'), T: simple.Node('5')},
	{F: simple.Node('2'), T: simple.Node('3')},
	{F: simple.Node('3'), T: simple.Node('6')},
	{F: simple.Node('4'), T: simple.Node('7')},
	{F: simple.Node('4'), T: simple.Node('5')},
	{F: simple.Node('5'), T: simple.Node('8')},
	{F: simple.Node('5'), T: simple.Node('6')},
	{F: simple.Node('6'), T: simple.Node('9')},
	{F: simple.Node('7'), T: simple.Node('8')},
	{F: simple.Node('8'), T: simple.Node('9')},
}

func sequence(line []byte, keypad map[byte]image.Point, times int) [][]byte {
	edges := directionEdges
	if len(keypad) == len(numericKeypad) {
		edges = numericEdges
	}
	g := simple.NewUndirectedGraph()
	for _, e := range edges {
		g.SetEdge(e)
	}

	// start with A
	variants := findAll(line, byte('A'), g)

	output := make([][]byte, len(variants))
	for i, nodes := range variants {
		for j := 0; j < len(nodes)-1; j++ {
			next := keypad[byte(nodes[j+1].ID())]
			current := keypad[byte(nodes[j].ID())]
			output[i] = append(output[i], symbols[next.Sub(current)])
		}
	}
	times -= 1
	if times == 0 {
		return output
	}

	next := [][]byte{}
	for _, mv := range output {
		key := fmt.Sprintf("%s/%d", mv, times)
		values, ok := cache[key]
		if !ok {
			values = sequence(mv, keypad, times)
			cache[key] = values
		}

		next = append(next, values...)
	}
	return next
}

var cache = map[string][][]byte{}

func findAll(line []byte, start byte, g *simple.UndirectedGraph) [][]graph.Node {
	o := [][]graph.Node{}
	if len(line) == 0 {
		o := append(o, []graph.Node{simple.Node(start)})
		return o
	}
	c := line[0]
	pt := path.DijkstraAllFrom(simple.Node(start), g)
	all, _ := pt.AllTo(int64(c))
	for _, a := range all {
		values := findAll(line[1:], c, g)
		for _, x := range values {
			o = append(o, append(a, x...))
		}
	}
	return o
}

// solve
func solvePart1(input []byte) string {
	lines := bytes.Split(bytes.Trim(input, "\n"), []byte("\n"))
	output := 0
	for _, line := range lines {
		variants := sequence(line, numericKeypad, 1)
		l := math.MaxInt64
		for _, v := range variants {
			seconds := sequence(v, directionalKeypad, 2)
			for _, s := range seconds {
				l = min(l, len(s))
			}
		}
		n, _ := strconv.Atoi(string(line)[:3])
		output += l * n
	}
	return strconv.Itoa(output)
}

// solve
func solvePart2(input []byte) string {
	reader := bytes.NewReader(input)
	var (
		line string
		err  error
	)
	for {
		_, err = fmt.Fscanln(reader, &line)
		if err != nil {
			if !errors.Is(err, io.EOF) {
				return err.Error()
			}
			break
		}
		log.Debug(line)
	}

	return strconv.Itoa(0)
}
