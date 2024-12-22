package main

import (
	"bytes"
	_ "embed"
	"image"
	"math"
	"strconv"

	"github.com/erka/aoc/pkg/log"
	"github.com/samber/lo"
	"gonum.org/v1/gonum/graph/path"
	"gonum.org/v1/gonum/graph/simple"
)

//go:embed input.txt
var input []byte

/*
* part 1: 270084
* part 2: 329431019997766
 */
func main() {
	log.Infof("part1: %s", solvePart1(input))
	log.Infof("part2: %s", solvePart2(input))
}

func init() {
	setup(numericKeypad, numericEdges)
	setup(directionalKeypad, directionEdges)
}

type ft struct {
	f byte
	t byte
}

var paths = map[ft][]string{}

// the best path where there is smallest direction changes?

func setup(keypad keypad, edges []simple.Edge) {
	g := simple.NewUndirectedGraph()
	for _, e := range edges {
		g.SetEdge(e)
	}
	nums := lo.Keys(keypad)
	for _, n := range nums {
		pt := path.DijkstraAllFrom(simple.Node(n), g)
		for j := range nums {
			if n == nums[j] {
				continue
			}
			variants, _ := pt.AllTo(int64(nums[j]))
			for _, nodes := range variants {
				output := []byte{}
				for jj := 0; jj < len(nodes)-1; jj++ {
					next := keypad[byte(nodes[jj+1].ID())]
					current := keypad[byte(nodes[jj].ID())]
					output = append(output, symbols[next.Sub(current)])
				}
				paths[ft{f: n, t: nums[j]}] = append(paths[ft{f: n, t: nums[j]}], string(output)+"A")
			}
		}
	}
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

var cache = map[string]int{}

func sequence(line []byte, times int) int {
	if len(line) == 0 {
		return 1
	}
	if times == 0 {
		return len(line)
	}

	key := string(line) + strconv.Itoa(times)
	if v, ok := cache[key]; ok {
		return v
	}

	// start with A
	current := byte('A')
	length := 0
	for _, next := range line {
		minimal := math.MaxInt
		if len(paths[ft{f: current, t: next}]) == 0 {
			length += 1
			continue
		}
		for _, variant := range paths[ft{f: current, t: next}] {
			d := sequence([]byte(variant), times-1)
			minimal = min(minimal, d)
		}
		length += minimal
		current = next
	}
	cache[key] = length
	return length
}

// solve
func solvePart1(input []byte) string {
	lines := bytes.Split(bytes.Trim(input, "\n"), []byte("\n"))
	output := 0
	for _, line := range lines {
		l := sequence(line, 3)
		n, _ := strconv.Atoi(string(line)[:3])
		output += l * n
	}
	return strconv.Itoa(output)
}

// solve
func solvePart2(input []byte) string {
	lines := bytes.Split(bytes.Trim(input, "\n"), []byte("\n"))
	output := 0
	for _, line := range lines {
		l := sequence(line, 26)
		n, _ := strconv.Atoi(string(line)[:3])
		output += l * n
	}
	return strconv.Itoa(output)
}
