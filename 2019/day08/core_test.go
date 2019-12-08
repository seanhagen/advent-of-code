package day08

import (
	"fmt"
	"testing"
)

func TestInvalidInput(t *testing.T) {
	tests := []struct {
		input string
		w     int
		h     int
		valid bool
	}{
		{"1234", 3, 2, false},
		{"123v", 2, 2, false},
		{"4321", 2, 2, true},
	}

	for _, tt := range tests {
		x := tt
		t.Run(fmt.Sprintf("test %v", x.input), func(t *testing.T) {
			_, err := NewImage(x.input, x.w, x.h)
			if x.valid && err != nil {
				t.Errorf("'%v' should be valid input bot got err: %v", x.input, err)
			}

			if !x.valid && err == nil {
				t.Errorf("should get error but got nil")
			}
		})
	}
}

func TestValidLayers(t *testing.T) {
	tests := []struct {
		input  string
		w      int
		h      int
		expect int
	}{
		{"123456789012", 3, 2, 1},
		{"122456789012", 3, 2, 2},
		{"112245789012", 3, 2, 4},
	}

	for _, tt := range tests {
		x := tt
		t.Run(fmt.Sprintf("test %v", x.input), func(t *testing.T) {
			img, err := NewImage(x.input, x.w, x.h)
			if err != nil {
				t.Fatalf("unable to create image: %v", err)
			}

			sml := img.FindSmallestNumZeroLayer()
			out := sml.OneByTwo()
			if out != x.expect {
				t.Errorf("wrong output, expected '%v', got '%v'", x.expect, out)
			}
		})
	}
}

// TestOutputImage ...
func TestOutputImage(t *testing.T) {
	tests := []struct {
		input  string
		w      int
		h      int
		output string
	}{
		{"0222112222120000", 2, 2, "01\n10"},
	}

	for _, tt := range tests {
		x := tt
		t.Run(fmt.Sprintf("test %v", x.input), func(t *testing.T) {
			img, err := NewImage(x.input, x.w, x.h)
			if err != nil {
				t.Fatalf("unable to create image: %v", err)
			}

			out := img.Output()
			if out != x.output {
				t.Errorf("wrong output\n expected:\n---\n%v\n---\n\ngot:\n---\n%v\n---\n", x.output, out)
			}
		})
	}
}
