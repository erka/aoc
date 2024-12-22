package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"slices"
	"strconv"
	"strings"

	"github.com/erka/aoc/pkg/log"

	"github.com/samber/lo"
)

//go:embed input.txt
var input []byte

/*
* part1: 511498
* part2: 284674q
 */
func main() {
	log.Infof("part1: %s", solvePart1(input))
	log.Infof("part2: %s", solvePart2(input))
}

func hashAlg(s string) int {
	currentValue := 0
	for _, c := range s {
		currentValue = ((currentValue + int(c)) * 17) % 256
	}
	return currentValue
}

// solve
func solvePart1(input []byte) string {
	scanner := bufio.NewScanner(bytes.NewBuffer(input))
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		vars := lo.Map[string, int](strings.Split(line, ","), func(v string, _ int) int {
			return hashAlg(v)
		})
		sum += lo.Sum[int](vars)
	}
	if err := scanner.Err(); err != nil {
		return err.Error()
	}
	return strconv.Itoa(sum)
}

type Box struct {
	id         int
	lensValues map[string]int
	lensOrder  []string
}

func (b *Box) putLens(label string, focal_length int) {
	if _, ok := b.lensValues[label]; !ok {
		b.lensOrder = append(b.lensOrder, label)
	}
	b.lensValues[label] = focal_length
}

func (b *Box) deleteLens(label string) {
	delete(b.lensValues, label)
	if i := slices.Index(b.lensOrder, label); i != -1 {
		b.lensOrder = append(b.lensOrder[:i], b.lensOrder[i+1:]...)
	}
}

func (b *Box) focusingPower() int {
	return (b.id + 1) * lo.Sum(
		lo.Map(b.lensOrder, func(item string, i int) int { return b.lensValues[item] * (i + 1) }),
	)
}
func newBox(id int) *Box {
	return &Box{
		id:         id,
		lensValues: make(map[string]int, 0),
		lensOrder:  make([]string, 0),
	}
}

// solve
func solvePart2(input []byte) string {
	scanner := bufio.NewScanner(bytes.NewBuffer(input))
	if !scanner.Scan() {
		panic("expected the input")
	}
	line := scanner.Text()
	if err := scanner.Err(); err != nil {
		return err.Error()
	}
	boxes := make(map[int]*Box, 256)
	commands := strings.Split(line, ",")
	for _, cmd := range commands {
		idx := strings.IndexAny(cmd, "-=")
		label := cmd[:idx]
		boxId := hashAlg(label)
		if _, ok := boxes[boxId]; !ok {
			boxes[boxId] = newBox(boxId)
		}
		switch cmd[idx] {
		case '=':
			focal_length := int(cmd[idx+1] - '0')
			boxes[boxId].putLens(label, focal_length)
		case '-':
			boxes[boxId].deleteLens(label)
		}
	}
	sum := 0
	for _, b := range boxes {
		sum += b.focusingPower()
	}

	return strconv.Itoa(sum)
}
