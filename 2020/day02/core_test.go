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
