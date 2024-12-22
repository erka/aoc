package day02

import (
	"bufio"
	"bytes"
	"log"
)

const (
	lose = 'X'
	draw = 'Y'
	win  = 'Z'
)

func Solve(input []byte) error {
	log.Println("total: ", solution(input))
	return nil
}

type Move byte

const (
	rockMove Move = 1 + iota
	paperMove
	scissorsMove
)

func parsePlayer1Move(input byte) Move {
	switch input {
	case 'A':
		return rockMove
	case 'B':
		return paperMove
	case 'C':
		return scissorsMove
	}
	return 0
}

type key struct {
	outcome byte
	move    Move
}

var rules = map[key]Move{
	{lose, rockMove}:     scissorsMove,
	{lose, paperMove}:    rockMove,
	{lose, scissorsMove}: paperMove,
	{draw, rockMove}:     rockMove,
	{draw, paperMove}:    paperMove,
	{draw, scissorsMove}: scissorsMove,
	{win, rockMove}:      paperMove,
	{win, paperMove}:     scissorsMove,
	{win, scissorsMove}:  rockMove,
}
var outcomes = map[byte]int{
	lose: 0,
	draw: 3,
	win:  6,
}

func movePoints(player1Move Move, expectedOutcome byte) int {
	return outcomes[expectedOutcome] + int(rules[key{expectedOutcome, player1Move}])
}

// X means you need to lose, Y means you need to end the round in a draw, and Z means you need to win
// 1 for Rock, 2 for Paper, and 3 for Scissors
// 0 if you lost, 3 if the round was a draw, and 6 if you won
func solution(input []byte) int {
	scanner := bufio.NewScanner(bytes.NewReader(input))
	points := 0
	for scanner.Scan() {
		line := scanner.Text()
		player1Move := parsePlayer1Move(line[0])
		expectedOutcome := line[2]
		points += movePoints(player1Move, expectedOutcome)
	}
	return points
}
