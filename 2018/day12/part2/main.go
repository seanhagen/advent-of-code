package main

import (
	"fmt"

	"github.com/seanhagen/advent-of-code/2018/day12"
)

/*
You realize that 20 generations aren't enough. After all,
these plants will need to last another 1500 years to even
reach your timeline, not to mention your future.

After fifty billion (50000000000) generations, what is the
sum of the numbers of all pots which contain a plant?
*/

const fiftybil = 50000000000

func main() {
	g := day12.SetupGame("../input.txt")
	g.TakeSteps(fiftybil)

	c := g.SumCurrent()
	fmt.Printf("sum: %v\n\n", c)
}
