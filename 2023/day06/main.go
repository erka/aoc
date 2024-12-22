package main

import (
	_ "embed"
	"fmt"
	"strconv"
)

//go:embed input.txt
var input []byte

type BestRace struct {
	Time     int
	Distance int
}

/*
* part 1: 5133600
* part 2: 40651271
 */
func main() {
	inputValue := []BestRace{
		{53, 313},
		{89, 1090},
		{76, 1214},
		{98, 1201},
	}
	fmt.Printf("part1: %s\n", solve(inputValue))
	fmt.Printf("part2: %s\n", solve([]BestRace{{
		53897698, 313109012141201,
	}}))
}

// solve
func solve(input []BestRace) string {
	out := 1
	for _, br := range input {
		options := 0
		for i := 1; i < br.Time; i++ {
			if (i * (br.Time - i)) > br.Distance {
				options++
			}
		}
		if options > 0 {
			out *= options
		}
	}
	return strconv.Itoa(out)
}
