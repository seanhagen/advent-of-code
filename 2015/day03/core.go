package day03

import (
	"github.com/seanhagen/advent-of-code/lib/facing"
)

// House ...
type House struct {
	received int
}

// Santa ...
type Santa struct {
	mover        *facing.Mover
	robot        *facing.Mover
	current      bool
	instructions []facing.Direction
}

func newHouse() interface{} {
	return House{}
}

// NewSanta ...
func NewSanta(input string, num int) (*Santa, error) {
	if num < 1 {
		num = 1
	}

	cnf := &facing.Config{
		Facing:    facing.North,
		NewObj:    newHouse,
		Type:      House{},
		NumMovers: num,
	}

	mover, err := facing.NewMover(cnf)
	if err != nil {
		return nil, err
	}

	mover.SetCurent(House{received: num})

	ins := facing.DirectionSliceFromString(input)

	return &Santa{instructions: ins, mover: mover}, nil
}

// Go ...
func (s *Santa) Go() error {
	for _, i := range s.instructions {
		s.mover.Move(i)
		s.mover.ModifyCurrent(func(i interface{}) interface{} {
			h := i.(House)
			h.received++
			return h
		})
	}
	return nil
}

// Visited ...
func (s *Santa) Visited() int {
	c := 0
	s.mover.Iterate(func(x, y int, i interface{}) {
		c++
	})
	return c
}
