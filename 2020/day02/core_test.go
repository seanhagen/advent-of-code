package day2

import (
	"fmt"
	"testing"
)

func TestValidatePassword(t *testing.T) {
	tests := []struct {
		min   int
		max   int
		l     string
		pw    string
		valid bool
	}{
		{1, 3, "a", "abcde", true},
		{1, 3, "b", "cdefg", false},
		{2, 9, "c", "ccccccccc", true},
		{3, 5, "d", "abdedfd", true},
		{1, 2, "b", "bbbbbbb", false},
	}

	for i, tt := range tests {
		x := tt
		t.Run(fmt.Sprintf("test_%v", i), func(t *testing.T) {
			t.Parallel()
			v := ValidatePassword(x.min, x.max, x.l, x.pw)
			if x.valid != v {
				t.Errorf("invalid output, expected '%v', got '%v'", x.valid, v)
			}
		})
	}
}

func TestNewValidPass(t *testing.T) {
	tests := []struct {
		min   int
		max   int
		l     string
		pw    string
		valid bool
	}{
		//1-3 a: abcde is valid: position 1 contains a and position 3 does not.
		{1, 3, "a", "abcde", true},
		//1-3 b: cdefg is invalid: neither position 1 nor position 3 contains b.
		{1, 3, "b", "cdefg", false},
		//2-9 c: ccccccccc is invalid: both position 2 and position 9 contain c.
		{2, 9, "c", "ccccccccc", false},
		// made up examples
		{3, 5, "d", "abdeffd", true},
		{1, 2, "b", "bbbbbbb", false},

		// testing input from actual file
		{14, 15, "r", "rkrbrvrrrgrczrz", true},
		{12, 13, "t", "trtjtttlnxnxx", false},
		{5, 8, "l", "qnwllfsl", false},
		{2, 15, "g", "xgtcjftlqqfwkggpf", false},
		{11, 16, "j", "jjjjjjjjjljjjjjjj", false},
		{8, 9, "b", "bbzbkbbvgcbb", false},
		{5, 6, "r", "dvkxrrsvrrksszsdr", false},
		{12, 13, "j", "jjjjjjjjjjjhdjjj", false},
	}

	for i, tt := range tests {
		x := tt
		t.Run(fmt.Sprintf("test_%v", i), func(t *testing.T) {
			// t.Parallel()
			v := NewValidPass(x.min, x.max, x.l, x.pw)
			if x.valid != v {
				t.Errorf("invalid output, expected '%v', got '%v'", x.valid, v)
			}
		})
	}
}
