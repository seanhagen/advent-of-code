package day12

import (
	"fmt"
)

var names = []string{
	"Io",
	"Europa",
	"Ganymede",
	"Callisto",
	"Metis",
	"Adrastea",
	"Thebe",
	"Themisto",
	"Leda",
	"Himalia",
	"Ersa",
	"Pandia",
	"Lysithea",
	"Elara",
	"Dia",
	"Carpo",
	"Valetudo",
	"Europie",
}

// Planet ...
type Planet struct {
	name string
	pos  Vec
	vel  Vec
}

// CreatePlanet takes an index and position string
func CreatePlanet(idx int, in string) (*Planet, error) {
	v, err := ParseVecString(in)
	if err != nil {
		return nil, err
	}

	name := fmt.Sprintf("moon %v", idx)
	if idx < len(names)-1 {
		name = names[idx]
	}

	return &Planet{name: name, pos: v}, nil
}

// Step ...
func (p *Planet) Step() {
	p.pos.Add(p.vel)
}

// Eq ...
func (p Planet) Eq(v Vec) bool {
	return p.pos.Eq(v)
}

// String ...
func (p Planet) String() string {
	return p.GoString()
}

// GoString ...
func (p Planet) GoString() string {
	return fmt.Sprintf("pos=<x=%3v, y=%3v, z=%3v>, vel=<x=%3v, y=%3v, z=%3v>", p.pos.x, p.pos.y, p.pos.z, p.vel.x, p.vel.y, p.vel.z)
}

// TotalEnergy ...
func (p Planet) TotalEnergy() int {
	a := p.pos.Energy()
	b := p.vel.Energy()

	return a * b
}
