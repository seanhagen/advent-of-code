package day08

import (
	"fmt"
	"testing"
)

func TestMemLength(t *testing.T) {
	tests := []struct {
		input  string
		length int
	}{
		{`""`, 0},
		{`"abc"`, 3},
		{`"aaa\"aaa"`, 7},
		{`"\x27"`, 1},
	}

	for i, x := range tests {
		tt := x
		t.Run(fmt.Sprintf("test %v", i), func(t *testing.T) {
			o := MemLength(tt.input)
			if o != tt.length {
				t.Errorf("wrong length for %v, expected %v got %v", tt.input, tt.length, o)
			}
		})
	}
}

func TestMemTotal(t *testing.T) {
	tests := []struct {
		input  []string
		expect int
	}{
		{[]string{`""`, `"abc"`, `"aaa\"aaa"`, `"\x27"`}, 12},
	}

	for i, x := range tests {
		tt := x
		t.Run(fmt.Sprintf("test %v", i), func(t *testing.T) {
			got := TotalMem(tt.input)
			if got != tt.expect {
				t.Errorf("wrong output, expected %v got %v", tt.expect, got)
			}
		})
	}
}
