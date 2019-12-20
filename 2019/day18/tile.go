package day18

import "fmt"

type tile struct {
	id    int
	key   string
	door  string
	coord coord
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
func (t tile) eq(t2 tile) bool {
	return t.hashKey() == t2.hashKey()
}

// hashKey ...
func (t tile) hashKey() string {
	return fmt.Sprintf("tile<%v, %v, %v, %v>", t.id, t.key, t.door, t.coord.hashKey())
}

// removeKey ...
func (t *tile) removeKey() {
	t.key = ""
}

// removeDoor ...
func (t *tile) removeDoor() {
	t.door = ""
}
