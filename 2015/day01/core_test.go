package day01

import (
	"fmt"
	"testing"
)

func TestGetFloor(t *testing.T) {
	tests := []struct {
		input string
		floor int
	}{
		{"(())", 0},
		{"()()", 0},
		{"(((", 3},
		{"(()(()(", 3},
		{"))(((((", 3},
		{"())", -1},
		{"))(", -1},
		{")))", -3},
		{")())())", -3},
	}

	for i, tt := range tests {
		x := tt
		t.Run(fmt.Sprintf("test %v", i), func(t *testing.T) {
			out := GetFloor(x.input)
			if out != x.floor {
				t.Errorf("wrong floor, expected '%v', got '%v'", x.floor, out)
			}
		})
	}
}

func TestPosEnterBasement(t *testing.T) {
	tests := []struct {
		input string
		pos   int
	}{
		{")", 1},
		{"()())", 5},
	}

	for i, tt := range tests {
		x := tt
		t.Run(fmt.Sprintf("test %v", i), func(t *testing.T) {
			out := PosGoIntoBasement(x.input)
			if out != x.pos {
				t.Errorf("wrong position, expected '%v', got '%v'", x.pos, out)
			}
		})
	}
}
