package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"strconv"

	"github.com/erka/aoc/pkg/log"
)

//go:embed input.txt
var input []byte

/*
* part1: 1588178
* part2: 3783758
 */
func main() {
	log.Infof("part1: %s", solvePart1(input))
	log.Infof("part2: %s", solvePart2(input))
}

func solvePart1(input []byte) string {
	i := 0
	var w, h, l int
	for _, c := range bytes.Split(input, []byte("\n")) {
		log.Debug(string(c))
		_, err := fmt.Sscanf(string(c), "%dx%dx%d", &w, &h, &l)
		if err != nil {
			continue
		}
		i += 2*w*h + 2*w*l + 2*h*l + min(w*h, w*l, h*l)
		log.Debug(w, h, l, i)
	}
	return strconv.Itoa(i)
}

func solvePart2(input []byte) string {
	i := 0
	var w, h, l int
	for _, c := range bytes.Split(input, []byte("\n")) {
		log.Debug(string(c))
		_, err := fmt.Sscanf(string(c), "%dx%dx%d", &w, &h, &l)
		if err != nil {
			continue
		}
		i += 2*min(l+h, h+w, l+w) + (w * l * h)
		log.Debug(w, h, l, i)
	}
	return strconv.Itoa(i)
}
