package day05

import (
	"fmt"
	"strings"
	"testing"
)

func TestNiceString(t *testing.T) {
	tests := []struct {
		input  string
		should bool
	}{
		{"ugknbfddgIcrmopn", true},
		{"aaa", true},
		{"aAa", true},
		{"jchzalrnumimnmhp", false},
		{"haegwjzuvuyypxyu", false},
		{"haegwjzuvuyypXyu", false},
		{"dvszwmarrgswjxmb", false},
	}

	for _, tt := range tests {
		x := tt
		t.Run(fmt.Sprintf("test %v", x.input), func(t *testing.T) {
			got := NiceString(x.input)
			if got != x.should {
				t.Errorf("wrong output, expected '%v' got '%v'", x.should, got)
			}
		})
	}
}

func TestNiceStringV2(t *testing.T) {
	tests := []struct {
		input  string
		should bool
	}{
		{"qjhvhtzxzqqjkmpb", true},
		{"xxyxx", true},
		{"xxYxx", true},
		{"uurcxstgmygtbstg", false},
		{"ieodomkazucvgmuy", false},
		{"iEodOmkazucvgmuy", false},
	}

	for _, tt := range tests {
		x := tt
		t.Run(fmt.Sprintf("test %v", x.input), func(t *testing.T) {
			got := NiceStringV2(x.input)
			if got != x.should {
				t.Errorf("wrong output, expected '%v' got '%v'", x.should, got)
			}
		})
	}
}

func TestNoOverlap(t *testing.T) {
	tests := []struct {
		input  string
		should bool
	}{
		{"xyxy", true},
		{"aabcdefgaa", true},
		{"aaa", false},
		{"ieodomkazucvgmuy", false},
	}

	for _, tt := range tests {
		x := tt
		t.Run(fmt.Sprintf("test %v", x.input), func(t *testing.T) {
			in := strings.Split(x.input, "")
			if got := twiceNoOverlap(in); got != x.should {
				t.Errorf("wrong output, expected '%v' got '%v'", x.should, got)
			}
		})
	}
}

func TestRepeatWithGap(t *testing.T) {
	tests := []struct {
		input  string
		should bool
	}{
		{"xyx", true},
		{"abcdefeghi", true},
		{"aaa", true},
		{"abc", false},
		{"higop", false},
		{"a", false},
		{"ab", false},
		{"uurcxstgmygtbstg", false},
	}

	for _, tt := range tests {
		x := tt
		t.Run(fmt.Sprintf("test %v", x.input), func(t *testing.T) {
			in := strings.Split(x.input, "")
			if got := repeatWithGap(in); got != x.should {
				t.Errorf("wrong output, expected '%v' got '%v'", x.should, got)
			}
		})
	}
}
