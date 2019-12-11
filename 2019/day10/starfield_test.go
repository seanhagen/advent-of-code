package day10

import (
	"fmt"
	"testing"
)

func TestOneToDestroy(t *testing.T) {
	input := `.#.
...`

	sx, sy := 1, 1

	sf, err := NewStarField(input)
	if err != nil {
		t.Fatalf("unable to create starfield: %v", err)
	}

	station, err := NewAsteroid(sx, sy)
	if err != nil {
		t.Fatalf("unable to create station: %v", err)
	}

	sf.SetStation(station)

	err = sf.LaserRotation(1)
	if err != nil {
		t.Fatalf("unable to run laser: %v", err)
	}

	if len(sf.toids) > 0 {
		t.Errorf("only one asteroid, should have been deleted")
	}

	if len(sf.destroyed) == 0 {
		t.Errorf("deleted asteroid should be in slice, isn't")
	}

	if len(sf.destroyed) != 1 {
		t.Errorf("should only be one asteroid, got: %v", len(sf.destroyed))
	}

}

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

func TestLaser(t *testing.T) {
	input := `.#....#####...#..
##...##.#####..##
##...#...#.#####.
..#.........###..
..#.#.....#....##`

	tests := []struct {
		numDestroy     int
		lastDestroyedX int
		lastDestroyedY int
	}{
		{1, 8, 1}, // test for first destroyed
		{2, 9, 0}, // test for second destroyed
		{3, 9, 1}, // etc..
		{4, 10, 0},
		{5, 9, 2},
		{6, 11, 1},
		{7, 12, 1},
		{8, 11, 2},
		{9, 15, 1},
		{10, 12, 2},
		{11, 13, 2},
		{12, 14, 2},
		{13, 15, 2},
		{14, 12, 3},
		{15, 16, 4},
		{16, 15, 4},
		{17, 10, 4},
		{18, 4, 4},
		{19, 2, 4},
		{20, 2, 3},
		{21, 0, 2},
		{22, 1, 2},
		{23, 0, 1},
		{24, 1, 1},
		{25, 5, 2},
		{26, 1, 0},
		{27, 5, 1},
		{28, 6, 1},
		{29, 6, 0},
		{30, 7, 0},
		{31, 8, 0},
		{32, 10, 1},
		{33, 14, 0},
		{34, 16, 1},
		{35, 13, 3},
		{36, 14, 3},
	}

	for i, tt := range tests {
		x := tt

		t.Run(fmt.Sprintf("test %v", i+1), func(t *testing.T) {
			sf, err := NewStarField(input)
			if err != nil {
				t.Fatalf("unable to create starfield: %v", err)
			}

			station, err := NewAsteroid(8, 3)
			if err != nil {
				t.Fatalf("unable to create station: %v", err)
			}

			sf.SetStation(station)

			err = sf.LaserRotation(x.numDestroy)
			if err != nil {
				t.Fatalf("unable to do laser rotation: %v", err)
			}

			if len(sf.destroyed) == 0 {
				t.Errorf("no asteroids in desteroyed list")
			}

			if len(sf.destroyed) != x.numDestroy {
				t.Fatalf("wrong number of asteroids destroyed, expected %v got %v", x.numDestroy, len(sf.destroyed))
			}

			d := sf.GetDestroyed(x.numDestroy)

			should, _ := NewAsteroid(x.lastDestroyedX, x.lastDestroyedY)
			if !should.Equals(d) {
				t.Errorf("wrong asteroid for destroyed %v\n\tgot: %v\n\texpected %v", x.numDestroy, d, should)

				found := false
				for i, a := range sf.toids {
					if a.Equals(should) {
						t.Errorf("found proper asteroid at idx %v in non-destroyed asteroids", i)
						found = true
						break
					}
				}

				if !found {
					for i, a := range sf.destroyed {
						if a.Equals(should) {
							t.Errorf("found proper asteroid at idx %v in destroyed asteroids", i)
							break
						}
					}
				}
			}
		})
	}
}

func TestSetStation(t *testing.T) {
	input := `.#..##.###...#######
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
###.##.####.##.#..##`

	sf, err := NewStarField(input)
	if err != nil {
		t.Errorf("unable to create starfield: %v", err)
	}

	a := sf.FindBest()
	if a == nil {
		t.Errorf("got nil asteroid from FindBest!")
	}

	preLen := len(sf.toids)

	sf.SetStation(a)

	postLen := len(sf.toids)

	if preLen == postLen {
		t.Errorf("setting station didn't remove asteroid, old len: %v, new len: %v", preLen, postLen)
	}

	if postLen != (preLen - 1) {
		t.Errorf("something weird happened, pre len: %v, post len: %v", preLen, postLen)
	}
}

func TestRemoveAll(t *testing.T) {
	input := `.#..##.###...#######
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
###.##.####.##.#..##`

	sf, err := NewStarField(input)
	if err != nil {
		t.Fatalf("unable to create starfield: %v", err)
	}

	a := sf.FindBest()
	if a == nil {
		t.Fatalf("got nil asteroid from FindBest!")
	}

	sf.SetStation(a)

	err = sf.LaserRotation(400)
	if err != nil {
		t.Fatalf("unable to do laser rotation: %v", err)
	}

	d := sf.GetDestroyed(200)
	expect, _ := NewAsteroid(8, 2)
	if !expect.Equals(d) {
		t.Errorf("destroyed in wrong order, expected asteroid 200 to be (%v,%v), got (%v,%v)", expect.X, expect.Y, d.X, d.Y)
	}
}
