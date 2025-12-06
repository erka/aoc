package main

import (
	"bytes"
	_ "embed"
	"log/slog"
	"strconv"

	_ "github.com/erka/aoc/pkg/xslog"
)

//go:embed input.txt
var input []byte

/*
* part 1: 6172481852142
* part 2: 10188206723429
 */
func main() {
	slog.Info("part1", slog.String("value", solvePart1(input)))
	slog.Info("part2", slog.String("value", solvePart2(input)))
}

// solve
func solvePart1(input []byte) string {
	return solve(input, func(o op, s int, e int, lines [][]byte) {
		var n int64
		for _, l := range lines[:len(lines)-2] {
			n = 0
			for i := s; i < e; i++ {
				if l[i] == ' ' {
					continue
				}
				n = n*10 + int64(l[i]-'0')
			}
			o.Apply(n)
		}
	})
}

// solve
func solvePart2(input []byte) string {
	return solve(input, func(o op, s int, e int, lines [][]byte) {
		var n int64
		for i := s; i < e; i++ {
			n = 0
			for _, l := range lines[:len(lines)-2] {
				if l[i] == ' ' {
					continue
				}
				n = n*10 + int64(l[i]-'0')
			}
			o.Apply(n)
		}
	})
}

type calculateFn func(o op, s int, e int, lines [][]byte)

func solve(input []byte, fn calculateFn) string {
	lines := bytes.Split(input, []byte{'\n'})
	ops := lines[len(lines)-2]
	var s int
	var grandTotal int64
	for e := 1; e < len(ops); e += 1 {
		if ops[e] == ' ' && e+1 < len(ops) {
			continue
		}
		// last data set
		if e+1 == len(ops) {
			e += 2
		}

		o := newOp(ops[s])
		fn(o, s, e-1, lines)
		s = e
		grandTotal += o.Value()
	}
	return strconv.FormatInt(grandTotal, 10)
}

func newOp(ops byte) op {
	if ops == '*' {
		return &opMul{val: 1}
	}
	return &opSum{}
}

type (
	op interface {
		Apply(int64)
		Value() int64
	}
	opSum struct {
		val int64
	}
)

func (o *opSum) Apply(n int64) {
	o.val += n
}

func (o *opSum) Value() int64 {
	return o.val
}

type opMul struct {
	val int64
}

func (o *opMul) Apply(n int64) {
	o.val *= n
}

func (o *opMul) Value() int64 {
	return o.val
}
