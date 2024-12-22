package main

import (
	"bytes"
	_ "embed"
	"strconv"

	"github.com/erka/aoc/pkg/log"
	"github.com/samber/lo"
)

//go:embed input.txt
var input []byte

/*
* part 1: 20332089158
* part 2: 2191
 */
func main() {
	log.Infof("part1: %s", solvePart1(input))
	log.Infof("part2: %s", solvePart2(input))
}

func mix(num int64, o int64) int64 {
	return num ^ o
}

func prune(num int64) int64 {
	return num % 16777216
}

func price(num int64) int {
	return int(num % 10)
}

// solve
func solvePart1(input []byte) string {
	lines := bytes.Split(bytes.Trim(input, "\n"), []byte("\n"))
	output := 0

	var secret int64
	for _, line := range lines {
		s, _ := strconv.Atoi(string(line))
		secret = int64(s)
		for i := 0; i < 2000; i++ {
			secret = next(secret)
		}
		output += int(secret)
	}

	return strconv.Itoa(output)
}

func next(secret int64) int64 {
	secret = prune(mix(secret, secret*64))
	secret = prune(mix(secret, secret/32))
	secret = prune(mix(secret, secret*2048))
	return secret
}

type exchange struct {
	price int
	delta int
}

// solve
func solvePart2(input []byte) string {
	lines := bytes.Split(bytes.Trim(input, "\n"), []byte("\n"))

	var secret int64
	exchanges := make(map[int][]exchange, len(lines))
	for b, line := range lines {
		s, _ := strconv.Atoi(string(line))
		secret = int64(s)
		oldPrice := price(secret)
		for i := 0; i < 2000; i++ {
			secret = next(secret)
			currentPrice := price(secret)
			exch := exchange{currentPrice, currentPrice - oldPrice}
			exchanges[b] = append(exchanges[b], exch)
			oldPrice = currentPrice
		}
	}

	history := map[[4]int]map[int]int{}

	for k, bs := range exchanges {
		for i := 3; i < len(bs); i++ {
			seq := [4]int{bs[i-3].delta, bs[i-2].delta, bs[i-1].delta, bs[i].delta}
			if _, ok := history[seq][k]; !ok {
				if _, ok := history[seq]; !ok {
					history[seq] = map[int]int{}
				}
				history[seq][k] += bs[i].price
			}
		}
	}

	best := 0
	var seq [4]int
	for k, n := range history {
		o := lo.Sum(lo.Values(n))
		if o > best {
			best = o
			seq = k
		}
	}
	log.Infof("best: %d %v", best, seq)
	return strconv.Itoa(best)
}
