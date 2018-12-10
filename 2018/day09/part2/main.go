package main

import (
	"fmt"
	"time"

	"github.com/seanhagen/advent-of-code-2018/day9"
)

/*
--- Part Two ---
Amused by the speed of your answer, the Elves are curious:

What would the new winning Elf's score be if the number of the last marble were 100 times larger?
*/

func main() {
	start := time.Now()

	o := day9.RunGame(428, 72061*100)
	end := time.Now()
	diff := end.Sub(start)

	fmt.Printf("running time: %v\n", diff)
	fmt.Printf("winning score: %v\n", o)
}
