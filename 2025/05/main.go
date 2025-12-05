package main

import (
	"bytes"
	"cmp"
	_ "embed"
	"log/slog"
	"slices"
	"strconv"
	"strings"

	"github.com/erka/aoc/pkg/iox"
	_ "github.com/erka/aoc/pkg/xslog"
	"github.com/samber/lo"
)

//go:embed input.txt
var input []byte

/*
* part 1: 558
* part 2: 344813017450467
 */
func main() {
	slog.Info("part1", slog.String("value", solvePart1(input)))
	slog.Info("part2", slog.String("value", solvePart2(input)))
}

// solve
func solvePart1(input []byte) string {
	lines := bytes.SplitN(input, []byte{'\n', '\n'}, 2)

	// ranges
	ranges, err := buildRanges(lines)
	if err != nil {
		return err.Error()
	}

	// flesh ingredients
	counter := 0
	for line := range iox.Lines(lines[1]) {
		v, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			return err.Error()
		}
		for _, r := range ranges {
			if r.contains(v) {
				counter += 1
				break
			}
		}
	}

	return strconv.Itoa(counter)
}

func buildRanges(lines [][]byte) ([]*Range, error) {
	ranges := []*Range{}
	for line := range iox.Lines(lines[0]) {
		in := strings.SplitN(line, "-", 2)
		mn, err := strconv.ParseInt(in[0], 10, 64)
		if err != nil {
			return nil, err
		}
		mx, err := strconv.ParseInt(in[1], 10, 64)
		if err != nil {
			return nil, err
		}
		ranges = append(ranges, &Range{min: mn, max: mx})
	}

	slices.SortFunc(ranges, func(a *Range, b *Range) int {
		if a.min == b.min {
			return cmp.Compare(a.max, b.max)
		}
		return cmp.Compare(a.min, b.min)
	})

	initial := make([]*Range, len(ranges))
	for {
		initial = initial[:len(ranges)]
		copy(initial, ranges)
		ranges = ranges[:0]
		r := initial[0]
		ranges = append(ranges, r)
		for _, n := range initial[1:] {
			// check of overlapping
			if r.contains(n.min) || r.contains(n.max) {
				r.min = min(r.min, n.min)
				r.max = max(r.max, n.max)
				continue
			}
			ranges = append(ranges, n)
			r = n
		}
		if len(ranges) == len(initial) {
			break
		}
	}

	return ranges, nil
}

// solve
func solvePart2(input []byte) string {
	lines := bytes.SplitN(input, []byte{'\n', '\n'}, 2)

	// ranges
	ranges, err := buildRanges(lines)
	if err != nil {
		return err.Error()
	}
	v := lo.SumBy(ranges, func(r *Range) int64 {
		return r.max - r.min + 1
	})
	return strconv.FormatInt(v, 10)
}

type Range struct {
	min int64
	max int64
}

func (r Range) contains(v int64) bool {
	return r.min <= v && r.max >= v
}
