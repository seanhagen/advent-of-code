package day06

import (
	"fmt"
	"testing"
)

func TestTotalOrbits(t *testing.T) {
	com := COM{planets: map[string]*Planet{}}
	a := Planet{com: &com, name: "A"}
	com.planets["A"] = &a

	i := a.TotalOrbits()
	if i != 1 {
		t.Errorf("a total orbits should be 1, got %v", i)
	}

	b := Planet{orbits: &a, name: "B", com: &com}
	a.children = append(a.children, &b)
	com.planets["B"] = &b

	i = b.TotalOrbits()
	if i != 2 {
		t.Errorf("b total orbits should be 2, got %v", i)
	}

	c := Planet{orbits: &a, name: "C", com: &com}
	a.children = append(a.children, &c)
	com.planets["C"] = &c

	d := Planet{orbits: &c, name: "D", com: &com}
	c.children = append(c.children, &d)
	com.planets["D"] = &d

	x := com.CalcTotalOrbits()
	if x != 8 {
		t.Errorf("com CalcTotalOrbits wrong output, expected 8 got '%v'", x)
	}
}

func TestAddPlanets(t *testing.T) {
	tests := []struct {
		input []string
		total int
	}{
		{[]string{"COM)B", "B)C", "C)D", "D)E", "E)F", "B)G", "G)H", "D)I", "E)J", "J)K", "K)L"}, 42},
		{[]string{"COM)A", "A)B", "A)C", "C)D"}, 8},
		{[]string{"C)D", "A)B", "A)C", "COM)A"}, 8},
	}

	for i, tt := range tests {
		x := tt
		t.Run(fmt.Sprintf("test %v", i), func(t *testing.T) {
			c := CreateOrbitalComputer()
			for _, p := range x.input {
				c.AddPlanet(p)
			}

			tc := c.CalcTotalOrbits()
			if tc != x.total {
				t.Errorf("wrong total orbits, expected '%v' got '%v'", x.total, tc)
			}
		})
	}
}

func TestMinOrbitalTransfers(t *testing.T) {
	tests := []struct {
		input []string
		total int
	}{
		{[]string{"COM)B", "B)C", "C)D", "D)E", "E)F", "B)G", "G)H", "D)I", "E)J", "J)K", "K)L", "K)YOU", "I)SAN"}, 4},
	}

	for i, tt := range tests {
		x := tt
		t.Run(fmt.Sprintf("test %v", i), func(t *testing.T) {
			c := CreateOrbitalComputer()
			for _, p := range x.input {
				c.AddPlanet(p)
			}

			tc := c.MinOrbitalTransfers()
			if tc != x.total {
				t.Errorf("wrong min transfers, expected '%v' got '%v'", x.total, tc)
			}
		})
	}
}
