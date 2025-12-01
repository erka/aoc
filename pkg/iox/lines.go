package iox

import (
	"bytes"
	"fmt"
	"io"
	"iter"
)

func Lines(input []byte) iter.Seq[string] {
	return func(yield func(string) bool) {
		reader := bytes.NewReader(input)

		var line string
		for {
			_, err := fmt.Fscanln(reader, &line)
			if err != nil {
				if err == io.EOF {
					return
				}
				panic(err)
			}

			if !yield(line) {
				return
			}
		}
	}
}
