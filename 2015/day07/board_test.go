package day07

import (
	"fmt"
	"testing"
)

func TestCreateBoard(t *testing.T) {
	brd, err := NewBoard()
	if err != nil {
		t.Errorf("error when creating board: %v", err)
	}
	if brd == nil {
		t.Errorf("expected board, got nil")
	}
}

func TestWireParsing(t *testing.T) {
	tests := []struct {
		input     string
		exType    OpType
		wireCheck string
		valid     bool
	}{
		{"123 -> x", OpIn, "x", true},
		{"x -> y", OpIn, "y", true},
		{"x AND y -> z", OpAnd, "z", true},
		{"x OR y -> z", OpOr, "z", true},
		{"x LSHIFT 2 -> z", OpLShift, "z", true},
		{"x RSHIFT 2 -> z", OpRShift, "z", true},
		{"NOT x -> z", OpNot, "z", true},

		{"->", OpNull, "z", false},
		{"-> z", OpNull, "z", false},
		{"AND OR", OpNull, "z", false},
		{"   a ->       b", OpNull, "z", false},
	}

	for i, x := range tests {
		tt := x
		t.Run(fmt.Sprintf("test %v %v", i, tt.input), func(t *testing.T) {
			brd, _ := NewBoard()
			err := brd.AddWire(tt.input)
			if err != nil {
				t.Fatalf("unable to add wire: %v", err)
			}

			w, ok := brd.wires[tt.wireCheck]
			if tt.valid && !ok {
				t.Errorf("expected wire '%v', got false for ok", tt.wireCheck)
			}

			if tt.valid && ok {
				if opt := w.input.Type(); opt != tt.exType {
					t.Errorf("wrong type for input operator, expected '%v' got '%v'", tt.exType, opt)
				}
			}

			if !tt.valid && ok {
				t.Errorf("shouldn't be valid, got wire: %#v", w)
			}
		})
	}
}

func TestInputToWire(t *testing.T) {
	tests := []struct {
		input  string
		name   string
		expect int16
	}{
		{"123 -> x", "x", 123},
		{"456 -> y", "y", 456},
	}

	for i, x := range tests {
		tt := x
		t.Run(fmt.Sprintf("test %v", i), func(t *testing.T) {
			b, _ := NewBoard()
			err := b.AddWire(tt.input)
			if err != nil {
				t.Fatalf("unable to add wire: %v", err)
			}

			out := b.GetWireValues()
			if len(out) == 0 {
				t.Fatalf("got 0-length map")
			}

			v, ok := out[tt.name]
			if !ok {
				t.Errorf("no value for wire %v", tt.name)
			} else {
				if v != tt.expect {
					t.Errorf("output not correct, expected %v got %v", tt.expect, out)
				}
			}
		})
	}
}
