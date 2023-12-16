package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"slices"
	"strconv"
	"strings"

	"github.com/erka/aoc/pkg/log"
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
	empty     = '.'
	splitterH = '-'
	splitterV = '|'
	mirrorL   = '/'
	mirrorR   = '\\'

	ud direction = 'u'
	dd direction = 'd'
	rd direction = 'r'
	ld direction = 'l'
)

type direction rune

type point struct {
	x, y int
}

type contraption struct {
	grid      [][]rune
	energized map[point][]direction
}

func (c *contraption) reset() {
	c.energized = map[point][]direction{}
}

func (c *contraption) energizedMax() int {
	value := 0
	for y := 0; y < len(c.grid[0]); y++ {
		c.energize(point{0, y}, dd)
		value = max(value, c.energizedTiles())
		c.reset()
		c.energize(point{0, len(c.grid[0]) - y - 1}, ud)
		value = max(value, c.energizedTiles())
		c.reset()
	}

	for x := 0; x < len(c.grid); x++ {
		c.energize(point{x, 0}, rd)
		value = max(value, c.energizedTiles())
		c.reset()
		c.energize(point{len(c.grid) - x - 1, 0}, ld)
		value = max(value, c.energizedTiles())
		c.reset()
	}
	return value
}

func (c *contraption) energize(beam point, dir direction) {
	for {
		if beam.x >= len(c.grid[0]) || beam.x < 0 ||
			beam.y < 0 || beam.y >= len(c.grid) {
			break
		}
		if dirs, ok := c.energized[beam]; !ok {
			c.energized[point{beam.x, beam.y}] = []direction{dir}
		} else {
			if slices.Contains(dirs, dir) {
				break
			}
			c.energized[point{beam.x, beam.y}] = append(dirs, dir)
		}
		switch c.grid[beam.x][beam.y] {
		case empty:
			switch dir {
			case ud:
				beam.x -= 1
			case dd:
				beam.x += 1
			case ld:
				beam.y -= 1
			case rd:
				beam.y += 1
			}
		case mirrorL:
			switch dir {
			case rd:
				dir = ud
				beam.x -= 1
			case ld:
				dir = dd
				beam.x += 1
			case dd:
				dir = ld
				beam.y -= 1
			case ud:
				dir = rd
				beam.y += 1
			}
		case mirrorR:
			switch dir {
			case rd:
				dir = dd
				beam.x += 1
			case ld:
				dir = ud
				beam.x -= 1
			case dd:
				dir = rd
				beam.y += 1
			case ud:
				dir = ld
				beam.y -= 1
			}
		case splitterH:
			switch dir {
			case ud, dd:
				c.energize(point{beam.x, beam.y + 1}, rd)
				dir = ld
				beam.y -= 1
			case rd:
				beam.y += 1
			case ld:
				beam.y -= 1
			}
		case splitterV:
			switch dir {
			case rd, ld:
				c.energize(point{beam.x + 1, beam.y}, dd)
				dir = ud
				beam.x -= 1
			case ud:
				beam.x -= 1
			case dd:
				beam.x += 1
			}
		}
	}
}

func (c *contraption) energizedTiles() int {
	return len(c.energized)
}
func (c *contraption) energizedString() string {
	b := strings.Builder{}
	b.WriteRune('\n')
	for x, cc := range c.grid {
		for y := range cc {
			if _, ok := c.energized[point{x, y}]; ok {
				b.WriteRune('#')
				continue
			}
			b.WriteRune('.')
		}
		b.WriteRune('\n')
	}
	return b.String()
}

func (c *contraption) String() string {
	b := strings.Builder{}
	b.WriteRune('\n')

	for x, cc := range c.grid {
		for y := range cc {
			b.WriteRune(c.grid[x][y])
		}
		b.WriteRune('\n')
	}
	return b.String()
}

// solve
func solvePart1(input []byte) string {
	scanner := bufio.NewScanner(bytes.NewBuffer(input))
	c := &contraption{
		energized: map[point][]direction{},
	}
	for scanner.Scan() {
		line := scanner.Text()
		c.grid = append(c.grid, []rune(line))
	}
	if err := scanner.Err(); err != nil {
		log.Errorf("failed to read input: %v", err)
		return ""
	}
	c.energize(point{0, 0}, rd)
	log.Debugf("%v", c.energizedString())
	return strconv.Itoa(c.energizedTiles())
}

// solve
func solvePart2(input []byte) string {
	scanner := bufio.NewScanner(bytes.NewBuffer(input))
	c := &contraption{}
	c.reset()
	for scanner.Scan() {
		line := scanner.Text()
		c.grid = append(c.grid, []rune(line))
	}
	if err := scanner.Err(); err != nil {
		log.Errorf("failed to read input: %v", err)
		return ""
	}
	log.Debugf("%v", c.energizedString())
	return strconv.Itoa(c.energizedMax())
}
