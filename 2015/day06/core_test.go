package day06

import (
	"fmt"
	"testing"
)

func sliceEq(a, b []int) bool {
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

func TestInsToCoords(t *testing.T) {
	tests := []struct {
		input string
		out   []int
	}{
		{"turn on 0,0 through 999,999", []int{0, 0, 999, 999}},
		{"toggle 0,0 through 999,0", []int{0, 0, 999, 0}},
		{"turn off 499,499 through 500,500", []int{499, 499, 500, 500}},
		{"turn off 99499,499 through -1,500", []int{9999, 499, 0, 500}},
	}

	for _, tt := range tests {
		x := tt
		t.Run(fmt.Sprintf("test %v", x.input), func(t *testing.T) {
			out := insToCoords(x.input)
			if len(out) != 4 {
				t.Fatalf("something's gone horribly wrong, got %v elements ( %#v )", len(out), out)
			}

			if !sliceEq(out, x.out) {
				t.Errorf("wrong output, expected %#v, got %#v", x.out, out)
			}
		})
	}
}

func TestNumLit(t *testing.T) {
	tests := []struct {
		input   string
		num     int
		initial int
	}{
		{"turn on 0,0 through 999,999", 1000 * 1000, 0},
		{"toggle 0,0 through 999,0", 1000, 0},
		{"turn off 499,499 through 500,500", (1000 * 1000) - 4, 1},

		{"turn on 0,0 through 999,999", 1000 * 1000, 1},
		{"toggle 0,0 through 999,0", (1000 * 1000) - 1000, 1},
		{"toggle 499,499 through 500,500", 4, 0},
	}

	for _, tt := range tests {
		x := tt
		t.Run(fmt.Sprintf("test initial %v %v", x.initial, x.input), func(t *testing.T) {
			lg := LightGrid{}
			lg.SetAll(x.initial)

			err := lg.ProcessInstruction(x.input)
			if err != nil {
				t.Fatalf("unable to proccess instruction: %v", err)
			}

			num := lg.CountOn()
			if num != x.num {
				t.Errorf("wrong count, expected '%v' got '%v'", x.num, num)
			}
		})
	}
}

func TestBrighness(t *testing.T) {
	/*
	   - turn on 0,0 through 0,0 would increase the total brightness by 1

	   - toggle 0,0 through 999,999 would increase the total brightness by 2000000
	*/

	tests := []struct {
		input   string
		num     int
		initial int
	}{
		{"turn on 0,0 through 0,0", 1, 0},
		{"turn on 0,0 through 0,0", 1000*1000 + 1, 1},
		{"toggle 0,0 through 999,999", 2 * 1000 * 1000, 0},
		{"toggle 0,0 through 999,999", 4 * 1000 * 1000, 2},
		{"turn off 0,0 through 999,999", 1000 * 1000, 2},
		{"turn off 0,0 through 999,999", 0, 0},
	}

	for _, tt := range tests {
		x := tt
		t.Run(fmt.Sprintf("test initial %v %v", x.initial, x.input), func(t *testing.T) {
			lg := LightGrid{}
			lg.SetAll(x.initial)

			err := lg.ProcessInstructionsV2(x.input)
			if err != nil {
				t.Fatalf("unable to proccess instruction: %v", err)
			}

			num := lg.TotalBrightness()
			if num != x.num {
				t.Errorf("wrong count, expected '%v' got '%v'", x.num, num)
			}
		})
	}
}
