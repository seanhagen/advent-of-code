package day16

import (
	"fmt"
	"strconv"
	"strings"
)

const phases = 100
const repeat = 10000

// FinalProcessing ...
func FinalProcessing(input string) string {
	input = strings.Replace(input, "\n", "", -1)
	input = strings.Repeat(input, repeat)

	bits := strings.Split(input, "")
	signal := make([]int, len(bits))
	for i := 0; i < len(bits); i++ {
		j, _ := strconv.Atoi(bits[i])
		signal[i] = j
	}
	otmp := strings.Join(bits[0:7], "")
	offset, _ := strconv.Atoi(otmp)

	for i := 0; i < phases; i++ {
		l := len(signal)
		out := make([]int, l)
		tmp := 0
		for i := l - 1; i >= offset; i-- {
			tmp += signal[i]
			tmp = tmp % 10
			out[i] = tmp
		}
		signal = out
	}

	outStr := ""
	for i := offset; i < offset+8; i++ {
		outStr = fmt.Sprintf("%v%v", outStr, signal[i])
	}

	return outStr
}

// FirstEight ...
func FirstEight(input string, phases int) string {
	input = strings.Replace(input, "\n", "", -1)
	tmp := ApplyPhases(input, phases)
	bits := strings.Split(tmp, "")
	return strings.Join(bits[:8], "")
}

// ApplyPhases ...
func ApplyPhases(input string, phases int) string {

	out := input
	outLen := len(input)

	// fmt.Printf("input: %v\nlen: %v\n", input, outLen)
	for p := 1; p <= phases; p++ {
		// fmt.Printf("phase %v\n", p)
		col := make([]int, outLen)
		for itr := 1; itr <= outLen; itr++ {
			// fmt.Printf("\titr: %v\n", itr)
			pat := getPattern(outLen, itr)
			// fmt.Printf("\tpattern: %#v (%v)\n", pat, len(pat))
			v := modify(input, pat)
			// fmt.Printf("\tv: %v\n\n", v)
			col[itr-1] = v
		}

		// fmt.Printf("col: %#v\n", col)
		out = ""
		for _, v := range col {
			out = fmt.Sprintf("%v%v", out, v)
		}
		// fmt.Printf("phase %v, out: %v\n", p, out)
		// if p == 2 {
		// 	os.Exit(1)
		// }

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
			fmt.Printf("unable to parse integer %v --> %v\n", v, err)
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
		// fmt.Printf("\t\tmodifing %v, val: %v, pat: %v, out: %v, sum now: %v\n", i, v, m, m*v, sum)
		idx++
	}

	i := abs(sum) % 10

	// out := strings.Split(strconv.Itoa(sum), "")
	// i, _ := strconv.Atoi(out[len(out)-1])
	return i
}

func abs(n int) int {
	x := int64(n)

	y := x >> 63
	return int((x ^ y) - y)
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
