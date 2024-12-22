package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func newSeed(input string) *Seed {
	seedNum, _ := strconv.Atoi(input)
	return &Seed{location: seedNum}
}

type Seed struct {
	location int
}

func (s *Seed) Transform(destination int, source int, rangeLen int) bool {
	if source <= s.location && s.location < source+rangeLen {
		s.location += destination - source
		return true
	}
	return false
}

//go:embed input.txt
var input []byte

/*
* part 1: 227653707
* part 2:
 */
func main() {
	fmt.Printf("part1: %s\n", solve(input))
	fmt.Printf("part2: %s\n", solvePart2(input))
}

// solve
func solve(input []byte) string {
	scanner := bufio.NewScanner(bytes.NewBuffer(input))
	scanner.Scan()
	line := scanner.Text()
	mapping := Transform(scanner)
	minValue := math.MaxInt64
	for _, seed := range strings.Fields(strings.TrimPrefix(line, "seeds: ")) {
		s := newSeed(seed)
		for _, m := range mapping {
			for _, t := range m {
				if s.Transform(t[0], t[1], t[2]) {
					break
				}
			}
		}
		minValue = min(minValue, s.location)
	}
	return strconv.Itoa(minValue)
}

// solve
func solvePart2(input []byte) string {
	scanner := bufio.NewScanner(bytes.NewBuffer(input))
	scanner.Scan()
	line := scanner.Text()
	mapping := Transform(scanner)
	minValue := math.MaxInt64
	values := strings.Fields(strings.TrimPrefix(line, "seeds: "))
	s := &Seed{}
	for i := 0; i < len(values); i += 2 {
		start, _ := strconv.Atoi(values[i])
		length, _ := strconv.Atoi(values[i+1])
		for i := 0; i < length; i++ {
			s.location = start + i
			for _, m := range mapping {
				for _, t := range m {
					if s.Transform(t[0], t[1], t[2]) {
						break
					}
				}
			}
			minValue = min(minValue, s.location)
		}
	}
	return strconv.Itoa(minValue)
}

func Transform(scanner *bufio.Scanner) [][][3]int {
	tables := [][][3]int{}
	scanner.Scan()
	line := scanner.Text()
	for scanner.Scan() {
		if line != "" {
			panic("here.." + line)
		}
		line = scanner.Text()
		if !strings.HasSuffix(line, " map:") {
			panic("map:" + line)
		}
		mapping := [][3]int{}
		for scanner.Scan() {
			line = scanner.Text()
			if line == "" {
				break
			}
			values := strings.SplitN(line, " ", 3)
			d, _ := strconv.Atoi(values[0])
			s, _ := strconv.Atoi(values[1])
			r, _ := strconv.Atoi(values[2])
			mapping = append(mapping, [3]int{d, s, r})
		}
		tables = append(tables, mapping)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return tables
}
