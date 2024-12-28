package main

import (
	_ "embed"
	"encoding/json"
	"regexp"
	"slices"
	"strconv"

	"github.com/erka/aoc/pkg/log"
	"github.com/samber/lo"
)

//go:embed input.txt
var input []byte

/*
* part 1: 119433
* part 2: 68466
 */
func main() {
	log.Infof("part1: %s", solvePart1(input))
	log.Infof("part2: %s", solvePart2(input))
}

// solve
func solvePart1(input []byte) string {
	re := regexp.MustCompile(`(-?\d+)`)
	matches := re.FindAllString(string(input), -1)
	output := lo.Sum(lo.Map(matches, func(s string, _ int) int {
		n, _ := strconv.Atoi(s)
		return n
	}))
	return strconv.Itoa(output)
}

// solve
func solvePart2(input []byte) string {
	var data any
	err := json.Unmarshal(input, &data)
	if err != nil {
		log.Errorf("error: %v", err)
	}
	output := sum(data)
	return strconv.Itoa(output)
}

func sum(in any) int {
	s := 0
	switch v := in.(type) {
	case float64:
		s += int(v)
	case int:
		s += v
	case []any:
		s += lo.SumBy(v, sum)
	case map[string]any:
		if slices.Contains(lo.Values(v), "red") {
			return 0
		}
		s += lo.SumBy(lo.Values(v), sum)
	}

	return s
}
