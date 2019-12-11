package facing

import (
	"fmt"
	"testing"
)

func TestDirFromString(t *testing.T) {
	tests := []struct {
		input string
		out   Direction
	}{
		{"v", South},
		{"V", South},
		{"<", West},
		{">", East},
		{"^", North},
		{"P", North},
		{"what", North},
	}

	for _, tt := range tests {
		x := tt
		t.Run(fmt.Sprintf("test %v", x.input), func(t *testing.T) {
			t.Parallel()
			out := DirectionFromString(x.input)
			if out != x.out {
				t.Errorf("wrong output, expected '%v' got '%v'", string(x.out), string(out))
			}
		})
	}
}

func TestDirSlicFromString(t *testing.T) {
	tests := []struct {
		input string
		out   []Direction
	}{
		{"^", []Direction{North}},
		{"^>v<", []Direction{North, East, South, West}},
		{"^v^v^v", []Direction{North, South, North, South, North, South}},
		{"what", []Direction{North, North, North, North}},
	}

	cmp := func(a, b []Direction) bool {
		if len(a) != len(b) {
			return false
		}
		for i := range a {
			if a[i] != b[i] {
				return false
			}
		}
		return true
	}

	for _, tt := range tests {
		x := tt
		t.Run(fmt.Sprintf("test %v", x.input), func(t *testing.T) {
			t.Parallel()
			out := DirectionSliceFromString(x.input)
			if !cmp(x.out, out) {
				t.Errorf("wrong output, expected %#v got %#v", x.out, out)
			}
		})
	}
}
