package day12

import (
	"fmt"
	"testing"
)

func TestStepsUntilReturn(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping long test in short mode")
	}

	tests := []struct {
		input []string
		steps int
	}{
		{[]string{"<x=-1, y=0, z=2>", "<x=2, y=-10, z=-7>", "<x=4, y=-8, z=8>", "<x=3, y=5, z=-1>"}, 2772},
		{[]string{"<x=-8, y=-10, z=0>", "<x=5, y=5, z=10>", "<x=2, y=-7, z=3>", "<x=9, y=-8, z=-3>"}, 4686774924},
	}

	for i, x := range tests {
		tt := x
		t.Run(fmt.Sprintf("test %v", i), func(t *testing.T) {
			sys := NewSystem()
			for _, s := range tt.input {
				err := sys.AddPlanet(s)
				if err != nil {
					t.Fatalf("unable to add planet '%v', error: %v", s, err)
				}
			}

			err := sys.SetupPairs()
			if err != nil {
				t.Fatalf("unable to setup pairs: %v", err)
			}

			steps := sys.DoStepsUntilReturned()

			if steps != tt.steps {
				t.Errorf("wrong number of steps, expected %v got %v", tt.steps, steps)
			}
		})
	}
}
