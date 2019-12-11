package facing

import (
	"fmt"
	"testing"
)

func TestTypeCheck(t *testing.T) {
	tests := []struct {
		a    interface{}
		b    NewObjFn
		good bool
	}{
		{"", func() interface{} { return "" }, true},
		{"", func() interface{} { return 1 }, false},
		{Mover{}, func() interface{} { return Mover{} }, true},
		{Mover{}, func() interface{} { return "" }, false},
	}

	for i, tt := range tests {
		x := tt
		t.Run(fmt.Sprintf("test %v", i), func(t *testing.T) {
			v := x.b()
			eq := typeEqual(x.a, v)
			if eq != x.good {
				t.Errorf("expected value (%#v) type %T same as (%#v) type %T to return '%v', got '%v'", x.a, x.a, v, v, x.good, eq)
			}
		})
	}
}

func TestMove(t *testing.T) {
	tests := []struct {
		input  string
		endX   int
		endY   int
		facing Direction
	}{
		{">", 1, 0, East},
		{"^>v<", 0, 0, West},
		{"^v^v^v^v^v", 0, 0, South},
	}

	fn := func() interface{} { return "" }
	cnf := Config{X: 0, Y: 0, Facing: North, NewObj: fn, Type: ""}

	for _, tt := range tests {
		z := tt
		t.Run(fmt.Sprintf("test %v", z.input), func(t *testing.T) {
			ins := DirectionSliceFromString(z.input)
			mv, err := NewMover(&cnf)
			if err != nil {
				t.Fatalf("unable to create mover: %v", err)
			}
			for _, i := range ins {
				mv.Move(i)
			}

			x, y := mv.Location()
			f := mv.Facing()

			if x != z.endX || y != z.endY {
				t.Errorf("mover not at correct coordinates, expected %v,%v -- got %v,%v", x, y, z.endX, z.endY)
			}

			if f != z.facing {
				t.Errorf("mover not facing correct direction, expected '%v' got '%v'", string(z.facing), string(f))
			}
		})
	}
}
