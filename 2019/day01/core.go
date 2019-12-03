package day1

import (
	"math"
)

// CalcFuel ...
func CalcFuel(in int) int {
	a := float64(in) / 3
	b := math.Floor(a)
	return int(b) - 2
}

func CalcFuelRecurse(in int) int {
	totalFuel := 0
	for {
		o := CalcFuel(in)
		if o <= 0 {
			return totalFuel
		}
		totalFuel += o
		in = o
	}

	return 0
}
