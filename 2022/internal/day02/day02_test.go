package day02

import (
	"testing"
)

const input = `A Y
B X
C Z
`

func TestSolution(t *testing.T) {
	result := solution([]byte(input))
	if result != 12 {
		t.Fatalf("experted 12 but got: %d", result)
	}

}
