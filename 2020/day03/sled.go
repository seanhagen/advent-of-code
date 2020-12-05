package day3

import (
	"fmt"
)

type Sled struct {
	numTrees       int
	px, py, vx, vy int
	g              *Grid
}

func NewSled(vx, vy int, gr string) (*Sled, error) {
	g, err := NewGrid(gr)
	if err != nil {
		return nil, err
	}

	if vx == 0 && vy == 0 {
		return nil, fmt.Errorf("need either x or y velocity, both can't be zero")
	}

	return &Sled{vx: vx, vy: vy, g: g}, nil
}

// NumTrees ...
func (s *Sled) NumTrees() int {
	return s.numTrees
}

// Run ...
func (s *Sled) Run() error {
	s.numTrees = 0
	s.px = 0
	s.py = 0

	for {
		if s.py >= s.g.Height() {
			break
		}

		if s.g.TreeAt(s.px, s.py) {
			s.numTrees++
		}

		s.px += s.vx
		s.py += s.vy
	}
	return nil
}
