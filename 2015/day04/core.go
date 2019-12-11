package day04

import (
	"crypto/md5"
	"fmt"
	"io"
	"strings"
)

func startZeroes(in string, i int) bool {
	if len(in) < i {
		return false
	}

	bits := strings.Split(in, "")
	good := true
	for _, v := range bits[:i] {
		good = good && v == "0"
	}
	return good
}

// FindHash returns the first int value that produces a hash that starts
// with five zeros when the key is paried with the input.
//
// For example, the input 'abcdef' will return (609043, "000001dbbfa3a5c83a2d506429c7b00e")
func FindHash(key string, start int) (int, string) {
	for i := 1; i > 0; i++ {
		h := md5.New()
		io.WriteString(h, fmt.Sprintf("%v%v", key, i))
		bits := h.Sum(nil)
		test := fmt.Sprintf("%x", bits)

		if startZeroes(test, start) {
			return i, test
		}
	}
	return 0, ""
}
