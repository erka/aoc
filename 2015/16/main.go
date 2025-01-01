package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"strconv"

	"github.com/erka/aoc/pkg/log"
	"github.com/samber/lo"
)

//go:embed input.txt
var input []byte

/*
* part 1: 103
* part 2:
 */
func main() {
	log.Infof("part1: %s", solvePart1(input))
	log.Infof("part2: %s", solvePart2(input))
}

var options = map[string]int{
	"children:":    3,
	"cats:":        7,
	"samoyeds:":    2,
	"pomeranians:": 3,
	"akitas:":      0,
	"vizslas:":     0,
	"goldfish:":    5,
	"trees:":       3,
	"cars:":        2,
	"perfumes:":    1,
}

// solve
func solvePart1(input []byte) string {
	aunts := parseInput(input)

	aunts = lo.Filter(aunts, func(aunt Aunt, _ int) bool {
		for opt, val := range aunt.attrs {
			if v, ok := options[opt]; ok && v != val {
				return false
			}
		}
		return true
	})

	return strconv.Itoa(aunts[0].num)
}

func parseInput(input []byte) []Aunt {
	lines := bytes.Split(bytes.Trim(input, "\n"), []byte("\n"))
	var aunt, val1, val2, val3 int
	var opt1, opt2, opt3 string
	aunts := lo.Map(lines, func(line []byte, _ int) Aunt {
		_, err := fmt.Sscanf(string(line), "Sue %d: %s %d, %s %d, %s %d",
			&aunt, &opt1, &val1, &opt2, &val2, &opt3, &val3)
		if err != nil {
			panic(err)
		}
		return Aunt{aunt, map[string]int{opt1: val1, opt2: val2, opt3: val3}}
	})
	return aunts
}

// solve
func solvePart2(input []byte) string {
	aunts := parseInput(input)

	aunts = lo.Filter(aunts, func(aunt Aunt, _ int) bool {
		for opt, val := range aunt.attrs {
			if v, ok := options[opt]; ok {
				switch opt {
				case "cats:", "trees:":
					if v >= val {
						return false
					}
				case "pomeranians:", "goldfish:":
					if v <= val {
						return false
					}
				default:
					if v != val {
						return false
					}
				}
			}
		}
		return true
	})

	return strconv.Itoa(aunts[0].num)
}

type Aunt struct {
	num   int
	attrs map[string]int
}
