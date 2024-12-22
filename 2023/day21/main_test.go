package main

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
	"gonum.org/v1/gonum/mat"
)

//go:embed example.txt
var example []byte

func TestSolvePart1(t *testing.T) {
	result := solve(example, 6, false)
	require.Equal(t, 16, result)
}

func TestSolvePart2(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{6, 16},
		{10, 50},
		{50, 1594},
		{100, 6536},
		{500, 167004},
		// {1000, 668697},
		// {5000, 16733044},
	}
	for _, tt := range tests {
		result := solve(example, tt.input, true)
		require.Equal(t, tt.expected, result)
	}
}

func abc(x, y []float64) (float64, float64, float64, error) {
	x = []float64{6, 10, 50, 100, 500, 1000, 5000}
	y = []float64{16, 50, 1594, 6536, 167004, 668697, 16733044}

	// Create matrices for the system of equations
	A := mat.NewDense(len(x), 3, nil)
	B := mat.NewDense(len(y), 1, y)

	// Fill matrix A with the values of x for each equation
	for i := 0; i < len(x); i++ {
		A.Set(i, 0, x[i]*x[i])
		A.Set(i, 1, x[i])
		A.Set(i, 2, 1)
	}

	// Create a matrix to hold the solution
	X := mat.NewDense(3, 1, nil)

	err := X.Solve(A, B)
	if err != nil {
		return 0, 0, 0, err
	}

	// Extract the solution
	return X.At(0, 0), X.At(1, 0), X.At(2, 0), nil
}
