package day12

import (
	"fmt"
	"testing"
)

func TestVecAdd(t *testing.T) {
	tests := []struct {
		x1 int
		xd int
		x2 int

		y1 int
		yd int
		y2 int

		z1 int
		zd int
		z2 int
	}{
		{0, 1, 1, 0, 1, 1, 0, 1, 1},
		{1, 1, 2, 1, 1, 2, 1, 1, 2},
		{1, -1, 0, 1, -1, 0, 1, -1, 0},
		{1, -2, -1, 2, 0, 2, 3, 3, 6}, // example from day 12 for Europa
	}

	for i, x := range tests {
		tt := x
		t.Run(fmt.Sprintf("test %v", i), func(t *testing.T) {
			a := Vec{tt.x1, tt.y1, tt.z1}
			b := Vec{tt.xd, tt.yd, tt.zd}
			c := Vec{tt.x2, tt.y2, tt.z2}

			a.Add(b)

			if !a.Eq(c) {
				t.Errorf("vec position wrong, expected %#v, got %#v ( after adding %#v )", c, a, b)
			}
		})
	}
}

func TestParseString(t *testing.T) {
	tests := []struct {
		input string
		valid bool
		out   Vec
	}{
		{"<x=-1, y=0, z=2>", true, Vec{-1, 0, 2}},
		{"x=-1, y=0, z=2", true, Vec{-1, 0, 2}},
		{"<x=-1,y=0,z=2>", true, Vec{-1, 0, 2}},
		{"z=2, x=3, y=-3", true, Vec{3, -3, 2}},
		{"<x=-1, y=0, z=2>", true, Vec{-1, 0, 2}},
		{"<x=2, y=-10, z=-7>", true, Vec{2, -10, -7}},
		{"<x=4, y=-8, z=8>", true, Vec{4, -8, 8}},
		{"<x=3, y=5, z=-1>", true, Vec{3, 5, -1}},

		{"z=2, z=2, z=3", false, Vec{}},
		{"a=2, z=3, y=-3", false, Vec{}},
		{"z=2, a=3, y=-3", false, Vec{}},
		{"z=2, x=3, a=-3", false, Vec{}},
		{"z=2, x=3, y-3", false, Vec{}},
		{"z=2, x=3", false, Vec{}},
		{"z=2", false, Vec{}},
		{"what", false, Vec{}},
	}

	for i, x := range tests {
		tt := x
		t.Run(fmt.Sprintf("test %v %v", i, tt.input), func(t *testing.T) {
			v, err := ParseVecString(tt.input)

			if tt.valid {
				if err != nil {
					t.Errorf("expected valid, got error '%v'", err)
				}

				if !v.Eq(tt.out) {
					t.Errorf("invalid output, expected %#v got %#v", tt.out, v)
				}
			} else {
				if err == nil {
					t.Errorf("expected error, got nil")
				}
			}
		})
	}
}

func TestVecEnergy(t *testing.T) {
	tests := []struct {
		x, y, z, total int
	}{
		{2, 1, 3, 6},
		{-2, 1, -3, 6},
		{1, 1, 1, 3},
		{-1, -1, -1, 3},
		{2, 0, 4, 6},
	}

	for i, x := range tests {
		tt := x
		t.Run(fmt.Sprintf("test %v", i), func(t *testing.T) {
			v := Vec{tt.x, tt.y, tt.z}
			if e := v.Energy(); e != tt.total {
				t.Errorf("wrong energy, expected %v got %v", tt.total, e)
			}
		})
	}
}
