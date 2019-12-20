package day18

import (
	"fmt"
	"strings"
)

// Map ...
type Map struct {
	data map[int]map[int]*tile

	doors map[string]*tile
	keys  map[string]*tile

	pathable grid

	xpos int
	ypos int

	standing *tile
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

	pathable := grid{}

	var xpos, ypos, idx int
	var standing *tile

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
					t.key = k
				}
				if d == strings.ToUpper(d) {
					dr := fmt.Sprintf("%v", d)
					doors[dr] = t
					t.door = dr
				}
			}

			pathable[t.hashKey()] = t
			row[x] = t

			if d == "@" {
				xpos = x
				ypos = y
				standing = t
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

	m := &Map{
		data:     data,
		pathable: pathable,
		keys:     keys,
		doors:    doors,
		standing: standing,
		xpos:     xpos,
		ypos:     ypos,
	}

	return m, nil
}

// getKeys ...
func (m Map) getableKeys() []string {
	keys := m.pathable.getableKeys(m.standing)
	out := []string{}
	for _, v := range keys {
		out = append(out, v.key)
	}
	return out
}

// copy ...
func (m Map) getCopy() grid {
	return m.pathable.getCopy()
}

func findByName(graph grid, name string) *tile {
	for n, t := range graph {
		if n == name {
			return t
		}
	}
	return nil
}

func findByDoorOrKey(graph grid, name string) *tile {
	for _, t := range graph {
		if t.key == name {
			return t
		}

		if t.door == name {
			return t
		}
	}
	return nil
}

func findByCoord(graph grid, x, y int) *tile {
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

	var start, end *tile
	if from == "@" {
		start = findByCoord(m.pathable, m.xpos, m.ypos)
	} else {
		start = findByName(m.pathable, from)
		if start == nil {
			start = findByDoorOrKey(m.pathable, from)
		}
	}

	if key == "@" {
		end = findByCoord(m.pathable, m.xpos, m.ypos)
	} else {
		end = findByName(m.pathable, key)
		if end == nil {
			end = findByDoorOrKey(m.pathable, key)
		}
	}

	return m.pathable.tileToTileSteps(start, end)
}
