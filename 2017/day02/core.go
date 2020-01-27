package day02

import (
	"strconv"
	"strings"
)

func getRowDiff(in []int) int {
	min, max := 9999999, -9999999
	for _, v := range in {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}

	return max - min
}

func GetChecksum(in string) int {
	in = strings.TrimSpace(in)
	lines := strings.Split(in, "\n")
	sum := 0

	for _, l := range lines {
		v := []int{}

		bits := strings.Split(l, "\t")
		for _, vv := range bits {
			if i, err := strconv.Atoi(vv); err == nil {
				v = append(v, i)
			}
		}

		sum += getRowDiff(v)
	}

	return sum
}
