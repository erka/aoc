package main

import (
	_ "embed"
	"log/slog"
	"strconv"

	"github.com/erka/aoc/pkg/iox"
	_ "github.com/erka/aoc/pkg/xslog"
)

//go:embed input.txt
var input []byte

/*
* part 1: 17405
* part 2: 171990312704598
 */
func main() {
	slog.Info("part1", slog.String("value", solvePart1(input)))
	slog.Info("part2", slog.String("value", solvePart2(input)))
}

// solve
func solvePart1(input []byte) string {
	var sum, l int
	var b, s byte
	for line := range iox.Lines(input) {
		b, s = line[0], line[1]
		l = len(line)
		for i := 1; i < l-1; i += 1 {
			switch {
			case line[i] > b:
				b, s = line[i], line[i+1]
			default:
				s = max(s, line[i])
			}
		}
		s = max(s, line[l-1])
		sum += int((b-'0')*10 + (s - '0'))
	}

	return strconv.Itoa(sum)
}

// solve
func solvePart2(input []byte) string {
	var sum, l, start int
	k := 12
	out := make([]byte, 0, k)
	for line := range iox.Lines(input) {
		out = out[:0]
		start = 0
		l = len(line)
		for picks := range k {
			maxB := byte('0' - 1)
			maxPos := start
			for i := start; i <= l-k+picks; i += 1 {
				b := line[i]
				if b > maxB {
					maxB = b
					maxPos = i
					if maxB == '9' {
						break
					}
				}
			}
			out = append(out, maxB)
			start = maxPos + 1
		}
		// sum
		val := 0
		for _, d := range out {
			val = val*10 + int(d-'0')
		}
		sum += val
	}

	return strconv.Itoa(sum)
}
