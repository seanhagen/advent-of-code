package day12

import (
	"strings"
)

// Row ...
type Row struct {
	data map[int]string
}

// getFromPosition ...
func (r Row) getFromPosition(pos, dist int) string {
	newPos := pos + dist
	x := r.data[newPos]
	if x == "" {
		r.data[newPos] = "."
	}
	return r.data[newPos]
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
	for k := range r.data {
		n[k] = "."
		for _, rule := range in {
			change := r.applyRule(rule, k)
			if change {
				n[k] = rule.output
			}
		}
	}
	return Row{data: n}
}

// getHashes ...
func (r Row) getHashes() map[int]string {
	n := map[int]string{}

	for k, v := range r.data {
		if v == "#" {
			n[k] = v
		}
	}

	return n
}

func createRowFromString(in string) Row {
	in = strings.Replace(in, "initial state: ", "", -1)
	bits := strings.Split(in, "")
	init := map[int]string{}

	for i := -1000; i <= 1000; i++ {
		init[i] = "."
		if i >= 0 && i < len(bits) {
			init[i] = bits[i]
		}
	}

	return Row{data: init}
}
