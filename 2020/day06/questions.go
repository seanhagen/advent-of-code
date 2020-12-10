package day6

import (
	"sort"
	"strings"

	"github.com/mpvl/unique"
)

func NumQuestions(in string) int {
	f := getFound(in)
	return len(f)
}

func EveryoneYes(in string) int {
	f := getAnswered(in)

	qs := []string{}
	for _, v := range f {
		qs = append(qs, v...)
	}

	sort.Strings(qs)
	unique.Strings(&qs)

	counts := map[string]int{}

	for _, p := range f {
		for _, v := range qs {
			if inAr(v, p) {
				counts[v]++
			}
		}
	}

	numPeople := len(f)
	count := 0

	for _, v := range counts {
		if v == numPeople {
			count++
		}
	}

	return count
}

// inAr ...
func inAr(a string, b []string) bool {
	for _, v := range b {
		if a == v {
			return true
		}
	}
	return false
}

func getFound(in string) map[string]int {
	found := map[string]int{}

	bits := strings.Split(in, "")

	for _, v := range bits {
		if v == "\n" {
			continue
		}
		found[v]++
	}

	return found
}

func getAnswered(in string) map[int][]string {
	found := map[int][]string{}
	bits := strings.Split(in, "\n")
	for i, v := range bits {
		found[i] = strings.Split(v, "")
	}

	return found
}
