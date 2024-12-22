package main

import (
	_ "embed"
	"math"
	"strconv"
	"strings"

	"github.com/erka/aoc/pkg/log"
	"github.com/samber/lo"
)

//go:embed input.txt
var input []byte

/*
* part 1: 233050
* part 2: 65601038650482
 */
func main() {
	log.Infof("part1: %s", solvePart1(input))
	log.Infof("part2: %s", solvePart2(input))
}

type Stone struct {
	next  *Stone
	value string
}

func arrangeStones(input []int, times int) int {
	hash := map[int]int{}
	for _, s := range input {
		hash[s] += 1
	}
	for i := 0; i < times; i++ {
		next := make(map[int]int, 0)
		for stone, value := range hash {
			switch {
			case stone == 0:
				next[1] += value
			case numLen(stone)%2 == 0:
				ln := numLen(stone)
				v := int(math.Pow(10, float64(ln/2)))
				left, right := stone/v, stone%v
				next[left] += value
				next[right] += value
			default:
				next[stone*2024] += value
			}
		}
		hash = next
	}

	return lo.Sum(lo.Values(hash))
}

func numLen(i int) int {
	return int(math.Log10(float64(i))) + 1
}

// solve
func solvePart1(input []byte) string {
	stones := lo.Map(strings.Fields(string(input)), func(item string, _ int) int {
		n, _ := strconv.Atoi(item)
		return n
	})
	output := arrangeStones(stones, 25)
	return strconv.Itoa(output)
}

// solve
func solvePart2(input []byte) string {
	stones := lo.Map(strings.Fields(string(input)), func(item string, _ int) int {
		n, _ := strconv.Atoi(item)
		return n
	})
	output := arrangeStones(stones, 75)
	return strconv.Itoa(output)
}
