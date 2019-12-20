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
	t    tile
	dist int
}

// bfs returns the number of steps to get from the starting tile to the ending tile.
// it returns -1 if there is no path that doesn't go through a currently locked door
func bfs(graph grid, start, end tile) int {
	visited := map[string]bool{}
	queue := []queueNode{{start, 0}}
	md := -1
	fx, fy := end.x(), end.y()

	vecs := [][]int{
		[]int{1, 0},
		[]int{-1, 0},
		[]int{0, 1},
		[]int{0, -1},
	}

	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]

		i, j, dist := n.t.x(), n.t.y(), n.dist

		if i == fx && j == fy {
			md = dist
			break
		}

		check := []tile{}
		for _, v := range vecs {
			x := i + v[0]
			y := j + v[1]
			nt := graph.tileAt(x, y)
			if nt != nil {
				check = append(check, *nt)
			}
		}

		for _, t := range check {
			if _, ok := visited[t.hashKey()]; ok {
				continue
			}
			visited[t.hashKey()] = true

			if t.door == "" {
				if t.key != "" && t.key != end.key {
					continue
				}

				queue = append(queue, queueNode{t, dist + 1})
			}
		}
	}

	return md
}

func revTopSort(graph grid, st string) []string {
	ord := []string{}
	g, ok := graph[st]
	if !ok {
		return ord
	}

	visited := map[string]bool{
		g.hashKey(): true,
	}

	vecs := [][]int{
		[]int{1, 0},
		[]int{-1, 0},
		[]int{0, 1},
		[]int{0, -1},
	}

	var fn func(t tile, name string)
	fn = func(t tile, name string) {
		check := []tile{}
		for _, v := range vecs {
			x := t.x() + v[0]
			y := t.y() + v[1]
			nt := graph.tileAt(x, y)
			if nt != nil {
				check = append(check, *nt)
			}
		}

		if len(check) == 0 {
			ord = append(ord, t.hashKey())
			return
		}

		valo := false
		for _, t := range check {
			n := t.hashKey()
			if _, ok := visited[n]; ok {
				continue
			}
			valo = true
			visited[n] = true

			fn(graph[n], n)
		}

		if !valo {
			return
		}
	}
	fn(g, st)
	return ord
}
