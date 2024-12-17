package main

import (
	_ "embed"
	"image"
	"strconv"

	"github.com/erka/aoc/pkg/log"
)

//go:embed input.txt
var input []byte

/*
* part1: 2565
* part2: 2639
 */
func main() {
	log.Infof("part1: %s", solvePart1(input))
	log.Infof("part2: %s", solvePart2(input))
}

func solvePart1(input []byte) string {
	santa := image.Pt(0, 0)
	seen := map[image.Point]struct{}{santa: {}}
	directions := map[byte]image.Point{
		'^': image.Pt(-1, 0),
		'<': image.Pt(0, -1),
		'>': image.Pt(0, 1),
		'v': image.Pt(1, 0),
	}
	for _, c := range input {
		santa = santa.Add(directions[c])
		if _, ok := seen[santa]; !ok {
			seen[santa] = struct{}{}
		}
	}
	return strconv.Itoa(len(seen))
}

func solvePart2(input []byte) string {
	santa := image.Pt(0, 0)
	robot := image.Pt(0, 0)
	seen := map[image.Point]struct{}{santa: {}}
	directions := map[byte]image.Point{
		'^': image.Pt(-1, 0),
		'<': image.Pt(0, -1),
		'>': image.Pt(0, 1),
		'v': image.Pt(1, 0),
	}
	for i, c := range input {
		if i%2 == 1 {
			santa = santa.Add(directions[c])
			if _, ok := seen[santa]; !ok {
				seen[santa] = struct{}{}
			}
		} else {
			robot = robot.Add(directions[c])
			if _, ok := seen[robot]; !ok {
				seen[robot] = struct{}{}
			}
		}
	}
	return strconv.Itoa(len(seen))
}
