package lib

// Abs returns the absolute value of the integer
func Abs(a int) int {
	if a < 0 {
		a *= -1
	}
	return a
}
