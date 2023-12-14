package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input []byte

/*
* part 1: 109596
* part 2: 96105
 */
func main() {
	fmt.Printf("part1: %s\n", solvePart1(input))
	fmt.Printf("part2: %s\n", solvePart2(input))
}

type ControlPanel [][]rune

func (c ControlPanel) tiltNorth() {
	d := len(c)
	for x := range c[0] {
		freeSpot := 0
		for y := 0; y < d; y++ {
			switch c[y][x] {
			case 'O':
				c[y][x] = '.'
				c[freeSpot][x] = 'O'
				freeSpot += 1
			case '#':
				freeSpot = y + 1
			}
		}
	}
}

func (c ControlPanel) tiltSouth() {
	d := len(c)
	for x := range c[0] {
		freeSpot := d - 1
		for y := d - 1; y >= 0; y-- {
			switch c[y][x] {
			case 'O':
				c[y][x] = '.'
				c[freeSpot][x] = 'O'
				freeSpot -= 1
			case '#':
				freeSpot = y - 1
			}
		}
	}
}

func (c ControlPanel) tiltWest() {
	for x := range c {
		freeSpot := 0
		for y := 0; y < len(c[0]); y++ {
			switch c[x][y] {
			case 'O':
				c[x][y] = '.'
				c[x][freeSpot] = 'O'
				freeSpot += 1
			case '#':
				freeSpot = y + 1
			}
		}
	}
}

func (c ControlPanel) tiltEast() {
	for x := range c {
		freeSpot := len(c[0]) - 1
		for y := len(c[0]) - 1; y >= 0; y-- {
			switch c[x][y] {
			case 'O':
				c[x][y] = '.'
				c[x][freeSpot] = 'O'
				freeSpot -= 1
			case '#':
				freeSpot = y - 1
			}
		}
	}
}

func (c ControlPanel) cycle() {
	c.tiltNorth()
	c.tiltWest()
	c.tiltSouth()
	c.tiltEast()
}

func (c ControlPanel) totalLoad() int {
	d := len(c)
	sum := 0
	for x := range c[0] {
		for y := 0; y < d; y++ {
			if c[y][x] == 'O' {
				sum += d - y
			}
		}
	}
	return sum
}

func (c ControlPanel) String() string {
	b := strings.Builder{}
	for _, r := range c {
		b.WriteString(string(r))
		b.WriteByte('\n')
	}
	return b.String()
}

// solve
func solvePart1(input []byte) string {
	scanner := bufio.NewScanner(bytes.NewBuffer(input))
	controlPanel := ControlPanel{}
	for scanner.Scan() {
		line := scanner.Text()
		controlPanel = append(controlPanel, []rune(line))
	}
	if err := scanner.Err(); err != nil {
		return err.Error()
	}
	controlPanel.tiltNorth()
	return strconv.Itoa(controlPanel.totalLoad())
}

// solve
func solvePart2(input []byte) string {
	scanner := bufio.NewScanner(bytes.NewBuffer(input))
	controlPanel := ControlPanel{}
	for scanner.Scan() {
		line := scanner.Text()
		controlPanel = append(controlPanel, []rune(line))
	}
	if err := scanner.Err(); err != nil {
		return err.Error()
	}
	vi, wi := 0, 0
	cycles := make(map[string]int)
	for i := 1; ; i++ {
		controlPanel.cycle()
		k := controlPanel.String()
		if v, ok := cycles[k]; ok {
			vi = v
			wi = i - v
			break
		}
		cycles[k] = i
	}
	for i := 0; i < (1_000_000_000-vi)%wi; i++ {
		controlPanel.cycle()
	}
	return strconv.Itoa(controlPanel.totalLoad())
}
