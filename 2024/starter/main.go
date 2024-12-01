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
* part 1:
* part 2:
 */
func main() {
	log.Infof("part1: %s", solvePart1(input))
	log.Infof("part2: %s", solvePart2(input))
}

// solve
func solvePart1(input []byte) string {
	reader := bytes.NewReader(input)
	var (
		line string
		err  error
	)
	for {
		_, err = fmt.Fscanln(reader, &line)
		if err != nil {
			if errors.Is(err, io.EOF) {
				return err.Error()
			}
			break
		}
		log.Debug(line)
	}

	return strconv.Itoa(0)
}

// solve
func solvePart2(input []byte) string {
	reader := bytes.NewReader(input)
	var (
		line string
		err  error
	)
	for {
		_, err = fmt.Fscanln(reader, &line)
		if err != nil {
			if errors.Is(err, io.EOF) {
				return err.Error()
			}
			break
		}
		log.Debug(line)
	}

	return strconv.Itoa(0)
}
