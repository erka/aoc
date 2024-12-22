package main

import (
	"fmt"
	"log"
	"os"

	"example.com/m/internal/day01"
	"example.com/m/internal/day02"
	"example.com/m/internal/day03"
	"example.com/m/internal/day04"
)

type SolveFunc func(input []byte) error

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Please provide the day to run")
	}
	day := os.Args[1]

	data, err := os.ReadFile(fmt.Sprintf("inputs/day%s.txt", day))
	if err != nil {
		log.Fatal(err)
	}

	log.SetPrefix("day" + day + " - ")
	solutions := map[string]SolveFunc{
		"01": day01.Solve,
		"02": day02.Solve,
		"03": day03.Solve,
		"04": day04.Solve,
	}

	if f, ok := solutions[day]; ok {
		f(data)
	}
}
