package day5

import (
	"fmt"
	"sort"
	"testing"
)

func TestChooseRow(t *testing.T) {
	ta := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	ta2 := []int{}
	for i := 0; i <= 127; i++ {
		ta2 = append(ta2, i)
	}

	tests := []struct {
		start []int
		end   []int
		path  string
	}{
		{ta, []int{1, 2, 3, 4, 5}, "F"},
		// 1,2,3,4,5,6,7,8,9,10 -> F: 1,2,3,4,5 B: 6,7,8,9,10 [F]
		// 1,2,3,4,5 -> F: 1,2  B: 3,4,5 [F]
		// 1,2 -> F: 1 B: 2 [F]
		// 1
		{ta, []int{1}, "FFF"},
		{ta, []int{6, 7, 8, 9, 10}, "B"},
		// 1,2,3,4,5,6,7,8,9,10 -> F: 1,2,3,4,5 B: 6,7,8,9,10 [B]
		// 6,7,8,9,10 -> F: 6,7 B: 8,9,10 [B]
		// 8,9,10 -> F: 8 B: 9,10 [B]
		// 9,10 -> F: 9 B: 10 [B]
		// 9
		{ta, []int{10}, "BBBB"},
		// 1,2,3,4,5,6,7,8,9,10 -> F: 1,2,3,4,5 B: 6,7,8,9,10 [B]
		// 6, 7, 8, 9, 10 -> F: 6,7 B: 8,9,10 [B]
		// 8,9, 10 -> F: 8 B: 9,10 [F]
		// 9, 10
		{ta, []int{8}, "BBF"},
		{ta, []int{9}, "BBBF"},
		// 1,2,3,4,5,6,7,8,9,10 -> F: 1,2,3,4,5 B: 6,7,8,9,10 [F]
		// 1,2,3,4,5 -> F: 1,2 B: 3, 4,5 [F]
		// 1,2 -> F: 1 B: 2 [B]
		// 2
		{ta, []int{2}, "FFB"},
		{ta, ta, ""},
		{[]int{1, 2}, []int{1}, "F"},
		{[]int{1, 2}, []int{2}, "B"},
		{ta2, []int{44}, "FBFBBFF"},
		{ta2, []int{70}, "BFFFBBF"},
		{ta2, []int{14}, "FFFBBBF"},
		{ta2, []int{102}, "BBFFBBF"},
	}

	for i, x := range tests {
		tt := x
		t.Run(fmt.Sprintf("test_%v", i), func(t *testing.T) {
			out := choose(tt.start, tt.path)
			if !eqa(out, tt.end) {
				t.Errorf("wrong output, expected '%v', got '%v'", tt.end, out)
			}
		})
	}
}

func TestChooseSeat(t *testing.T) {
	ta := []int{0, 1, 2, 3, 4, 5, 6, 7}
	tests := []struct {
		start []int
		end   []int
		path  string
	}{
		{ta, []int{5}, "RLR"},
		{ta, []int{7}, "RRR"},
		{ta, []int{4}, "RLL"},
	}

	for i, x := range tests {
		tt := x
		t.Run(fmt.Sprintf("test_%v", i), func(t *testing.T) {
			out := choose(tt.start, tt.path)
			if !eqa(out, tt.end) {
				t.Errorf("wrong output, expected '%v', got '%v'", tt.end, out)
			}
		})
	}
}

func TestGetRowSeat(t *testing.T) {
	tests := []struct {
		path     string
		row, col int
	}{
		{"BFFFBBFRRR", 70, 7},
		{"FFFBBBFRRR", 14, 7},
		{"BBFFBBFRLL", 102, 4},
	}

	for i, x := range tests {
		tt := x
		t.Run(fmt.Sprintf("test_%v", i), func(t *testing.T) {
			r, c := GetRowSeat(tt.path)
			if r != tt.row {
				t.Errorf("wrong row, expected %v got %v", tt.row, r)
			}
			if c != tt.col {
				t.Errorf("wrong col, expected %v got %v", tt.col, c)
			}
		})
	}
}

func TestGetSeatID(t *testing.T) {
	tests := []struct {
		row, col, id int
	}{
		{70, 7, 567},
		{14, 7, 119},
		{102, 4, 820},
	}

	for i, x := range tests {
		tt := x
		t.Run(fmt.Sprintf("test_%v", i), func(t *testing.T) {
			id := GetSeatID(tt.row, tt.col)
			if id != tt.id {
				t.Errorf("wrong seat id, expected %v got %v", tt.id, id)
			}
		})
	}
}

func eqa(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	sort.Ints(a)
	sort.Ints(b)

	for i, aa := range a {
		if aa != b[i] {
			return false
		}
	}
	return true
}
