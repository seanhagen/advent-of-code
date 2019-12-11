package day11

type Color int

const White Color = 1
const Black Color = 0

// ToInt ...
func (c Color) ToInt() int {
	return int(c)
}

type Facing rune

const Up Facing = '^'
const Down Facing = 'v'
const Left Facing = '<'
const Right Facing = '>'

const left = 0
const right = 1

var dir = map[Facing][2]int{
	Up:    {0, 1},
	Down:  {0, -1},
	Left:  {-1, 0},
	Right: {1, 0},
}
