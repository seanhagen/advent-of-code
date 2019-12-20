package day18

import (
	"strings"
)

// inputToData ...
func inputToData(i string) map[int]map[int]string {
	out := map[int]map[int]string{}
	i = strings.TrimSpace(i)
	if i == "" {
		return out
	}

	x, y := 0, 0
	tmp := map[int]string{}
	for _, v := range strings.Split(i, "") {
		if v == "\n" {
			x = 0
			out[y] = tmp
			y++
			tmp = map[int]string{}
			continue
		}
		tmp[x] = v
		x++
	}
	out[y] = tmp
	return out
}

type queueNode struct {
	t    *tile
	dist int
}

// bfs returns the number of steps to get from the starting tile to the ending tile.
// it returns -1 if there is no path that doesn't go through a currently locked door
func bfs(graph map[string]*tile, start, end *tile) int {
	if start == nil || end == nil {
		return -1
	}

	visited := map[string]bool{}
	queue := []queueNode{{start, 0}}
	md := -1
	fx, fy := end.x(), end.y()

	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]

		i, j, dist := n.t.x(), n.t.y(), n.dist

		if i == fx && j == fy {
			md = dist
			break
		}

		for _, nei := range n.t.neighbours {
			if _, ok := visited[nei.b.hashKey()]; ok {
				continue
			}

			visited[nei.b.hashKey()] = true

			if nei.b.door == "" {
				if end.key != "" && nei.b.key != "" && end.key != nei.b.key {
					// if we're on a tile with a key
					// and it's not the same key as our ending key
					// return -1 because this key should be picked up on this
					// path before the stated end tile
					// fmt.Printf("unable to get to %v from %v because tile %v has a key and it's in the way\n", end.hashKey(), start.hashKey(), nei.b.hashKey())
					//return md
					continue
				}

				// if a key is in the way, consider it an invalid path
				// the key should be picked up first before continuing on this path
				queue = append(queue, queueNode{nei.b, dist + 1})
				// } else {
				// 	fmt.Printf("unable to get to %v from %v because of door %v\n", end.hashKey(), start.hashKey(), nei.b.hashKey())
			}
		}
	}

	return md
}

func revTopSort(graph map[string]*tile, st string) []string {
	ord := []string{}
	g := graph[st]
	visited := map[string]bool{
		g.hashKey(): true,
	}

	var fn func(v *tile, name string)
	fn = func(v *tile, name string) {
		if v == nil {
			return
		}

		if len(v.neighbours) == 0 {
			ord = append(ord, v.hashKey())
			return
		}

		valo := false
		for _, edge := range v.neighbours {
			n := edge.b.hashKey()
			if _, ok := visited[n]; ok {
				continue
			}
			valo = true
			visited[n] = true

			// fmt.Printf("\t%v\n", n)
			fn(graph[n], n)
		}

		if !valo {
			return
		}
	}
	fn(g, st)
	return ord
}
