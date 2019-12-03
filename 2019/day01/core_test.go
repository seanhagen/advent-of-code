package day1

import "testing"

func TestCore(t *testing.T) {
	data := map[int]int{
		12:     2,
		14:     2,
		1969:   654,
		100756: 33583,
	}

	for in, shouldOut := range data {
		out := CalcFuel(in)
		if out != shouldOut {
			t.Errorf("unexpected output '%v', expected '%v'", out, shouldOut)
		}
	}
}

func TestCalcFuelRecurse(t *testing.T) {
	data := map[int]int{
		14:     2,
		1969:   966,
		100756: 50346,
	}

	for in, shouldOut := range data {
		out := CalcFuelRecurse(in)
		if out != shouldOut {
			t.Errorf("unexpected output '%v', expected '%v'", out, shouldOut)
		}
	}
}
