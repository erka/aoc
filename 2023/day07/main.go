package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input []byte

/*
* part 1: 248569531
* part 2:
 */
func main() {
	fmt.Printf("part1: %s\n", solve(input, part1, handType))
	fmt.Printf("part2: %s\n", solve(input, part2, handTypeJoker))
}

type HandType int
type Hits map[string]int

const (
	FiveOfKind HandType = iota
	FourOfKind
	FullHouse
	ThreeOfKind
	TwoPair
	OnePair
	HighCard
)

type Camel struct {
	hand     string
	bid      int
	handType HandType
}

func hits(hand string) Hits {
	m := make(Hits, 0)
	for i := range hand {
		m[hand[i:i+1]] += 1
	}
	return m
}

func handType(hits Hits) HandType {
	if len(hits) == 5 {
		return HighCard
	}
	if len(hits) == 1 {
		return FiveOfKind
	}
	if len(hits) == 2 {
		for _, h := range hits {
			if h == 4 {
				return FourOfKind
			}
		}
		return FullHouse
	}

	if len(hits) == 3 {
		for _, h := range hits {
			if h == 3 {
				return ThreeOfKind
			}
		}
		return TwoPair
	}
	return OnePair
}

func handTypeJoker(hits Hits) HandType {
	jokers := 0
	for h, v := range hits {
		if h == "J" {
			jokers = v
			break
		}
	}

	if jokers == 0 || jokers == 5 {
		return handType(hits)
	}

	if len(hits) == 2 {
		return FiveOfKind
	}

	if len(hits) == 3 {
		for _, h := range hits {
			if h+jokers == 4 {
				return FourOfKind
			}
		}
		return FullHouse
	}

	if len(hits) == 4 {
		for k, h := range hits {
			if k == "J" {
				continue
			}
			if h+jokers == 3 {
				return ThreeOfKind
			}
		}
		return TwoPair
	}

	return OnePair
}

func (c *Camel) stronger(o *Camel, index string) bool {
	if c.handType != o.handType {
		return c.handType > o.handType
	}
	for i := 0; i < len(c.hand); i++ {
		a := strings.Index(index, c.hand[i:i+1])
		b := strings.Index(index, o.hand[i:i+1])
		if a != b {
			return a > b
		}
	}
	return false
}

func NewCamel(s string, f func(Hits) HandType) (*Camel, error) {
	values := strings.SplitN(s, " ", 2)
	handType := f(hits(values[0]))
	bid, err := strconv.Atoi(values[1])
	if err != nil {
		return nil, err
	}
	return &Camel{values[0], bid, handType}, nil
}

func part1(camels []*Camel) {
	sort.Slice(camels, func(i int, j int) bool {
		return camels[i].stronger(camels[j], "AKQJT98765432")
	})
}

func part2(camels []*Camel) {
	sort.Slice(camels, func(i int, j int) bool {
		return camels[i].stronger(camels[j], "AKQT98765432J")
	})
}

// solve
func solve(input []byte, f func(camels []*Camel), handType func(hits Hits) HandType) string {
	scanner := bufio.NewScanner(bytes.NewBuffer(input))
	camels := make([]*Camel, 0)
	for scanner.Scan() {
		line := scanner.Text()
		c, err := NewCamel(line, handType)
		if err != nil {
			panic(err)
		}
		camels = append(camels, c)
	}
	if err := scanner.Err(); err != nil {
		return err.Error()
	}
	f(camels)
	sum := 0
	for i, c := range camels {
		sum += c.bid * (i + 1)
	}
	return strconv.Itoa(sum)
}
