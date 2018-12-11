package day11

import (
	"testing"
)

func TestLargestSubArray(t *testing.T) {
	tests := []struct {
		serial int
		x      int
		y      int
		total  int
	}{
		{18, 33, 45, 29},
		{42, 21, 61, 30},
	}

	for _, tt := range tests {
		g := SetupGrid(tt.serial)
		var x, y, power int
		x, y, power = g.FindLargest(3)

		if x != tt.x {
			t.Errorf("wrong x coord: expected %v got %v", tt.x, x)
		}

		if y != tt.y {
			t.Errorf("wrong y coord: expected %v got %v", tt.y, y)
		}

		if power != tt.total {
			t.Errorf("wrong power level: expected %v got %v", tt.total, power)
		}
	}
}
