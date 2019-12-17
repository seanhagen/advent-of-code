package main

import (
	"fmt"
	"os"

	"github.com/seanhagen/advent-of-code/2019/day14"
	"github.com/seanhagen/advent-of-code/lib"
)

/*
--- Part Two ---

After collecting ORE for a while, you check your cargo hold: 1 trillion (1000000000000) units of
ORE.

With that much ore, given the examples above:

 - The 13312 ORE-per-FUEL example could produce 82892753 FUEL.
 - The 180697 ORE-per-FUEL example could produce 5586022 FUEL.
 - The 2210736 ORE-per-FUEL example could produce 460664 FUEL.

Given 1 trillion ORE, what is the maximum amount of FUEL you can produce?

*/

func main() {
	data := []string{}
	lib.LoadAndLoop("../input.txt", func(in string) error {
		data = append(data, in)
		return nil
	})

	nf, err := day14.CreateNanofactory(data)
	if err != nil {
		fmt.Printf("unable to create nanofactory: %v\n", err)
		os.Exit(1)
	}
	o := nf.CalcTrillionOre()
	fmt.Printf("answer: %v\n", o)
}
