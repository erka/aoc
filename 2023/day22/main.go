package main

import (
	"bufio"
	"bytes"
	"cmp"
	_ "embed"
	"fmt"
	"slices"
	"strconv"

	"github.com/erka/aoc/pkg/log"
	"github.com/samber/lo"
)

//go:embed input.txt
var input []byte

/*
* part 1:
* part 2:
 */
func main() {
	log.Infof("part1: %s", solvePart1(input))
	log.Infof("part2: %s", solvePart2(input))
}

const (
	X = 0
	Y = 1
	Z = 2
)

type Point struct {
	x, y, z int
}

type Brick struct {
	p1 Point
	p2 Point
}

func (b *Brick) MoveToZ(z int) {
	b.p2.z = z + b.p2.z - b.p1.z
	b.p1.z = z
}

func (b *Brick) OverlapsXY(a *Brick) bool {
	return max(a.p1.x, b.p1.x) <= min(a.p2.x, b.p2.x) &&
		max(a.p1.y, b.p1.y) <= min(a.p2.y, b.p2.y)
}
func (b *Brick) SupportedBy(a *Brick) bool {
	return b.OverlapsXY(a) && a.p2.z+1 == b.p1.z
}

func sortByMinZ(a *Brick, b *Brick) int {
	return cmp.Compare(a.p1.z, b.p1.z)
}

func create(input []byte) []*Brick {
	scanner := bufio.NewScanner(bytes.NewBuffer(input))
	bricks := []*Brick{}
	for scanner.Scan() {
		line := scanner.Text()
		log.Debug(line)
		brick := &Brick{}
		_, err := fmt.Sscanf(line, "%d,%d,%d~%d,%d,%d",
			&brick.p1.x, &brick.p1.y, &brick.p1.z,
			&brick.p2.x, &brick.p2.y, &brick.p2.z,
		)
		if err != nil {
			log.Error(err)
		}
		bricks = append(bricks, brick)
	}
	if err := scanner.Err(); err != nil {
		log.Errorf("failed to read input: %v", err)
		panic("")
	}
	slices.SortFunc(bricks, sortByMinZ)
	return bricks
}

func fall(bricks []*Brick) int {
	log.Debug("bricks ", len(bricks))
	fallen := 0
	for i, b := range bricks {
		z := 0
		for j := 0; j < i; j++ {
			if b.OverlapsXY(bricks[j]) {
				log.Debugf("%v overlaps %v", strconv.QuoteRune('A'+rune(i)), strconv.QuoteRune('A'+rune(j)))
				z = max(z, bricks[j].p2.z+1)
			}
		}
		if b.p1.z != z {
			b.MoveToZ(z)
			fallen += 1
		}
	}
	return fallen
}

func relations(bricks []*Brick) (map[int][]int, map[int][]int) {
	supports := make(map[int][]int)
	supported_by := make(map[int][]int)
	for j, upper := range bricks {
		for i, lower := range bricks {
			if upper.SupportedBy(lower) {
				log.Debug(i, " supports ", j)
				if _, ok := supports[i]; !ok {
					supports[i] = []int{}
				}
				supports[i] = append(supports[i], j)

				if _, ok := supported_by[j]; !ok {
					supported_by[j] = []int{}
				}
				supported_by[j] = append(supported_by[j], i)

			}
		}
	}
	return supports, supported_by
}

// solve
func solvePart1(input []byte) string {
	bricks := create(input)
	fall(bricks)
	supports, supported_by := relations(bricks)
	result := 0
	for i := range bricks {
		disintegrated := true
		for _, k := range supports[i] {
			if len(supported_by[k]) == 1 {
				disintegrated = false
				break
			}
		}
		if disintegrated {
			result += 1
		}
	}
	return strconv.Itoa(result)
}

// solve
func solvePart2(input []byte) string {
	bricks := create(input)
	fall(bricks)
	supports, supported_by := relations(bricks)
	disintegrated := lo.Filter[*Brick](bricks, func(item *Brick, i int) bool {
		for _, k := range supports[i] {
			return len(supported_by[k]) == 1
		}
		return false
	})
	result := 0

	for _, d := range disintegrated {
		index := slices.Index(bricks, d)
		result += fall(bricks[index+1:])
	}

	return strconv.Itoa(result)
}
