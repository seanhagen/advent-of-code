package main

import (
	"fmt"
	"os"

	"github.com/seanhagen/advent-of-code/2019/lib2019"
	"github.com/seanhagen/advent-of-code/lib"
)

/*

--- Day 17: Set and Forget ---

An early warning system detects an incoming solar flare and automatically activates the ship's
electromagnetic shield. Unfortunately, this has cut off the Wi-Fi for many small robots that,
unaware of the impending danger, are now trapped on exterior scaffolding on the unsafe side of the
shield. To rescue them, you'll have to act quickly!

The only tools at your disposal are some wired cameras and a small vacuum robot currently asleep at
its charging station. The video quality is poor, but the vacuum robot has a needlessly bright LED
that makes it easy to spot no matter where it is.

An Intcode program, the Aft Scaffolding Control and Information Interface (ASCII, your puzzle
input), provides access to the cameras and the vacuum robot.  Currently, because the vacuum robot is
asleep, you can only access the cameras.

Running the ASCII program on your Intcode computer will provide the current view of the scaffolds.
This is output, purely coincidentally, as ASCII code: 35 means #, 46 means ., 10 starts a new line
of output below the current one, and so on. (Within a line, characters are drawn left-to-right.)

In the camera output, # represents a scaffold and . represents open space. The vacuum robot is
visible as ^, v, <, or > depending on whether it is facing up, down, left, or right respectively.
When drawn like this, the vacuum robot is always on a scaffold; if the vacuum robot ever walks off
of a scaffold and begins tumbling through space uncontrollably, it will instead be visible as X.

In general, the scaffold forms a path, but it sometimes loops back onto itself.  For example,
suppose you can see the following view from the cameras:

..#..........
..#..........
#######...###
#.#...#...#.#
#############
..#...#...#..
..#####...^..


Here, the vacuum robot, ^ is facing up and sitting at one end of the scaffold near the bottom-right
of the image. The scaffold continues up, loops across itself several times, and ends at the top-left
of the image.

The first step is to calibrate the cameras by getting the alignment parameters of some well-defined
points.  Locate all scaffold intersections; for each, its alignment parameter is the distance
between its left edge and the left edge of the view multiplied by the distance between its top edge
and the top edge of the view.  Here, the intersections from the above image are marked O:

..#..........
..#..........
##O####...###
#.#...#...#.#
##O###O###O##
..#...#...#..
..#####...^..


For these intersections:

	- The top-left intersection is 2 units from the left of the image and 2 units from the top of the image, so its alignment parameter is 2 * 2 = 4.

	- The bottom-left intersection is 2 units from the left and 4 units from the top, so its alignment parameter is 2 * 4 = 8.

	- The bottom-middle intersection is 6 from the left and 4 from the top, so its alignment parameter is 24.

	- The bottom-right intersection's alignment parameter is 40.

To calibrate the cameras, you need the sum of the alignment parameters.  In the above example, this
is 76.

Run your ASCII program. What is the sum of the alignment parameters for the scaffold intersections?

*/

func main() {
	f := lib.LoadInput("../input.txt")
	p, err := lib2019.ReadProgram(f)
	if err != nil {
		fmt.Printf("unable to create program: %v\n", err)
		os.Exit(1)
	}

	data := [][]string{}
	tmp := []string{}

	p.SetOutputFn(func(i int) bool {
		v := fmt.Sprintf("%v", string(i))
		if v == "\n" {
			data = append(data, tmp)
			tmp = []string{}
			return false
		}
		tmp = append(tmp, v)
		return false
	})

	err = p.Run()

	data = data[:len(data)-1]
	locs := map[int]map[int]bool{}

	checks := []string{"#", "^", "v", "<", ">"}
	fn := func(i string) bool {
		for _, v := range checks {
			if v == i {
				return true
			}
		}
		return false
	}

	endpoints := map[int]map[int]bool{}

	for j, row := range data {
		for i, d := range row {
			if d == "." {
				fmt.Printf(".")
				continue
			}

			count := 0

			// can check north if j > 0
			if j > 0 {
				n := data[j-1][i]
				if fn(n) {
					count++
				}
			}

			// can check south if j < len(data)-1
			if j < len(data)-1 {
				s := data[j+1][i]
				if fn(s) {
					count++
				}
			}

			// can check east if i < len(row) -1
			if i < len(row)-1 {
				e := data[j][i+1]
				if fn(e) {
					count++
				}
			}

			// can check west if i > 0
			if i > 0 {
				w := data[j][i-1]
				if fn(w) {
					count++
				}
			}

			// count > 2, then an intersection
			if count > 2 {
				x, ok := locs[j]
				if !ok {
					x = map[int]bool{}
				}
				x[i] = true
				locs[j] = x
			}

			// count == 1 is an endpoint ( start or end )
			if count == 1 {
				tmp, ok := endpoints[j]
				if !ok {
					tmp = map[int]bool{}
				}
				tmp[i] = true
				endpoints[j] = tmp
			}

			fmt.Printf("%v", count)
		}
		fmt.Printf("\n")
	}

	for y, t := range endpoints {
		for x := range t {
			fmt.Printf("endpoint at %v, %v\n", x, y)
		}
	}

	sum := 0
	for y, v := range locs {
		for x := range v {
			sum += x * y
		}
	}

	fmt.Printf("answer: %v\n", sum)
}
