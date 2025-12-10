package iox

import (
	"bytes"
	"iter"
	"strings"
)

func Lines(input []byte) iter.Seq[string] {
	return func(yield func(string) bool) {
		for line := range bytes.Lines(input) {
			if !yield(strings.TrimSpace(string(line))) {
				return
			}
		}
	}
}
