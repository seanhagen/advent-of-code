package day10

import (
	"math"
	"strings"

	geo "github.com/paulmach/go.geo"
)

const TypeEmpty = "."
const TypeAsteroid = "#"

type Asteroid struct {
	X   float64
	Y   float64
	See int
}

type StarField struct {
	toids []*geo.Point
}

func NewStarField(input string) (*StarField, error) {
	rows := strings.Split(input, "\n")
	a := []*geo.Point{}

	for y, r := range rows {
		cols := strings.Split(r, "")
		for x, v := range cols {
			if v == TypeAsteroid {
				a = append(a, geo.NewPoint(float64(x), float64(y)))
			}
		}
	}

	return &StarField{a}, nil
}

func inPointSlice(a *geo.Point, sl []*geo.Point) bool {
	for _, v := range sl {
		if a.Equals(v) {
			return true
		}
	}
	return false
}

// FindBest finds the best asteroid for the base and a count
// of how many other asteroids it can see
func (sf StarField) FindBest() *Asteroid {
	done := []*Asteroid{}

	// fmt.Printf("there are %v points total\n", len(sf.toids))

	for _, p1 := range sf.toids {
		// if p1.X() != 4 || p1.Y() != 2 {
		// 	continue
		// }

		// p1 := geo.NewPoint(a.x, a.y)

		pointsUsed := []*geo.Point{}
		lines := []*geo.Line{}

		for _, p2 := range sf.toids {
			if !p1.Equals(p2) {
				l := geo.NewLine(p1, p2)
				lines = append(lines, l)
				pointsUsed = append(pointsUsed, p2)
			}
		}

		// spew.Dump(pointsUsed, lines)
		// os.Exit(1)

		blockingPoints := []*geo.Point{}
		blockedLines := []*geo.Line{}
		blockedPoints := []*geo.Point{}

		for _, l := range lines {
			for _, p2 := range pointsUsed {
				if !p2.Equals(l.A()) && !p2.Equals(l.B()) && onLine(p2, l) {
					// fmt.Printf("line %v->%v blocked by %v\n", l.A(), l.B(), p2)
					blockedLines = append(blockedLines, l)
					blockingPoints = append(blockingPoints, p2)
					blockedPoints = append(blockedPoints, l.B())
					goto lineBlocked
				}
			}

		lineBlocked:

			// // fmt.Printf("checking line: %#v\n", l)
			// for _, b := range sf.toids {
			// 	if a != b {
			// 		p1 := geo.NewPoint(b.x, b.y)
			// 		//p1.Equals(point *Point)

			// 		if !p1.Equals(l.A()) && !p1.Equals(l.B()) && onLine(p1, l) {
			// 			fmt.Printf("line %v->%v is blocked by %v\n", l.A(), l.B(), p1)

			// 			blockedLines = append(blockedLines, l)
			// 			blockingPoints = append(blockingPoints, p1)
			// 		}
			// 	}
			// }
		}

		see := len(sf.toids) - len(blockedPoints) - 1

		// fmt.Printf("point %v can see %v other points\n", p1, see)
		// fmt.Printf("\n\n")

		done = append(done, &Asteroid{p1.X(), p1.Y(), see})
		// spew.Dump(blockingPoints, blockedLines)
		// os.Exit(1)
	}

	// spew.Dump(done)
	// os.Exit(1)

	most := 0
	var best *Asteroid

	for _, v := range done {
		if v.See > most {
			// fmt.Printf("astroid %#v can see more than %v\n", v, most)
			best = v
			most = v.See
		}
	}

	// fmt.Printf("best: %#v\n", best)

	return best
}

func inAsteroidSlice(a *geo.Point, sl []*geo.Point) bool {
	for _, v := range sl {
		if a.Equals(v) {
			return true
		}
	}
	return false
}

func onLine(c *geo.Point, l *geo.Line) bool {
	a := l.A()
	b := l.B()

	ax := a.X()
	ay := a.Y()

	bx := b.X()
	by := b.Y()

	cx := c.X()
	cy := c.Y()

	eps := 0.0001

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
