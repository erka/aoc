package main

import (
	"bytes"
	_ "embed"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/erka/aoc/pkg/log"
)

//go:embed input.txt
var input []byte

/*
* part 1: 999
* part 2: 6099
 */
func main() {
	log.Infof("part1: %s", solvePart1(input))
	log.Infof("part2: %s", solvePart2(input))
}

// solve
func solvePart1(input []byte) string {
	reader := bytes.NewReader(input)
	var (
		line     string
		err      error
		counter  int
		position = 50
		ticks    int
	)
	for {
		_, err = fmt.Fscanln(reader, &line)
		if err != nil {
			if !errors.Is(err, io.EOF) {
				return err.Error()
			}
			break
		}

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
	reader := bytes.NewReader(input)
	var (
		line     string
		err      error
		counter  int
		position = 50
		ticks    int
	)
	for {
		_, err = fmt.Fscanln(reader, &line)
		if err != nil {
			if !errors.Is(err, io.EOF) {
				return err.Error()
			}
			break
		}
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
