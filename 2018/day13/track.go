package day13

import (
	"io"
	"strings"

	"github.com/seanhagen/advent-of-code/lib"
)

// Mine ...
type Mine struct {
	carts []*Cart

	tracks [][]string
}

// SetupMine ...
func SetupMine(path string) *Mine {
	f := lib.LoadInput(path)

	m := &Mine{
		carts:  []*Cart{},
		tracks: [][]string{},
	}

	y := 0
	err := lib.LoopOverLines(f, func(line []byte) error {
		l := string(line)
		bits := strings.Split(l, "")

		for x, v := range bits {
			c := CreateCart(v, x, y)
			if c != nil {
				m.carts = append(m.carts, c)
			}
		}

		y++
		return nil
	})

	if err != io.EOF {
		panic(err)
	}

	return m
}
