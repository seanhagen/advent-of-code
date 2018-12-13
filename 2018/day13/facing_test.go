package day13

import "testing"

func TestStringToFacing(t *testing.T) {
	tests := []struct {
		f Facing
		v string
	}{
		{FacingNorth, "^"},
		{FacingEast, ">"},
		{FacingSouth, "v"},
		{FacingWest, "<"},
	}

	for _, tt := range tests {
		f := StringToFacing(tt.v)
		if tt.f != f {
			t.Errorf("got wrong value for '%v'. expected '%#v', got '%#v'", tt.v, tt.f, f)
		}
	}
}

func TestFacingNext(t *testing.T) {
	tests := []struct {
		f  Facing
		s  string
		in string
	}{
		// no change
		{FacingNorth, "^", "|"},
		{FacingEast, ">", "-"},
		// north -> east
		{FacingEast, "^", "/"},
		// east -> south
		{FacingSouth, ">", "\\"},
		// south -> west
		{FacingWest, "v", "/"},
		// west -> north
		{FacingNorth, "<", "\\"},
	}

	for _, tt := range tests {
		f := StringToFacing(tt.s)
		f.Next(tt.in)
		if f != tt.f {
			t.Errorf("wrong facing. start '%v', input: '%v', now: '%v'", tt.s, tt.in, f)
		}
	}
}
