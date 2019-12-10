package day10

import (
	"fmt"
	"testing"
)

func TestFirstExample(t *testing.T) {
	tests := []struct {
		input string
		posX  float64
		posY  float64
		see   int
	}{
		{
			`.#..#
.....
#####
....#
...##`, 3, 4, 8,
		},
		{`......#.#.
#..#.#....
..#######.
.#.#.###..
.#..#.....
..#....#.#
#..#....#.
.##.#..###
##...#..#.
.#....####`, 5, 8, 33},
		{`#.#...#.#.
.###....#.
.#....#...
##.#.#.#.#
....#.#.#.
.##..###.#
..#...##..
..##....##
......#...
.####.###.`, 1, 2, 35},
		{`.#..#..###
####.###.#
....###.#.
..###.##.#
##.##.#.#.
....###..#
..#.#..#.#
#..#.#.###
.##...##.#
.....#.#..`, 6, 3, 41},
		{`.#..##.###...#######
##.############..##.
.#.######.########.#
.###.#######.####.#.
#####.##.#.##.###.##
..#####..#.#########
####################
#.####....###.#.#.##
##.#################
#####.##.###..####..
..######..##.#######
####.##.####...##..#
.#####..#.######.###
##...#.##########...
#.##########.#######
.####.#.###.###.#.##
....##.##.###..#####
.#.#.###########.###
#.#.#.#####.####.###
###.##.####.##.#..##`, 11, 13, 210},
	}

	for i, tt := range tests {
		x := tt
		t.Run(fmt.Sprintf("test %v", i), func(t *testing.T) {
			sf, err := NewStarField(x.input)
			if err != nil {
				t.Errorf("unable to create starfield: %v", err)
			}

			a := sf.FindBest()
			if a == nil {
				t.Errorf("got nil asteroid from FindBest!")
			}

			if a != nil {
				if a.X != x.posX || a.Y != x.posY {
					t.Errorf("wrong coords for asteroid, expected x: %v, y: %v -- got x: %v, y: %v", x.posX, x.posY, a.X, a.Y)
				}

				if a.See != x.see {
					t.Errorf("wrong count, expected %v got %v", x.see, a.See)
				}
			}

		})
	}
}
