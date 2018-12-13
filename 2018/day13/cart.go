package day13

var nextTurn = []string{
	"left",
	"straight",
	"right",
}

// Cart ...
type Cart struct {
	facing Facing

	prevInput string

	prevTurn int

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
		c.prevTurn++
		n := len(nextTurn) % c.prevTurn
		nt := nextTurn[n]
		c.facing.Turn(nt)

		c.xvel, c.yvel = c.facing.Velocity()
	} else {
		c.facing.Next(in)
		c.xvel, c.yvel = c.facing.Velocity()
	}
	// fmt.Printf("\t%#v\n", c)
}

// Move ...
func (c *Cart) Move() {
	c.x += c.xvel
	c.y += c.yvel
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
