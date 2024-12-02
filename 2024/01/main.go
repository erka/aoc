package main

import (
	"bytes"
	_ "embed"
	"errors"
	"fmt"
	"io"
	"slices"
	"strconv"

	"github.com/erka/aoc/pkg/log"
)

//go:embed input.txt
var input []byte

/*
 * part1: 1320851
 * part2: 26859182
 */
func main() {
	log.Infof("part1: %s", solvePart1(input))
	log.Infof("part2: %s", solvePart2(input))
}

// solve
func solvePart1(input []byte) string {
	lnums, rnums, err := readInput(input)
	if err != nil {
		return err.Error()
	}

	output := 0
	for i := range lnums {
		output += max(lnums[i], rnums[i]) - min(lnums[i], rnums[i])
	}

	return strconv.Itoa(output)
}

// solve
func solvePart2(input []byte) string {
	lnums, rnums, err := readInput(input)
	if err != nil {
		return err.Error()
	}

	output := 0
	for _, lv := range lnums {
		for _, rv := range rnums {
			if lv == rv {
				output += lv
			}
			if rv > lv {
				break
			}
		}
	}

	return strconv.Itoa(output)
}

// readInput reads the input data and returns two sorted slices
func readInput(input []byte) ([]int, []int, error) {
	reader := bytes.NewReader(input)
	var left, right int

	var lnums, rnums []int

	for {
		_, err := fmt.Fscan(reader, &left, &right)
		if err != nil {
			if !errors.Is(err, io.EOF) {
				return nil, nil, err
			}
			break
		}
		lnums = append(lnums, left)
		rnums = append(rnums, right)
	}

	slices.Sort(lnums)
	slices.Sort(rnums)
	return lnums, rnums, nil
}
