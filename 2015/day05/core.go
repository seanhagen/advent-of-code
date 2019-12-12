package day05

import (
	"strings"
)

func isVowel(i string) bool {
	return strings.Contains("aeiou", i)
}

func isForbidden(i string) bool {
	return i == "ab" ||
		i == "cd" ||
		i == "pq" ||
		i == "xy"
}

// NiceString ...
func NiceString(in string) bool {
	in = strings.ToLower(in)

	countVowels := 0
	doubleLetter := false
	hasForbidden := false

	bits := strings.Split(in, "")

	for i, v := range bits {
		if isVowel(v) {
			countVowels++
		}

		if i > 0 && bits[i-1] == v {
			doubleLetter = true
		}

		if i > 0 {
			if c := bits[i-1] + v; isForbidden(c) {
				hasForbidden = true
			}
		}
	}

	return (countVowels >= 3) && doubleLetter && !hasForbidden
}

func twiceNoOverlap(bits []string) bool {
	pairs := map[string][]int{}

	for i, v := range bits {
		if i == 0 {
			continue
		}

		x := bits[i-1] + v
		ar, ok := pairs[x]
		if !ok {
			ar = []int{}
		}
		ar = append(ar, i-1)
		pairs[x] = ar
	}

	for p, locs := range pairs {
		if len(locs) < 2 {
			delete(pairs, p)
		}
	}

	if len(pairs) < 1 {
		return false
	}

	isGood := true
	for _, locs := range pairs {
		if len(locs) == 1 {
			continue
		}

		for i, v := range locs {
			if i == 0 {
				continue
			}

			prev := locs[i-1]
			if prev+2 > v {
				isGood = false
			}
		}
	}

	return isGood
}

func repeatWithGap(bits []string) bool {
	if len(bits) < 3 {
		return false
	}

	if len(bits) == 3 {
		return bits[0] == bits[2]
	}

	for i, v := range bits {
		if i < 2 {
			continue
		}

		if bits[i-2] == v {
			return true
		}
	}

	return false
}

// NiceStringV2 ...
func NiceStringV2(in string) bool {
	in = strings.ToLower(in)
	bits := strings.Split(in, "")

	tw := twiceNoOverlap(bits)
	rg := repeatWithGap(bits)

	return tw && rg
}
