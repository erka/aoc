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
* part 1:
* part 2:
 */
func main() {
	slog.Info("part1", slog.String("value", solvePart1(input)))
	slog.Info("part2", slog.String("value", solvePart2(input)))
}

// solve
func solvePart1(input []byte) string {
	for line := range iox.Lines(input) {
		slog.Debug(line)
	}

	return strconv.Itoa(0)
}

// solve
func solvePart2(input []byte) string {
	for line := range iox.Lines(input) {
		slog.Debug(line)
	}
	return strconv.Itoa(0)
}
