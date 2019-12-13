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

func TestTwoLetterWires(t *testing.T) {
	tests := []struct {
		input      string
		wireChecks []string
	}{
		{"123 -> x", []string{"x"}},
		{"x -> y", []string{"y"}},
		{"x AND y -> z", []string{"x", "y", "z"}},
		{"xy AND yz -> zb", []string{"xy", "yz", "zb"}},
		{"x AND y -> xy", []string{"x", "y", "xy"}},
		{"x OR y -> z", []string{"x", "y", "z"}},
		{"1 OR y -> z", []string{"y", "z"}},
		{"x OR 1 -> z", []string{"x", "z"}},
		{"1 OR 2 -> z", []string{"z"}},
		{"x LSHIFT 2 -> z", []string{"x", "z"}},
		{"x RSHIFT 2 -> z", []string{"x", "z"}},
		{"NOT x -> z", []string{"x", "z"}},
		{"NOT 1 -> z", []string{"z"}},
	}
	for i, x := range tests {
		tt := x
		t.Run(fmt.Sprintf("test %v %v", i, tt.input), func(t *testing.T) {
			brd, _ := NewBoard()
			err := brd.AddWire(tt.input)
			if err != nil {
				t.Fatalf("unable to add wire: %v", err)
			}

			for _, v := range tt.wireChecks {
				if _, ok := brd.wires[v]; !ok {
					t.Errorf("unable to find wire '%v'", v)
				}
			}
		})
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
		{"x AND y -> xy", OpAnd, "xy", true},
		{"1 AND y -> z", OpAnd, "z", true},
		{"x AND 1 -> z", OpAnd, "z", true},
		{"2 AND 1 -> z", OpAnd, "z", true},
		{"x OR y -> z", OpOr, "z", true},
		{"1 OR y -> z", OpOr, "z", true},
		{"x OR 2 -> z", OpOr, "z", true},
		{"1 OR 2 -> z", OpOr, "z", true},
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

			if tt.valid {
				if err != nil {
					t.Fatalf("unable to add wire: %v", err)
				}

				w, ok := brd.wires[tt.wireCheck]
				if !ok {
					t.Fatalf("expected wire '%v', got false for ok", tt.wireCheck)
				}

				if opt := w.input.Type(); opt != tt.exType {
					t.Errorf("wrong type for input operator, expected '%v' got '%v'", tt.exType, opt)
				}
			}

			if !tt.valid && err == nil {
				t.Errorf("shouldn't be valid, error is nil")
			}
		})
	}
}

func TestInputToWire(t *testing.T) {
	tests := []struct {
		input  string
		name   string
		expect uint16
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

func TestDayExample(t *testing.T) {
	tests := []struct {
		wires  []string
		values map[string]uint16
	}{
		{[]string{
			"123 -> x",
			"456 -> y",
			"x AND y -> d",
			"x OR y -> e",
			"x LSHIFT 2 -> f",
			"y RSHIFT 2 -> g",
			"NOT x -> h",
			"NOT y -> i",
			"d AND e -> z",
			"1 AND y -> m",
		}, map[string]uint16{
			"d": 72,
			"e": 507,
			"f": 492,
			"g": 114,
			"h": 65412,
			"i": 65079,
			"x": 123,
			"y": 456,
			"z": 72,
			"m": 0,
		}},
	}

	for i, x := range tests {
		tt := x
		t.Run(fmt.Sprintf("test %v", i), func(t *testing.T) {
			b, _ := NewBoard()
			for _, v := range tt.wires {
				err := b.AddWire(v)
				if err != nil {
					t.Fatalf("unable to add wire '%v', reason: %v", v, err)
				}
			}

			out := b.GetWireValues()

			for k, v := range tt.values {
				x, ok := out[k]
				if !ok {
					t.Errorf("no value for wire '%v'", k)
				}

				if ok && x != v {
					t.Errorf("incorrect value for wire '%v', expected '%v' got '%v'", k, v, x)
				}
			}
		})
	}
}
