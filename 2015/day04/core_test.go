package day04

import (
	"fmt"
	"testing"
)

func TestFindHash(t *testing.T) {
	tests := []struct {
		input  string
		answer int
	}{
		{"abcdef", 609043},
		{"pqrstuv", 1048970},
	}

	for _, tt := range tests {
		x := tt
		t.Run(fmt.Sprintf("test %v", x.input), func(t *testing.T) {
			i, _ := FindHash(x.input)
			if i != x.answer {
				fmt.Printf("wrong answer -- expected %v, got %v", x.answer, i)
			}
		})
	}
}
