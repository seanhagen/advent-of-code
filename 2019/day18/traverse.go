package day18

import "fmt"

type path struct {
	name     string
	node     *tile
	steps    int
	parent   *path
	children []path

	keysFound []string

	final bool

	mapNow grid
}

func printWithIndent(ind int, format string, args ...interface{}) {
	if ind > 0 {
		for i := 0; i < ind; i++ {
			fmt.Printf("....")
		}
	}
	fmt.Printf(format, args...)
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
	printWithIndent(ind, "gettable keys from node: %v\n", tmpk)

	// for all those keys
	for _, k := range keys {
		ng := tmpGrid.getCopy()
		d := ng.tileToTileSteps(in.node, k)

		if in.parent != nil {
			printWithIndent(ind, "parent: %v\n", in.name)
		}

		printWithIndent(ind, "acting on key %v\n", k.key)
		printWithIndent(ind, "distance from node to key %v is %v\n", k.key, d)

		kn := k.key
		kf := []string{kn}
		if len(in.keysFound) > 0 {
			kf = append(kf, in.keysFound...)
		}
		ng.Print(ind, in.node.x(), in.node.y())
		printWithIndent(ind, "remove key %v\n", kn)

		ng2 := ng.removeKey(kn)
		keysLeft := ng2.keys()
		gettableKeys := ng2.getableKeys(k)
		// ng.Print(ind, k.x(), k.y())

		fmt.Printf("\n")

		k.key = ""

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

		if len(keysLeft) > 0 && len(gettableKeys) > 0 {
			np = traverse(ind+1, np)
		}

		in.children = append(in.children, np)
	}

	return in
}
