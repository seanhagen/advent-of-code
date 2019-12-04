package day04

import (
	"regexp"
	"strconv"
	"strings"
)

// var adjacent = regexp.MustCompile(`(.)\1+`)
var adjacent = regexp.MustCompile(`(.)\\1+`)

// FindPasswords ...
func FindPasswords(min, max int) []int {
	out := []int{}

	for i := min; i <= max; i++ {
		if MeetsCriterea(i) {
			out = append(out, i)
		}
	}

	return out
}

// GroupMatch ...
func GroupMatch(input []int) []int {
	out := []int{}

	for _, v := range input {
		if lenSmallestMatchingGroup(v) == 2 {
			out = append(out, v)
		}
	}

	return out
}

func lenSmallestMatchingGroup(input int) int {
	v := strconv.Itoa(input)

	matches := []string{}

	bits := strings.Split(v, "")

	var last string
	var collector string

	for _, b := range bits {
		if b != last && last != "" {
			// collector = collector + b
			if len(collector+last) > 1 {
				matches = append(matches, collector+last)
			}
			collector = ""
			last = b
			continue
		}

		if b == last {
			collector = collector + b
		}

		last = b
	}

	if strings.Contains(collector, last) {
		collector = collector + last
	}

	if collector != "" && len(collector) > 1 {
		matches = append(matches, collector)
	}

	lenb := len(bits)
	for _, m := range matches {
		if len(m) < lenb {
			lenb = len(m)
		}
	}

	return lenb
}

// MeetsCriterea ...
func MeetsCriterea(input int) bool {
	v := strconv.Itoa(input)

	if len(v) != 6 {
		return false
	}

	if !adjacentEqual(v) {
		return false
	}

	if !alwaysIncrease(v) {
		return false
	}

	return true
}

func adjacentEqual(in string) bool {
	return strings.Contains(in, "00") ||
		strings.Contains(in, "11") ||
		strings.Contains(in, "22") ||
		strings.Contains(in, "33") ||
		strings.Contains(in, "44") ||
		strings.Contains(in, "55") ||
		strings.Contains(in, "66") ||
		strings.Contains(in, "77") ||
		strings.Contains(in, "88") ||
		strings.Contains(in, "99")
}

func alwaysIncrease(in string) bool {
	bits := strings.Split(in, "")

	var last int
	ai := true

	for i := 0; i < len(bits); i++ {
		j, err := strconv.Atoi(bits[i])
		if err != nil {
			return false
		}
		if last > j {
			ai = false
		}
		last = j
	}

	return ai
}
