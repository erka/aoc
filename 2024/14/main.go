package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"image"
	"strconv"

	"github.com/erka/aoc/pkg/log"
	"github.com/samber/lo"
)

//go:embed input.txt
var input []byte

/*
* part 1: 229868730
* part 2: 7861
 */
func main() {
	boundary := image.Rect(0, 0, 101, 103)
	log.Infof("part1: %s", solvePart1(input, boundary))
	log.Infof("part2: %s", solvePart2(input, boundary))
}

type Robot struct {
	position image.Point
	velocity image.Point
}

func (r *Robot) move(boundary image.Rectangle, times int) {
	next := r.position.Add(r.velocity.Mul(times))
	if !next.In(boundary) {
		// teleport
		next = next.Mod(boundary)
	}
	r.position = next
}

// solve
func solvePart1(input []byte, boundary image.Rectangle) string {
	robots := lo.Map(bytes.FieldsFunc(input, func(r rune) bool { return r == '\n' }), func(line []byte, _ int) *Robot {
		var p, v image.Point
		_, _ = fmt.Sscanf(string(line), "p=%d,%d v=%d,%d", &p.X, &p.Y, &v.X, &v.Y)
		return &Robot{
			position: p,
			velocity: v,
		}
	})
	for _, r := range robots {
		r.move(boundary, 100)
	}
	quadrants := map[image.Rectangle]int{
		image.Rect(0, 0, boundary.Dx()/2, boundary.Dy()/2):                               0,
		image.Rect(boundary.Dx()/2+1, 0, boundary.Max.X, boundary.Dy()/2):                0,
		image.Rect(0, boundary.Dy()/2+1, boundary.Dx()/2, boundary.Max.Y):                0,
		image.Rect(boundary.Dx()/2+1, boundary.Dy()/2+1, boundary.Max.X, boundary.Max.Y): 0,
	}
	for _, r := range robots {
		for q := range quadrants {
			if r.position.In(q) {
				quadrants[q] += 1
				break
			}
		}
	}
	output := 1
	for _, v := range quadrants {
		output *= v
	}
	return strconv.Itoa(output)
}

// solve
func solvePart2(input []byte, boundary image.Rectangle) string {
	robots := lo.Map(bytes.FieldsFunc(input, func(r rune) bool { return r == '\n' }), func(line []byte, _ int) *Robot {
		var p, v image.Point
		_, _ = fmt.Sscanf(string(line), "p=%d,%d v=%d,%d", &p.X, &p.Y, &v.X, &v.Y)
		return &Robot{
			position: p,
			velocity: v,
		}
	})
	i := 1
	for ; ; i++ {
		seen := map[image.Point]struct{}{}
		for _, r := range robots {
			r.move(boundary, 1)
			seen[r.position] = struct{}{}
		}

		if len(seen) == len(robots) {
			break
		}
	}
	return strconv.Itoa(i)
}
