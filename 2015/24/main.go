package main

import (
	"bytes"
	_ "embed"
	"math"
	"slices"
	"strconv"

	"github.com/erka/aoc/pkg/log"
	"github.com/samber/lo"
)

//go:embed input.txt
var input []byte

/*
* part 1: 11846773891
* part 2: 80393059
 */
func main() {
	log.Infof("part1: %s", solvePart1(input))
	log.Infof("part2: %s", solvePart2(input))
}

// solve
func solvePart1(input []byte) string {
	return solve(input, 3)
}

// solve
func solvePart2(input []byte) string {
	return solve(input, 4)
}

func solve(input []byte, groups int) string {
	lines := bytes.Split(bytes.Trim(input, "\n"), []byte("\n"))

	nums := lo.Map(lines, func(item []byte, _ int) int {
		n, _ := strconv.Atoi(string(item))
		return n
	})
	weight := lo.Sum(nums) / groups

	log.Debugf("w: %d nums: %v", weight, nums)
	combinations := findCombinations(nums, weight)
	minimalPackages := lo.Reduce(combinations, func(acc int, item []int, _ int) int {
		return min(acc, len(item))
	}, math.MaxInt)

	log.Debugf("minimal packages: %v", minimalPackages)
	output := lo.Map(combinations, func(a []int, _ int) int64 {
		return lo.Reduce(a, func(acc int64, item int, _ int) int64 { return acc * int64(item) }, int64(1))
	})
	output = lo.Filter(output, func(item int64, _ int) bool {
		return item > 0
	})
	slices.Sort(output)
	return strconv.FormatInt(output[0], 10)
}

func findCombinations(nums []int, target int) [][]int {
	var result [][]int
	var temp []int
	var find func(int, int)

	find = func(start, sum int) {
		if sum == target {
			comb := make([]int, len(temp))
			copy(comb, temp)
			result = append(result, comb)
			return
		}

		for i := start; i < len(nums); i++ {
			if sum+nums[i] <= target {
				temp = append(temp, nums[i])
				find(i+1, sum+nums[i])
				temp = temp[:len(temp)-1]
			}
		}
	}

	find(0, 0)
	return result
}
