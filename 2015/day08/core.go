package day08

import (
	"fmt"
	"strconv"
	"strings"
)

func MemLength(in string) int {
	bits := strings.Split(in, "")
	tmp := bits[1 : len(bits)-1]

	if len(tmp) == 0 {
		return 0
	}

	out := []string{}

	idx := 0
	for {
		if idx > len(tmp)-1 {
			break
		}

		v := tmp[idx]
		if v == "\\" && tmp[idx+1] == "x" {
			a := tmp[idx+2] + tmp[idx+3]
			i, err := strconv.ParseInt(a, 16, 0)
			if err != nil {
				return 0
			}
			v = fmt.Sprintf("%c", i)
			idx += 3
		} else if v == "\\" {
			n := tmp[idx+1]

			if n == `\` || n == `"` {
				v = n
				idx += 1
			}
		}
		out = append(out, v)
		idx++
	}
	return len(out)
}

func TotalMem(in []string) int {
	s1, s2 := 0, 0

	for _, v := range in {
		l1 := len(v)
		l2 := MemLength(v)

		s1 += l1
		s2 += l2
	}

	return s1 - s2
}

func EncodeLength(in string) int {
	o := strconv.Quote(in)
	return len(o)
}

func TotalEncodeMem(in []string) int {
	s1, s2 := 0, 0

	for _, v := range in {
		l1 := len(v)
		l2 := EncodeLength(v)

		s1 += l1
		s2 += l2
	}

	return s2 - s1
}
