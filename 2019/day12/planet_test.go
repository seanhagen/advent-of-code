package day12

import (
	"fmt"
	"testing"
)

func TestCalcVel(t *testing.T) {
	tests := []struct {
		p1  *Planet
		p1s Vec

		p2  *Planet
		p2s Vec
	}{
		// planets both at 0,0,0 shouldn't move
		{&Planet{}, Vec{}, &Planet{}, Vec{}},

		{&Planet{pos: Vec{-1, 0, 2}}, Vec{0, -1, 3}, &Planet{pos: Vec{2, -10, 7}}, Vec{1, -9, 6}},
	}

	for i, x := range tests {
		tt := x
		t.Run(fmt.Sprintf("test %v", i), func(t *testing.T) {
			CalculateVelocity(tt.p1, tt.p2)
			tt.p1.Step()
			tt.p2.Step()

			if !tt.p1.Eq(tt.p1s) {
				t.Errorf("planet 1 not at right position, expected %#v, got %#v", tt.p1s, tt.p1.pos)
			}

			if !tt.p2.Eq(tt.p2s) {
				t.Errorf("planet 2 not at right position, expected %#v, got %#v", tt.p2s, tt.p2.pos)
			}

		})
	}
}

func TestPlanetTotalEnergy(t *testing.T) {
	tests := []struct {
		pos   Vec
		vel   Vec
		total int
	}{
		{Vec{}, Vec{}, 0},
		{Vec{1, 1, 1}, Vec{1, 1, 1}, 9},
		{Vec{2, 1, 3}, Vec{3, 2, 1}, 36},
		{Vec{1, 8, 0}, Vec{1, 1, 3}, 45},
		{Vec{3, 6, 1}, Vec{3, 2, 3}, 80},
		{Vec{2, 0, 4}, Vec{1, 1, 1}, 18},
	}

	for i, x := range tests {
		tt := x
		t.Run(fmt.Sprintf("test %v", i), func(t *testing.T) {
			p := Planet{"", tt.pos, tt.vel}
			if e := p.TotalEnergy(); e != tt.total {
				t.Errorf("wrong total energy, expected %v got %v", tt.total, e)
			}
		})
	}
}
