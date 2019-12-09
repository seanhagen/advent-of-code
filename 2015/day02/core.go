package day02

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func inputToIntSlict(input string) []int {
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
	return parts
}

func GetRequiredSqFt(input string) int {
	parts := inputToIntSlict(input)
	sides := []int{
		parts[0] * parts[1],
		parts[1] * parts[2],
		parts[0] * parts[2],
	}

	smallest := getSmallest(sides)

	return smallest + (sides[0] * 2) + (sides[1] * 2) + (sides[2] * 2)
}

func GetRibbonLen(input string) int {
	parts := inputToIntSlict(input)

	sort.Slice(parts, func(i, j int) bool {
		if parts[i] < parts[j] {
			return true
		}
		return false
	})

	a := parts[0]
	b := parts[1]
	c := parts[2]

	return (a * 2) + (b * 2) + (a * b * c)
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
