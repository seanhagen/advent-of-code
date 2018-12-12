package day12

import (
	"fmt"
	"testing"
)

func TestInitialState(t *testing.T) {
	expect := " 0: ...#..#.#..##......###...###..........."
	g := SetupGame("./test.txt")
	out := g.OutputCurrent()

	if expect != out {
		t.Errorf("wrong output!\n\texpect:\t%v\n\tgot:\t%v\n\n", expect, out)
	}
}

func TestInitialSum(t *testing.T) {
	g := SetupGame("./test.txt")

	expect := 145
	out := g.SumCurrent()

	if expect != out {
		t.Errorf("wrong sum! expected %v, got %v", expect, out)
	}
}

func TestSecondGeneration(t *testing.T) {
	expect := " 1: ...#...#....#.....#..#..#..#..........."
	g := SetupGame("./test.txt")
	g.Step()

	out := g.OutputCurrent()
	if expect != out {
		t.Errorf("wrong output!\n\texpect:\t%v\n\tgot:\t%v\n\n", expect, out)
	}
}

func TestAllSteps(t *testing.T) {
	tests := []struct {
		should string
		steps  int
	}{
		{should: " 2: ...##..##...##....#..#..#..##..........", steps: 2},
		{should: "10: ..#.#..#...#.##....##..##..##..##......", steps: 10},
		{should: "20: .#....##....#####...#######....#.#..##.", steps: 20},
	}

	for _, tt := range tests {
		g := SetupGame("./test.txt")
		g.TakeSteps(tt.steps)
		got := g.OutputCurrent()
		if got != tt.should {
			t.Errorf("wrong output for step %v\nexpected:\t%v\ngot:\t\t%v", tt.steps, tt.should, got)
		}
	}
}

func TestFullOutput(t *testing.T) {
	should := ` 0: ...#..#.#..##......###...###...........
 1: ...#...#....#.....#..#..#..#...........
 2: ...##..##...##....#..#..#..##..........
 3: ..#.#...#..#.#....#..#..#...#..........
 4: ...#.#..#...#.#...#..#..##..##.........
 5: ....#...##...#.#..#..#...#...#.........
 6: ....##.#.#....#...#..##..##..##........
 7: ...#..###.#...##..#...#...#...#........
 8: ...#....##.#.#.#..##..##..##..##.......
 9: ...##..#..#####....#...#...#...#.......
10: ..#.#..#...#.##....##..##..##..##......
11: ...#...##...#.#...#.#...#...#...#......
12: ...##.#.#....#.#...#.#..##..##..##.....
13: ..#..###.#....#.#...#....#...#...#.....
14: ..#....##.#....#.#..##...##..##..##....
15: ..##..#..#.#....#....#..#.#...#...#....
16: .#.#..#...#.#...##...#...#.#..##..##...
17: ..#...##...#.#.#.#...##...#....#...#...
18: ..##.#.#....#####.#.#.#...##...##..##..
19: .#..###.#..#.#.#######.#.#.#..#.#...#..
20: .#....##....#####...#######....#.#..##.`

	g := SetupGame("./test.txt")
	got := g.OutputCurrent()

	for i := 1; i <= 20; i++ {
		g.Step()
		got = fmt.Sprintf("%v\n%v", got, g.OutputCurrent())
	}

	if should != got {
		t.Errorf("wrong output!\nexpected:\n%v\n\ngot:\n%v\n\n", should, got)
	}
}

func TestStep20Count(t *testing.T) {
	expect := 325

	g := SetupGame("./test.txt")
	g.TakeSteps(20)

	got := g.SumCurrent()

	if got != expect {
		t.Errorf("got wrong answer! expected: %v, got: %v", expect, got)
	}
}
