package day12

import (
	"strings"
)

// Row ...
type Row struct {
	data []string
	d    map[int]string
}

// getFromPosition ...
func (r Row) getFromPosition(pos, dist int) string {
	newPos := pos + dist
	x := r.d[newPos]
	if x == "" {
		r.d[newPos] = "."
	}
	return r.d[newPos]
}

// applyRule ...
func (r Row) applyRule(in Rule, pos int) bool {
	x := []string{
		r.getFromPosition(pos, -2),
		r.getFromPosition(pos, -1),
		r.getFromPosition(pos, 0),
		r.getFromPosition(pos, 1),
		r.getFromPosition(pos, 2),
	}
	return in.matches(x)
}

// getNext ...
func (r Row) getNext(in []Rule) Row {
	n := map[int]string{}
	for k := range r.d {
		n[k] = "."
		change := false
		to := ""
		for _, rule := range in {
			change = r.applyRule(rule, k)
			if change {
				to = rule.output
			}
		}
		if change {
			n[k] = to
		}
	}

	return Row{d: n}
}

func createRowFromString(in string) Row {
	in = strings.Replace(in, "initial state: ", "", -1)
	bits := strings.Split(in, "")

	init := map[int]string{
		-3: ".",
		-2: ".",
		-1: ".",
	}

	for i, v := range bits {
		init[i] = v
	}

	initial := []string{".", ".", "."}
	initial = append(initial, bits...)
	for i := 1; i <= 10; i++ {
		initial = append(initial, ".")
		init[i+len(in)] = "."
	}

	return Row{data: initial, d: init}
}
