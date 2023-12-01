package day01

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"sort"
	"strconv"
)

const emptyString = ""

func Solve(input []byte) error {
	log.Println("total for the first: ", solution(bytes.NewReader(input), 1))
	log.Println("total for the three first: ", solution(bytes.NewReader(input), 3))
	return nil
}
func solution(input io.Reader, top int) int {
	scanner := bufio.NewScanner(input)
	elfsCalories := make([]int, 0)
	currentElfCalories := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line != emptyString {
			i, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal(err)
			}
			currentElfCalories += i
			continue
		}
		elfsCalories = append(elfsCalories, currentElfCalories)
		currentElfCalories = 0
	}
	elfsCalories = append(elfsCalories, currentElfCalories)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	sort.Ints(elfsCalories)
	l := len(elfsCalories)
	total := 0
	for i := 1; i <= top && l > 0; i++ {
		total += elfsCalories[l-i]
	}
	return total
}
