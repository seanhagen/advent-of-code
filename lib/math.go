package lib

// Abs returns the absolute value of the integer
func Abs(a int) int {
	if a < 0 {
		a *= -1
	}
	return a
}

// slightly faster
func Abs2(n int) int {
	x := int64(n)
	y := x >> 63
	return int((x ^ y) - y)
}
