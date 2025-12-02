package main

import (
	_ "embed"
	"log/slog"
	"math"
	"strconv"
	"strings"

	_ "github.com/erka/aoc/pkg/xslog"
)

//go:embed input.txt
var input []byte

/*
* part 1: 19128774598
* part 2: 21932258645
 */
func main() {
	slog.Info("part1", slog.String("value", solvePart1(input)))
	slog.Info("part2", slog.String("value", solvePart2(input)))
}

// solve
func solvePart1(input []byte) string {
	invalid := make(map[int64]struct{})
	for rng := range strings.SplitSeq(strings.TrimSpace(string(input)), ",") {
		vals := strings.SplitN(rng, "-", 2)
		start, err := strconv.ParseInt(vals[0], 10, 64)
		if err != nil {
			return err.Error()
		}
		end, err := strconv.ParseInt(vals[1], 10, 64)
		if err != nil {
			return err.Error()
		}

		m := (len(vals[0]) + 1) / 2
		d := pow10[m]

		b := start / d
		for {
			e := b*d + b
			b += 1
			if e > end {
				break
			}
			if digits(e)%2 != 0 || e < start {
				continue
			}

			invalid[e] = struct{}{}
		}

	}

	var sum int64
	for n := range invalid {
		sum += n
	}
	return strconv.FormatInt(sum, 10)
}

// solve
func solvePart2(input []byte) string {
	invalid := make(map[int64]struct{})
	for rng := range strings.SplitSeq(strings.TrimSpace(string(input)), ",") {
		vals := strings.SplitN(rng, "-", 2)
		start, err := strconv.ParseInt(vals[0], 10, 64)
		if err != nil {
			return err.Error()
		}
		end, err := strconv.ParseInt(vals[1], 10, 64)
		if err != nil {
			return err.Error()
		}
		minL := digits(start)
		maxL := digits(end)
		var e int64
		for L := minL; L <= maxL; L++ {
			for base := 1; base*2 <= L; base++ {
				if L%base != 0 {
					continue
				}

				k := L / base
				d := pow10[base]

				bStart := pow10[base-1]
				bEnd := pow10[base] - 1

				for b := bStart; b <= bEnd; b++ {
					e = 0
					for range k {
						e = e*d + b
					}

					if e > end {
						break
					}
					if e >= start {
						invalid[e] = struct{}{}
					}
				}
			}
		}
	}
	var sum int64
	for n := range invalid {
		sum += n
	}
	return strconv.FormatInt(sum, 10)
}

var pow10 = [...]int64{
	1,
	10,
	100,
	1000,
	10000,
	100000,
	1000000,
	10000000,
	100000000,
	1000000000,
	10000000000,
}

func digits(n int64) int {
	if n == 0 {
		return 1
	}
	return int(math.Log10(float64(n))) + 1
}
