package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"strconv"

	"github.com/erka/aoc/pkg/log"
	"github.com/samber/lo"
)

//go:embed input.txt
var input []byte

/*
* part 1: 2640
* part 2: 1102
 */
func main() {
	log.Infof("part1: %s", solvePart1(input, 2503))
	log.Infof("part2: %s", solvePart2(input, 2503))
}

// solve
func solvePart1(input []byte, seconds int) string {
	reindeers := reindeers(input)
	longest := 0
	for _, r := range reindeers {
		d := r.move(seconds)
		longest = max(d, longest)
	}
	return strconv.Itoa(longest)
}

// solve
func solvePart2(input []byte, seconds int) string {
	reindeers := reindeers(input)
	for s := 1; s < seconds; s++ {
		longest := 0
		for _, r := range reindeers {
			longest = max(longest, r.move(s))
		}
		for _, r := range reindeers {
			if longest == r.distance {
				r.extra += 1
			}
		}
	}
	winner := lo.MaxBy(reindeers, func(a *reindeer, b *reindeer) bool {
		return a.extra > b.extra
	})
	return strconv.Itoa(winner.extra)
}

func reindeers(input []byte) []*reindeer {
	lines := bytes.Split(bytes.Trim(input, "\n"), []byte("\n"))
	reindeers := make([]*reindeer, len(lines))
	for i, line := range lines {
		r := &reindeer{}
		fmt.Sscanf(string(line), "%s can fly %d km/s for %d seconds, but then must rest for %d seconds.", &r.name, &r.speed, &r.duration, &r.rest)
		reindeers[i] = r
	}
	return reindeers
}

type reindeer struct {
	name     string
	speed    int
	duration int
	rest     int
	extra    int
	distance int
}

func (r *reindeer) move(seconds int) int {
	r.distance = seconds/(r.duration+r.rest)*r.speed*r.duration + min(r.duration, (seconds%(r.duration+r.rest)))*r.speed
	return r.distance
}
