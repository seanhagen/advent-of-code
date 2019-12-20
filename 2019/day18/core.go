package day18

/*
@ -> a -> b -> c -> e -> d -> f
                 -> d -> e -> f

*/

import (
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

	return score(graph)

	// steps := 100000

	// graph = traverse(0, graph)

	// kgr := gr.keys()
	// keys := []string{"@"}
	// for _, k := range kgr {
	//  keys = append(keys, k.key)
	// }
	// found := findPaths(0, graph, keys)
	// for _, v := range found {
	// 	if v.steps < steps {
	// 		steps = v.steps
	// 	}
	// 	// fmt.Printf("%v step path: %v\n", v.steps, v.path)
	// }

	// return steps
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

		fp := strings.Join(parents, " -> ")

		printWithIndent(ind+1, "%v step path: %v\n", in.steps, fp)

		if hasAll(parents, required) {
			out = append(out, foundPath{fp, in.steps})
		}

		return out
	}

	for _, c := range in.children {
		x := findPaths(ind+1, c, required)
		out = append(out, x...)
	}

	return out
}

func hasAll(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	sort.Strings(a)
	sort.Strings(b)

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
