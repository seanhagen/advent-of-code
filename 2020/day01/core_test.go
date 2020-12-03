package day1

import (
	"fmt"
	"sort"
	"testing"
)

func TestFindEntries(t *testing.T) {
	tests := []struct {
		nums []int
		goal int
		a    int
		b    int
		fine bool
	}{
		{[]int{1721, 979, 366, 299, 675, 1456}, 2020, 1721, 299, true},
		{[]int{1010, 1010}, 2020, 1010, 1010, true},
		{[]int{2020, 0}, 2020, 2020, 0, true},
		{[]int{2, 3, 4, 2, 5}, 4, 2, 2, true},
		{[]int{3, 5, 6, 2, 1}, 4, 3, 1, true},
		{[]int{1, 1}, 2, 1, 1, true},
		{[]int{}, 1, 0, 0, false},
		{[]int{2021, 0}, 2020, 0, 0, false},
		{[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 1, 0, 0, false},
	}

	for i, tt := range tests {
		x := tt
		t.Run(fmt.Sprintf("test_%v", i), func(t *testing.T) {
			t.Parallel()
			a, b, err := FindEntries(x.nums, x.goal)
			if !x.fine && err != nil {
				return
			}

			if x.fine && err != nil {
				t.Fatalf("received error when should have been fine: %v", err)
			}
			if !x.fine && err == nil {
				t.Fatalf("received no error but should have")
			}
			var agood, bgood bool

			if a == x.a || a == x.b {
				agood = true
			}

			if b == x.b || b == x.a {
				if x.b == x.a && a == b {
					bgood = true
				}
				if x.b != x.a && a != b {
					bgood = true
				}
			}

			if agood != true || bgood != true {
				t.Errorf("incorrect answer, expected (%v,%v), got (%v,%v)", x.a, x.b, a, b)
			}
		})
	}
}

func TestFindThree(t *testing.T) {
	tests := []struct {
		nums  []int
		found []int
		sumto int
		prod  int
		fine  bool
	}{
		{[]int{1721, 979, 366, 299, 675, 1456}, []int{979, 366, 675}, 2020, 241861950, true},
		{[]int{1, 2, 1, 3, 1}, []int{1, 1, 1}, 3, 1, true},
		{[]int{}, []int{}, 1, 0, false},
	}

	eq := func(a, b []int) bool {
		if len(a) != len(b) {
			return false
		}
		sort.Ints(a)
		sort.Ints(b)
		for i, v := range a {
			if v != b[i] {
				return false
			}
		}
		return true
	}

	for i, tt := range tests {
		x := tt
		t.Run(fmt.Sprintf("test_%v", i), func(t *testing.T) {
			t.Parallel()
			o, err := FindThreeEntries(x.nums, x.sumto)
			if x.fine && err != nil {
				t.Fatalf("expected answer, got error: %v", err)
			}

			if !x.fine && err == nil {
				t.Fatalf("expected error, got answer?")
			}

			if !eq(o, x.found) {
				t.Fatalf("wrong answer, expected %v got %v", x.found, o)
			}

			ans := 0
			for _, v := range o {
				if ans == 0 {
					ans = v
				} else {
					ans *= v
				}
			}
			if ans != x.prod {
				t.Errorf("wrong product, expected %v got %v", x.prod, ans)
			}
			t.Logf("found items: %v, product: %v", o, ans)
		})
	}
}
