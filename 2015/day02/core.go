package day02

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetRequiredSqFt(input string) int {
	bits := strings.Split(input, "x")

	parts := make([]int, 3)

	for i, v := range bits {
		j, err := strconv.Atoi(v)
		if err != nil {
			fmt.Printf("bad input to strconv.Atoi(%v): %v", v, err)
			os.Exit(1)
		}
		parts[i] = j
	}

	sides := []int{
		parts[0] * parts[1],
		parts[1] * parts[2],
		parts[0] * parts[2],
	}

	smallest := getSmallest(sides)

	return smallest + (sides[0] * 2) + (sides[1] * 2) + (sides[2] * 2)
}

func getSmallest(in []int) int {
	o := 0
	for i, v := range in {
		if i == 0 {
			o = v
			continue
		}
		if v < o {
			o = v
		}
	}
	return o
}
