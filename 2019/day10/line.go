package day10

import "math"

const rotateDeg = 0.00001
const eps = 0.00001

// Line is a straight line, drawn between two asteroids
type Line struct {
	a *Asteroid
	b *Asteroid
}

// NewLine ...
func NewLine(a, b *Asteroid) *Line {
	return &Line{a, b}
}

// OnLine ...
func (a Line) OnLine(p1 *Asteroid) bool {
	ax := a.a.X
	ay := a.a.Y

	bx := a.b.X
	by := a.b.Y

	cx := p1.X
	cy := p1.Y

	cp := (cy-ay)*(bx-ax) - (cx-ax)*(by-ay)
	if math.Abs(cp) > eps {
		return false
	}

	dp := (cx-ax)*(bx-ax) + (cy-ay)*(by-ay)
	if dp < 0 {
		return false
	}

	slba := (bx-ax)*(bx-ax) + (by-ay)*(by-ay)
	if dp > slba {
		return false
	}

	return true
}

// Rotate ...
func (a Line) Rotate() {
	cos := math.Cos(rotateDeg)
	sinP := math.Sin(rotateDeg)
	sinN := -1 * sinP

	x1 := a.a.X
	y1 := a.a.Y

	x2 := a.b.X
	y2 := a.b.Y

	xt := x2 - x1
	yt := y2 - y1

	/*
	   Matrix math!
	    [x3] [cosθ  sinθ][x2 - x1] [x1]
	    [y3]=[-sinθ cosθ][y2 - y1]+[y1]


	    [a b]   [e]
	    [c d] * [f]

	    [a*e + b*f]
	    [c*e + d*f]

	    [cosO * xt + sinO*yt]
	    [-sinO * xt + cosO*yt]
	*/

	x3 := (cos*xt + sinN*yt) + x1
	y3 := (sinP*xt + cos*yt) + y1

	a.b.X = x3
	a.b.Y = y3
}
