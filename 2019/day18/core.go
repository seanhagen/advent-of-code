package day18

/*
@ -> a -> b -> c -> e -> d -> f
                 -> d -> e -> f

*/

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
)

// SolveP1 ...
func (m Map) SolveP1() int {
	gr := m.getCopy()

	graph := path{
		name:     "@",
		node:     m.standing,
		steps:    0,
		parent:   nil,
		children: []path{},
		mapNow:   gr.getCopy(),
	}
	fmt.Printf("starting at %v, here's a random path: \n\n", m.standing.hashKey())
	graph = traverse(0, graph)

	// current := graph
	st := 0
	// for {
	// 	if len(current.children) <= 0 {
	// 		break
	// 	}
	// 	current = pickRandom(current.children)
	// 	fmt.Printf("%v -> %v\n", current.node.hashKey(), current.steps)
	// 	st += current.steps
	// }

	// fmt.Printf("\n\n")
	// fmt.Printf("total steps: %v\n", st)

	kgr := gr.keys()
	keys := []string{"@"}
	for _, k := range kgr {
		keys = append(keys, k.key)
	}

	//printPathTo(graph)
	found := findPaths(0, graph, keys)
	for _, v := range found {
		fmt.Printf("%v step path: %v\n", v.steps, v.path)
	}

	return st
}

func pickRandom(in []path) path {
	if len(in) == 1 {
		return in[0]
	}

	min := 0
	max := len(in) - 1
	p := rand.Intn(max-min) + min
	return in[p]
}

func printPathTo(in path) {
	if len(in.children) == 0 {
		parents := []string{}

		p := in.parent

		for p != nil {
			parents = append(parents, p.name)
			p = p.parent
		}

		for i := len(parents)/2 - 1; i >= 0; i-- {
			opp := len(parents) - 1 - i
			parents[i], parents[opp] = parents[opp], parents[i]
		}

		parents = append(parents, in.name)

		fmt.Printf("path: %v\n", strings.Join(parents, " -> "))
		fmt.Printf("final: %v, steps: %v\n", in.name, in.steps)
		// fmt.Printf("parents: %#v\n", parents)
		return
	}

	for _, c := range in.children {
		printPathTo(c)
	}
}

type foundPath struct {
	path  string
	steps int
}

func findPaths(ind int, in path, required []string) []foundPath {
	out := []foundPath{}

	printWithIndent(ind, "finding path for %v\n", in.name)

	if len(in.children) == 0 {
		p := in.parent
		parents := []string{}
		for p != nil {
			parents = append(parents, p.name)
			p = p.parent
		}

		for i := len(parents)/2 - 1; i >= 0; i-- {
			opp := len(parents) - 1 - i
			parents[i], parents[opp] = parents[opp], parents[i]
		}

		parents = append(parents, in.name)

		printWithIndent(ind+1, "%v step path: %v\n", in.steps, strings.Join(parents, " -> "))

		if hasAll(parents, required) {
			// fmt.Printf("doing thing\n")
			out = append(out, foundPath{strings.Join(parents, " -> "), in.steps})
		}
		// fmt.Printf("final: %v, steps: %v\n", in.name, in.steps)
		// fmt.Printf("parents: %#v\n\n---------\n\n", parents)

		return out
	}

	for _, c := range in.children {
		x := findPaths(ind+1, c, required)
		out = append(out, x...)
	}

	return out
}

func hasAll(a, b []string) bool {
	sort.Strings(a)
	sort.Strings(b)

	// fmt.Printf("parents:  %#v\n", a)
	// fmt.Printf("required: %#v\n", b)
	if len(a) != len(b) {
		// fmt.Printf("not same length:\na: %#v\nb: %#v\n", a, b)
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			// fmt.Printf("value i not same: %v != %v\n", a[i], b[i])
			return false
		}
	}

	return true
}
