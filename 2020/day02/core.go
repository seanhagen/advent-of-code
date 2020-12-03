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
