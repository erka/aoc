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
* part 1: 999
* part 2: 6099
 */
func main() {
	slog.Info("part1", slog.String("value", solvePart1(input)))
	slog.Info("part2", slog.String("value", solvePart2(input)))
}

// solve
func solvePart1(input []byte) string {
	var (
		counter  int
		position = 50
		ticks    int
		err      error
	)

	for line := range iox.Lines(input) {
		slog.Debug(line)
		ticks, err = strconv.Atoi(line[1:])
		if err != nil {
			return err.Error()
		}
		if line[0] == 'L' {
			ticks *= -1
		}
		position = (position + ticks%100 + 100) % 100
		if position == 0 {
			counter += 1
		}
	}

	return strconv.Itoa(counter)
}

// solve
func solvePart2(input []byte) string {
	var (
		err      error
		counter  int
		position = 50
		ticks    int
	)

	for line := range iox.Lines(input) {
		ticks, err = strconv.Atoi(line[1:])
		if err != nil {
			return err.Error()
		}

		if line[0] == 'L' {
			ticks *= -1
		}

		switch {
		case position == 0:
			counter += abs(ticks / 100)
		case position+ticks <= 0:
			counter += abs((position+ticks)/100) + 1
		case position+ticks > 99:
			counter += (position + ticks) / 100
		}

		position = (position + ticks%100 + 100) % 100
	}

	return strconv.Itoa(counter)
}

func abs[T ~int](a T) T {
	return max(a, -a)
}
