package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"

	"github.com/erka/aoc/pkg/log"
	"github.com/samber/lo"
)

//go:embed input.txt
var input []byte

/*
* part 1: 1,5,0,1,7,4,1,0,3
* part 2: 47910079998866
 */
func main() {
	log.Infof("part1: %s", solvePart1(input))
	log.Infof("part2: %s", solvePart2(input))
}

type opcode int

const (
	adv opcode = iota
	bxl
	bst
	jnz
	bxc
	out
	bdv
	cdv
)

const (
	A = 0
	B = 1
	C = 2
)

var template = `Register A: %d 
Register B: %d 
Register C: %d`

func combo(operand int, regs [3]int) int {
	switch operand {
	case 0, 1, 2, 3:
		return operand
	case 4:
		return regs[A]
	case 5:
		return regs[B]
	case 6:
		return regs[C]
	default:
		panic(fmt.Sprintf("invalid operand: %d", operand))
	}
}

// solve
func solvePart1(input []byte) string {
	regs := [3]int{}

	di := strings.Split(string(input), "\n\n")
	fmt.Sscanf(di[0], template, &regs[0], &regs[1], &regs[2])
	prog := lo.Map(
		strings.FieldsFunc(strings.TrimLeft(di[1], "Program: "), func(r rune) bool { return r == ',' }),
		func(s string, i int) int { v, _ := strconv.Atoi(s); return v },
	)

	log.Debug(regs, prog)
	output, _ := solve(prog, regs)
	return output
}

func solve(prog []int, regs [3]int) (string, [3]int) {
	output := []int{}

	for i := 0; i < len(prog)-1; i += 2 {
		op := prog[i]
		operand := prog[i+1]
		switch opcode(op) {
		case adv:
			regs[A] = regs[A] >> combo(operand, regs)
		case bdv:
			regs[B] = regs[A] >> combo(operand, regs)
		case cdv:
			regs[C] = regs[A] >> combo(operand, regs)
		case bxl:
			regs[B] ^= operand
		case bst:
			regs[B] = combo(operand, regs) % 8
		case bxc:
			regs[B] = regs[B] ^ regs[C]
		case jnz:
			if regs[A] != 0 {
				i = combo(operand, regs) - 2
			}
		case out:
			operand = combo(operand, regs)
			output = append(output, operand%8)
		}
	}
	return strings.Join(lo.Map(output, func(n int, _ int) string { return strconv.Itoa(n) }), ","), regs
}

// solve - what a hell
func solvePart2(input []byte) string {
	di := strings.Split(string(input), "\n\n")
	p := strings.TrimLeft(di[1], "Program: ")
	prog := lo.Map(
		strings.FieldsFunc(p, func(r rune) bool { return r == ',' }),
		func(s string, i int) int { v, _ := strconv.Atoi(s); return v },
	)
	a := 0
	for n := len(prog) - 1; n >= 0; n -= 1 {
		pref := strings.Join(lo.Map(prog[n:], func(item int, index int) string {
			return strconv.Itoa(item)
		}), ",")
		a <<= 3
		log.Info(pref)
		for {
			o, _ := solve(prog, [3]int{a, 0, 0})
			if strings.HasPrefix(o, pref) {
				break
			}
			a++
		}
	}

	return strconv.Itoa(a)
}
