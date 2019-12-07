package day02

import (
	"fmt"
	"testing"
)

func TestGetRequiredSqFt(t *testing.T) {
	tests := []struct {
		input string
		out   int
	}{
		{"2x3x4", 58},
		{"1x1x10", 43},
	}

	for i, tt := range tests {
		x := tt
		t.Run(fmt.Sprintf("test %v", i), func(t *testing.T) {
			out := GetRequiredSqFt(x.input)
			if out != x.out {
				t.Errorf("invalid output, expected '%v', got '%v'", x.out, out)
			}
		})
	}
}
