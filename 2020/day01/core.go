package day1

import (
	"fmt"
	"sort"
)

func FindEntries(in []int, sumto int) (int, int, error) {
	var a, b int = -1, -1

	if sumto != 0 && len(in) == 0 {
		return a, b, fmt.Errorf("can't get to %v with no numbers in list", sumto)
	}

	for i, v := range in {
		x := v
		if v < 0 && sumto > 0 {
			x = x * -1
		}
		need := sumto - x
		found := 0
		for j, w := range in {
			if i == j {
				continue
			}
			if w == need {
				found = w
				break
			}
		}

		if found == need {
			a = x
			b = found
			break
		}
	}

	if a == -1 && b == -1 && sumto != a+b {
		return a, b, fmt.Errorf("numbers summing to %v not found in list", sumto)
	}

	return a, b, nil
}

// copied from reddit megathread
func FindThreeEntries(in []int, sumto int) ([]int, error) {
	sort.Ints(in)

	for i := 0; i < len(in)-2; i++ {
		r := len(in) - 1
		l := i + 1
		for l < r {
			a, b, c := in[i], in[l], in[r]
			if a+b+c == sumto {
				return []int{a, b, c}, nil
			} else {
				if a+b+c < sumto {
					l++
				} else {
					r--
				}
			}
		}
	}
	return []int{}, fmt.Errorf("unable to find %v as sum of three items in input list", sumto)
}
