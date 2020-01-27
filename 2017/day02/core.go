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

func runFunc(in string, fn func([]int) int) int {
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

		sum += fn(v)
	}

	return sum
}

func GetChecksum(in string) int {
	return runFunc(in, getRowDiff)
}

func getDivisible(in []int) int {
	a, b := 0, 1

	for i, v := range in {
		for ii, vv := range in {
			if i == ii {
				continue
			}
			if v%vv == 0 {
				a = v
				b = vv
				goto done
			}
		}
	}
done:

	return a / b
}

func GetDivisibleSum(in string) int {
	return runFunc(in, getDivisible)
}
