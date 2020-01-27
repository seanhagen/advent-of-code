package day01

import (
	"fmt"
	"strconv"
	"strings"
)

func SolveCaptcha(in string) int {
	in = strings.TrimSpace(in)
	bits := strings.Split(in, "")
	parts := []string{}

	for i, v := range bits {
		if i != len(bits)-1 {
			if bits[i] == bits[i+1] {
				parts = append(parts, v)
			}
		} else {
			if bits[i] == bits[0] {
				parts = append(parts, v)
			}
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
