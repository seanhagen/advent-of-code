package day12

import (
	"fmt"
	"testing"
)

func TestCalcVel(t *testing.T) {
	tests := []struct {
		p1  *Planet
		p1s Planet

		p2  *Planet
		p2s Planet
	}{
		// planets both at 0,0,0 shouldn't move
		{&Planet{}, Planet{}, &Planet{}, Planet{}},

		{&Planet{pos: Vec{-1, 0, 2}}, Planet{pos: Vec{0, -1, 3}, vel: Vec{1, -1, 1}}, &Planet{pos: Vec{2, -10, 7}}, Planet{pos: Vec{1, -9, 6}, vel: Vec{-1, 1, -1}}},
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

		{Vec{8, 12, 9}, Vec{7, 3, 0}, 290},
		{Vec{13, 16, 3}, Vec{3, 11, 5}, 608},
		{Vec{29, 11, 1}, Vec{3, 7, 4}, 574},
		{Vec{16, 13, 23}, Vec{7, 1, 1}, 468},
	}

	for i, x := range tests {
		tt := x
		t.Run(fmt.Sprintf("test %v", i), func(t *testing.T) {
			p := Planet{pos: tt.pos, vel: tt.vel}
			if e := p.TotalEnergy(); e != tt.total {
				t.Errorf("wrong total energy, expected %v got %v", tt.total, e)
			}
		})
	}
}

func TestPlanetAtStart(t *testing.T) {
	tests := []struct {
		apos Vec
		bpos Vec

		avel Vec
		bvel Vec

		expect bool
	}{
		{Vec{}, Vec{}, Vec{}, Vec{}, true},
		{Vec{1, 1, 1}, Vec{1, 1, 1}, Vec{}, Vec{}, true},
		{Vec{1, 1, 1}, Vec{}, Vec{1, 1, 1}, Vec{}, false},
	}

	for i, x := range tests {
		tt := x
		t.Run(fmt.Sprintf("test %v", i), func(t *testing.T) {
			p := Planet{pos: tt.apos, initialPos: tt.bpos, vel: tt.avel, initialVel: tt.bvel}

			if got := p.AtStart(); got != tt.expect {
				t.Errorf("wrong output, expected %v got %v", tt.expect, got)
			}
		})
	}
}
