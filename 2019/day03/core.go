package day03

import (
	"math"
	"strconv"
	"strings"

	geo "github.com/paulmach/go.geo"
)

/*
line: R8,U5,L5,D3
Start: 0,0 ( all start here )
   R8: 8,0
   U5: 8,5
   L5: 3,5
   D3: 3,2

line: U7,R6,D4,L4
Start: 0,0
   U7: 0,7
   R6: 6,7
   D4: 6,3
   L4: 2,3
*/

type Breadboard struct {
	wires []*geo.Path
}

type Point struct {
	X float64
	Y float64
}

// DistFromOrigin ...
func (p Point) DistFromOrigin() int {
	x := int(math.Abs(p.X))
	y := int(math.Abs(p.Y))
	return x + y
}

func (b *Breadboard) ReadWire(in string) {
	var x, y float64

	w := geo.NewPath()
	w.Push(geo.NewPoint(x, y))

	parts := strings.Split(in, ",")

	for _, v := range parts {
		movX, movY := parse(v)
		x += movX
		y += movY

		w.Push(geo.NewPoint(x, y))
	}

	b.wires = append(b.wires, w)
}

func parse(in string) (x, y float64) {
	var dir, tmp string

	for i, v := range in {
		if i == 0 {
			dir = string(v)
			continue
		}
		tmp += string(v)
	}

	f, err := strconv.ParseFloat(tmp, 64)
	if err != nil {
		return 0, 0
	}

	switch dir {
	case "U":
		return 0, f
	case "D":
		return 0, f * -1
	case "L":
		return f * -1, 0
	case "R":
		return f, 0
	}

	return 0, 0
}

// Intersections ...
func (b Breadboard) Intersections() []Point {
	out := []Point{}
	wireA := b.wires[0]
	wireB := b.wires[1]

	points, _ := wireA.IntersectionPath(wireB)

	for _, p := range points {
		x := p.X()
		y := p.Y()
		if x == 0 && y == 0 {
			continue
		}

		if math.IsInf(x, 64) || math.IsInf(x, 64) {
			continue
		}

		out = append(out, Point{x, y})
	}

	return out
}

// Steps ...
func (b Breadboard) Steps() int {
	wireA := b.wires[0]
	wireB := b.wires[1]

	points, _ := wireA.IntersectionPath(wireB)

	dist := 99999999999999999

	for _, p := range points {
		x := p.X()
		y := p.Y()
		if x == 0 && y == 0 {
			continue
		}

		if math.IsInf(x, 64) || math.IsInf(x, 64) {
			continue
		}

		d1 := wireA.Measure(p)
		d2 := wireB.Measure(p)

		if dist > int(d1)+int(d2) {
			dist = int(d1 + d2)
		}
	}
	return dist
}

// ManDist ...
func (b Breadboard) ManDist() int {
	out := 999999999999999999

	points := b.Intersections()

	for _, p := range points {
		d := p.DistFromOrigin()
		if d < out {
			out = d
		}
	}

	return out
}
