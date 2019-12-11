package facing

import (
	"fmt"
	"sort"
)

// NewObjFn is the signature of the function used by a Mover to create
// a new object when it moves
type NewObjFn func() interface{}

// Config is used when creating a new Mover.
type Config struct {
	// StartX is the starting X position of all movers
	StartX int
	// StartY is the starting Y position of all movers
	StartY int
	// Facing is what direction all movers face at the start
	Facing Direction
	// NewObj is a function used to create new things when moving
	NewObj NewObjFn
	// Type is used to check that methods are creating the right type of thing
	Type interface{}
	// NumMovers controls how many movers are created
	NumMovers int
}

// Mover is an object that starts at a given X,Y coordinate (usually 0,0) and then processes
// instructions to move north, south, east, or west. These instructions can come in one of two forms:
//   1) instructions to turn left or right ( after which it'll move forward in the direction it now faces )
//   2) instructions to move north, south, east, or west ( no turning, just go )
type Mover struct {
	things map[int]map[int]interface{}

	numMvrs int
	curMvr  int
	mvrs    []*mvr

	lastX int
	lastY int

	newObj NewObjFn
	inType interface{}
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

	x, y, f := cnf.StartX, cnf.StartY, cnf.Facing
	th := map[int]map[int]interface{}{
		x: map[int]interface{}{
			y: cnf.NewObj(),
		},
	}

	numMvr := cnf.NumMovers
	if numMvr <= 0 {
		numMvr = 1
	}

	mvrs := make([]*mvr, numMvr)
	for i := 0; i < numMvr; i++ {
		mvrs[i] = &mvr{
			x:      x,
			xdir:   Vectors[f][0],
			y:      y,
			ydir:   Vectors[f][1],
			facing: f,
		}
	}

	return &Mover{
		things: th,

		numMvrs: numMvr,
		curMvr:  0,
		mvrs:    mvrs,

		newObj: cnf.NewObj,
		inType: cnf.Type,
	}, nil
}

// cur ...
func (m Mover) cur() *mvr {
	return m.mvrs[m.curMvr]
}

// incr ...
func (m *Mover) incr() {
	if m.numMvrs > 1 {
		m.curMvr++
		if m.curMvr == m.numMvrs {
			m.curMvr = 0
		}
	}
}

// setLast ...
func (m *Mover) setLast() {
	mvr := m.cur()
	m.lastX = mvr.x
	m.lastY = mvr.y
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
