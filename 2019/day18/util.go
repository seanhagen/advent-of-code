package day18

import (
	"strings"
)

// inputToData ...
func inputToData(i string) map[int]map[int]string {
	out := map[int]map[int]string{}
	i = strings.TrimSpace(i)
	if i == "" {
		return out
	}

	x, y := 0, 0
	tmp := map[int]string{}
	for _, v := range strings.Split(i, "") {
		if v == "\n" {
			x = 0
			out[y] = tmp
			y++
			tmp = map[int]string{}
			continue
		}
		tmp[x] = v
		x++
	}
	out[y] = tmp
	return out
}
