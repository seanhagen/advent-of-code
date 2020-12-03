package day3

import (
	"fmt"
	"testing"
)

func TestTreeAt(t *testing.T) {
	tests := []struct {
		in string //input
		c  int    //column to check
		t  bool   //has a tree in column to check
	}{
		{"..", 1, false},
		{"##", 1, true},
		{"..", 3, false},
		{"##", 3, true},
		{".#", 3, false},
		{".#", 4, true},
		{"...#", 1, false},
		{"...#", 4, true},
		{"...#", 5, false},
		{"...#", 8, true},
	}

	for i, tt := range tests {
		x := tt
		t.Run(fmt.Sprintf("test_%v", i), func(t *testing.T) {
			r := NewRow(x.in)
			h := r.TreeAt(x.c)
			if h != x.t {
				t.Errorf("in Row(%v) at position %v, expected %v, got %v", x.in, x.c, x.t, h)
			}
		})
	}
}
