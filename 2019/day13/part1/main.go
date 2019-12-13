package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/seanhagen/advent-of-code/2019/day13"
	"github.com/seanhagen/advent-of-code/2019/lib2019"
)

/*

--- Day 13: Care Package ---

As you ponder the solitude of space and the ever-increasing three-hour roundtrip for messages
between you and Earth, you notice that the Space Mail Indicator Light is blinking.  To help keep you
sane, the Elves have sent you a care package.

It's a new game for the ship's arcade cabinet! Unfortunately, the arcade is all the way on the other
end of the ship. Surely, it won't be hard to build your own - the care package even comes with
schematics.

The arcade cabinet runs Intcode software like the game the Elves sent (your puzzle input). It has a
primitive screen capable of drawing square tiles on a grid.  The software draws tiles to the screen
with output instructions: every three output instructions specify the x position (distance from the
left), y position (distance from the top), and tile id. The tile id is interpreted as follows:

	- 0 is an empty tile.  No game object appears in this tile.

	- 1 is a wall tile.  Walls are indestructible barriers.

	- 2 is a block tile.  Blocks can be broken by the ball.

	- 3 is a horizontal paddle tile.  The paddle is indestructible.

	- 4 is a ball tile.  The ball moves diagonally and bounces off objects.

For example, a sequence of output values like 1,2,3,6,5,4 would draw a horizontal paddle tile (1
tile from the left and 2 tiles from the top) and a ball tile (6 tiles from the left and 5 tiles from
the top).

Start the game. How many block tiles are on the screen when the game exits?

*/

func main() {
	code, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		fmt.Printf("unable to read input: %v\n", err)
		os.Exit(1)
	}

	prog, err := lib2019.FromString(string(code))
	if err != nil {
		fmt.Printf("unable to create program: %v\n", err)
		os.Exit(1)
	}

	err = prog.Run()
	if err != nil && err != lib2019.ErrHalt {
		fmt.Printf("error while running program: %v\n", err)
		os.Exit(1)
	}

	out := prog.GetOutputs()
	g, err := day13.NewGame(out)
	if err != nil {
		fmt.Printf("Unable to create game: %v\n", err)
		os.Exit(1)
	}

	c := g.CountTileType(day13.BlockTile)
	fmt.Printf("answer: %v\n", c)

}
