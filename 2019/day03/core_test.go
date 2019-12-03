package day03

import (
	"fmt"
	"testing"
)

func TestPathing(t *testing.T) {
	b := Breadboard{}
	b.ReadWire("R8,U5,L5,D3")
	b.ReadWire("U7,R6,D4,L4")

	md := b.ManDist()

	if md != 6 {
		t.Errorf("Wrong value, expected 5, got '%v'", md)
	}
}

func TestManDists(t *testing.T) {
	tests := []struct {
		wireA string
		wireB string
		dist  int
	}{
		{"R75,D30,R83,U83,L12,D49,R71,U7,L72", "U62,R66,U55,R34,D71,R55,D58,R83", 159},
		{"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51", "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7", 135},
	}

	for i, tt := range tests {
		x := tt
		t.Run(fmt.Sprintf("test%v", i), func(t *testing.T) {
			t.Parallel()

			b := Breadboard{}
			b.ReadWire(x.wireA)
			b.ReadWire(x.wireB)
			md := b.ManDist()
			if md != x.dist {
				t.Errorf("wrong manhattan distance, expected '%v', got '%v'", x.dist, md)
			}
		})
	}
}

func TestSteps(t *testing.T) {
	tests := []struct {
		wireA string
		wireB string
		steps int
	}{
		{"R8,U5,L5,D3", "U7,R6,D4,L4", 30},
		{"R75,D30,R83,U83,L12,D49,R71,U7,L72", "U62,R66,U55,R34,D71,R55,D58,R83", 610},
		{"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51", "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7", 410},
	}

	for i, tt := range tests {
		x := tt
		t.Run(fmt.Sprintf("test%v", i), func(t *testing.T) {
			t.Parallel()

			b := Breadboard{}
			b.ReadWire(x.wireA)
			b.ReadWire(x.wireB)
			steps := b.Steps()
			if steps != x.steps {
				t.Errorf("wrong steps, expected '%v', got '%v'", x.steps, steps)
			}
		})
	}
}
