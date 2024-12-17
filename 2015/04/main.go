package main

import (
	"bytes"
	"crypto/md5"
	_ "embed"
	"encoding/hex"
	"strconv"
	"strings"

	"github.com/erka/aoc/pkg/log"
)

//go:embed input.txt
var input []byte

/*
* part1: 254575
* part2: 1038736
 */
func main() {
	log.Infof("part1: %s", solvePart1(input))
	log.Infof("part2: %s", solvePart2(input))
}

func solvePart1(input []byte) string {
	for i := 0; ; i += 1 {
		data := string(bytes.Trim(input, "\n")) + strconv.Itoa(i)
		v := md5.Sum([]byte(data))
		d := hex.EncodeToString(v[:])
		if strings.HasPrefix(d, "00000") {
			return strconv.Itoa(i)
		}
	}
}

func solvePart2(input []byte) string {
	for i := 0; ; i += 1 {
		data := string(bytes.Trim(input, "\n")) + strconv.Itoa(i)
		v := md5.Sum([]byte(data))
		d := hex.EncodeToString(v[:])
		if strings.HasPrefix(d, "000000") {
			return strconv.Itoa(i)
		}
	}
}
