package day18

import "fmt"

type tile struct {
	id         int
	key        string
	door       string
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
	if t2 == nil {
		return false
	}
	return t.hashKey() == t2.hashKey()
}

// addEdge ...
func (t *tile) addEdge(t2 *tile) {
	e := edge{t, t2}
	t.neighbours[e.hashKey()] = e
}

// hashKey ...
func (t tile) hashKey() string {
	return fmt.Sprintf("tile<%v, %v, %v, %v>", t.id, t.key, t.door, t.coord.hashKey())
}

// removeKey ...
func (t *tile) removeKey() {
	if t.key == "" {
		return
	}

	for str, edg := range t.neighbours {
		edg.b.removeNeighbourKey(t.hashKey())
		delete(t.neighbours, str)
		t.addEdge(edg.b)
	}
	t.key = ""
}

// removeNeighbourKey ...
func (t *tile) removeNeighbourKey(in string) {
	// fmt.Printf("neighbour %v, removing key %v\n", t.hashKey(), in)
	for str, edg := range t.neighbours {
		if edg.b.hashKey() == in {
			delete(t.neighbours, str)
			edg.b.key = ""
			t.addEdge(edg.b)
		}
	}
}

// removeDoor ...
func (t *tile) removeDoor() {
	if t.door == "" {
		return
	}

	for str, edg := range t.neighbours {
		edg.b.removeNeighbourDoor(t.hashKey())
		delete(t.neighbours, str)
		t.addEdge(edg.b)
	}
	t.door = ""
}

// removeNeighbourDoor ...
func (t *tile) removeNeighbourDoor(in string) {
	// fmt.Printf("neighbour %v, removing door %v\n", t.hashKey(), in)
	for str, edg := range t.neighbours {
		if edg.b.hashKey() == in {
			delete(t.neighbours, str)
			edg.b.door = ""
			t.addEdge(edg.b)
		}
	}
}
