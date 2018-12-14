package day13

// Facing ...
type Facing int

const (
	// FacingNull is for an invalid direction
	FacingNull Facing = iota
	// FacingNorth ...
	FacingNorth
	// FacingEast ...
	FacingEast
	// FacingSouth ...
	FacingSouth
	// FacingWest ...
	FacingWest
)

// Turn ...
func (f *Facing) Turn(dir string) {
	if dir == "left" {
		f.turnLeft()
	}

	if dir == "right" {
		f.turnRight()
	}
}

// turnRight ...
func (f *Facing) turnRight() {
	switch *f {
	case FacingNorth:
		*f = FacingEast
	case FacingEast:
		*f = FacingSouth
	case FacingSouth:
		*f = FacingWest
	case FacingWest:
		*f = FacingNorth
	}
}

// turnLeft ...
func (f *Facing) turnLeft() {
	switch *f {
	case FacingNorth:
		*f = FacingWest
	case FacingEast:
		*f = FacingNorth
	case FacingSouth:
		*f = FacingEast
	case FacingWest:
		*f = FacingSouth
	}
}

// String ...
func (f Facing) String() string {
	switch f {
	case FacingNorth:
		return "^"
	case FacingEast:
		return ">"
	case FacingSouth:
		return "v"
	case FacingWest:
		return "<"
	}
	return ""
}

// Velocity ...
func (f Facing) Velocity() (int, int) {
	x := 0
	y := 0

	switch f {
	case FacingNorth:
		y = -1
	case FacingEast:
		x = 1
	case FacingSouth:
		y = 1
	case FacingWest:
		x = -1
	}

	return x, y
}

// Next ...
func (f *Facing) Next(in string) {
	switch in {
	case "/":
		switch *f {
		case FacingNorth:
			*f = FacingEast
		case FacingSouth:
			*f = FacingWest
		case FacingEast:
			*f = FacingNorth
		case FacingWest:
			*f = FacingSouth
		}
	case "\\":
		switch *f {
		case FacingNorth:
			*f = FacingWest
		case FacingSouth:
			*f = FacingEast
		case FacingEast:
			*f = FacingSouth
		case FacingWest:
			*f = FacingNorth
		}
	}
}

func StringToFacing(in string) Facing {
	switch in {
	case "^":
		return FacingNorth
	case ">":
		return FacingEast
	case "v":
		return FacingSouth
	case "<":
		return FacingWest
	default:
		return FacingNull
	}
}
