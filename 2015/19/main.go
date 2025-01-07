package main

import (
	"bytes"
	_ "embed"
	"math/rand"
	"slices"
	"strconv"
	"strings"

	"github.com/erka/aoc/pkg/log"
	"github.com/samber/lo"
)

//go:embed input.txt
var input []byte

/*
* part 1: 509
* part 2: 195
 */
func main() {
	log.Infof("part1: %s", solvePart1(input))
	log.Infof("part2: %s", solvePart2(input))
}

// solve
func solvePart1(input []byte) string {
	lines := bytes.Split(bytes.Trim(input, "\n"), []byte("\n\n"))
	replacements := lo.Map(bytes.Split(lines[0], []byte("\n")), func(item []byte, _ int) [2]string {
		parts := bytes.Split(item, []byte(" => "))
		return [2]string{string(parts[0]), string(parts[1])}
	})
	molecules := map[string]struct{}{}
	inputMolecule := lines[1]
	for _, replacement := range replacements {
		for i := 0; i < len(inputMolecule); i++ {
			if string(inputMolecule[i:i+len(replacement[0])]) == replacement[0] {
				molecule := append(append([]byte{}, inputMolecule[:i]...), []byte(replacement[1])...)
				molecule = append(molecule, inputMolecule[i+len(replacement[0]):]...)
				molecules[string(molecule)] = struct{}{}
			}
		}
	}
	return strconv.Itoa(len(molecules))
}

// solve
func solvePart2(input []byte) string {
	lines := bytes.Split(bytes.Trim(input, "\n"), []byte("\n\n"))
	replacements := lo.Map(bytes.Split(lines[0], []byte("\n")), func(item []byte, _ int) [2]string {
		parts := bytes.Split(item, []byte(" => "))
		return [2]string{string(parts[0]), string(parts[1])}
	})
	targetMolecule := string(lines[1])
	steps := 0
	tried := slices.Clone(replacements)
	for targetMolecule != "e" {
		i := rand.Intn(len(tried))
		replacement := replacements[i]
		if strings.Contains(targetMolecule, replacement[1]) {
			targetMolecule = strings.Replace(targetMolecule, replacement[1], replacement[0], 1)
			steps += 1
			tried = slices.Clone(replacements)
		}
	}
	return strconv.Itoa(steps)
}
