package day10

import (
	"fmt"
	"math"

	"github.com/gofrs/uuid"
)

// Asteroid conatins the X,Y coordinates of where an asteroid is on the map,
// and will also contain how many asteroids it can see when returned from
// the `StarField.FindBest` method
type Asteroid struct {
	id  uuid.UUID
	X   float64
	Y   float64
	See int
}

// NewAsteroid returns an asteroid for the given integer x,y coordinates
func NewAsteroid(x, y int) (*Asteroid, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	return &Asteroid{id: id, X: float64(x), Y: float64(y)}, nil
}

// Equals ...
func (a Asteroid) Equals(a1 *Asteroid) bool {
	cX := math.Abs(a.X-a1.X) < eps
	cY := math.Abs(a.Y-a1.Y) < eps

	return cX && cY || a.id == a1.id
}

// Distance ...
func (a Asteroid) Distance(a1 *Asteroid) float64 {
	d0 := a1.X - a.X
	d1 := a1.Y - a.Y
	return math.Sqrt(d0*d0 + d1*d1)
}

// LaserLine ...
func (a *Asteroid) LaserLine(len float64) (*Line, error) {
	x2 := a.X
	y2 := a.Y - len
	// fmt.Printf("laser line starts at %v,%v\n", a.X, a.Y)
	// fmt.Printf("laser line ends at %v,%v\n", x2, y2)

	id, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	p2 := &Asteroid{id, x2, y2, 0}
	return &Line{a, p2}, nil
}

// AngleTo ...
func (a Asteroid) AngleTo(b *Asteroid) float64 {
	x := a.X - b.X
	y := a.Y - b.Y
	return math.Atan2(y, x)
}

// String ...
func (a Asteroid) String() string {
	return fmt.Sprintf("Asteroid{id: %v, X: %v, Y: %v, See: %v}", a.id.String(), a.X, a.Y, a.See)
}

// GoString ...
func (a Asteroid) GoString() string {
	return fmt.Sprintf("Asteroid{id: %v, X: %v, Y: %v, See: %v}", a.id.String(), a.X, a.Y, a.See)
}
