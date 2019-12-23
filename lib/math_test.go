package lib

import (
	"fmt"
	"testing"
)

func TestAbs(t *testing.T) {
	tests := []struct {
		in  int
		out int
	}{
		{-1, 1},
		{-2, 2},
		{1, 1},
		{3, 3},
	}

	for i, x := range tests {
		tt := x
		t.Run(fmt.Sprintf("test %v", i), func(t *testing.T) {
			out := Abs(tt.in)
			if out != tt.out {
				t.Errorf("wrong answer, expected %v got %v", tt.out, out)
			}
		})
	}
}

func BenchmarkAbs(t *testing.B) {
	for n := 0; n < t.N; n++ {
		Abs(n * -1)
	}
}

func BenchmarkAbs2(t *testing.B) {
	for n := 0; n < t.N; n++ {
		Abs2(n * -1)
	}
}
