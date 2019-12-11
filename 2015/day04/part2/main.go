package main

import (
	"fmt"

	"github.com/seanhagen/advent-of-code/2015/day04"
)

/*
Part Two:

Now find one that starts with six zeroes

*/

func main() {
	i, _ := day04.FindHash("ckczppom", 6)
	fmt.Printf("answer: %v\n", i)
}
