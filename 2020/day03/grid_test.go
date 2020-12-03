package day3

import (
	"fmt"
	"testing"
)

func TestNewGrid(t *testing.T) {
	tests := []struct {
		input  string
		numRow int
		valid  bool
	}{
		{"...", 1, true},
		{"...\n...", 2, true},
		{"...\n....", 2, false},
	}

	for i, tt := range tests {
		x := tt
		t.Run(fmt.Sprintf("test_%v", i), func(t *testing.T) {
			g, err := NewGrid(x.input)
			if !x.valid && err == nil {
				t.Fatalf("shouldn have error, got valid Grid?")
			}

			if x.valid {
				if err != nil {
					t.Fatalf("should have valid Grid, got error: %v", err)
				}
				if g == nil {
					t.Fatalf("should have valid Grid, got null")
				}

				lr := len(g.rows)
				if lr != x.numRow {
					t.Errorf("expected %v rows, got %v", x.numRow, lr)
				}
			}
		})
	}
}

func TestTreeAtNilGrid(t *testing.T) {
	var g *Grid
	b := g.TreeAt(1, 1)
	if b == true {
		t.Errorf("shouldn't have tree for nill Grid")
	}
}

func TestGridTreeAt(t *testing.T) {
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
		input string
		x     int
		y     int
		have  bool
	}{
		{"..\n..", 0, 0, false},
		{"##\n##", 0, 0, true},
		{"##\n##", 0, 2, false},
		{"##\n##", 8, 0, true},
		{"##\n##", 8, 1, true},
		{ex, 0, 0, false},
		{ex, 2, 0, true},
		{ex, 0, 1, true},
		{ex, 10, 3, true},
		{ex, 10, 10, true},
	}

	for i, tt := range tests {
		x := tt
		t.Run(fmt.Sprintf("test_%v", i), func(t *testing.T) {
			g, err := NewGrid(x.input)
			if err != nil {
				t.Fatalf("unable to create Grid: %v", err)
			}

			h := g.TreeAt(x.x, x.y)
			if h != x.have {
				t.Errorf("wrong TreeAt response, expected %v, got %v", x.have, h)
			}
		})
	}
}

func TestGridHeight(t *testing.T) {
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

	h := 11

	g, err := NewGrid(ex)
	if err != nil {
		t.Fatalf("unable to create Grid: %v", err)
	}

	hh := g.Height()
	if h != hh {
		t.Errorf("wrong height, expected %v got %v", h, hh)
	}
}
