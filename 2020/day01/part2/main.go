package main

import (
	"fmt"
	"os"
	"strconv"

	day1 "github.com/seanhagen/advent-of-code/2020/day01"
	"github.com/seanhagen/advent-of-code/lib"
)

/*
--- Part Two ---
The Elves in accounting are thankful for your help; one of them even offers you a
starfish coin they had left over from a past vacation. They offer you a second one
if you can find three numbers in your expense report that meet the same criteria.

Using the above example again, the three entries that sum to 2020 are 979, 366, and
675. Multiplying them together produces the answer, 241861950.

In your expense report, what is the product of the three entries that sum to 2020?
*/

func main() {
	nums := []int{}
	err := lib.LoadAndLoop("../input.txt", func(l string) error {
		i, err := strconv.Atoi(l)
		if err != nil {
			return err
		}
		nums = append(nums, i)
		return nil
	})
	if err != nil {
		fmt.Printf("Unable to load input.txt: %v\n", err)
		os.Exit(1)
	}

	out, err := day1.FindThreeEntries(nums, 2020)
	if err != nil {
		fmt.Printf("Unable to find entries: %v\n", err)
		os.Exit(1)
	}

	ans := 1
	for _, v := range out {
		ans *= v
	}
	fmt.Printf("answer: %v\n\n", ans)
}
