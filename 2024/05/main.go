package main

import (
	"bytes"
	_ "embed"
	"slices"
	"strconv"
	"strings"

	"github.com/erka/aoc/pkg/log"
	"github.com/samber/lo"
)

//go:embed input.txt
var input []byte

type Rule struct {
	X, Y string
}
type RuleSet map[string][]string

/*
* part 1: 5129
* part 2: 4077
 */
func main() {
	log.Infof("part1: %s", solvePart1(input))
	log.Infof("part2: %s", solvePart2(input))
}

func process(input []byte) []int {
	data := bytes.Split(input, []byte{'\n', '\n'})

	rulesSet := buildRuleSet(parseRules(data[0]))

	return lo.Map(bytes.Split(data[1], []byte{'\n'}), func(line []byte, index int) int {
		if len(line) == 0 {
			return 0
		}
		return evaluateLine(line, rulesSet)
	})
}

func parseRules(data []byte) []Rule {
	return lo.Map(bytes.Split(data, []byte{'\n'}), func(item []byte, index int) Rule {
		values := strings.Split(string(item), "|")
		return Rule{
			X: values[0],
			Y: values[1],
		}
	})
}

func buildRuleSet(rules []Rule) RuleSet {
	return lo.Reduce(rules, func(agg RuleSet, item Rule, index int) RuleSet {
		if _, ok := agg[item.X]; !ok {
			agg[item.X] = []string{}
		}
		agg[item.X] = append(agg[item.X], item.Y)
		return agg
	}, RuleSet{})
}

func evaluateLine(line []byte, rulesSet RuleSet) int {
	pages := strings.Split(string(line), ",")
	for i, page := range pages {
		if pageRules, ok := rulesSet[page]; ok {
			for _, rule := range pageRules {
				if slices.Contains(pages[0:i], rule) {
					slices.SortFunc(pages, func(a string, b string) int {
						if p, ok := rulesSet[a]; ok {
							if slices.Contains(p, b) {
								return -1
							}
						}
						return 1
					})
					return -mustInt(pages)
				}
			}
		}
	}
	return mustInt(pages)
}

func mustInt(pages []string) int {
	v, err := strconv.Atoi(pages[len(pages)/2])
	if err != nil {
		panic(err)
	}
	return v
}

// solve
func solvePart1(input []byte) string {
	output := lo.Sum(lo.Filter(process(input), func(item int, index int) bool { return item > 0 }))
	return strconv.Itoa(output)
}

// solve
func solvePart2(input []byte) string {
	output := -lo.Sum(lo.Filter(process(input), func(item int, index int) bool { return item < 0 }))
	return strconv.Itoa(output)
}
