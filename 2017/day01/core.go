package day01

import (
	"fmt"
	"strconv"
	"strings"
)

func solver(in string, dist int) int {

	in = strings.TrimSpace(in)
	bits := strings.Split(in, "")
	parts := []string{}

	for i, v := range bits {
		p := i + dist
		if p >= len(bits) {
			p -= len(bits)
		}

		if bits[i] == bits[p] {
			parts = append(parts, v)
		}
	}

	output := 0
	if len(parts) <= 0 {
		return output
	}

	for _, v := range parts {
		if i, err := strconv.Atoi(v); err == nil {
			output += i
		} else {
			fmt.Printf("unable to convert: %v\n", err)
		}
	}

	return output

}

func SolveCaptcha(in string) int {
	return solver(in, 1)
}

func SolveCaptchaFiveStep(in string) int {
	dist := len(in) / 2
	return solver(in, dist)
}
