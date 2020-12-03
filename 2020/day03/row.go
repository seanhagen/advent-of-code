package day3

import (
	"strings"
)

const Empty = "."
const Tree = "#"

type Row struct {
	l     int
	trees []string
}

func NewRow(s string) Row {
	ch := strings.Split(s, "")
	return Row{len(ch), ch}
}

//  TreeAt ...
func (r Row) TreeAt(x int) bool {
	if x < r.l {
		// fmt.Printf("x(%v) < r.l(%v) -> %v\n", x, r.l, r.trees[x])
		return r.trees[x] == Tree
	}

	y := (x % r.l) - 1
	if y < 0 {
		y = r.l - 1
	}
	// fmt.Printf("y(%v) < r.l(%v) -> %v\n", y, r.l, r.trees[y])
	return r.trees[y] == Tree
}
