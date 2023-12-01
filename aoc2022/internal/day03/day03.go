package day03

import (
	"bufio"
	"bytes"
	"log"
	"strings"
)

func Solve(input []byte) error {
	log.Println("total: ", solutionPart2(input))
	return nil
}

func solutionPart1(input []byte) int {
	scanner := bufio.NewScanner(bytes.NewReader(input))
	points := 0
	var l int
	var firstCompartment, secondCompartment string
	for scanner.Scan() {
		line := scanner.Text()
		l = len(line)
		firstCompartment = line[:l/2]
		secondCompartment = line[l/2:]
		shareItems := make(map[rune]byte, 0)
		for _, r := range firstCompartment {
			if strings.ContainsRune(secondCompartment, r) {
				shareItems[r]++
			}
		}
		for k := range shareItems {
			points += itemPriority(k)
		}

	}
	return points
}

func solutionPart2(input []byte) int {
	scanner := bufio.NewScanner(bytes.NewReader(input))
	points := 0
	var rucksack1, rucksack2, rucksack3 string
	for scanner.Scan() {
		rucksack1 = scanner.Text()
		if !scanner.Scan() {
			log.Fatal("expected the second line")
		}
		rucksack2 = scanner.Text()
		if !scanner.Scan() {
			log.Fatal("expected the third line")
		}
		rucksack3 = scanner.Text()
		shareItems := make(map[rune]byte, 0)
		for _, r := range rucksack1 {
			if strings.ContainsRune(rucksack2, r) && strings.ContainsRune(rucksack3, r) {
				shareItems[r]++
			}
		}
		for k := range shareItems {
			points += itemPriority(k)
		}
	}
	return points
}

func itemPriority(k rune) int {
	if k >= 'A' && k <= 'Z' {
		return int(k - 'A' + 27)
	}
	return int(k - 'a' + 1)
}
