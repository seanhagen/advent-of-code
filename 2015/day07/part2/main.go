package main

import (
	"fmt"
	"os"

	"github.com/seanhagen/advent-of-code/2015/day07"
	"github.com/seanhagen/advent-of-code/lib"
)

/*
--- Part Two ---

Now, take the signal you got on wire a, override wire b to that signal, and reset the other wires (including
wire a). What new signal is ultimately provided to wire a?
*/

func main() {
	brd, err := day07.NewBoard()
	if err != nil {
		fmt.Printf("unable to create wire board: %v\n", err)
		os.Exit(1)
	}

	file := "../input.txt"
	// file := "../sorted.txt"
	i := 0
	err = lib.LoadAndLoop(file, func(in string) error {
		i++
		return brd.AddWire(in)
	})
	fmt.Printf("processed %v lines\n", i)

	if err != nil {
		fmt.Printf("unable to load input: %v\n", err)
		os.Exit(1)
	}

	a := brd.GetWire("a").Value()
	if a == nil {
		fmt.Printf("no value for a?\n")
		os.Exit(1)
	}

	value := int(*a)

	brd.Reset()

	brd.SetWireValue("b", value)

	na := brd.GetWire("a").Value()
	if na == nil {
		fmt.Printf("no value for a?\n")
		os.Exit(1)
	}

	fmt.Printf("new a value: %v\n", *na)
}
