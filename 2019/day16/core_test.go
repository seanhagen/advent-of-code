package day16

import (
	"fmt"
	"testing"
)

func TestFinalProcessing(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		{"03036732577212944063491565474664", "84462026"},
		{"02935109699940807407585447034323", "78725270"},
		{"03081770884921959731165446850517", "53553731"},
	}

	for i, x := range tests {
		tt := x
		t.Run(fmt.Sprintf("test %v", i+1), func(t *testing.T) {
			out := FinalProcessing(tt.input)
			if out != tt.output {
				t.Errorf("wrong output, expected %#v got %#v", tt.output, out)
			}
		})
	}
}

func TestFirstEight(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		{"80871224585914546619083218645595", "24176176"},
		{"19617804207202209144916044189917", "73745418"},
		{"69317163492948606335995924319873", "52432133"},
	}

	for i, x := range tests {
		tt := x
		t.Run(fmt.Sprintf("test %v", i+1), func(t *testing.T) {
			t.Parallel()
			out := FirstEight(tt.input, 100)
			if out != tt.output {
				t.Errorf("wrong output, expected %#v got %#v", tt.output, out)
			}
		})
	}
}

func TestApplyPhases(t *testing.T) {
	tests := []struct {
		input  string
		phases int
		output string
	}{
		{"12345678", 1, "48226158"},
		{"12345678", 2, "34040438"},
		{"12345678", 3, "03415518"},
		{"12345678", 4, "01029498"},
	}

	for i, x := range tests {
		tt := x
		t.Run(fmt.Sprintf("test %v", i), func(t *testing.T) {
			t.Parallel()
			out := ApplyPhases(tt.input, tt.phases)
			if out != tt.output {
				t.Errorf("wrong output for %v phases, expected %#v got %#v", tt.phases, tt.output, out)
			}
		})
	}
}

func TestModify(t *testing.T) {
	tests := []struct {
		input   string
		pattern []int
		expect  int
	}{
		{"98765", []int{1, 2, 3}, 2},
		{"12345678", []int{1, 0, -1, 0, 1, 0, -1, 0}, 4},
		{"12345678", []int{0, 1, 1, 0, 0, -1, -1, 0}, 8},
	}

	for i, x := range tests {
		tt := x
		t.Run(fmt.Sprintf("test %v", i), func(t *testing.T) {
			t.Parallel()
			out := modify(tt.input, tt.pattern)
			if out != tt.expect {
				t.Errorf("wrong output, expected %v got %v", tt.expect, out)
			}
		})
	}
}

func TestGetPattern(t *testing.T) {
	tests := []struct {
		length int
		itr    int
		out    []int
	}{
		{8, 1, []int{1, 0, -1, 0, 1, 0, -1, 0}},
		{8, 2, []int{0, 1, 1, 0, 0, -1, -1, 0}},
		{8, 3, []int{0, 0, 1, 1, 1, 0, 0, 0}},
		{8, 4, []int{0, 0, 0, 1, 1, 1, 1, 0}},
		{8, 5, []int{0, 0, 0, 0, 1, 1, 1, 1}},
		{8, 6, []int{0, 0, 0, 0, 0, 1, 1, 1}},
		{8, 7, []int{0, 0, 0, 0, 0, 0, 1, 1}},
		{8, 8, []int{0, 0, 0, 0, 0, 0, 0, 1}},
	}

	for _, x := range tests {
		tt := x
		t.Run(fmt.Sprintf("test %v", tt.itr), func(t *testing.T) {
			t.Parallel()
			out := getPattern(tt.length, tt.itr)
			if !sliceEq(tt.out, out) {
				t.Errorf("wrong output\n\texpected %#v\n\tgot:     %#v", tt.out, out)
			}
		})
	}
}

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
