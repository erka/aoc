package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"image"
	"strconv"

	"github.com/erka/aoc/pkg/log"
)

//go:embed input.txt
var input []byte

/*
* part1: 1413675
* part2: 1399772
 */
func main() {
	log.Infof("part1: %s", solvePart1(input))
	log.Infof("part2: %s", solvePart2(input))
}

var dirs = map[byte]image.Point{
	'^': image.Pt(0, -1),
	'>': image.Pt(1, 0),
	'<': image.Pt(-1, 0),
	'v': image.Pt(0, 1),
	'[': image.Pt(1, 0),
	']': image.Pt(-1, 0),
}

// solve
func solvePart1(input []byte) string {
	in := bytes.Split(input, []byte("\n\n"))
	area := bytes.FieldsFunc(in[0], func(r rune) bool { return r == '\n' })
	movements := in[1]
	warehouse := make(map[image.Point]byte)
	var robot image.Point
	for y, line := range area {
		for x, c := range line {
			p := image.Pt(x, y)
			warehouse[p] = c
			if c == '@' {
				robot = p
			}
		}
	}
	log.Debug("robot position:", robot)
	for _, m := range movements {
		dir, ok := dirs[m]
		if !ok {
			continue
		}

		next := robot.Add(dir)
		switch warehouse[next] {
		case '.':
			warehouse[next] = '@'
			warehouse[robot] = '.'
			robot = next
		case 'O':
			nn := next
			for {
				nn = nn.Add(dir)
				if warehouse[nn] == '.' {
					warehouse[nn] = 'O'
					warehouse[next] = '@'
					warehouse[robot] = '.'
					robot = next
					break
				} else if warehouse[nn] == '#' {
					break
				}
			}
		}
	}

	output := 0
	for p, c := range warehouse {
		if c == 'O' {
			output += 100*p.Y + p.X
		}
	}

	return strconv.Itoa(output)
}

// solve
func solvePart2(input []byte) string {
	in := bytes.Split(input, []byte("\n\n"))
	area := bytes.FieldsFunc(in[0], func(r rune) bool { return r == '\n' })
	movements := in[1]
	warehouse := make(map[image.Point]byte)
	var robot image.Point
	for y, line := range area {
		x := 0
		for _, c := range line {
			p := image.Pt(x, y)
			switch c {
			case '#', '.':
				warehouse[p] = c
				warehouse[p.Add(dirs['>'])] = c
			case '@':
				warehouse[p] = c
				warehouse[p.Add(dirs['>'])] = '.'
			case 'O':
				warehouse[p] = '['
				warehouse[p.Add(dirs['>'])] = ']'

			}
			if c == '@' {
				robot = p
			}
			x += 2
		}
	}

	for _, m := range movements {
		dir, ok := dirs[m]
		if !ok {
			continue
		}

		next := robot.Add(dir)
		switch warehouse[next] {
		case '.':
			warehouse[next] = '@'
			warehouse[robot] = '.'
			robot = next
		case '[', ']':
			boxes := map[image.Point]byte{
				next:                            warehouse[next],
				next.Add(dirs[warehouse[next]]): warehouse[next.Add(dirs[warehouse[next]])],
			}
			queue := []image.Point{next, next.Add(dirs[warehouse[next]])}

			for len(queue) > 0 {
				b := queue[0]
				queue = queue[1:]
				if _, box := boxes[b]; !box {
					continue
				}
				p := b.Add(dir)
				switch warehouse[p] {
				case '.':
					if m == '<' || m == '>' {
						queue = nil
					}
				case '#':
					boxes = nil
					queue = nil
				case '[', ']':
					boxes[p] = warehouse[p]
					boxes[p.Add(dirs[warehouse[p]])] = warehouse[p.Add(dirs[warehouse[p]])]
					queue = append(queue, p, p.Add(dirs[warehouse[p]]))
				}
			}

			if len(boxes) > 0 {
				for b := range boxes {
					warehouse[b] = '.'
				}
				for b, c := range boxes {
					warehouse[b.Add(dir)] = c
				}
				warehouse[next] = '@'
				warehouse[robot] = '.'
				robot = next
			}
		}
	}

	output := 0
	for p, c := range warehouse {
		if c == '[' {
			output += 100*p.Y + p.X
		}
	}
	return strconv.Itoa(output)
}

func printWarehouse(area [][]byte, warehouse map[image.Point]byte) {
	for y, line := range area {
		x := 0
		for range line {
			p := image.Pt(x, y)
			fmt.Print(string(warehouse[p]))
			fmt.Print(string(warehouse[p.Add(dirs['>'])]))
			x += 2
		}
		fmt.Println("")
	}
}
