package day12

import (
	"io"

	"github.com/seanhagen/advent-of-code/lib"
)

// Game ...
type Game struct {
	current *Generation
	history map[int]*Generation

	rules []Rule

	diff       int
	cg         int
	currentSum int
}

// setup ...
func (g *Game) setup(in string) {
	g.current = createGeneration(in)
	g.history = map[int]*Generation{}
	g.rules = []Rule{}
	g.diff = 0
	g.cg = 0
	g.currentSum = g.current.count()
}

// addRule ...
func (g *Game) addRule(in string) {
	r := createRule(in)
	g.rules = append(g.rules, r)
}

// Step ...
func (g *Game) Step() {
	g.history[g.current.id] = g.current
	old := g.current.count()
	g.current = g.current.Next(g.rules)
	cur := g.current.count()

	diff := cur - old

	if diff == g.diff {
		g.cg++
	} else {
		g.cg = 0
	}
	g.diff = diff
	g.currentSum = cur
}

// TakeSteps ...
func (g *Game) TakeSteps(i int) {
	for j := 0; j < i; j++ {
		if g.cg > 10 {
			g.currentSum += g.diff * (i - j)
			j = i + 1
		} else {
			g.Step()
		}
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
	return g.currentSum //g.current.count()
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
