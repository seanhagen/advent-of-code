package main

import (
	"fmt"
	"os"
	"strings"

	day6 "github.com/seanhagen/advent-of-code/2020/day06"
	"github.com/seanhagen/advent-of-code/lib"
)

/*
As you finish the last group's customs declaration, you notice that you misread one word in the
instructions:

You don't need to identify the questions to which anyone answered "yes"; you need to identify the
questions to which everyone answered "yes"!

Using the same example as above:

abc

a
b
c

ab
ac

a
a
a
a

b
This list represents answers from five groups:

In the first group, everyone (all 1 person) answered "yes" to 3 questions: a, b, and c.
In the second group, there is no question to which everyone answered "yes".
In the third group, everyone answered yes to only 1 question, a. Since some people did not answer
"yes" to b or c, they don't count.
In the fourth group, everyone answered yes to only 1 question, a.
In the fifth group, everyone (all 1 person) answered "yes" to 1 question, b.
In this example, the sum of these counts is 3 + 0 + 1 + 1 + 1 = 6.


*/

func main() {
	in, err := lib.GetString("../input.txt")
	if err != nil {
		fmt.Printf("Unable to load input file: %v\n", err)
		os.Exit(1)
	}

	data := strings.Split(in, "\n\n")
	numQs := []int{}
	for _, v := range data {
		n := day6.EveryoneYes(v)
		numQs = append(numQs, n)
	}

	s := sum(numQs)
	fmt.Printf("sum: %v\n", s)

}

func sum(in []int) int {
	out := 0
	for _, v := range in {
		out += v
	}
	return out
}
