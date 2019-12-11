package facing

import "strings"

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
	North: {0, 1},
	South: {0, -1},
	West:  {-1, 0},
	East:  {1, 0},
}

// Turn is used to indicate what direction an object that uses Direction should
// change it's orientation to face
type Turn int

// Left means turn left
const Left Turn = 0

// Right means turn right
const Right Turn = 1

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
