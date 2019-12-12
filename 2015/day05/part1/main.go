package main

import (
	"fmt"
	"io"
	"os"

	"github.com/seanhagen/advent-of-code/2015/day05"
	"github.com/seanhagen/advent-of-code/lib"
)

/*
--- Day 5: Doesn't He Have Intern-Elves For This? ---

Santa needs help figuring out which strings in his text file are naughty or nice.

A nice string is one with all of the following properties:

 - It contains at least three vowels (aeiou only), like aei, xazegov, or aeiouaeiouaeiou.
 - It contains at least one letter that appears twice in a row, like xx, abcdde (dd), or aabbccdd (aa, bb, cc, or dd).
 - It does not contain the strings ab, cd, pq, or xy, even if they are part of one of the other requirements.

For example:

 - -ugknbfddgicrmopn is nice because it has at least three vowels
 (u...i...o...), a double letter (...dd...), and none of the disallowed
 substrings.

 - aaa is nice because it has at least three vowels and a double letter, even
 though the letters used by different rules overlap.

 - jchzalrnumimnmhp is naughty because it has no double letter.

 - haegwjzuvuyypxyu is naughty because it contains the string xy.

 - dvszwmarrgswjxmb is naughty because it contains only one vowel.

How many strings are nice?

*/

func main() {
	f := lib.LoadInput("../input.txt")
	countNice := 0

	err := lib.LoopOverLines(f, func(in []byte) error {
		if day05.NiceString(string(in)) {
			countNice++
		}
		return nil
	})

	if err != io.EOF {
		fmt.Printf("unable to loop over lines: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("found %v nice strings\n", countNice)
}
