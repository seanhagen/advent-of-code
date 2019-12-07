package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/seanhagen/advent-of-code/2015/day01"
)

/*
--- Part Two ---
Now, given the same instructions, find the position of the first character that causes him
to enter the basement (floor -1). The first character in the instructions has position 1, the
second character has position 2, and so on.

For example:

  - ) causes him to enter the basement at character position 1.
  - ()()) causes him to enter the basement at character position 5.

What is the position of the character that causes Santa to first enter the basement?

*/

func main() {
	raw, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		fmt.Printf("unable to load file: %v\n", err)
		os.Exit(1)
	}

	pos := day01.PosGoIntoBasement(string(raw))
	fmt.Printf("pos: %v\n", pos)

}
