package day3

import (
	"fmt"
	"strings"
)

type Grid struct {
	rows []Row
}

func NewGrid(s string) (*Grid, error) {
	if s == "" {
		return nil, fmt.Errorf("got empty string for Grid")
	}
	r := strings.Split(s, "\n")
	if len(r) == 0 {
		return nil, fmt.Errorf("no rows in input")
	}
	l := len(r[0])

	rows := []Row{}

	for i, v := range r {
		rr := NewRow(v)
		rows = append(rows, rr)
		if i == 0 {
			continue
		}
		if len(v) != l {
			return nil, fmt.Errorf("line %v of input not same length as first line", i+1)
		}
	}

	return &Grid{rows: rows}, nil
}

// TreeAt ...
func (g *Grid) TreeAt(x, y int) bool {
	if g == nil {
		return false
	}
	nr := len(g.rows) - 1
	if y > nr {
		return false
	}

	r := g.rows[y]
	return r.TreeAt(x)
}

// Height ...
func (g *Grid) Height() int {
	if g == nil || g.rows == nil {
		return 0
	}
	return len(g.rows)
}
