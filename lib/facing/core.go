package facing

import (
	"fmt"
	"strings"
)

// Direction is a rune indicating what direction an object is facing
type Direction rune

// North is the same as Up in some of the puzzles
const North Direction = '^'

// South is the same as Down in some of the puzzles
const South Direction = 'v'

// West is the same as Left in some of the puzzles
const West Direction = '<'

// East is the same as Right in some of the puzzles
const East Direction = '>'

var dirNames = map[Direction]string{
	North: "North",
	South: "South",
	East:  "East",
	West:  "West",
}

var oppositeDirs = map[Direction]Direction{
	North: South,
	South: North,
	East:  West,
	West:  East,
}

// current -> want = turn
var turnDur = map[Direction]map[Direction]Turn{
	North: map[Direction]Turn{
		South: Right,
		East:  Right,
		West:  Left,
	},
	South: map[Direction]Turn{
		East:  Left,
		West:  Right,
		North: Right,
	},
	East: map[Direction]Turn{
		North: Left,
		South: Right,
		West:  Right,
	},
	West: map[Direction]Turn{
		South: Left,
		North: Right,
		East:  Right,
	},
}

// GoString ...
func (d Direction) GoString() string {
	return fmt.Sprintf("%v", dirNames[d])
}

// String ...
func (d Direction) String() string {
	return d.GoString()
}

// DirectionFromString takes a string and returns the direction. Defaults to North if the string
// doesn't match one of the defined directions
func DirectionFromString(i string) Direction {
	switch i {
	case "v":
		fallthrough
	case "V":
		return South

	case "<":
		return West

	case ">":
		return East
	}

	return North
}

// DirectionSliceFromString takes a string and treats it as a list of directions
func DirectionSliceFromString(i string) []Direction {
	bits := strings.Split(i, "")
	out := []Direction{}
	for _, v := range bits {
		out = append(out, DirectionFromString(v))
	}
	return out
}

// Vectors is a map of Direction -> x,y direction, so Vectors[Up] returns {0,1}.
// {0,1} means 0 movement in the X plane, and +1 movement in the Y plane.
var Vectors = map[Direction][2]int{
	North: {0, -1},
	South: {0, 1},
	West:  {-1, 0},
	East:  {1, 0},
}

// Turn is used to indicate what direction an object that uses Direction should
// change it's orientation to face
type Turn int

// NopTurn means don't turn, just go straight
const NopTurn Turn = -1

// Left means turn left
const Left Turn = 0

// Right means turn right
const Right Turn = 1

// TurnNames maps the turn to a string name
var TurnNames = map[Turn]string{
	NopTurn: "Neither",
	Left:    "Left",
	Right:   "Right",
}

// TurnTowards figures out what direction to turn in order to _eventually_ be facing in the right direction.
// So if currently facing north and want to be facing south, can turn left or right, so will return
// right ( right turns preferred ). If facing east and want to be facing south, will turn right.
// If facing west and want to be facing south, will turn left.
func TurnTowards(cur, new Direction) Turn {
	if cur == new {
		return NopTurn
	}
	return turnDur[cur][new]
}

// TurnTo takes the current Direction an object is facing, and the instruction
// on whether to turn Right or Left, and returns the new Direction of the object
func TurnTo(cur Direction, in Turn) Direction {
	switch in {
	case Left:
		return turnLeft(cur)
	default:
		return turnRight(cur)
	}
}

func turnRight(cur Direction) Direction {
	switch cur {
	case North:
		return East
	case South:
		return West
	case West:
		return North
	}
	// default current direction is Right
	return South
}
func turnLeft(cur Direction) Direction {
	switch cur {
	case North:
		return West
	case South:
		return East
	case West:
		return South
	}
	// default is Right
	return North
}
