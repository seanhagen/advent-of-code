package day01

import "strings"

func GetFloor(input string) int {
	fl := 0
	ins := strings.Split(input, "")
	for _, v := range ins {
		switch v {
		case "(":
			fl++
		case ")":
			fl--
		default:
		}
	}
	return fl
}

func PosGoIntoBasement(input string) int {
	fl := 0
	ins := strings.Split(input, "")
	for pos, v := range ins {
		switch v {
		case "(":
			fl++
		case ")":
			fl--
		default:
		}
		if fl < 0 {
			return pos + 1
		}
	}
	return -1
}
