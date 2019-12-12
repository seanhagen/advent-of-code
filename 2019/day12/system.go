package day12

import (
	"bytes"
	"fmt"
)

// System ...
type System struct {
	planets []*Planet

	pairs [][]*Planet
}

// NewSystem ...
func NewSystem() *System {
	return &System{planets: []*Planet{}, pairs: [][]*Planet{}}
}

// AddPlanet ...
func (s *System) AddPlanet(in string) error {
	p, err := CreatePlanet(len(s.planets), in)
	if err != nil {
		return err
	}

	s.planets = append(s.planets, p)
	return nil
}

// SetupPairs ...
func (s *System) SetupPairs() error {
	// cribbed from https://github.com/mxschmitt/golang-combinations/blob/master/combinations.go
	pairs := [][]*Planet{}
	ssmax := 2

	length := uint(len(s.planets))

	for sb := 1; sb < (1 << length); sb++ {
		subset := []*Planet{}

		for obj := uint(0); obj < length; obj++ {
			if (sb>>obj)&1 == 1 {
				subset = append(subset, s.planets[obj])
			}
		}

		if len(subset) == ssmax {
			pairs = append(pairs, subset)
		}
	}

	s.pairs = pairs

	return nil
}

// PrintPairs ...
func (s System) PrintPairs() {
	for _, pair := range s.pairs {
		fmt.Printf("%v <-> %v\n", pair[0].name, pair[1].name)
	}
}

// Output ...
func (s System) Output() string {
	buf := bytes.NewBufferString("")

	for _, p := range s.planets {
		fmt.Fprintf(buf, "%v\n", p)
	}

	return buf.String()
}

// DoSteps ...
func (s *System) DoSteps(st int) {
	for i := 0; i < st; i++ {
		for _, pair := range s.pairs {
			CalculateVelocity(pair[0], pair[1])
		}

		for _, p := range s.planets {
			p.Step()
		}
		// fmt.Printf("done step %v\n", i)
		// spew.Dump(s.planets)
	}
}

// CalculateTotalEnergy ...
func (s *System) CalculateTotalEnergy() int {
	sum := 0
	for _, p := range s.planets {
		sum += p.TotalEnergy()
	}
	return sum
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lcm(a, b int, integers ...int) int {
	res := a * b / gcd(a, b)
	for i := 0; i < len(integers); i++ {
		res = lcm(res, integers[i])
	}
	return res
}

// DoStepsUntilReturned ...
func (s *System) DoStepsUntilReturned() int {
	// cribbed from https://github.com/XorZy/Aoc_2019_Day12/blob/master/Program_PartB.cs because
	// I couldn't figure it out ( brute force would.... uh, definitely not work )
	i := 0

	xc, yc, zc := -1, -1, -1

	for {
		for _, pair := range s.pairs {
			CalculateVelocity(pair[0], pair[1])
		}

		for _, p := range s.planets {
			p.Step()
		}

		i++
		if xc == -1 && s.AllPlanets(xZero) {
			xc = i
		}

		if yc == -1 && s.AllPlanets(yZero) {
			yc = i
		}

		if zc == -1 && s.AllPlanets(zZero) {
			zc = i
		}

		if xc > 0 && yc > 0 && zc > 0 {
			break
		}
	}

	test := lcm(xc, yc, zc)
	return test * 2
}

func xZero(p *Planet) bool {
	return p.vel.x == 0
}

func yZero(p *Planet) bool {
	return p.vel.y == 0
}

func zZero(p *Planet) bool {
	return p.vel.z == 0
}

// AllPlanets ...
func (s System) AllPlanets(fn func(p *Planet) bool) bool {
	o := true
	for _, p := range s.planets {
		v := fn(p)
		o = o && v
	}
	return o
}
