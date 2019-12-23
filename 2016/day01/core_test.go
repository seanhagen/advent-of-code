package day01

import (
	"fmt"
	"testing"
)

func TestAnswer(t *testing.T) {
	tests := []struct {
		input  string
		answer int
	}{
		{"R2, L3", 5},
		{"R2, R2, R2", 2},
		{"R5, L5, R5, R3", 12},
	}

	for i, x := range tests {
		tt := x
		t.Run(fmt.Sprintf("test %v", i), func(t *testing.T) {
			a, err := answerFromDirString(tt.input)
			if err != nil {
				t.Fatalf("unable to get solver: %v", err)
			}

			out, err := a.Process()
			if err != nil {
				t.Fatalf("unable to process answer: %v", err)
			}

			if out != tt.answer {
				t.Errorf("got wrong answer, expected '%v' got '%v'", tt.answer, out)
			}
		})
	}
}

func TestFirstTwice(t *testing.T) {
	tests := []struct {
		input  string
		answer int
	}{
		{"R8, R4, R4, R8", 4},
	}

	for i, x := range tests {
		tt := x
		t.Run(fmt.Sprintf("test %v", i), func(t *testing.T) {
			a, err := answerFromDirString(tt.input)
			if err != nil {
				t.Fatalf("unable to get solver: %v", err)
			}

			out, err := a.FirstTwice()
			if err != nil {
				t.Fatalf("unable to process answer: %v", err)
			}

			if out != tt.answer {
				t.Errorf("got wrong answer, expected '%v' got '%v'", tt.answer, out)
			}
		})
	}
}
