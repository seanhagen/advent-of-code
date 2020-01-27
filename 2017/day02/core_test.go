package day02

import (
	"fmt"
	"testing"
)

func TestGetRowDiff(t *testing.T) {
	tests := []struct {
		input []int
		out   int
	}{
		{[]int{5, 1, 9, 5}, 8},
		{[]int{7, 5, 3}, 4},
		{[]int{2, 4, 6, 8}, 6},
	}

	for i, x := range tests {
		tt := x
		t.Run(fmt.Sprintf("test %v", i), func(t *testing.T) {
			out := getRowDiff(tt.input)
			if out != tt.out {
				t.Errorf("wrong output, expected '%v', got '%v'", tt.out, out)
			}
		})
	}
}

func TestGetChecksum(t *testing.T) {
	tests := []struct {
		input string
		out   int
	}{
		{"5\t1\t9\t5\n7\t5\t3\n2\t4\t6\t8", 18},
	}

	for i, x := range tests {
		tt := x
		t.Run(fmt.Sprintf("test %v", i), func(t *testing.T) {
			out := GetChecksum(tt.input)
			if out != tt.out {
				t.Errorf("wrong output, expected '%v' got '%v'", tt.out, out)
			}
		})
	}
}
