package day3

import (
	"fmt"
	"strings"
)

type Grid struct {
	rows []Row
	l    int
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

	return &Grid{rows: rows, l: l}, nil
}

// TreeAt ...
func (g *Grid) TreeAt(x, y int) bool {
	if g == nil {
		return false
	}
	nr := len(g.rows)
	if y >= nr {
		return false
	}

	r := g.rows[y]

	l := len(r.trees)
	nx := (x%l + l) % l

	z := r.TreeAt(nx)
	return z
}

// PrintTreeAt ...
func (g *Grid) PrintTreeAt(x, y int) bool {
	ret := false
	l := len(g.rows[0].trees)
	x = (x%l + l) % l
	for i, v := range g.rows {
		for j, vv := range v.trees {
			if i == y && j == x {
				if vv == Tree {
					fmt.Printf("X")
					ret = true
				} else {
					fmt.Printf("0")
				}

			} else {
				fmt.Printf("%v", vv)
			}
		}
		fmt.Printf("\n")
	}
	return ret
}

// Height ...
func (g *Grid) Height() int {
	if g == nil || g.rows == nil {
		return 0
	}
	return len(g.rows)
}
