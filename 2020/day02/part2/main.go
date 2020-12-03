package main

import (
	"fmt"
	"strings"

	day2 "github.com/seanhagen/advent-of-code/2020/day02"
	"github.com/seanhagen/advent-of-code/lib"
)

/*
--- Part Two ---
While it appears you validated the passwords correctly, they don't seem to be what
the Official Toboggan Corporate Authentication System is expecting.

The shopkeeper suddenly realizes that he just accidentally explained the password
policy rules from his old job at the sled rental place down the street! The Official
Toboggan Corporate Policy actually works a little differently.

Each policy actually describes two positions in the password, where 1 means the first
character, 2 means the second character, and so on. (Be careful; Toboggan Corporate
Policies have no concept of "index zero"!) Exactly one of these positions must contain
the given letter. Other occurrences of the letter are irrelevant for the purposes of
policy enforcement.

Given the same example list from above:

1-3 a: abcde is valid: position 1 contains a and position 3 does not.
1-3 b: cdefg is invalid: neither position 1 nor position 3 contains b.
2-9 c: ccccccccc is invalid: both position 2 and position 9 contain c.

How many passwords are valid according to the new interpretation of the policies?
*/

const f = "%d %d %s %s"

func main() {
	num := 0
	p := func(i string) (int, int, string, string, error) {
		var min, max int
		var letter, pw string

		i = strings.Replace(i, "-", " ", 1)
		i = strings.Replace(i, ":", " ", 1)

		_, err := fmt.Sscanf(i, f, &min, &max, &letter, &pw)
		if err != nil {
			return min, max, letter, pw, err
		}
		return min, max, letter, pw, nil
	}

	lib.LoadAndLoop("../input.txt", func(in string) error {
		min, max, letter, pw, err := p(in)
		if err != nil {
			return err
		}
		v := day2.NewValidPass(min, max, letter, pw)
		fmt.Printf("line %v, valid: %v\n", in, v)
		if v {
			num++
		}
		return nil
	})

	fmt.Printf("valid passwords: %v\n", num)
}
