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

type ScratchCard struct {
	id      string
	winning []string
	own     []string
}

func NewScratchCard(line string) *ScratchCard {
	values := strings.SplitN(line, ": ", 2)
	c := &ScratchCard{
		id: values[0],
	}
	cards := strings.SplitN(values[1], " | ", 2)
	c.winning = strings.Fields(cards[0])
	c.own = strings.Fields(cards[1])
	return c
}

func (c *ScratchCard) Points() int {
	value := c.Matches()
	if value > 1 {
		value = 1 << (value - 1)
	}
	return value
}

func (c *ScratchCard) Matches() int {
	value := 0
	for _, w := range c.winning {
		if slices.Contains(c.own, w) {
			value++
		}
	}
	return value
}

/*
* part 1: 25174
* part 2: 6420979
 */
func main() {
	fmt.Printf("part1: %s\n", solve(input))
	fmt.Printf("part2: %s\n", solveCopies(input))
}

// solve
func solve(input []byte) string {
	scanner := bufio.NewScanner(bytes.NewBuffer(input))
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		c := NewScratchCard(line)
		sum += c.Points()
	}
	if err := scanner.Err(); err != nil {
		return err.Error()
	}
	return strconv.Itoa(sum)
}

// solve
func solveCopies(input []byte) string {
	scanner := bufio.NewScanner(bytes.NewBuffer(input))
	i := 1
	copies := make(map[int]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		c := NewScratchCard(line)
		copies[i] += 1
		for j := 1; j <= c.Matches(); j++ {
			copies[i+j] += copies[i]
		}
		i++
	}
	if err := scanner.Err(); err != nil {
		return err.Error()
	}
	sum := 0
	for _, v := range copies {
		sum += v
	}
	return strconv.Itoa(sum)
}
