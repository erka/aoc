package main

import (
	"bytes"
	_ "embed"
	"image"
	"regexp"
	"strconv"

	"github.com/erka/aoc/pkg/log"
	"github.com/samber/lo"
)

//go:embed input.txt
var input []byte

/*
* part 1: 400410
* part 2: 15343601
 */
func main() {
	log.Infof("part1: %s", solvePart1(input))
	log.Infof("part2: %s", solvePart2(input))
}

var rx = regexp.MustCompile(`^(turn on|toggle|turn off) (\d+),(\d+) through (\d+),(\d+)$`)

// solve
func solvePart1(input []byte) string {
	lines := bytes.Split(bytes.Trim(input, "\n"), []byte("\n"))
	area := make([][]bool, 1000)
	for i := 0; i < 1000; i++ {
		area[i] = make([]bool, 1000)
	}
	var s, e image.Point
	for _, line := range lines {
		data := rx.FindStringSubmatch(string(line))
		s.X, _ = strconv.Atoi(data[2])
		s.Y, _ = strconv.Atoi(data[3])
		e.X, _ = strconv.Atoi(data[4])
		e.Y, _ = strconv.Atoi(data[5])
		switch data[1] {
		case "turn on":
			for x := s.X; x <= e.X; x++ {
				for y := s.Y; y <= e.Y; y++ {
					area[x][y] = true
				}
			}
		case "turn off":
			for x := s.X; x <= e.X; x++ {
				for y := s.Y; y <= e.Y; y++ {
					area[x][y] = false
				}
			}
		case "toggle":
			for x := s.X; x <= e.X; x++ {
				for y := s.Y; y <= e.Y; y++ {
					area[x][y] = !area[x][y]
				}
			}
		default:
			panic("unknown command")
		}
	}

	output := len(lo.Filter(lo.Flatten(area), func(v bool, _ int) bool { return v }))

	return strconv.Itoa(output)
}

// solve
func solvePart2(input []byte) string {
	lines := bytes.Split(bytes.Trim(input, "\n"), []byte("\n"))
	area := make([][]int, 1000)
	for i := 0; i < 1000; i++ {
		area[i] = make([]int, 1000)
	}
	var s, e image.Point
	for _, line := range lines {
		data := rx.FindStringSubmatch(string(line))
		s.X, _ = strconv.Atoi(data[2])
		s.Y, _ = strconv.Atoi(data[3])
		e.X, _ = strconv.Atoi(data[4])
		e.Y, _ = strconv.Atoi(data[5])
		switch data[1] {
		case "turn on":
			for x := s.X; x <= e.X; x++ {
				for y := s.Y; y <= e.Y; y++ {
					area[x][y] += 1
				}
			}
		case "turn off":
			for x := s.X; x <= e.X; x++ {
				for y := s.Y; y <= e.Y; y++ {
					area[x][y] = max(0, area[x][y]-1)
				}
			}
		case "toggle":
			for x := s.X; x <= e.X; x++ {
				for y := s.Y; y <= e.Y; y++ {
					area[x][y] += 2
				}
			}
		default:
			panic("unknown command")
		}
	}

	output := lo.Sum(lo.Flatten(area))

	return strconv.Itoa(output)
}
