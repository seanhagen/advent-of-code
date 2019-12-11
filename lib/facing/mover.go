package facing

import (
	"fmt"
	"reflect"
	"sort"
)

// NewObjFn is the signature of the function used by a Mover to create
// a new object when it moves
type NewObjFn func() interface{}

// Mover is an object that starts at a given X,Y coordinate (usually 0,0) and then processes
// instructions to move north, south, east, or west. These instructions can come in one of two forms:
//   1) instructions to turn left or right ( after which it'll move forward in the direction it now faces )
//   2) instructions to move north, south, east, or west ( no turning, just go )
type Mover struct {
	things map[int]map[int]interface{}

	x    int
	xdir int
	y    int
	ydir int

	facing Direction

	newObj NewObjFn
	inType interface{}
}

// Config is used when creating a new Mover. X & Y are the coordinates of where it starts, and Facing is
// what direction it starts facing ( default is North ). NewObj is NewObjFn used to create new things
// when it moves to stand/rest on.
type Config struct {
	X      int
	Y      int
	Facing Direction
	NewObj NewObjFn
	Type   interface{}
}

func defaultNewObj() interface{} {
	return "."
}

var defaultType string

func typeEqual(a, b interface{}) bool {
	t1 := reflect.TypeOf(a)
	t2 := reflect.TypeOf(b)
	return t1 == t2
}

// NewMover ...
func NewMover(cnf *Config) (*Mover, error) {
	if cnf == nil {
		cnf = &Config{
			Facing: North,
			NewObj: defaultNewObj,
			Type:   defaultType,
		}
	}

	if !typeEqual(cnf.Type, cnf.NewObj()) {
		return nil, fmt.Errorf("new object function doesn't return same type as config.Type")
	}

	x, y, f := cnf.X, cnf.Y, cnf.Facing
	th := map[int]map[int]interface{}{
		x: map[int]interface{}{
			y: cnf.NewObj(),
		},
	}

	return &Mover{
		things: th,

		x:    cnf.X,
		xdir: Vectors[f][0],
		y:    cnf.Y,
		ydir: Vectors[f][1],

		facing: f,

		newObj: cnf.NewObj,
		inType: cnf.Type,
	}, nil
}

// Turn takes a direction to turn, then moves forward in the new direction
func (m *Mover) Turn(i Turn) {
	f := TurnTo(m.facing, i)
	m.facing = f
	m.moveForward()
}

// Move takes a direction to face, then moves forward
func (m *Mover) Move(f Direction) {
	m.facing = f
	m.moveForward()
}

// moveForward ...
func (m *Mover) moveForward() {
	f := m.facing
	xdir, ydir := Vectors[f][0], Vectors[f][1]
	m.xdir = xdir
	m.ydir = ydir
	m.x += xdir
	m.y += ydir

	row, ok := m.things[m.x]
	if !ok {
		row = map[int]interface{}{}
	}

	t, ok := row[m.y]
	if !ok {
		t = m.newObj()
		row[m.y] = t
		m.things[m.x] = row
	}
}

// GetAt will return the thing stored at X,Y. It will create a new thing
// by calling the configured NewObj function if nothing exists at that spot.
func (m Mover) GetAt(x, y int) interface{} {
	row, ok := m.things[x]
	if !ok {
		row = map[int]interface{}{}
	}

	t, ok := row[y]
	if !ok {
		t = m.newObj()
	}

	return t
}

// GetCurrent ...
func (m Mover) GetCurrent() interface{} {
	return m.GetAt(m.x, m.y)
}

// SetAt sets the thing at x,y to n. Returns an error if the type of n
// isn't the same as the type passed in in Config.Type
func (m *Mover) SetAt(x, y int, n interface{}) error {
	if !typeEqual(m.inType, n) {
		return fmt.Errorf("value %#v type %T isn't same type as %T", n, n, m.inType)
	}

	row, ok := m.things[m.x]
	if !ok {
		row = map[int]interface{}{}
	}

	row[m.y] = n
	m.things[m.x] = row
	return nil
}

// SetCurent ...
func (m *Mover) SetCurent(n interface{}) error {
	return m.SetAt(m.x, m.y, n)
}

// ModifyAt takes an x,y coordinate and a function. The function will be passed
// the current thing at x,y, and the return value will become the new thing at x,y.
// Will return an error if the value returned from fn isn't the same type as the
// value passed in.
func (m *Mover) ModifyAt(x, y int, fn func(interface{}) interface{}) error {
	t := m.GetAt(x, y)

	newT := fn(t)
	if !typeEqual(t, newT) {
		return fmt.Errorf("modify func returns different type")
	}

	return m.SetAt(x, y, newT)
}

// ModifyCurrent ...
func (m *Mover) ModifyCurrent(fn func(interface{}) interface{}) error {
	return m.ModifyAt(m.x, m.y, fn)
}

// Iterate ...
func (m *Mover) Iterate(fn func(int, int, interface{})) {
	for x, row := range m.things {
		for y, t := range row {
			fn(x, y, t)
		}
	}
}

// Print ...
func (m Mover) Print(fn func(interface{})) {
	xkeys := []int{}
	ykeys := []int{}
	for i, row := range m.things {
		xkeys = append(xkeys, i)
		for j := range row {
			ykeys = append(ykeys, j)
		}
	}

	sort.Ints(xkeys)
	sort.Ints(ykeys)

	minX := xkeys[0] - 1
	maxX := xkeys[len(xkeys)-1] + 1

	minY := ykeys[0] - 1
	maxY := ykeys[len(ykeys)-1] + 1

	// spew.Dump(minX, maxX, minY, maxY)

	for j := maxY; j >= minY; j-- {
		for i := minX; i <= maxX; i++ {
			t := m.GetAt(i, j)
			fn(t)
		}
		fmt.Printf("  \n")
	}

}
