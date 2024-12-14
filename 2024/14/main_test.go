package main

import (
	_ "embed"
	"image"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:embed example.txt
var example []byte

func TestRobotTeleport(t *testing.T) {
	boundary := image.Rect(0, 0, 11, 7)
	robot := &Robot{
		position: image.Pt(2, 4),
		velocity: image.Pt(2, -3),
	}
	nextPositions := []image.Point{
		image.Pt(4, 1),
		image.Pt(6, 5),
		image.Pt(8, 2),
		image.Pt(10, 6),
		image.Pt(1, 3),
	}
	for _, n := range nextPositions {
		robot.move(boundary, 1)
		require.Equal(t, n, robot.position)
	}
}

func TestSolvePart1(t *testing.T) {
	boundary := image.Rect(0, 0, 11, 7)
	result := solvePart1(example, boundary)
	require.Equal(t, "12", result)
}

func TestSolvePart2(t *testing.T) {
	boundary := image.Rect(0, 0, 11, 7)
	result := solvePart2(example, boundary)
	require.Equal(t, "1", result)
}

func BenchmarkSolvePart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solvePart1(input, image.Rect(0, 0, 101, 103))
	}
}

func BenchmarkSolvePart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solvePart2(input, image.Rect(0, 0, 101, 103))
	}
}
