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
* part 1: 307
* part 2: 160
 */
func main() {
	log.Infof("part1: %s", solvePart1(input))
	log.Infof("part2: %s", solvePart2(input))
}

// solve
func solvePart1(input []byte) string {
	regs := map[byte]int{'a': 0, 'b': 0}
	return solve(input, regs)
}

// solve
func solvePart2(input []byte) string {
	regs := map[byte]int{'a': 1, 'b': 0}
	return solve(input, regs)
}

func solve(input []byte, regs map[byte]int) string {
	lines := bytes.Split(bytes.Trim(input, "\n"), []byte("\n"))
	ops := lo.Map(lines, func(l []byte, _ int) Op {
		op := Op{}
		values := bytes.Fields(l)
		op.kind = string(values[0])
		if op.kind == "jmp" {
			op.offset, _ = strconv.Atoi(string(values[1]))
			return op
		}
		op.reg = values[1][0]
		if len(values) == 3 {
			op.offset, _ = strconv.Atoi(string(values[2]))
		}
		return op
	})
	for i := 0; i < len(ops); i++ {
		op := ops[i]
		switch op.kind {
		case "hlf":
			regs[op.reg] /= 2
		case "tpl":
			regs[op.reg] *= 3
		case "inc":
			regs[op.reg] += 1
		case "jmp":
			i += op.offset - 1
		case "jie":
			if regs[op.reg]%2 == 0 {
				i += op.offset - 1
			}
		case "jio":
			if regs[op.reg] == 1 {
				i += op.offset - 1
			}
		}
	}
	return strconv.Itoa(regs['b'])
}

type Op struct {
	kind   string
	reg    byte
	offset int
}
