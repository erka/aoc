package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"strconv"

	"github.com/erka/aoc/pkg/log"
	"gonum.org/v1/gonum/mat"
)

//go:embed input.txt
var input []byte

/*
* part 1:
* part 2:
 */
func main() {
	log.Infof("part1: %s", solvePart1(input))
	log.Infof("part2: %s", solvePart2(input))
}

// solve
func solvePart1(input []byte) string {
	lines := bytes.Split(bytes.Trim(input, "\n"), []byte("\n"))
	ingredients := make([]*ingredient, len(lines))
	for i, line := range lines {
		ingredient := &ingredient{}
		ingredient.UnmarshalText(line)
		ingredients[i] = ingredient
	}

	A := mat.NewDense(4, len(ingredients), nil)
	for i := 0; i < len(ingredients); i++ {
		x := ingredients[i]
		A.SetCol(i, []float64{float64(x.capacity), float64(x.durability), float64(x.flavor), float64(x.texture)})
	}
	actual := make([]float64, 4)
	X := mat.NewVecDense(4, actual)
	score := 0
	combinations := combinations(len(ingredients), 100)
	for _, v := range combinations {
		B := mat.NewVecDense(len(ingredients), v)
		X.MulVec(A, B)
		score = max(score, mul(X))
	}

	return strconv.Itoa(score)
}

func combinations(x, target int) [][]float64 {
	var results [][]float64
	combination := make([]float64, x)
	var backtrack func(index, remaining int)

	backtrack = func(index, remaining int) {
		if index == x-1 {
			combination[index] = float64(remaining)
			if remaining >= 0 {
				temp := make([]float64, x)
				copy(temp, combination)
				results = append(results, temp)
			}
			return
		}
		for i := 0; i <= remaining; i++ {
			combination[index] = float64(i)
			backtrack(index+1, remaining-i)
		}
	}
	backtrack(0, target)
	return results
}

func val(f float64) int {
	return max(0, int(f))
}

// solve
func solvePart2(input []byte) string {
	lines := bytes.Split(bytes.Trim(input, "\n"), []byte("\n"))
	ingredients := make([]*ingredient, len(lines))
	for i, line := range lines {
		ingredient := &ingredient{}
		ingredient.UnmarshalText(line)
		ingredients[i] = ingredient
	}

	A := mat.NewDense(4, len(ingredients), nil)
	for i := 0; i < len(ingredients); i++ {
		x := ingredients[i]
		A.SetCol(i, []float64{float64(x.capacity), float64(x.durability), float64(x.flavor), float64(x.texture)})
	}
	actual := make([]float64, 4)
	X := mat.NewVecDense(4, actual)
	score := 0
	combinations := combinations(len(ingredients), 100)
	calories := make([]float64, len(ingredients))
	for i, c := range ingredients {
		calories[i] = float64(c.calories)
	}
	C := mat.NewVecDense(len(ingredients), calories)
	for _, v := range combinations {
		B := mat.NewVecDense(len(ingredients), v)
		X.MulVec(A, B)
		B.MulElemVec(C, B)
		if mat.Sum(B) != 500 {
			continue
		}
		score = max(score, mul(X))
	}

	return strconv.Itoa(score)
}

func mul(rma *mat.VecDense) int {
	o := 1
	rm := rma.RawVector()
	for i := 0; i < rm.N; i++ {
		o *= val(rm.Data[i*rm.Inc])
	}
	return o
}

type ingredient struct {
	name                                            string
	capacity, durability, flavor, texture, calories int
}

func (i *ingredient) UnmarshalText(text []byte) error {
	_, err := fmt.Sscanf(string(text), "%s capacity %d, durability %d, flavor %d, texture %d, calories %d",
		&i.name, &i.capacity, &i.durability, &i.flavor, &i.texture, &i.calories)
	return err
}
