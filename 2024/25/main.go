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
* part 1: 3338
 */
func main() {
	log.Infof("part1: %s", solvePart1(input))
}

// solve
func solvePart1(input []byte) string {
	schematics := lo.Map(bytes.Split(bytes.Trim(input, "\n"), []byte("\n\n")), func(item []byte, _ int) schematic {
		s := schematic{}
		s.kind = key
		if bytes.Equal(item[:6], []byte("#####\n")) {
			s.kind = lock
		}

		for i := 6; i < len(item)-6; i++ {
			if item[i] == '#' {
				s.pins[i%6] += 1
			}
		}
		return s
	})

	locks := lo.Filter(schematics, func(s schematic, _ int) bool {
		return s.kind == lock
	})

	keys := lo.Filter(schematics, func(s schematic, _ int) bool {
		return s.kind == key
	})
	output := 0

	for _, l := range locks {
		for _, k := range keys {
			match := true
			for i := 0; i < pins; i++ {
				if l.pins[i]+k.pins[i] > pins {
					match = false
					break
				}
			}
			if match {
				output += 1
			}
		}
	}

	return strconv.Itoa(output)
}

const (
	lock = "lock"
	key  = "key"
	pins = 5
)

type schematic struct {
	kind string
	pins [pins]int
}
