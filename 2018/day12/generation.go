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

// count ...
func (g Generation) count() int {
	data := g.pots.getHashes()
	sum := 0

	// fmt.Printf("generation %v: %v pots have plants, ", g.id, len(data))
	// fmt.Printf("plants in pots [")
	// for k := range data {
	// 	fmt.Printf("%v ", k)
	// }
	// fmt.Printf("]\n")

	for i := range data {
		sum += i
	}

	return sum
}

// output ...
func (g Generation) output() string {
	first := 100
	max := -100
	keys := []int{}

	for i, v := range g.pots.data {
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
	for i := -3; i < 36; i++ {
		t = fmt.Sprintf("%v%v", t, g.pots.data[i])
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
