package day13

var nextTurn = map[int]string{
	2: "left",
	1: "straight",
	0: "right",
}

// Cart ...
type Cart struct {
	facing Facing
	turn   int

	xvel int
	yvel int

	x int
	y int
}

// Process ...
func (c *Cart) Process(in string) {
	// fmt.Printf("%#v\t=>\tprocess(%v)\t=>\t", c, in)
	c.Move()

	if in == "+" {
		c.turn--
		if c.turn < 0 {
			c.turn = 2
		}
		nt := nextTurn[c.turn]
		// fmt.Printf("need to turn '%v'\n", nt)
		c.facing.Turn(nt)

		c.xvel, c.yvel = c.facing.Velocity()
	} else {
		c.facing.Next(in)
		c.xvel, c.yvel = c.facing.Velocity()
	}
	// fmt.Printf("\t%s - %#v\n", c.facing, c)
}

// Move ...
func (c *Cart) Move() {
	c.x += c.xvel
	c.y += c.yvel
}

// Under ...
func (c Cart) Under() string {
	switch c.facing {
	case FacingSouth:
		fallthrough
	case FacingNorth:
		return "|"
	case FacingEast:
		fallthrough
	case FacingWest:
		return "-"
	}
	return ""
}

// NextPos ...
func (c Cart) NextPos() (int, int) {
	return c.x + c.xvel, c.y + c.yvel
}

// CreateCart ...
func CreateCart(in string, x, y int) *Cart {
	f := StringToFacing(in)
	xvel, yvel := f.Velocity()

	switch f {
	case FacingNull:
		return nil
	default:
		return &Cart{
			facing: f,
			x:      x,
			y:      y,
			xvel:   xvel,
			yvel:   yvel,
		}
	}
}
