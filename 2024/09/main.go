package main

import (
	_ "embed"
	"slices"
	"strconv"

	"github.com/erka/aoc/pkg/log"
	"github.com/samber/lo"
)

//go:embed input.txt
var input []byte

/*
* part 1: 6241633730082
* part 2: 6265268809555
 */
func main() {
	log.Infof("part1: %s", solvePart1(input))
	log.Infof("part2: %s", solvePart2(input))
}

func num(b byte) int {
	return int(b - '0')
}

// solve
func solvePart1(input []byte) string {
	lastBlock := len(input) - 2
	size := num(input[lastBlock])
	output := 0
	index := 0
	for i := 0; i < lastBlock; i += 2 {
		for head := num(input[i]); head > 0; head -= 1 {
			output += i / 2 * index
			index += 1
		}
		for space := num(input[i+1]); space > 0; space-- {
			output += lastBlock / 2 * index
			index += 1
			size -= 1
			if size < 1 {
				lastBlock -= 2
				size = num(input[lastBlock])
			}
		}
	}
	// adding remaining
	for ; size > 0; size -= 1 {
		val := int(lastBlock/2) * index
		output += val
		index += 1
	}

	return strconv.Itoa(output)
}

// solve
func solvePart2(input []byte) string {
	input[len(input)-1] = '0'
	blocks := []int{}
	for i := 0; i < len(input); i += 2 {
		for head := num(input[i]); head > 0; head -= 1 {
			blocks = append(blocks, i/2)
		}

		for space := num(input[i+1]); space > 0; space -= 1 {
			blocks = append(blocks, -1)
		}
	}

	for e := len(input) - 2; e > 2; e -= 2 {
		size := num(input[e])
		file := e / 2
		fileIndex := slices.Index(blocks, file)
		for i := 0; i < fileIndex; {
			spaceStart := slices.Index(blocks[i:], -1)
			if spaceStart == -1 || spaceStart+i >= fileIndex {
				break
			}
			spaceStart = spaceStart + i
			spaceEnd := spaceStart
			for k := spaceStart; blocks[k] == -1 && spaceEnd-spaceStart < size; k += 1 {
				spaceEnd += 1
			}
			if spaceEnd-spaceStart < size {
				i = spaceEnd
				continue
			}
			for k := 0; k < size; k += 1 {
				blocks[spaceStart+k] = file
				blocks[fileIndex+k] = -2
			}
			break
		}
	}

	output := lo.Reduce(blocks, func(agg int, item int, index int) int {
		if item < 0 {
			return agg
		}
		return agg + item*index
	}, 0)

	return strconv.Itoa(output)
}
