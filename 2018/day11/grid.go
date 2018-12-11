package day11

import (
	"fmt"
)

const gridSize = 300

type Grid struct {
	Serial int

	grid [][]*FuelCell
}

func SetupGrid(serial int) *Grid {
	g := &Grid{Serial: serial}
	g.grid = make([][]*FuelCell, gridSize)

	for i := 0; i < gridSize; i++ {
		g.grid[i] = make([]*FuelCell, gridSize)
		for j := 0; j < gridSize; j++ {
			fc := CreateFuelCell(j, i, g.Serial)
			g.grid[i][j] = fc
		}
	}

	return g
}

// CellAt ...
func (g Grid) CellAt(x, y int) *FuelCell {
	if y >= gridSize {
		return nil
	}

	if x >= gridSize {
		return nil
	}

	return g.grid[y][x]
}

// FindLargest ...
func (g *Grid) FindLargest(size int) (int, int, int) {
	var x, y, totalPower int

	for i := 0; i < gridSize-size; i++ {
		for j := 0; j < gridSize-size; j++ {
			sum := g.SumPowers(j, i, size)
			if sum > totalPower {
				x = j
				y = i
				totalPower = sum
				// fmt.Printf("largets power so far: <%v,%v> -> %v\n", x, y, totalPower)
			}
		}
	}

	return x, y, totalPower
}

// FindLargest ...
func (g *Grid) FindLargestGrid() (int, int, int, int) {
	var x, y, power, size int

	for s := 2; s < gridSize; s++ {
		fmt.Printf("checking sub grid of size %vx%v\n", s, s)
		for i := 0; i < gridSize-s; i++ {
			for j := 0; j < gridSize-s; j++ {
				sum := g.SumPowers(j, i, s)
				if sum > power {
					x = j
					y = i
					power = sum
					size = s
				}
			}
		}
		fmt.Printf("largest so far: <%v,%v>x%v -> %v\n", x, y, size, power)
	}

	return x, y, power, size
}

// SumPowers ...
func (g Grid) SumPowers(x, y, s int) int {
	sum := 0
	for i := y; i < y+s; i++ {
		for j := x; j < x+s; j++ {
			sum += g.grid[i][j].power
		}
	}
	return sum
}

func maxKadane(list []*FuelCell) (int, []*FuelCell) {
	maxSoFar := list[0].power
	maxEndingHere := list[0].power
	subArrayStart := 0
	subArrayEnd := 0

	for i, x := range list {
		maxEndingHere = max(x.power, maxEndingHere+x.power)
		if maxEndingHere == x.power {
			subArrayStart = i
		}

		maxSoFar = max(maxSoFar, maxEndingHere)
		if maxSoFar == maxEndingHere {
			subArrayEnd = i
		}
	}

	start := subArrayStart
	end := subArrayEnd + 1
	if start > end {
		start = subArrayEnd + 1
		end = subArrayStart
	}

	if end-start <= 10 {
		fmt.Printf("start: %v, end: %v, diff: %v, len: %v\n", start, end, end-start, len(list))
	}

	return subArrayStart, list[start:end]
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
