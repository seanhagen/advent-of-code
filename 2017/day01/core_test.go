package day01

import (
	"fmt"
	"testing"
)

func TestSolveCaptcha(t *testing.T) {
	tests := []struct {
		input string
		out   int
	}{
		{"1122", 3},
		{"1111", 4},
		{"1234", 0},
		{"91212129", 9},
	}

	for i, x := range tests {
		tt := x
		t.Run(fmt.Sprintf("test %v", i), func(t *testing.T) {
			out := SolveCaptcha(tt.input)
			if out != tt.out {
				t.Errorf("wrong output, expected '%v' got '%v'", tt.out, out)
			}
		})
	}
}

func TestSolveCaptchaFiveStep(t *testing.T) {
	tests := []struct {
		input string
		out   int
	}{
		{"1212", 6},
		{"1221", 0},
		{"123425", 4},
		{"123123", 12},
		{"12131415", 4},
	}

	for i, x := range tests {
		tt := x
		t.Run(fmt.Sprintf("test %v", i), func(t *testing.T) {
			out := SolveCaptchaFiveStep(tt.input)
			if out != tt.out {
				t.Errorf("wrong output, expected '%v' got '%v'", tt.out, out)
			}
		})
	}
}
