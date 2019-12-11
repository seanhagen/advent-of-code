package day03

import (
	"fmt"
	"testing"
)

func TestDeliveryCounts(t *testing.T) {
	// - > delivers presents to 2 houses: one at the starting location, and one to the
	// east.

	//   - ^>v< delivers presents to 4 houses in a square, including twice to the house
	// at his starting/ending location.

	//   - ^v^v^v^v^v delivers a bunch of presents to some very lucky children at only 2
	// houses.

	tests := []struct {
		input  string
		expect int
	}{
		{">", 2},
		{"^>v<", 4},
		{"^v^v^v^v^v", 2},
	}

	for _, tt := range tests {
		x := tt
		t.Run(fmt.Sprintf("test %v", x.input), func(t *testing.T) {
			sta, err := NewSanta(x.input, 1)
			if err != nil {
				t.Fatalf("unable to create santa: %v", err)
			}

			err = sta.Go()
			if err != nil {
				t.Fatalf("unable to deliver presents: %v", err)
			}

			d := sta.Visited()
			if d != x.expect {
				t.Errorf("wrong number of houses visited, expected %v got %v", x.expect, d)
			}
		})
	}
}

func TestDeliveryCountWithRobot(t *testing.T) {
	/*
	   ^v delivers presents to 3 houses, because Santa goes north, and then Robo-Santa goes south.

	   ^>v< now delivers presents to 3 houses, and Santa and Robo-Santa end up back where they started.

	   ^v^v^v^v^v now delivers presents to 11 houses, with Santa going one direction and Robo-Santa going the other.
	*/
	tests := []struct {
		input  string
		expect int
	}{
		{"^v", 3},
		{"^>v<", 3},
		{"^v^v^v^v^v", 11},
	}

	for _, tt := range tests {
		x := tt
		t.Run(fmt.Sprintf("test %v", x.input), func(t *testing.T) {
			sta, err := NewSanta(x.input, 2)
			if err != nil {
				t.Fatalf("unable to create santa: %v", err)
			}

			err = sta.Go()
			if err != nil {
				t.Fatalf("unable to deliver presents: %v", err)
			}

			d := sta.Visited()
			if d != x.expect {
				t.Errorf("wrong number of houses visited, expected %v got %v", x.expect, d)
			}
		})
	}

}
