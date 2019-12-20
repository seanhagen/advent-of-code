package day18

import (
	"fmt"
	"strings"
)

type grid map[string]*tile

func inSlice(a string, s []string) bool {
	for _, v := range s {
		if a == v {
			return true
		}
	}
	return false
}

// Print ...
func (g grid) Print(ind, atx, aty int) {
	minX, maxX, minY, maxY := 1000, -1000, 1000, -1000
	for _, t := range g {
		x, y := t.x(), t.y()
		if x > maxX {
			maxX = x
		}
		if x < minX {
			minX = x
		}
		if y > maxY {
			maxY = y
		}
		if y < minY {
			minY = y
		}
	}
	minX--
	maxX++
	minY--
	maxY++

	for j := minY; j <= maxY; j++ {
		printWithIndent(ind, " ")
		for i := minX; i <= maxX; i++ {
			t := g.tileAt(i, j)
			if t == nil {
				fmt.Printf("#")
			} else {
				if i == atx && j == aty {
					fmt.Printf("@")
				} else {
					if t.door != "" {
						fmt.Printf("%v", t.door)
					} else if t.key != "" {
						fmt.Printf("%v", t.key)
					} else {
						fmt.Printf(".")
					}
				}
			}
		}
		fmt.Printf("\n")
	}
}

// tileAt ...
func (g grid) tileAt(x, y int) *tile {
	for _, t := range g {
		if t.x() == x && t.y() == y {
			return t
		}
	}
	return nil
}

// tileToTileSteps ...
func (g grid) tileToTileSteps(start, end *tile) int {
	if start.eq(end) {
		return 0
	}
	ng := g.getCopy()
	return bfs(ng, start, end)
}

// getTileOfKey ...
func (g grid) getTileOfKey(in string) *tile {
	for _, t := range g {
		if t.key == in {
			return t
		}
	}
	return nil
}

// foundKeys ...
func (g grid) removeKey(in string) grid {
	fmt.Printf("removing key %v\n", in)
	out := g.getCopy()
	for k, t := range out {
		if t.key == in {
			// fmt.Printf("tile %v matches key\n", t.hashKey())
			delete(out, k)
			t.removeKey()
			out[t.hashKey()] = t
		}
	}
	in = strings.ToUpper(in)
	for k, t := range out {
		if t.door == in {
			// fmt.Printf("tile %v matches key door!\n", t.hashKey())
			delete(out, k)
			t.removeDoor()
			out[t.hashKey()] = t
		}
	}
	return out
}

// getableKeys ...
func (g grid) getableKeys(from *tile) []*tile {
	gettable := []*tile{}
	// g.Print()
	for _, t := range g.keys() {
		dist := g.tileToTileSteps(from, t)
		if dist >= 0 {
			gettable = append(gettable, t)
		}
	}
	return gettable
}

// keys ...
func (g grid) keys() []*tile {
	out := []*tile{}
	for _, t := range g {
		if t.key != "" {
			out = append(out, t)
		}
	}
	return out
}

// doors ...
func (g grid) doors() []*tile {
	out := []*tile{}
	for _, t := range g {
		if t.door != "" {
			out = append(out, t)
		}
	}
	return out
}

// copy ...
func (g grid) getCopy() grid {
	// fmt.Printf("copying grid:\n")
	// g.Print(0, -1, -1)

	out := make(grid)
	for n, t := range g {
		// fmt.Printf("copying tile %v\n", n)
		nei := map[string]edge{}
		for nn, e := range t.neighbours {
			nei[nn] = e
		}

		nt := &tile{
			id:         t.id,
			key:        t.key,
			door:       t.door,
			coord:      t.coord,
			neighbours: nei,
		}

		out[n] = nt
	}
	return out
}
