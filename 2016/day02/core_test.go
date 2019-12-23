package day02

import (
	"fmt"
	"testing"
)

func TestAnswer(t *testing.T) {
	tests := []struct {
		input string
		out   int
	}{
		{`ULL
RRDDD
LURDL
UUUUD`, 1985},
	}

	for i, x := range tests {
		tt := x
		t.Run(fmt.Sprintf("test %v", i), func(t *testing.T) {
			s, err := CreateSolver(tt.input)
			if err != nil {
				t.Fatalf("unable to create solver: %v", err)
			}
			out := s.Solve()
			if out != tt.out {
				t.Errorf("wrong output, expected %v got %v", tt.out, out)
			}
		})
	}
}
