package day04

import (
	"bufio"
	"bytes"
	"log"
	"regexp"
	"strconv"
)

type comparator func(start, end, x, y int) bool

func Solve(input []byte) error {
	log.Println("total fully: ", solution(input, fully))
	log.Println("total overlap: ", solution(input, overlap))
	return nil
}

func solution(input []byte, f comparator) int {
	scanner := bufio.NewScanner(bytes.NewReader(input))
	points := 0
	r := regexp.MustCompile(`^(?P<elf1start>\d+)-(?P<elf1end>\d+),(?P<elf2start>\d+)-(?P<elf2end>\d+)$`)
	var err error
	for scanner.Scan() {
		line := scanner.Text()
		m := r.FindStringSubmatch(line)
		result := make(map[string]int)

		for i, name := range r.SubexpNames() {
			if i != 0 && name != "" {
				result[name], err = strconv.Atoi(m[i])
				if err != nil {
					log.Fatal("failed to parse area point")
				}
			}
		}
		if f(result["elf1start"], result["elf1end"], result["elf2start"], result["elf2end"]) || f(result["elf2start"], result["elf2end"], result["elf1start"], result["elf1end"]) {
			points++
		}

	}
	return points
}

func fully(start, end, k, l int) bool {
	return start <= k && end >= l
}

func overlap(start, end, k, l int) bool {
	return (start <= l && end >= l) || (start <= k && end >= k)
}
