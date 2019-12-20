package day18

import "fmt"

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
