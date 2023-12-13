package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"strconv"
)

//go:embed input.txt
var input []byte

/*
* part 1: 28895
* part 2: 31603
 */
func main() {
	fmt.Printf("part1: %s\n", solvePart1(input))
	fmt.Printf("part2: %s\n", solvePart2(input))
}

func findHorizontalReflection(mirror []string, skip int) int {
	for i := 1; i < len(mirror); i++ {
		if mirror[i-1] == mirror[i] && i != skip {
			found := true
			j := i - 2
			k := i + 1
			for j >= 0 && k < len(mirror) {
				if mirror[j] != mirror[k] {
					found = false
					break
				}
				j--
				k++
			}

			if found {
				return i
			}
		}
	}
	// fmt.Println("no horizontal match")
	return -1
}

func findReflection(mirror []string) int {
	// fmt.Println(strings.Join(mirror, "\n"))
	// fmt.Println("")
	num := findHorizontalReflection(mirror, -1)
	if num > -1 {
		return num * 100
	}
	return findHorizontalReflection(swap(mirror), -1)
}

func diff(a, b string) int {
	diff := 0
	for k := 0; k < len(a); k += 1 {
		if a[k] != b[k] {
			diff += 1
		}
	}
	return diff
}

func newMirror(mirror []string, i, j int) []string {
	newmirror := append([]string{}, mirror[:i]...)
	newmirror = append(newmirror, mirror[j])
	newmirror = append(newmirror, mirror[i+1:]...)
	return newmirror
}

func findSmudge(mirror []string) int {
	// horizontal
	for i := 0; i < len(mirror)-1; i++ {
		for j := i + 1; j < len(mirror); j++ {
			diff := diff(mirror[i], mirror[j])
			if diff == 1 {
				old := findHorizontalReflection(mirror, -1)
				newmirror := newMirror(mirror, i, j)
				value := findHorizontalReflection(newmirror, old)
				if value != -1 {
					return value * 100
				}
			}
		}
	}
	mirror = swap(mirror)
	//vertical
	for i := 0; i < len(mirror)-1; i++ {
		for j := i + 1; j < len(mirror); j++ {
			diff := diff(mirror[i], mirror[j])
			if diff == 1 {
				newmirror := newMirror(mirror, i, j)
				old := findHorizontalReflection(mirror, -1)
				value := findHorizontalReflection(newmirror, old)
				if value != -1 {
					return value
				}
			}
		}
	}
	panic("oops")
}

func swap(mirror []string) []string {
	n := []string{}
	for i := 0; i < len(mirror[0]); i++ {
		l := ""
		for j := len(mirror) - 1; j >= 0; j-- {
			l += string(mirror[j][i])
		}
		n = append(n, l)
	}
	return n
}

// solve
func solvePart1(input []byte) string {
	scanner := bufio.NewScanner(bytes.NewBuffer(input))
	mirror := []string{}
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			sum += findReflection(mirror)
			mirror = []string{}
			continue
		}
		mirror = append(mirror, line)
	}
	sum += findReflection(mirror)
	if err := scanner.Err(); err != nil {
		return err.Error()
	}
	return strconv.Itoa(sum)
}

// solve
func solvePart2(input []byte) string {
	scanner := bufio.NewScanner(bytes.NewBuffer(input))
	mirror := []string{}
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			sum += findSmudge(mirror)
			mirror = []string{}
			continue
		}
		mirror = append(mirror, line)
	}
	sum += findSmudge(mirror)
	if err := scanner.Err(); err != nil {
		return err.Error()
	}
	return strconv.Itoa(sum)
}
