package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"slices"
	"strconv"
)

//go:embed input.txt
var input []byte

const (
	N = 'n'
	E = 'e'
	W = 'w'
	S = 's'
)

type Point struct {
	x, y int
}

/*
* part 1: 6701
* part 2: 303
 */
func main() {
	fmt.Println("part1:", solve(input))
	fmt.Printf("part2: %s\n", solvePart2(input))
}

func path(field [][]rune) []Point {
	var start Point
	for x, row := range field {
		for y := range row {
			if field[x][y] == 'S' {
				start = Point{x, y}
			}
		}
	}
	next := Point{start.x, start.y}
	direction := W
	points := []Point{start}

	for {
		points = append(points, next)
		nx := Point{next.x, next.y}

		switch field[next.x][next.y] {
		case 'L':
			switch direction {
			case W:
				direction = N
				nx.x -= 1
			case S:
				direction = E
				nx.y += 1
			}
		case 'J':
			switch direction {
			case S:
				direction = W
				nx.y -= 1
			case E:
				direction = N
				nx.x -= 1
			}
		case 'F', 'S':
			switch direction {
			case N:
				direction = E
				nx.y += 1
			case W:
				direction = S
				nx.x += 1
			}

		case '7':
			switch direction {
			case E:
				nx.x += 1
				direction = S
			case N:
				nx.y -= 1
				direction = W
			}

		case '|':
			switch direction {
			case N:
				nx.x -= 1
			case S:
				nx.x += 1
			}
		case '-':
			switch direction {
			case E:
				nx.y += 1
			case W:
				nx.y -= 1
			}
		}
		next = nx
		if start == next {
			break
		}
	}
	return points
}

// solve
func solve(input []byte) string {
	scanner := bufio.NewScanner(bytes.NewBuffer(input))
	field := make([][]rune, 0)
	for scanner.Scan() {
		line := scanner.Text()
		field = append(field, []rune(line))
	}
	if err := scanner.Err(); err != nil {
		return err.Error()
	}
	points := path(field)
	return strconv.Itoa(len(points) / 2)
}

func solvePart2(input []byte) string {
	scanner := bufio.NewScanner(bytes.NewBuffer(input))
	field := make([][]rune, 0)
	for scanner.Scan() {
		line := scanner.Text()
		field = append(field, []rune(line))
	}
	if err := scanner.Err(); err != nil {
		return err.Error()
	}
	n := 0
	points := path(field)
	for x, row := range field {
		inside := false
		for y := range row {
			if slices.Contains(points, Point{x, y}) {
				if crossed(field[x][y]) {
					inside = !inside
				}
				continue
			}
			if inside {
				n++
			}
		}
	}
	return strconv.Itoa(n)
}

func crossed(char rune) bool {
	return slices.Contains([]rune{'|', '7', 'F', 'S'}, char)
}
