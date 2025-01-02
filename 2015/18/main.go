package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"image"
	"os"
	"os/exec"
	"slices"
	"strconv"

	"github.com/erka/aoc/pkg/log"
	"github.com/samber/lo"
)

//go:embed input.txt
var input []byte

/*
* part 1: 768
* part 2: 781
 */
func main() {
	log.Infof("part1: %s", solvePart1(input, 100))
	log.Infof("part2: %s", solvePart2(input, 100))
}

var directions = []image.Point{
	image.Pt(1, 0),
	image.Pt(0, 1),
	image.Pt(-1, 0),
	image.Pt(0, -1),
	image.Pt(1, 1),
	image.Pt(-1, 1),
	image.Pt(-1, -1),
	image.Pt(1, -1),
}

// solve
func solvePart1(input []byte, times int) string {
	lines := bytes.Split(bytes.Trim(input, "\n"), []byte("\n"))
	lines = slices.Clone(lines)
	area := image.Rect(0, 0, len(lines[0]), len(lines))

	next := make([][]byte, area.Dx())
	for x := 0; x < area.Dx(); x++ {
		next[x] = make([]byte, area.Dy())
	}
	for i := 0; i < times; i++ {
		for x := 0; x < area.Dx(); x++ {
			for y := 0; y < area.Dy(); y++ {
				neighborsOn := 0
				for _, d := range directions {
					n := image.Pt(y, x).Add(d)
					if !n.In(area) {
						continue
					}
					if lines[n.Y][n.X] == '#' {
						neighborsOn++
					}
				}
				if neighborsOn == 3 || (neighborsOn == 2 && lines[x][y] == '#') {
					next[x][y] = '#'
				} else {
					next[x][y] = '.'
				}
			}
		}
		for x := 0; x < area.Dx(); x++ {
			for y := 0; y < area.Dy(); y++ {
				lines[x][y] = next[x][y]
			}
		}
	}
	output := lo.SumBy(lines, func(item []byte) int {
		return lo.Count(item, '#')
	})
	return strconv.Itoa(output)
}

// solve
func solvePart2(input []byte, times int) string {
	{
		lines := bytes.Split(bytes.Trim(input, "\n"), []byte("\n"))
		lines = slices.Clone(lines)
		area := image.Rect(0, 0, len(lines[0]), len(lines))

		next := make([][]byte, area.Dx())
		for x := 0; x < area.Dx(); x++ {
			next[x] = make([]byte, area.Dy())
		}
		for i := 0; i < times; i++ {
			for x := 0; x < area.Dx(); x++ {
				for y := 0; y < area.Dy(); y++ {
					if (x == 0 || x == area.Dx()-1) && (y == 0 || y == area.Dy()-1) {
						next[x][y] = '#'
						continue
					}
					neighborsOn := 0
					for _, d := range directions {
						n := image.Pt(y, x).Add(d)
						if !n.In(area) {
							continue
						}
						if lines[n.Y][n.X] == '#' {
							neighborsOn++
						}
					}
					if neighborsOn == 3 || (neighborsOn == 2 && lines[x][y] == '#') {
						next[x][y] = '#'
					} else {
						next[x][y] = '.'
					}
				}
			}
			for x := 0; x < area.Dx(); x++ {
				for y := 0; y < area.Dy(); y++ {
					lines[x][y] = next[x][y]
				}
			}
			c := exec.Command("clear")
			c.Stdout = os.Stdout
			c.Run()
			for _, l := range lines {
				fmt.Println(string(l))
			}
		}
		output := lo.SumBy(lines, func(item []byte) int {
			return lo.Count(item, '#')
		})
		return strconv.Itoa(output)
	}
}
