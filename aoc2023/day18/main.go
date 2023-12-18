package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"image"
	"regexp"
	"strconv"

	"github.com/erka/aoc/pkg/log"
)

//go:embed input.txt
var input []byte

/*
 *	part1: 46394
 *	part2: 201398068194715
 */
func main() {
	log.Infof("part1: %s", solvePart1(input))
	log.Infof("part2: %s", solvePart2(input))
}

type record struct {
	direction string
	meters    int
}
type digPlan []record

func (plan digPlan) area() int {
	start, perimeter := image.Point{0, 0}, 0
	points := []image.Point{start}
	for _, r := range plan {
		n := start.Add(instructions[r.direction].Mul(r.meters))
		perimeter += r.meters
		start = n
		points = append(points, n)
	}
	return prick(shoelace(points), perimeter)
}

// prick's theorem
func prick(inside int, perimeter int) int {
	return (inside+perimeter)/2 + 1
}

// shoelace algorithm
func shoelace(points []image.Point) int {
	sum := 0
	p0 := points[len(points)-1]
	for _, p1 := range points {
		sum += p0.X*p1.Y - p0.Y*p1.X
		p0 = p1
	}
	return sum
}

var instructions = map[string]image.Point{
	"U": {0, -1}, "R": {1, 0}, "D": {0, 1}, "L": {-1, 0},
	"3": {0, -1}, "0": {1, 0}, "1": {0, 1}, "2": {-1, 0},
}

// solve
func solvePart1(input []byte) string {
	var re = regexp.MustCompile(`([URDL]) (\d+) .+`)
	scanner := bufio.NewScanner(bytes.NewBuffer(input))
	plan := make(digPlan, 0)
	for scanner.Scan() {
		line := scanner.Text()
		values := re.FindStringSubmatch(line)
		if len(values) != 3 {
			panic(line)
		}
		volume, err := strconv.Atoi(values[2])
		if err != nil {
			panic(err)
		}
		plan = append(plan, record{values[1], volume})
	}
	if err := scanner.Err(); err != nil {
		log.Errorf("failed to read input: %v", err)
		return ""
	}

	return strconv.Itoa(plan.area())
}

// solve
func solvePart2(input []byte) string {
	var re = regexp.MustCompile(`.+ \(#(\w{5})(\d)\)`)
	scanner := bufio.NewScanner(bytes.NewBuffer(input))
	plan := make(digPlan, 0)
	for scanner.Scan() {
		line := scanner.Text()
		values := re.FindStringSubmatch(line)
		if len(values) != 3 {
			panic(line)
		}
		volume, err := strconv.ParseInt(values[1], 16, strconv.IntSize)
		if err != nil {
			panic(err)
		}
		plan = append(plan, record{values[2], int(volume)})
	}
	if err := scanner.Err(); err != nil {
		log.Errorf("failed to read input: %v", err)
		return ""
	}
	return strconv.Itoa(plan.area())
}
