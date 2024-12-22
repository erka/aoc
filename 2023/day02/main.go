package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//go:embed input.txt
var input []byte

var cubes = set{
	"red":   12,
	"green": 13,
	"blue":  14,
}

type set map[string]int

type game struct {
	id   int
	sets []set
}

func (g *game) isPossible(cubeLimits set) bool {
	for _, s := range g.sets {
		for color, limit := range cubeLimits {
			if s[color] > limit {
				return false
			}
		}
	}
	return true
}

func (g *game) power() int {
	fset := set{}
	for _, s := range g.sets {
		for color, num := range s {
			if num > fset[color] {
				fset[color] = num
			}
		}
	}
	power := 1
	for _, v := range fset {
		power *= v
	}
	return power
}

func newGame(line string) (*game, error) {
	values := strings.SplitN(line, ": ", 2)

	gid, err := strconv.Atoi(strings.TrimPrefix(values[0], "Game "))
	if err != nil {
		return nil, err
	}
	g := &game{id: gid, sets: []set{}}
	sets := strings.Split(values[1], "; ")
	for _, s := range sets {
		gset := set{}
		cubes := strings.Split(s, ", ")
		for _, c := range cubes {
			cube := strings.SplitN(c, " ", 2)
			cubeNum, err := strconv.Atoi(cube[0])
			if err != nil {
				return nil, err
			}
			gset[cube[1]] = cubeNum
		}
		g.sets = append(g.sets, gset)
	}
	return g, nil
}

/*
* part 1: 2679
* part 2: 77607
 */
func main() {
	possible, power := solve(input, cubes)
	fmt.Printf("possible: %s, power: %s\n", possible, power)
}

// solve
func solve(input []byte, table map[string]int) (string, string) {
	scanner := bufio.NewScanner(bytes.NewBuffer(input))
	var sum, power int
	for scanner.Scan() {
		line := scanner.Text()
		g, err := newGame(line)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		if g.isPossible(table) {
			sum += g.id
		}
		power += g.power()
	}

	if err := scanner.Err(); err != nil {
		return err.Error(), err.Error()
	}

	return strconv.Itoa(sum), strconv.Itoa(power)
}
