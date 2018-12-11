package day11

import "testing"

func TestFuelLevels(t *testing.T) {
	tests := []struct {
		x           int
		y           int
		serial      int
		powerShould int
	}{
		{122, 79, 57, -5},
		{217, 196, 39, 0},
		{101, 153, 71, 4},
	}

	for _, tt := range tests {
		fc := CreateFuelCell(tt.x, tt.y, tt.serial)

		if fc.power != tt.powerShould {
			t.Errorf("Wrong power level, expected %v got %v", tt.powerShould, fc.power)
		}
	}
}
