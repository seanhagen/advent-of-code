package day4

import (
	"errors"
	"fmt"
	"image/color"
	"strings"
)

// ParseFile parses the input file and normalizes it to one line per passport
func ParseFile(in string) []string {
	out := []string{}
	bits := strings.Split(in, "\n")
	tmp := []string{}
	l := ""
	for _, v := range bits {
		if v == "" {
			l = strings.TrimSpace(l)
			tmp = append(tmp, l)
			l = ""
			continue
		}

		if l == "" {
			l = v
		} else {
			l = fmt.Sprintf("%v %v", l, v)
		}
	}
	tmp = append(tmp, l)

	for _, v := range tmp {
		if v != "" {
			out = append(out, v)
		}
	}

	return out
}

var errInvalidFormat = errors.New("invalid format")

func ParseHexColorFast(s string) (c color.RGBA, err error) {
	c.A = 0xff

	if s[0] != '#' {
		return c, errInvalidFormat
	}

	hexToByte := func(b byte) byte {
		switch {
		case b >= '0' && b <= '9':
			return b - '0'
		case b >= 'a' && b <= 'f':
			return b - 'a' + 10
		case b >= 'A' && b <= 'F':
			return b - 'A' + 10
		}
		err = errInvalidFormat
		return 0
	}

	switch len(s) {
	case 7:
		c.R = hexToByte(s[1])<<4 + hexToByte(s[2])
		c.G = hexToByte(s[3])<<4 + hexToByte(s[4])
		c.B = hexToByte(s[5])<<4 + hexToByte(s[6])
	case 4:
		c.R = hexToByte(s[1]) * 17
		c.G = hexToByte(s[2]) * 17
		c.B = hexToByte(s[3]) * 17
	default:
		err = errInvalidFormat
	}
	return
}