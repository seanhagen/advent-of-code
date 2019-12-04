package day04

import (
	"fmt"
	"testing"
)

func TestMeetsCriterea(t *testing.T) {
	tests := []struct {
		input  int
		expect bool
	}{
		{111111, true},
		{123345, true},
		{223450, false},
		{123789, false},
		{1, false},
		{12, false},
		{123, false},
		{1234, false},
		{12345, false},
		{123456, false},
		{122345, true},
		{1234567, false},
	}

	for _, tt := range tests {
		x := tt
		t.Run(fmt.Sprintf("test-%v", x.input), func(t *testing.T) {
			t.Parallel()
			out := MeetsCriterea(x.input)
			if out != x.expect {
				t.Errorf("Unexpected output, expected '%v', got '%v'", x.expect, out)
			}
		})
	}
}

func TestAdjacentEqual(t *testing.T) {
	tests := []struct {
		input  string
		expect bool
	}{
		{"01", false},
		{"00", true},
		{"91", false},
		{"919", false},
		{"9199", true},
	}

	for _, tt := range tests {
		x := tt
		t.Run(fmt.Sprintf("test-%v", x.input), func(t *testing.T) {
			t.Parallel()
			out := adjacentEqual(x.input)
			if out != x.expect {
				t.Errorf("Unexpected output, expected '%v', got '%v'", x.expect, out)
			}
		})
	}
}

func TestAlwaysIncrease(t *testing.T) {
	tests := []struct {
		input  string
		expect bool
	}{
		{"11", true},
		{"12", true},
		{"19", true},
		{"01", true},
		{"91", false},
	}

	for _, tt := range tests {
		x := tt
		t.Run(fmt.Sprintf("test-%v", x.input), func(t *testing.T) {
			out := alwaysIncrease(x.input)
			if out != x.expect {
				t.Errorf("wrong output, expected '%v' got '%v'", x.expect, out)
			}
		})
	}
}

func TestLenSMG(t *testing.T) {
	tests := []struct {
		input int
		len   int
	}{
		{112233, 2},
		{123444, 3},
		{111122, 2},
	}

	for _, tt := range tests {
		x := tt
		t.Run(fmt.Sprintf("test-%v", x.input), func(t *testing.T) {
			t.Parallel()
			l := lenSmallestMatchingGroup(x.input)
			if l != x.len {
				t.Errorf("wrong output, expected '%v', got '%v'", x.len, l)
			}
		})
	}
}

func TestFindPasswords(t *testing.T) {
	tests := []struct {
		min int
		max int
		num int
	}{
		{111110, 111112, 2},
		{111110, 111111, 1},
		{111110, 111110, 0},
	}

	for _, tt := range tests {
		x := tt
		t.Run(fmt.Sprintf("test-%v-%v", x.min, x.max), func(t *testing.T) {
			t.Parallel()
			out := FindPasswords(x.min, x.max)
			if len(out) != x.num {
				t.Errorf("wrong number of passwords, expected '%v' got '%v'", x.num, out)
			}
		})
	}
}
