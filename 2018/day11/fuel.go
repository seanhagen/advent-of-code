package day11

import "math"

type FuelCell struct {
	x int
	y int

	power int
}

func CreateFuelCell(x, y, serial int) *FuelCell {
	rackID := x + 10

	pl := rackID * y
	pl += serial
	pl *= rackID
	pl = digit(pl, 3)
	pl -= 5

	return &FuelCell{x: x, y: y, power: pl}
}

func digit(num, place int) int {
	r := num % int(math.Pow(10, float64(place)))
	return r / int(math.Pow(10, float64(place-1)))
}
