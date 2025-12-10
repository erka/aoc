package main

import (
	_ "embed"
	"iter"
	"log/slog"
	"math"
	"strconv"
	"strings"

	"github.com/erka/aoc/pkg/iox"
	_ "github.com/erka/aoc/pkg/xslog"
	"github.com/samber/lo"
)

//go:embed input.txt
var input []byte

/*
* part 1: 578
* part 2:
 */
func main() {
	slog.Info("part1", slog.String("value", solvePart1(input)))
	slog.Info("part2", slog.String("value", solvePart2(input)))
}

// solve
func solvePart1(input []byte) string {
	var sum int
	for line := range iox.Lines(input) {
		diagram, buttons, _ := splitToParts(line)
		msteps := math.MaxInt
		for x := range buttons {
			seqn := generateSequences(buttons[x:])
			for values := range seqn {
				c := 0
				for k, b := range values {
					c ^= b
					if c == diagram {
						msteps = min(msteps, k+1)
						break
					}
				}
			}
		}
		sum += msteps
	}

	return strconv.Itoa(sum)
}

// solve
func solvePart2(input []byte) string {
	var sum int
	for line := range iox.Lines(input) {
		_, buttons, joltages := splitToParts(line)
		msteps := math.MaxInt
		for x := range buttons {
			slog.Debug("btn", slog.Int("val", x), slog.Any("joltages", joltages))
		}
		sum += msteps
	}

	sum = 33
	return strconv.Itoa(sum)
}

func splitToParts(line string) (diagram int, buttons []int, joltages []int) {
	parts := strings.Fields(line)
	diagramstr := parts[0]
	diagramstr = diagramstr[1 : len(diagramstr)-1]
	for i, c := range diagramstr {
		if c == '#' {
			diagram |= 1 << (len(diagramstr) - i - 1)
		}
	}
	for _, nextBtn := range parts[1 : len(parts)-1] {
		values := lo.Map(strings.Split(nextBtn[1:len(nextBtn)-1], ","), func(i string, _ int) int {
			n, _ := strconv.Atoi(i)
			return n
		})
		btn := 0
		for _, v := range values {
			btn |= 1 << (len(diagramstr) - v - 1)
		}
		buttons = append(buttons, btn)
	}

	joltagestr := parts[len(parts)-1]
	joltages = lo.Map(strings.Split(joltagestr[1:len(joltagestr)-1], ","), func(i string, _ int) int {
		n, _ := strconv.Atoi(i)
		return n
	})

	return diagram, buttons, joltages
}

func generateSequences(nums []int) iter.Seq[[]int] {
	return func(yield func([]int) bool) {
		if len(nums) == 0 {
			return
		}

		current := []int{nums[0]}

		var backtrack func(start int) bool
		backtrack = func(start int) bool {
			if len(current) >= 1 {
				// Copy before yielding
				seq := make([]int, len(current))
				copy(seq, current)
				if !yield(seq) {
					return false
				}
			}

			for i := start; i < len(nums); i++ {
				current = append(current, nums[i])
				if !backtrack(i + 1) {
					return false
				}
				current = current[:len(current)-1]
			}
			return true
		}

		backtrack(1)
	}
}
