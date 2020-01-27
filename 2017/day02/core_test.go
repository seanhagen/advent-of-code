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

func TestGetDivisible(t *testing.T) {
	tests := []struct {
		input []int
		out   int
	}{
		{[]int{5, 9, 2, 8}, 4},
		{[]int{9, 4, 7, 3}, 3},
		{[]int{3, 8, 6, 5}, 2},
	}

	for i, x := range tests {
		tt := x
		t.Run(fmt.Sprintf("test %v", i), func(t *testing.T) {
			out := getDivisible(tt.input)
			if out != tt.out {
				t.Errorf("wrong output, expected '%v' got '%v'", tt.out, out)
			}
		})
	}
}

func TestGetDivisibleSum(t *testing.T) {
	tests := []struct {
		input string
		out   int
	}{
		{"5\t9\t2\t8\n9\t4\t7\t3\n3\t8\t6\t5", 9},
	}

	for i, x := range tests {
		tt := x
		t.Run(fmt.Sprintf("test %v", i), func(t *testing.T) {
			out := GetDivisibleSum(tt.input)
			if out != tt.out {
				t.Errorf("wrong output, expected '%v' got '%v'", tt.out, out)
			}
		})
	}
}

/*
5 9 2 8
9 4 7 3
3 8 6 5

In the first row, the only two numbers that evenly divide are 8 and 2; the result of this division is 4.
In the second row, the two numbers are 9 and 3; the result is 3.
In the third row, the result is 2.
In this example, the sum of the results would be 4 + 3 + 2 = 9.
*/
