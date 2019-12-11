package day10

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

// TypeAsteroid contains the character used to denote an asteroid in a map
const TypeAsteroid = "#"

// StarField contains all the asteroids from a map
type StarField struct {
	toids     []*Asteroid
	station   *Asteroid
	destroyed []*Asteroid

	height int
	width  int
}

// NewStarField takes a map string and returns a struct
// that contains all the asteroids found on the map.
func NewStarField(input string) (*StarField, error) {
	rows := strings.Split(input, "\n")
	a := []*Asteroid{}

	h, w := 0, 0

	for y, r := range rows {
		cols := strings.Split(r, "")
		for x, v := range cols {
			if v == TypeAsteroid {
				ast, err := NewAsteroid(x, y)
				if err != nil {
					return nil, err
				}

				a = append(a, ast)
			}
			w = x
		}
		h = y
	}

	return &StarField{toids: a, destroyed: []*Asteroid{}, height: h, width: w}, nil
}

// FindBest finds the best asteroid for the base and a count
// of how many other asteroids it can see
func (sf StarField) FindBest() *Asteroid {
	done := []*Asteroid{}

	for _, p1 := range sf.toids {
		pointsUsed := []*Asteroid{}
		lines := []*Line{}

		for _, p2 := range sf.toids {
			if !p1.Equals(p2) {
				l := NewLine(p1, p2)
				lines = append(lines, l)
				pointsUsed = append(pointsUsed, p2)
			}
		}

		blockedPoints := []*Asteroid{}
		for _, l := range lines {
			for _, p2 := range pointsUsed {
				if !p2.Equals(l.a) && !p2.Equals(l.b) && l.OnLine(p2) {
					blockedPoints = append(blockedPoints, l.b)
					goto lineBlocked
				}
			}

		lineBlocked:
		}

		see := len(sf.toids) - len(blockedPoints) - 1
		p1.See = see
		done = append(done, p1)
	}

	most := 0
	var best *Asteroid

	for _, v := range done {
		if v.See > most {
			best = v
			most = v.See
		}
	}
	return best
}

// SetStation ...
func (sf *StarField) SetStation(st *Asteroid) {
	sf.station = st
	toids := []*Asteroid{}
	for _, a := range sf.toids {
		if !st.Equals(a) {
			toids = append(toids, a)
		}
	}
	sf.toids = toids
}

// GetDestroyed ...
func (sf *StarField) GetDestroyed(n int) *Asteroid {
	if n > len(sf.destroyed) {
		return nil
	}

	// fmt.Printf("get destroyed(%v), len: %v\n", n, len(sf.destroyed)) //

	return sf.destroyed[n-1]
}

// LR2 ...
func (sf *StarField) LaserRotation(destroy int) error {
	if sf.station == nil {
		return fmt.Errorf("set a station before calling this function")
	}

	if destroy < 1 {
		return fmt.Errorf("argument must be greater than or equal to 1")
	}

	sorted := map[float64]map[float64]*Asteroid{}

	for _, a := range sf.toids {
		z := sf.station.AngleTo(a)
		d := sf.station.Distance(a)

		ang, ok := sorted[z]
		if !ok {
			ang = map[float64]*Asteroid{}
		}

		ang[d] = a
		sorted[z] = ang
	}

	xdr, ydr := 0, -1
	if sf.station.Y == 0 {
		xdr, ydr = 1, 0
	}

	if sf.station.Y == 0 && sf.station.X == float64(sf.width) {
		xdr, ydr = 0, 1
	}

	startX, startY := sf.station.X, sf.station.Y

	var current *Asteroid
	for {
		startX += float64(xdr)
		startY += float64(ydr)
		if startY <= 0 {
			ydr *= -1
		}

		found := false
		for _, a := range sf.toids {
			if a.X == startX && a.Y == startY {
				current = a
				found = true
				break
			}
		}
		if found {
			break
		}
	}

	angle := sf.station.AngleTo(current)

	destroyed := []*Asteroid{}
	toids := []*Asteroid{}
	for _, a := range sf.toids {
		if !sf.station.Equals(a) {
			// remove station from list of asteroids
			toids = append(toids, a)
		}
	}

	for i := 1; i <= destroy; i++ {
		// remove current from toids
		tmp := []*Asteroid{}
		for _, a := range toids {
			if !current.Equals(a) {
				tmp = append(tmp, a)
			}
		}
		toids = tmp
		// add current to destroyed
		destroyed = append(destroyed, current)

		// find next asteroid set by finding next biggest angle
		angChk := []float64{}
		for a := range sorted {
			if a > angle {
				angChk = append(angChk, a)
			}
		}
		//   (if no next biggest, set angle to -pi then check again)
		if len(angChk) == 0 {
			angle = -1 * math.Pi
			for a := range sorted {
				if a > angle {
					angChk = append(angChk, a)
				}
			}
		}

		// if still zero, no angles left to check!
		if len(angChk) == 0 {
			break
		}

		sort.Float64s(angChk)

		angle = angChk[0]
		aset := sorted[angle]

		dist := 1000000.0
		for d := range aset {
			if d < dist {
				dist = d
			}
		}

		// spew.Dump(aset, dist)
		// os.Exit(1)

		current = aset[dist]
		delete(aset, dist)

		// if set now has zero asteroids, remove set ( delete(sorted[angle]))
		if len(aset) == 0 {
			delete(sorted, angle)
		}

	}

	// spew.Dump(sorted, current, angle, dist, toids, destroyed)

	sf.toids = toids
	sf.destroyed = append(sf.destroyed, destroyed...)

	// os.Exit(1)

	return nil
}

// Reset ...
func (sf *StarField) Reset() {
	asteroids := []*Asteroid{}

	asteroids = append(asteroids, sf.destroyed...)
	asteroids = append(asteroids, sf.station)

	sf.station = nil
	sf.destroyed = []*Asteroid{}
	sf.toids = asteroids
}

// func remove(s []int, i int) []int {
// 	s[i] = s[len(s)-1]
// 	// We do not need to put s[i] at the end, as it will be discarded anyway
// 	return s[:len(s)-1]
// }
