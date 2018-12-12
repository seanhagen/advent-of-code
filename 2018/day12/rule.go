package day12

import (
	"strings"
)

// Rule ...
type Rule struct {
	conditions map[int]string
	match      string
	output     string
}

// apply ...
func (r Rule) matches(in []string) bool {
	matches := false
	if in[2] == r.match {
		if in[0] == r.conditions[-2] &&
			in[1] == r.conditions[-1] &&
			in[3] == r.conditions[1] &&
			in[4] == r.conditions[2] {
			matches = true
		}
		// if matches {
		// 	return r.output
		// }
	}
	// return "."
	return matches
}

func createRule(in string) Rule {
	bits := strings.Split(in, " => ")
	conditions := strings.Split(bits[0], "")
	return Rule{
		conditions: map[int]string{
			-2: conditions[0],
			-1: conditions[1],
			1:  conditions[3],
			2:  conditions[4],
		},
		match:  conditions[2],
		output: bits[1],
	}
}
