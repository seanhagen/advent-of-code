package day12

import (
	"fmt"
	"io"

	"github.com/seanhagen/advent-of-code/lib"
)

// Game ...
type Game struct {
	current *Generation
	history map[int]*Generation

	rules []Rule
}

// setup ...
func (g *Game) setup(in string) {
	g.current = createGeneration(in)
	g.history = map[int]*Generation{}
	g.rules = []Rule{}
}

// addRule ...
func (g *Game) addRule(in string) {
	r := createRule(in)
	g.rules = append(g.rules, r)
}

// Step ...
func (g *Game) Step() {
	g.history[g.current.id] = g.current

	fmt.Printf("sum: %v\n", g.current.count())

	g.current = g.current.Next(g.rules)
}

// TakeSteps ...
func (g *Game) TakeSteps(i int) {
	for j := 0; j < i; j++ {
		g.Step()
	}
}

// OutputCurrent ...
func (g Game) OutputCurrent() string {
	if g.current == nil {
		return ""
	}
	return g.current.output()
}

// CountCurrent ...
func (g Game) SumCurrent() int {
	return g.current.count()
}

// SetupGame ...
func SetupGame(path string) *Game {
	f := lib.LoadInput(path)

	g := &Game{}

	cnt := 0
	err := lib.LoopOverLines(f, func(line []byte) error {
		l := string(line)

		switch cnt {
		case 0:
			g.setup(l)
		case 1:
		default:
			g.addRule(l)
		}
		cnt++

		return nil
	})

	if err != io.EOF {
		panic(err)
	}

	return g
}
