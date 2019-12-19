package day18

import (
	"fmt"
	"strings"
)

type coord [2]int

// hashKey ...
func (c coord) hashKey() string {
	return fmt.Sprintf("coord{%v,%v}", c[0], c[1])
}

// eq ...
func (c coord) eq(x, y int) bool {
	return c[0] == x && c[1] == y
}

// x ...
func (c coord) x() int {
	return c[0]
}

// y ...
func (c coord) y() int {
	return c[1]
}

type edge struct {
	a *tile
	b *tile
}

// hashKey ...
func (e edge) hashKey() string {
	if e.a == nil && e.b == nil {
		return fmt.Sprintf("edge<nil <-> nil>")
	}

	if e.a == nil && e.b != nil {
		return fmt.Sprintf("edge<nil <-> %v>", e.b.hashKey())
	}

	if e.a != nil && e.b == nil {
		return fmt.Sprintf("edge<%v <-> nil>", e.a.hashKey())
	}

	return fmt.Sprintf("edge<%v <-> %v>", e.a.hashKey(), e.b.hashKey())
}

type tile struct {
	id         int
	key        *string
	door       *string
	coord      coord
	neighbours map[string]edge
}

// x ...
func (t tile) x() int {
	return t.coord.x()
}

// y ...
func (t tile) y() int {
	return t.coord.y()
}

// eq ...
func (t tile) eq(t2 *tile) bool {
	return t.hashKey() == t2.hashKey()
}

// addEdge ...
func (t *tile) addEdge(t2 *tile) {
	e := edge{t, t2}
	t.neighbours[e.hashKey()] = e
}

// hashKey ...
func (t tile) hashKey() string {
	var key, door string
	if t.key != nil {
		key = fmt.Sprintf("%v", *t.key)
	}
	if t.door != nil {
		door = fmt.Sprintf("%v", *t.door)
	}
	return fmt.Sprintf("tile<%v, %v, %v, %v>", t.id, key, door, t.coord.hashKey())
}

// Map ...
type Map struct {
	data map[int]map[int]*tile

	doors map[string]*tile
	keys  map[string]*tile

	pathable map[string]*tile

	xpos int
	ypos int
}

// NewMap ...
func NewMap(input string) (*Map, error) {
	if input == "" {
		return nil, fmt.Errorf("bad input, blank string")
	}
	base := inputToData(input)

	data := map[int]map[int]*tile{}
	doors := map[string]*tile{}
	keys := map[string]*tile{}

	pathable := map[string]*tile{}

	var xpos, ypos, idx int

	for y, tr := range base {
		row, ok := data[y]
		if !ok {
			row = map[int]*tile{}
		}

		for x, d := range tr {
			if d == "#" {
				continue
			}

			t, ok := row[x]
			if !ok {
				t = &tile{id: idx, coord: coord{x, y}, neighbours: map[string]edge{}}
				idx++
			}

			if d != "." && d != "@" {
				if d == strings.ToLower(d) {
					k := fmt.Sprintf("%v", d)
					keys[k] = t
					t.key = &k
				}
				if d == strings.ToUpper(d) {
					dr := fmt.Sprintf("%v", d)
					doors[dr] = t
					t.door = &dr
				}
			}

			pathable[t.hashKey()] = t
			row[x] = t

			if d == "@" {
				xpos = x
				ypos = y
			}

			row[x] = t
		}
		data[y] = row
	}

	for y, row := range data {
		for x, tile := range row {
			if upr, ok := data[y-1]; ok {
				// there's a row above this one, see if there's a tile
				if upt, ok := upr[x]; ok {
					// there is a tile above this one, add an edge
					tile.addEdge(upt)
				}
			}

			if left, ok := row[x-1]; ok {
				// there's a tile to the left, add an edge
				tile.addEdge(left)
			}

			if right, ok := row[x+1]; ok {
				// there's a tile to the right, add an edge
				tile.addEdge(right)
			}

			if dwnr, ok := data[y+1]; ok {
				// there's a row below this one, see if there's a tile
				if dwnt, ok := dwnr[x]; ok {
					// there is, add edge
					tile.addEdge(dwnt)
				}
			}
		}
	}

	// fmt.Printf("pathable:\n")
	// for n := range pathable {
	// 	fmt.Printf("n: %v\n", n)
	// }
	// fmt.Printf("\n\n")

	m := &Map{
		data:     data,
		pathable: pathable,
		keys:     keys,
		doors:    doors,
		xpos:     xpos,
		ypos:     ypos,
	}

	return m, nil
}

// copy ...
func (m Map) copy() map[string]*tile {
	out := map[string]*tile{}

	for n, t := range m.pathable {
		tmp := *t
		nei := map[string]edge{}
		for nn, e := range tmp.neighbours {
			nei[nn] = e
		}

		nt := &tile{
			id:         tmp.id,
			coord:      tmp.coord,
			neighbours: nei,
		}

		if tmp.key != nil {
			var k string
			k = fmt.Sprintf("%v", *tmp.key)
			nt.key = &k
		}
		if tmp.door != nil {
			var d string
			d = fmt.Sprintf("%v", *tmp.door)
			nt.door = &d
		}

		out[n] = nt
	}

	return out
}

func findByName(graph map[string]*tile, name string) *tile {
	for n, t := range graph {
		if n == name {
			return t
		}
	}
	return nil
}

func findByDoorOrKey(graph map[string]*tile, name string) *tile {
	for _, t := range graph {
		if t.key != nil && *t.key == name {
			return t
		}

		if t.door != nil && *t.door == name {
			return t
		}
	}
	return nil
}

func findByCoord(graph map[string]*tile, x, y int) *tile {
	for _, t := range graph {
		if t.coord.x() == x && t.coord.y() == y {
			return t
		}
	}
	return nil
}

// stepsToKey ...
func (m Map) stepsToKey(from, key string) int {
	if from == key {
		return 0
	}

	ng := m.copy()

	var start, end *tile

	if from == "@" {
		start = findByCoord(ng, m.xpos, m.ypos)
	} else {
		start = findByName(ng, from)
		if start == nil {
			start = findByDoorOrKey(ng, from)
		}
	}

	if key == "@" {
		end = findByCoord(ng, m.xpos, m.ypos)
	} else {
		end = findByName(ng, key)
		if end == nil {
			end = findByDoorOrKey(ng, key)
		}
	}

	return bfs(ng, start, end)
}

type queueNode struct {
	t    *tile
	dist int
}

// bfs returns the number of steps to get from the starting tile to the ending tile.
// it returns -1 if there is no path that doesn't go through a currently locked door
func bfs(graph map[string]*tile, start, end *tile) int {
	if start == nil || end == nil {
		fmt.Printf("start or end nil\n")
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

			if nei.b.door == nil {
				queue = append(queue, queueNode{nei.b, dist + 1})
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

			fmt.Printf("\t%v\n", n)
			fn(graph[n], n)
		}

		if !valo {
			return
		}
	}

	fmt.Printf("starting with \n%v\n\n", st)
	fn(g, st)
	return ord
}

// keysRequiredForDoor ...
func (m Map) keysRequiredForDoor(door string) int {
	return 0
}

// AllKeySteps ...
func (m Map) AllKeySteps() int {
	return 0
}
