package main

import (
	"fmt"
	"sort"

	day5 "github.com/seanhagen/advent-of-code/2020/day05"
	"github.com/seanhagen/advent-of-code/lib"
)

/*
--- Part Two ---
Ding! The "fasten seat belt" signs have turned on. Time to find your seat.

It's a completely full flight, so your seat should be the only missing boarding pass in your list.
However, there's a catch: some of the seats at the very front and back of the plane don't exist on
this aircraft, so they'll be missing from your list as well.

Your seat wasn't at the very front or back, though; the seats with IDs +1 and -1 from yours will be
in your list.

*/

func main() {
	rows := []int{}
	for i := 12; i <= 122; i++ {
		rows = append(rows, i)
	}

	found := map[int][]int{}

	lib.LoadAndLoop("../input.txt", func(line string) error {
		row, seat := day5.GetRowSeat(line)
		// id := day5.GetSeatID(row, seat)
		// if id > hid {
		//   hid = id
		// }

		r, ok := found[row]
		if !ok {
			r = []int{}
		}

		r = append(r, seat)
		found[row] = r

		return nil
	})

	var row, seat int

	for _, v := range rows {
		r, ok := found[v]
		if ok {
			sort.Ints(r)
			if len(r) != 8 {
				row = v
				for i := 1; i < len(r)-1; i++ {
					vv := r[i-1]
					if vv+1 != r[i] {
						seat = vv + 1
					}
				}
				break
			}
		}
	}

	fmt.Printf("Row: %v, Seat: %v\n", row, seat)
	id := day5.GetSeatID(row, seat)

	fmt.Printf("seat id: %v\n", id)
}
