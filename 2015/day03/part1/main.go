package main

import (
	"fmt"
	"os"

	"github.com/seanhagen/advent-of-code/2015/day03"
	"github.com/seanhagen/advent-of-code/lib"
)

/*

--- Day 3: Perfectly Spherical Houses in a Vacuum ---
Santa is delivering presents to an infinite two-dimensional grid of houses.

He begins by delivering a present to the house at his starting location, and
then an elf at the North Pole calls him via radio and tells him where to move
next. Moves are always exactly one house to the north (^), south (v), east (>),
or west (<). After each move, he delivers another present to the house at his
new location.

However, the elf back at the north pole has had a little too much eggnog, and so
his directions are a little off, and Santa ends up visiting some houses more
than once. How many houses receive at least one present?

For example:

 - > delivers presents to 2 houses: one at the starting location, and one to the
 east.

 - ^>v< delivers presents to 4 houses in a square, including twice to the house
 at his starting/ending location.

 - ^v^v^v^v^v delivers a bunch of presents to some very lucky children at only 2
 houses.


*/

func main() {
	ins, err := lib.GetString("../input.txt")
	if err != nil {
		fmt.Printf("unable to load instructions: %v\n", err)
		os.Exit(1)
	}

	sta, err := day03.NewSanta(ins, 1)
	if err != nil {
		fmt.Printf("unable to create Santa: %v\n", err)
		os.Exit(1)
	}

	err = sta.Go()
	if err != nil {
		fmt.Printf("error delivering presents: %v\n", err)
	}

	v := sta.Visited()
	fmt.Printf("santa visited %v houses\n", v)
}
