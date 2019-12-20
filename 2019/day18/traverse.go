package day18

import (
	"fmt"
	"math"
	"sort"
	"strings"
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

var memoizedScore = map[int]map[int]map[string]int{}

func getMS(x, y int, k string) int {
	xv, ok := memoizedScore[x]
	if !ok {
		xv = map[int]map[string]int{}
	}

	yv, ok := xv[y]
	if !ok {
		yv = map[string]int{}
	}
	sv, ok := yv[k]
	if !ok {
		return -1
	}
	return sv
}

func setMS(x, y int, k string, v int) {
	xv, ok := memoizedScore[x]
	if !ok {
		xv = map[int]map[string]int{}
	}

	yv, ok := xv[y]
	if !ok {
		yv = map[string]int{}
	}
	yv[k] = v
	xv[y] = yv
	memoizedScore[x] = xv
}

func score(in path) int {
	curkey := in.node
	curgrid := in.mapNow.getCopy()
	keysFound := strings.Join(in.keysFound, "")
	if v := getMS(curkey.x(), curkey.y(), keysFound); v >= 0 {
		// fmt.Printf("returning memoized score for %v, %v, %v ==> %v\n", curkey.x(), curkey.y(), keysFound, v)
		return v
	}

	scores := []int{}
	reachable := curgrid.getableKeys(curkey)

	if len(reachable) == 0 {
		// fmt.Printf("no more reachable keys\n")
		return 0
	}

	for _, k := range reachable {
		grid := curgrid.getCopy()
		kyc := []string{}
		for _, v := range in.keysFound {
			kyc = append(kyc, v)
		}
		kyc = append(kyc, k.key)
		sort.Strings(kyc)
		kn := k.key
		ng := grid.removeKey(kn)

		dist := ng.tileToTileSteps(curkey, k)

		np := path{
			name:      k.key,
			node:      k,
			steps:     in.steps + dist,
			parent:    &in,
			children:  []path{},
			keysFound: kyc,
			mapNow:    ng.getCopy(),
		}
		// fmt.Printf("checking score for keys %v -> %v, dist: %3d, found: %v\n", curkey.key, k.key, dist, in.keysFound)

		ns := score(np) + dist
		scores = append(scores, ns)
	}
	smallest := int(math.MaxInt64 - 1)
	for _, s := range scores {
		if s < smallest {
			smallest = s
		}
	}
	// fmt.Printf("scores: %v --> %v\n", scores, smallest)

	setMS(curkey.x(), curkey.y(), keysFound, smallest)
	return smallest
}

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
