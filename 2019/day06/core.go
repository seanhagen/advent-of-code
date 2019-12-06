package day06

import (
	"strings"
)

// Person ...
type Person struct {
	orbiting *Planet
	name     string
}

// Planet ...
type Planet struct {
	com *COM

	children []*Planet

	orbitName string
	orbits    *Planet

	name string
}

// TotalOrbits ...
func (p Planet) TotalOrbits() int {
	o := 1

	x := p.orbits
	for {
		if x == nil {
			break
		}
		x = x.orbits
		o++
	}

	return o
}

type COM struct {
	planets  map[string]*Planet
	children []*Planet

	people map[string]*Person
}

// CalcTotalOrbits ...
func (c COM) CalcTotalOrbits() int {
	t := 0
	for _, p := range c.planets {
		i := p.TotalOrbits()
		t += i
	}
	return t
}

// AddPlanet ...
func (c *COM) AddPlanet(in string) {
	bits := strings.Split(in, ")")
	orbits := bits[0]
	name := bits[1]

	if name == "YOU" || name == "SAN" {
		p := Person{name: name}

		if o, ok := c.planets[orbits]; ok {
			p.orbiting = o
		} else {
			o := &Planet{name: orbits, com: c}
			c.planets[orbits] = o
			p.orbiting = o
		}

		c.people[name] = &p
		return
	}

	var planet *Planet
	if e, ok := c.planets[name]; ok {
		planet = e
		planet.orbitName = orbits
	} else {
		planet = &Planet{name: name, com: c, orbitName: orbits}
	}
	c.planets[name] = planet

	if orbits != "COM" {
		o, ok := c.planets[orbits]
		if !ok {
			o = &Planet{name: orbits, com: c}
			c.planets[orbits] = o
		}

		if o.orbits == nil {
			if f, ok := c.planets[o.orbitName]; ok {
				o.orbits = f
			}
		}

		planet.orbits = o
		planet.orbitName = orbits

		o.children = append(o.children, planet)
	} else {
		c.children = append(c.children, planet)
	}
}

func CreateOrbitalComputer() *COM {
	return &COM{
		planets:  map[string]*Planet{},
		children: []*Planet{},
		people:   map[string]*Person{},
	}
}

/*
        2   3       5   6   7
        G - H       J - K - L
       /      3    /
COM - B - C - D - E - F
      1   2    \  4   5
                I
                4

1+2+2+3+3+4+4+5+5+6+7=42
*/

// RouteMeToSanta ...
func (c COM) MinOrbitalTransfers() int {
	me, ok := c.people["YOU"]
	if !ok {
		return 0
	}

	san, ok := c.people["SAN"]
	if !ok {
		return 0
	}

	youPath := []*Planet{}
	sanPath := []*Planet{}

	p := me.orbiting
	for {
		youPath = append(youPath, p)
		if p.orbits == nil {
			break
		}
		p = p.orbits
	}

	p = san.orbiting
	for {
		sanPath = append(sanPath, p)
		if p.orbits == nil {
			break
		}
		p = p.orbits
	}

	var longest, other []*Planet
	if len(sanPath) > len(youPath) {
		longest = sanPath
		other = youPath
	} else {
		longest = youPath
		other = sanPath
	}

	j := len(other)
	x := len(longest)
	for x >= 0 && j >= 0 {
		x--
		j--
		if longest[x].name != other[j].name {
			break
		}
	}
	return (j + 1) + (x + 1)
}
