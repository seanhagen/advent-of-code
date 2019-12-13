package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/seanhagen/advent-of-code/2019/day13"
	"github.com/seanhagen/advent-of-code/2019/lib2019"
	"github.com/seanhagen/advent-of-code/lib"
)

/*
--- Part Two ---
The game didn't run because you didn't put in any quarters. Unfortunately, you did not bring any
quarters. Memory address 0 represents the number of quarters that have been inserted; set it to 2 to
play for free.

The arcade cabinet has a joystick that can move left and right. The software reads the position of
the joystick with input instructions:

 - If the joystick is in the neutral position, provide 0.
 - If the joystick is tilted to the left, provide -1.
 - If the joystick is tilted to the right, provide 1.

The arcade cabinet also has a segment display capable of showing a single number that represents the
player's current score. When three output instructions specify X=-1, Y=0, the third output
instruction is not a tile; the value instead specifies the new score to show in the segment display.
For example, a sequence of output values like -1,0,12345 would show 12345 as the player's current
score.

Beat the game by breaking all the blocks. What is your score after the last block is broken?
*/

func main() {
	f := lib.LoadInput("../input.txt")
	prog, err := lib2019.ReadProgram(f)
	if err != nil {
		fmt.Printf("unable to create program: %v\n", err)
		os.Exit(1)
	}

	prog.AddInput(2)

	score := 0
	for {
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
		score = g.GetScore()

		// // uncomment to see the game played!
		// g.Print()
		// if v, ok := clear[runtime.GOOS]; ok {
		// 	time.Sleep(time.Millisecond * 10)
		// 	v()
		// }

		cb := g.CountTileType(day13.BlockTile)
		if cb == 0 {
			break
		}

		padX := g.GetPaddlePosition()
		baX := g.GetBallPosition()

		input := 0
		if padX > baX {
			input = -1
		}
		if padX < baX {
			input = 1
		}
		prog.AddInput(input)
		prog.Unhalt()
	}

	fmt.Printf("score: %v\n", score)
}

// only for when printing gameplay
var clear map[string]func() //create a map for storing clear funcs

func init() {
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
