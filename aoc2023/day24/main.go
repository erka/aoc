package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"math/big"
	"strconv"

	"github.com/erka/aoc/pkg/log"
)

//go:embed input.txt
var input []byte

/*
* part 1:
* 17122 low
* part 2:
 */
func main() {
	log.Infof("part1: %s", solvePart1(input, 200000000000000, 400000000000000))
	log.Infof("part2: %s", solvePart2(input))
}

type Area struct {
	xyMin, xyMax int64
}

func (a Area) Contains(p point) bool {
	return a.xyMin <= p.x && p.x <= a.xyMax && a.xyMin <= p.y && p.y <= a.xyMax
}

type point struct {
	x, y int64
}

type hailstone struct {
	px, py, vx, vy int64
}

func (h hailstone) c() *big.Int {
	c1 := new(big.Int).Mul(big.NewInt(h.vy), big.NewInt(h.px))
	c2 := new(big.Int).Mul(big.NewInt(h.vx), big.NewInt(h.py))
	return new(big.Int).Sub(c1, c2)
}

func (h hailstone) hasFuturePoint(p point) bool {
	return (p.x-h.px)*h.vx >= 0 && (p.y-h.py)*h.vy >= 0
}

// collideForward returns future cross point and true if future collide is happens
func (h hailstone) collideForward(o hailstone) (point, bool) {
	a1, b1, c1 := big.NewInt(h.vy), big.NewInt(-h.vx), h.c()
	a2, b2, c2 := big.NewInt(o.vy), big.NewInt(-o.vx), o.c()

	if new(big.Int).Mul(a1, b2).Cmp(new(big.Int).Mul(b1, a2)) == 0 {
		log.Debug("line do not intersect")
		//If cp = 0, the lines do not intersect.
		return point{}, false
	}
	//x := (c1*b2 - c2*b1) / (a1*b2 - a2*b1)
	x := new(big.Int).Div(
		new(big.Int).Sub(
			new(big.Int).Mul(c1, b2),
			new(big.Int).Mul(c2, b1),
		),
		new(big.Int).Sub(
			new(big.Int).Mul(a1, b2),
			new(big.Int).Mul(a2, b1),
		),
	)
	//y := (c2*a1 - c1*a2) / (a1*b2 - a2*b1)
	y := new(big.Int).Div(
		new(big.Int).Sub(
			new(big.Int).Mul(c2, a1),
			new(big.Int).Mul(c1, a2),
		),
		new(big.Int).Sub(
			new(big.Int).Mul(a1, b2),
			new(big.Int).Mul(a2, b1),
		),
	)
	cross := point{x.Int64(), y.Int64()}
	return cross, h.hasFuturePoint(cross) && o.hasFuturePoint(cross)
}

// solve
func solvePart1(input []byte, areaMin, areaMax int64) string {
	scanner := bufio.NewScanner(bytes.NewBuffer(input))
	hail := []hailstone{}
	for scanner.Scan() {
		line := scanner.Text()
		log.Debug(line)
		var px, py, pz, vx, vy, vz int64
		_, err := fmt.Sscanf(
			line,
			"%d, %d, %d @ %d, %d, %d",
			&px, &py, &pz, &vx, &vy, &vz,
		)
		if err != nil {
			log.Error(err)
		}
		hail = append(hail, hailstone{px: px, py: py, vx: vx, vy: vy})
	}
	if err := scanner.Err(); err != nil {
		log.Errorf("failed to read input: %v", err)
		return ""
	}

	area := Area{xyMin: areaMin, xyMax: areaMax}
	counter, inarea, t := 0, 0, 0
	for i, a := range hail {
		for _, b := range hail[i+1:] {
			t++
			log.Debug("\n")
			log.Debug(a)
			log.Debug(b)
			crossPoint, collided := a.collideForward(b)
			if collided {
				counter += 1
			}

			if collided && area.Contains(crossPoint) {
				inarea += 1
			}
		}
	}
	log.Infof("t %v total: %v in area: %v", t, counter, inarea)
	return strconv.Itoa(inarea)
}

// solve
func solvePart2(input []byte) string {
	scanner := bufio.NewScanner(bytes.NewBuffer(input))
	for scanner.Scan() {
		line := scanner.Text()
		log.Debug(line)
	}
	if err := scanner.Err(); err != nil {
		return err.Error()
	}
	return strconv.Itoa(0)
}
