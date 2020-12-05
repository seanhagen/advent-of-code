package day3

import (
	"fmt"
	"testing"
)

func TestNewSled(t *testing.T) {
	ex := `..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`

	tests := []struct {
		input  string
		vx, vy int
		valid  bool
	}{
		{ex, 0, 0, false},
		{"", 1, 1, false},
		{ex, 1, 1, true},
	}

	for i, x := range tests {
		tt := x
		t.Run(fmt.Sprintf("test_%v", i), func(t *testing.T) {
			s, err := NewSled(tt.vx, tt.vy, tt.input)

			if tt.valid {
				if err == nil && s == nil {
					t.Fatalf("should be valid but got no sled")
				}
				if err != nil {
					t.Fatalf("should be valid but got error: %v", err)
				}
			}

			if !tt.valid {
				if err == nil {
					t.Fatalf("should be invalid but got no error?")
				}
				if s != nil {
					t.Fatalf("should be invalid but got sled?")
				}
			}
		})
	}
}

func TestRunSled(t *testing.T) {
	ex := `..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`

	tests := []struct {
		input  string
		vx, vy int
		num    int
		valid  bool
	}{
		{ex, 3, 1, 7, true},
		{ex, 0, 1, 3, true},
	}

	for i, x := range tests {
		tt := x
		t.Run(fmt.Sprintf("test_%v", i), func(t *testing.T) {
			s, err := NewSled(tt.vx, tt.vy, tt.input)
			if err != nil {
				t.Fatalf("unable to create Sled: %v", err)
			}
			if s == nil {
				t.Fatalf("no error but no sled?")
			}

			err = s.Run()

			if tt.valid {
				if err != nil {
					t.Fatalf("error while running sled: %v", err)
				}
				nt := s.NumTrees()
				if nt != tt.num {
					t.Errorf("wrong output, expected %v got %v", tt.num, nt)
				}
			} else {
				if err == nil {
					t.Errorf("expected error, got none")
				}
			}
		})
	}
}
