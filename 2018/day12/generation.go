package day12

import (
	"fmt"
	"sort"
)

// Generation ...
type Generation struct {
	id   int
	pots Row
}

func createGeneration(in string) *Generation {
	p := createRowFromString(in)
	gen := &Generation{
		id:   0,
		pots: p,
	}
	return gen
}

// output ...
func (g Generation) output() string {
	first := 100
	max := -100
	keys := []int{}

	for i, v := range g.pots.d {
		keys = append(keys, i)
		if v == "#" && i < first {
			first = i
		}
		if v == "#" && i > max {
			max = i
		}
	}
	sort.Ints(keys)

	t := ""
	for i := first - 3; i < max+12; i++ {
		x := g.pots.d[i]
		// fmt.Printf("d[%v] = %v\n", i, x)
		if x == "" {
			t = fmt.Sprintf("%v.", t)
		} else {
			t = fmt.Sprintf("%v%v", t, g.pots.d[i])
		}
	}

	t = fmt.Sprintf("%2d: %v", g.id, t)
	return t
}

// Next ...
func (g Generation) Next(rules []Rule) *Generation {
	x := &Generation{id: g.id + 1, pots: g.pots}
	x.pots = x.pots.getNext(rules)
	return x
}
