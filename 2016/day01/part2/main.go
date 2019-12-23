package main

import (
	"fmt"
	"os"

	"github.com/seanhagen/advent-of-code/2016/day01"
)

/*

Then, you notice the instructions continue on the back of the
Recruiting Document. Easter Buny HQ is actually at the first location
you visit twice.

For example, if your instructions are R8, R4, R4, R8, the first
location you visit twice is 4 blocks away, due East.

How many blocks away is the first location you visit twice?

*/

func main() {
	ans, err := day01.NewAnswer("../input.txt")
	if err != nil {
		fmt.Printf("unable to get solver: %v\n")
		os.Exit(1)
	}
	out, err := ans.FirstTwice()
	if err != nil {
		fmt.Printf("unable to process answer: %v\n")
		os.Exit(1)
	}
	fmt.Printf("%v blocks away\n", out)
}
