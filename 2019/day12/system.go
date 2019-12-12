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
