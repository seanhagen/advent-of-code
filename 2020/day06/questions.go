package day6

import (
	"strings"
)

func NumQuestions(in string) int {
	found := map[string]int{}

	bits := strings.Split(in, "")

	for _, v := range bits {
		if v == "\n" {
			continue
		}
		found[v]++
	}

	return len(found)
}
