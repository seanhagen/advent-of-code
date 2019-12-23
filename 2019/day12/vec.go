package day12

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/seanhagen/advent-of-code/lib"
)

// Vec represents a 3d vector of x,y,z
type Vec struct {
	x int
	y int
	z int
}

// ParseVecString ...
func ParseVecString(in string) (Vec, error) {
	in = strings.Replace(in, "<", "", 1)
	in = strings.Replace(in, ">", "", 1)
	in = strings.Replace(in, " ", "", -1)
	parts := strings.Split(in, ",")

	pos := Vec{}
	if len(parts) != 3 {
		return pos, fmt.Errorf("need x,y,z coords, not enough parts")
	}

	xset, yset, zset := false, false, false
	for i, p := range parts {
		bits := strings.Split(p, "=")
		if len(bits) != 2 {
			return pos, fmt.Errorf("coord %v not formatted correctly, should be like 'x=0', got %#v", i, p)
		}

		v, err := strconv.Atoi(bits[1])
		if err != nil {
			return pos, fmt.Errorf("unable to parse coordinate: %v", err)
		}

		switch strings.ToLower(bits[0]) {
		case "x":
			xset = true
			pos.x = v
		case "y":
			yset = true
			pos.y = v
		case "z":
			zset = true
			pos.z = v
		default:
			return pos, fmt.Errorf("only x,y or z can be set, got '%v'", bits[0])
		}
	}

	if !xset {
		return pos, fmt.Errorf("need x coordinate")
	}

	if !yset {
		return pos, fmt.Errorf("need y coordinate")
	}

	if !zset {
		return pos, fmt.Errorf("need z coordinate")
	}

	return pos, nil
}

// Add ...
func (v *Vec) Add(v2 Vec) {
	v.x += v2.x
	v.y += v2.y
	v.z += v2.z
}

// Eq ...
func (v Vec) Eq(v2 Vec) bool {
	return v.x == v2.x &&
		v.y == v2.y &&
		v.z == v2.z
}

// Energy ...
func (v Vec) Energy() int {
	return lib.Abs(v.x) + lib.Abs(v.y) + lib.Abs(v.z)
}
