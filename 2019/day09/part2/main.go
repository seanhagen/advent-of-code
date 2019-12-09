package main

import (
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"
	l2019 "github.com/seanhagen/advent-of-code/2019/lib2019"
	"github.com/seanhagen/advent-of-code/lib"
)

/*
--- Part Two ---
You now have a complete Intcode computer.

Finally, you can lock on to the Ceres distress signal! You just need to boost your sensors using
the BOOST program.

The program runs in sensor boost mode by providing the input instruction the value 2. Once run, it
will boost the sensors automatically, but it might take a few seconds to complete the operation on
slower hardware. In sensor boost mode, the program will output a single value: the coordinates of the
distress signal.

Run the BOOST program in sensor boost mode. What are the coordinates of the distress signal?

*/

func main() {
	f := lib.LoadInput("../input.txt")
	prog, err := l2019.ReadProgram(f)
	if err != nil {
		fmt.Printf("unable to read program from file: %v\n", err)
		os.Exit(1)
	}

	prog.AddInput(2)

	err = prog.Run()
	if err != nil {
		fmt.Printf("unable to run program: %v\n", err)
		os.Exit(1)
	}

	out := prog.GetOutputs()
	fmt.Printf("output:\n")
	spew.Dump(out)
}
