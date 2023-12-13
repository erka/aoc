package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input []byte

/*
 * part1: 7350
 * part2: 200097286528151
 */
func main() {
	fmt.Printf("part1: %s\n", solvePart1(input))
	fmt.Printf("part2: %s\n", solvePart2(input))
}

func rowFromLine(line string, copies int) *Row {
	values := strings.SplitN(line, " ", 2)
	springs := values[0]

	digits := strings.Split(values[1], ",")
	groups := make([]int, 0, len(digits))
	for _, digit := range digits {
		d, err := strconv.Atoi(digit)
		if err != nil {
			panic(err)
		}
		groups = append(groups, d)
	}

	os := springs
	og := groups
	for i := 1; i < copies; i++ {
		springs = springs + "?" + os
		groups = append(groups, og...)
	}

	// normalize
	springs = strings.Trim(springs, ".")
	for strings.Contains(springs, "..") {
		springs = strings.ReplaceAll(springs, "..", ".")
	}

	return &Row{
		springs: []rune(springs),
		groups:  groups,
	}
}

type Row struct {
	springs []rune
	groups  []int
}

func count(x map[string]int, springs []rune, groups []int) int {
	key := fmt.Sprintf("%s-%+v", string(springs), groups)
	if v, ok := x[key]; ok {
		return v
	}
	if len(groups) == 0 {
		if slices.Contains(springs, '#') {
			return 0
		}
		return 1
	}

	if len(springs) == 0 {
		return 0
	}

	if springs[0] == '.' {
		for len(springs) > 0 && springs[0] == '.' {
			springs = springs[1:]
		}
		x[key] = count(x, springs, groups)
		return x[key]
	}

	if springs[0] == '?' {
		x[key] = count(x, springs[1:], groups) +
			count(x, append([]rune{'#'}, springs[1:]...), groups)
		return x[key]
	}

	group, groups := groups[0], groups[1:]
	if len(springs) < group {
		return 0
	}

	if slices.Contains(springs[0:group], '.') {
		return 0
	}

	if len(groups) > 0 {
		if len(springs) < group+1 || springs[group] == '#' {
			return 0
		}
		x[key] = count(x, springs[group+1:], groups)
		return x[key]
	}
	x[key] = count(x, springs[group:], groups)
	return x[key]
}

func (r *Row) Arrangements() int {
	springs := r.springs
	groups := r.groups
	// reverse if it's easier
	if springs[len(springs)-1] == '#' {
		slices.Reverse(springs)
		slices.Reverse(groups)
	}

	for len(springs) > 0 && springs[0] == '#' {
		x := min(groups[0]+1, len(springs))
		springs = springs[x:]
		groups = groups[1:]
	}
	return count(map[string]int{}, springs, groups)
}

// solve
func solvePart1(input []byte) string {
	scanner := bufio.NewScanner(bytes.NewBuffer(input))
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		sum += rowFromLine(line, 1).Arrangements()
	}
	if err := scanner.Err(); err != nil {
		return err.Error()
	}
	return strconv.Itoa(sum)
}

// solve
func solvePart2(input []byte) string {
	scanner := bufio.NewScanner(bytes.NewBuffer(input))
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		sum += rowFromLine(line, 5).Arrangements()
	}
	if err := scanner.Err(); err != nil {
		return err.Error()
	}
	return strconv.Itoa(sum)
}
