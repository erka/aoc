package main

import (
	"bufio"
	"bytes"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var (
	digits         = map[string]int{}
	wordsAndDigits = map[string]int{
		"three": 3,
		"seven": 7,
		"eight": 8,
		"four":  4,
		"five":  5,
		"nine":  9,
		"one":   1,
		"two":   2,
		"six":   6,
	}
)

/*
* part 1:  54630
* part 2:  54770
 */
func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	println("part 1: ", solve(input, digits))
	println("part 2: ", solve(input, wordsAndDigits))
}

// solve
func solve(input []byte, table map[string]int) string {
	scanner := bufio.NewScanner(bytes.NewBuffer(input))
	var sum int
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		first, last := -1, 0
		for i, v := range line {
			if unicode.IsDigit(v) {
				last = int(v - '0')
				if first == -1 {
					first = last
				}
			}
			for k, w := range table {
				if strings.HasPrefix(line[i:], k) {
					last = w
					if first == -1 {
						first = last
					}
				}
			}
		}
		sum += int(first*10 + last)
	}

	if err := scanner.Err(); err != nil {
		return err.Error()
	}

	return strconv.Itoa(sum)
}
