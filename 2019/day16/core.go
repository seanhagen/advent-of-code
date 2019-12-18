package day16

import (
	"fmt"
	"strconv"
	"strings"
)

// ApplyPhases ...
func ApplyPhases(input string, phases int) string {
	out := input
	outLen := len(input)

	for p := 1; p <= phases; p++ {
		// fmt.Printf("phase %v\n", p)
		col := make([]int, outLen)
		for itr := 1; itr <= outLen; itr++ {
			// fmt.Printf("\titr: %v\n", itr)
			pat := getPattern(outLen, itr)
			// fmt.Printf("\tpattern: %#v\n", pat)
			v := modify(input, pat)
			// fmt.Printf("\tv: %v\n\n", v)
			col[itr-1] = v
		}

		out = ""
		for _, v := range col {
			out = fmt.Sprintf("%v%v", out, v)
		}
		// fmt.Printf("col: %v -> %#v\n\n", col, out)
		input = out
	}
	return out
}

func modify(input string, pattern []int) int {
	bits := strings.Split(input, "")
	vals := []int{}

	for _, v := range bits {
		i, err := strconv.Atoi(v)
		if err != nil {
			return 0
		}
		vals = append(vals, i)
	}

	sum := 0
	idx := 0
	for _, v := range vals {
		if idx == len(pattern) {
			idx = 0
		}
		m := pattern[idx]
		sum += m * v

		idx++
	}

	out := strings.Split(strconv.Itoa(sum), "")
	i, _ := strconv.Atoi(out[len(out)-1])
	return i
}

func getPattern(length, itr int) []int {
	base := []int{0, 1, 0, -1}

	out := make([]int, length+1)

	idx := 0
	pos := 0
	for i := 0; i < length; i++ {
		for j := 0; j < itr; j++ {
			if idx == len(base) {
				idx = 0
			}

			out[pos] = base[idx]
			if pos == length {
				goto done
			}
			pos++
		}
		idx++
	}

done:
	return out[1:]
}
