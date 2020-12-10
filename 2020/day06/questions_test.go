package day6

import (
	"fmt"
	"testing"
)

/*

abc

a
b
c

ab
ac

a
a
a
a

b
This list represents answers from five groups:

The first group contains one person who answered "yes" to 3 questions: a, b, and c.
The second group contains three people; combined, they answered "yes" to 3 questions: a, b, and c.
The third group contains two people; combined, they answered "yes" to 3 questions: a, b, and c.
The fourth group contains four people; combined, they answered "yes" to only 1 question, a.
The last group contains one person who answered "yes" to only 1 question, b.
*/
func TestNumQuestions(t *testing.T) {

	tests := []struct {
		input string
		num   int
	}{
		{"abc", 3},
		{`a
b
c`, 3},
		{`ab
ac`, 3},
		{`a
a
a
a`, 1},
		{"b", 1},
	}

	for i, x := range tests {
		tt := x
		t.Run(fmt.Sprintf("test_%v", i), func(t *testing.T) {
			n := NumQuestions(tt.input)
			if n != tt.num {
				t.Errorf("wrong number of questions, expected %v got %v", tt.num, n)
			}
		})
	}
}

func TestEveryoneYes(t *testing.T) {
	tests := []struct {
		input string
		num   int
	}{
		{"abc", 3},
		{`a
b
c`, 0},
		{`ab
ac`, 1},
		{`a
a
a
a`, 1},
		{"b", 1},
	}
	for i, x := range tests {
		tt := x
		t.Run(fmt.Sprintf("test_%v", i), func(t *testing.T) {
			n := EveryoneYes(tt.input)
			if n != tt.num {
				t.Errorf("wrong number of questions, expected %v got %v", tt.num, n)
			}
		})
	}
}
