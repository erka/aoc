package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"math"
	"strconv"

	"github.com/erka/aoc/pkg/log"
)

//go:embed input.txt
var input []byte

/*
* part 1: 16076
* part 2: 2797
 */
func main() {
	log.Infof("part1: %s", solvePart1(input))
	log.Infof("part2: %s", solvePart2(input))
}

func solvePart1(input []byte) string {
	regs := solvePart(input, nil)
	return strconv.Itoa(int(regs["a"]))
}

// solve
func solvePart(input []byte, overrides map[string]uint16) map[string]uint16 {
	regs := map[string]uint16{}
	lines := bytes.Split(bytes.Trim(input, "\n"), []byte("\n"))

	var out, op, o1, o2 string

	for len(lines) > 0 {
		line := lines[0]
		lines = lines[1:]

		_, err := fmt.Sscanf(string(line), "%s -> %s", &o1, &out)
		o2 = o1
		op = "SET"
		if err != nil {
			_, err = fmt.Sscanf(string(line), "NOT %s -> %s", &o1, &out)
			o2 = o1
			op = "NOT"
			if err != nil {
				_, err = fmt.Sscanf(string(line), "%s %s %s -> %s", &o1, &op, &o2, &out)
				if err != nil {
					panic(err.Error() + string(line))
				}
			}
		}

		d1, ok1 := combo(regs, o1)
		d2, ok2 := combo(regs, o2)
		if !ok1 || !ok2 {
			// command it not ready
			lines = append(lines, line)
			continue
		}
		switch op {
		case "SET":
			if override, ok := overrides[out]; ok {
				d1 = override
			}
			regs[out] = d1
		case "NOT":
			regs[out] = math.MaxUint16 ^ d1
			continue
		case "AND":
			regs[out] = d1 & d2
		case "OR":
			regs[out] = d1 | d2
		case "LSHIFT":
			regs[out] = d1 << d2
		case "RSHIFT":
			regs[out] = d1 >> d2
		default:
			panic("unknown op:" + op)
		}
	}

	return regs
}

func combo(regs map[string]uint16, o1 string) (uint16, bool) {
	if d1, err := strconv.Atoi(o1); err == nil {
		return uint16(d1), true
	}
	v, ok := regs[o1]
	return v, ok
}

// solve
func solvePart2(input []byte) string {
	regs := solvePart(input, nil)
	regs = solvePart(input, map[string]uint16{"b": regs["a"]})
	return strconv.Itoa(int(regs["a"]))
}
