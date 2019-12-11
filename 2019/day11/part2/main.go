package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/seanhagen/advent-of-code/2019/day11"
)

/*
--- Part Two ---

You're not sure what it's trying to paint, but it's definitely not a
registration identifier. The Space Police are getting impatient.

Checking your external ship cameras again, you notice a white panel marked
"emergency hull painting robot starting panel". The rest of the panels are still
black, but it looks like the robot was expecting to start on a white panel, not
a black one.

Based on the Space Law Space Brochure that the Space Police attached to one of
your windows, a valid registration identifier is always eight capital
letters. After starting the robot on a single white panel instead, what
registration identifier does it paint on your hull?

*/

func main() {
	code, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		fmt.Printf("unable to read input: %v\n", err)
		os.Exit(1)
	}

	robo, err := day11.NewRobot(string(code))
	if err != nil {
		fmt.Printf("Unable to create robot: %v\n", err)
		os.Exit(1)
	}

	robo.ColorPanel(0, 0, day11.White)

	err = robo.Run()
	if err != nil {
		fmt.Printf("Error running robot: %v\n", err)
		os.Exit(1)
	}

	robo.Print()
}
