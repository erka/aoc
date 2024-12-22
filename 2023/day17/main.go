// the author of the solution: https://github.com/mnml/aoc/tree/main/2023/17 . My respect to you
package main

import (
	"bufio"
	"bytes"
	"container/heap"
	_ "embed"
	"image"
	"strconv"

	"github.com/erka/aoc/pkg/log"
)

//go:embed input.txt
var input []byte

/*
* part 1: <1073
* part 2:
 */
func main() {
	log.Infof("part1: %s", solvePart1(input))
	log.Infof("part2: %s", solvePart2(input))
}

type hqi[T any] struct {
	v T
	p int
}

type Queue[T any] []hqi[T]

func (q Queue[_]) Len() int           { return len(q) }
func (q Queue[_]) Less(i, j int) bool { return q[i].p < q[j].p }
func (q Queue[_]) Swap(i, j int)      { q[i], q[j] = q[j], q[i] }
func (q *Queue[T]) Push(x any)        { *q = append(*q, x.(hqi[T])) }
func (q *Queue[_]) Pop() (x any)      { x, *q = (*q)[len(*q)-1], (*q)[:len(*q)-1]; return x }
func (q *Queue[T]) GPush(v T, p int)  { heap.Push(q, hqi[T]{v, p}) }
func (q *Queue[T]) GPop() (T, int)    { x := heap.Pop(q).(hqi[T]); return x.v, x.p }

type State struct {
	Pos image.Point
	Dir image.Point
}

func solve(input []byte, pathMin, pathMax int) int {
	scanner := bufio.NewScanner(bytes.NewBuffer(input))
	grid := map[image.Point]int{}
	end := image.Pt(0, 0)
	for y := 0; scanner.Scan(); y += 1 {
		for x, w := range scanner.Text() {
			grid[image.Pt(x, y)] = int(w - '0')
			end.X, end.Y = x, y
		}
	}
	if err := scanner.Err(); err != nil {
		log.Errorf("failed to read input: %v", err)
		return -1
	}

	queue, seen := Queue[State]{}, map[State]struct{}{}
	queue.GPush(State{image.Point{0, 0}, image.Point{1, 0}}, 0) // go east
	queue.GPush(State{image.Point{0, 0}, image.Point{0, 1}}, 0) // go south
	result := 0
	for queue.Len() > 0 {
		state, heat := queue.GPop()
		if state.Pos == end {
			result = heat
			break
		}
		if _, ok := seen[state]; ok {
			continue
		}
		seen[state] = struct{}{}
		for _, d := range []image.Point{
			{state.Dir.Y, state.Dir.X}, {-state.Dir.Y, -state.Dir.X},
		} {
			for i := pathMin; i <= pathMax; i++ {
				n := state.Pos.Add(d.Mul(i))
				if _, ok := grid[n]; ok {
					h := 0
					for j := 1; j <= i; j++ {
						h += grid[state.Pos.Add(d.Mul(j))]
					}
					queue.GPush(State{n, d}, heat+h)
				}
			}
		}
	}
	return result
}

// solve
func solvePart1(input []byte) string {
	return strconv.Itoa(solve(input, 1, 3))
}

// solve
func solvePart2(input []byte) string {
	return strconv.Itoa(solve(input, 4, 10))
}
