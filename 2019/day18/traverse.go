package day18

import (
	"fmt"
	"sort"
)

type path struct {
	name     string
	node     tile
	steps    int
	parent   *path
	children []path

	keysFound []string

	final bool

	mapNow grid
}

func strSliceEq(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func printWithIndent(ind int, format string, args ...interface{}) {
	if ind > 0 {
		for i := 0; i < ind; i++ {
			fmt.Printf("..")
		}
	}
	fmt.Printf(format, args...)
}

// var memory

// traverse ...
func traverse(ind int, in path) path {
	if in.final {
		return in
	}

	tmpGrid := in.mapNow.getCopy()
	// get all the keys from current position
	keys := tmpGrid.getableKeys(in.node)
	tmpk := []string{}
	for _, k := range keys {
		tmpk = append(tmpk, k.key)
	}
	sort.Strings(tmpk)

	printWithIndent(ind, "gettable keys from %v: %v\tgot: %v\n", in.name, tmpk, in.keysFound)

	// for all those keys
	for _, k := range keys {
		// printWithIndent(ind, "node: %v, key: %v, gettable: %v\n", in.name, k.key, tmpk)

		ng := tmpGrid.getCopy()
		d := ng.tileToTileSteps(in.node, k)

		// printWithIndent(ind, "distance from node to key %v is %v\n", k.key, d)

		kn := k.key
		kf := []string{kn}
		if len(in.keysFound) > 0 {
			kf = append(kf, in.keysFound...)
		}
		sort.Strings(kf)
		// ng.Print(ind, in.node.x(), in.node.y())
		// printWithIndent(ind, "remove key %v\n", kn)

		ng2 := ng.removeKey(kn)
		keysLeft := ng2.keys()
		gettableKeys := ng2.getableKeys(k)
		// ng.Print(ind, k.x(), k.y())

		// fmt.Printf("\n")

		k.key = ""

		if len(keysLeft) > 0 && len(gettableKeys) > 0 {
			np := path{
				name:      kn,
				node:      k,
				parent:    &in,
				steps:     in.steps + d,
				children:  []path{},
				mapNow:    ng2.getCopy(),
				final:     len(keysLeft) <= 0,
				keysFound: kf,
			}
			np2 := traverse(ind+1, np)
			in.children = append(in.children, np2)
		}

		if len(keysLeft) == 0 {
			np := path{
				name:      kn,
				node:      k,
				parent:    &in,
				steps:     in.steps + d,
				children:  []path{},
				mapNow:    ng2.getCopy(),
				final:     len(keysLeft) <= 0,
				keysFound: kf,
			}
			in.children = append(in.children, np)
		}

	}

	return in
}
