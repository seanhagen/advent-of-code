package day10

import (
	"math"
	"testing"
)

func TestLaserRotate(t *testing.T) {
	newX := 3.06999883331
	newY := 9.99965000287

	p1, err := NewAsteroid(3, 3)
	if err != nil {
		t.Fatalf("unable to create asteroid: %v", err)
	}
	p2, err := NewAsteroid(3, 10)
	if err != nil {
		t.Fatalf("unable to create asteroid: %v", err)
	}

	l := Line{p1, p2}
	l.Rotate()

	x3 := math.Abs(l.b.X - p2.X)
	y3 := math.Abs(l.b.Y - p2.Y)

	if x3 > eps || y3 > eps {
		t.Errorf("invalid new coords, expected (%v,%v), got (%v,%v)", newX, newY, l.b.X, l.b.Y)
	}
}
