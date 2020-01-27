package main

import (
	"fmt"
	"os"

	"github.com/seanhagen/advent-of-code/2017/day02"
	"github.com/seanhagen/advent-of-code/lib"
)

func main() {
	input, err := lib.GetString("../input.txt")
	if err != nil {
		fmt.Printf("unable to load input: %v\n", err)
		os.Exit(1)
	}

	checksum := day02.GetDivisibleSum(input)

	fmt.Printf("got sum: %v\n", checksum)
}
