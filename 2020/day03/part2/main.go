package main

import (
	"fmt"
	"os"

	day3 "github.com/seanhagen/advent-of-code/2020/day03"
	"github.com/seanhagen/advent-of-code/lib"
)

/*
Time to check the rest of the slopes - you need to minimize the probability of a sudden arboreal
stop, after all.

Determine the number of trees you would encounter if, for each of the following slopes, you start
at the top-left corner and traverse the map all the way to the bottom:

Right 1, down 1.
Right 3, down 1. (This is the slope you already checked.)
Right 5, down 1.
Right 7, down 1.
Right 1, down 2.

In the above example, these slopes would find 2, 7, 3, 4, and 2 tree(s) respectively; multiplied
together, these produce the answer 336.

What do you get if you multiply together the number of trees encountered on each of the listed
slopes?

*/

func main() {
	hill := ""
	lib.LoadAndLoop("../input.txt", func(ln string) error {
		if hill == "" {
			hill = fmt.Sprintf("%v", ln)
		} else {
			hill = fmt.Sprintf("%v\n%v", hill, ln)
		}

		return nil
	})

	pairs := []struct {
		vx, vy int
	}{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	res := 1
	for _, v := range pairs {
		s, err := day3.NewSled(v.vx, v.vy, hill)
		if err != nil {
			fmt.Printf("Unable to create sled: %v\n", err)
			os.Exit(1)
		}

		if err = s.Run(); err != nil {
			fmt.Printf("Error encountered while running sled: %v\n", err)
			os.Exit(1)
		}
		nt := s.NumTrees()
		fmt.Printf("Path intersects %v trees\n\n", nt)
		res *= nt
	}

	fmt.Printf("product: %v\n", res)

}
