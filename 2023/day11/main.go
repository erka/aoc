package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

//go:embed input.txt
var input []byte

type Galaxy struct {
	Row, Col int
}

func (p *Galaxy) Dist(o *Galaxy) int {
	x := p.Col - o.Col
	if x <= 0 {
		x *= -1
	}
	y := p.Row - o.Row
	if y <= 0 {
		y *= -1
	}
	return x + y
}

/*
* part 1: 9724940
* part 2: 569052586852
 */
func main() {
	fmt.Printf("part1: %s\n", solve(input, 2))
	fmt.Printf("part2: %s\n", solve(input, 1_000_000))
}

func expand(universe [][]string, galaxies []*Galaxy, factor int) {
	factor = factor - 1 //one row already in the place
	emptyRows := make([]int, 0)
	emptyCols := make([]int, 0)

	for r, row := range universe {
		if !slices.Contains(row, "#") {
			emptyRows = append(emptyRows, r)
		}
	}

	for c := 0; c < len(universe[0]); c++ {
		empty := true
		for r := 0; r < len(universe); r++ {
			if universe[r][c] == "#" {
				empty = false
				break
			}
		}
		if empty {
			emptyCols = append(emptyCols, c)
		}
	}

	for _, galaxy := range galaxies {
		deltaRow := 0
		for _, row := range emptyRows {
			if galaxy.Row > row {
				deltaRow += factor
			}
		}
		deltaCol := 0
		for _, Col := range emptyCols {
			if galaxy.Col > Col {
				deltaCol += factor
			}
		}
		galaxy.Row += deltaRow
		galaxy.Col += deltaCol
	}
}

// solve
func solve(input []byte, factor int) string {
	scanner := bufio.NewScanner(bytes.NewBuffer(input))
	universe := make([][]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		universe = append(universe, strings.Split(line, ""))
	}

	galaxies := make([]*Galaxy, 0)
	for x, row := range universe {
		for y, char := range row {
			if char == "#" {
				galaxies = append(galaxies, &Galaxy{x, y})
			}
		}
	}

	expand(universe, galaxies, factor)

	distances := make([]int, 0)
	for i, galaxy := range galaxies {
		for j := i + 1; j < len(galaxies); j++ {
			distances = append(distances, galaxy.Dist(galaxies[j]))
		}
	}
	if err := scanner.Err(); err != nil {
		return err.Error()
	}
	return strconv.Itoa(lo.Sum(distances))
}
