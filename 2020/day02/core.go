package day2

import (
	"strings"
)

func ValidatePassword(min, max int, letter, pw string) bool {
	c := 0
	bits := strings.Split(pw, "")
	for _, v := range bits {
		if v == letter {
			c++
		}
	}
	if c >= min && c <= max {
		return true
	}
	return false
}

func NewValidPass(f, s int, ch, pw string) bool {
	x := 0
	bits := strings.Split(pw, "")
	f -= 1
	s -= 1

	l := len(pw) - 1
	if f > l || s > l {
		return false
	}

	if f <= l && bits[f] == ch {
		x++
	}

	if s <= l && bits[s] == ch {
		x++
	}

	return x == 1
}
