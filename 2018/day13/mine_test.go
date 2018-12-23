package day13

import (
	"testing"
)

func TestTrackOnly(t *testing.T) {
	expect := `/---\
|   |
\---/
`

	m := SetupMine("track-only.txt")

	got := m.PrintTrack()

	if expect != got {
		t.Errorf("wrong track!\nexpected:\n%v\n\ngot:\n%v\n", expect, got)
	}
}

func TestSimpleTrack(t *testing.T) {
	expect := `/---\
|   |
\---/
`

	m := SetupMine("simple-track.txt")

	got := m.PrintTrack()

	if expect != got {
		t.Errorf("wrong track!\nexpected:\n%v\n\ngot:\n%v\n", expect, got)
	}
}

func TestSimpleTrackWithCart(t *testing.T) {
	expect := `/>--\
|   |
\---/
`

	m := SetupMine("simple-track.txt")

	got := m.Print()

	if expect != got {
		t.Errorf("wrong track!\nexpected:\n%v\n\ngot:\n%v\n", expect, got)
	}
}

func TestSimpleTrackOneStep(t *testing.T) {
	expect := `/->-\
|   |
\---/
`

	m := SetupMine("simple-track.txt")
	m.Step()

	got := m.Print()
	if expect != got {
		t.Errorf("wrong track!\nexpected:\n%v\n\ngot:\n%v\n", expect, got)
	}
}

func TestSimpleCollision(t *testing.T) {
	expect := `---X---
`

	m := SetupMine("simple-collision.txt")
	m.Step()
	m.Step()
	m.Step()

	got := m.Print()
	if expect != got {
		t.Errorf("wrong collision output:\nexpected:\t%v\ngot:\t\t%v\n", expect, got)
	}

	c, x, y := m.CheckCollision()

	if c != true {
		t.Error("should be a collision, got false")
	}

	if x != 3 && y != 0 {
		t.Errorf("wrong position, expected <%v, %v> got <%v, %v>", 3, 0, x, y)
	}
}

func TestEachExampleStep(t *testing.T) {
	tests := map[int]string{
		1: `/-->\
|   |  /----\
| /-+--+-\  |
| | |  | |  |
\-+-/  \->--/
  \------/
`,

		2: `/---v
|   |  /----\
| /-+--+-\  |
| | |  | |  |
\-+-/  \-+>-/
  \------/
`,

		3: `/---\
|   v  /----\
| /-+--+-\  |
| | |  | |  |
\-+-/  \-+->/
  \------/
`,

		4: `/---\
|   |  /----\
| /->--+-\  |
| | |  | |  |
\-+-/  \-+--^
  \------/
`,

		5: `/---\
|   |  /----\
| /-+>-+-\  |
| | |  | |  ^
\-+-/  \-+--/
  \------/
`,

		6: `/---\
|   |  /----\
| /-+->+-\  ^
| | |  | |  |
\-+-/  \-+--/
  \------/
`,

		7: `/---\
|   |  /----<
| /-+-->-\  |
| | |  | |  |
\-+-/  \-+--/
  \------/
`,

		8: `/---\
|   |  /---<\
| /-+--+>\  |
| | |  | |  |
\-+-/  \-+--/
  \------/
`,

		9: `/---\
|   |  /--<-\
| /-+--+-v  |
| | |  | |  |
\-+-/  \-+--/
  \------/
`,

		10: `/---\
|   |  /-<--\
| /-+--+-\  |
| | |  | v  |
\-+-/  \-+--/
  \------/
`,

		11: `/---\
|   |  /<---\
| /-+--+-\  |
| | |  | |  |
\-+-/  \-<--/
  \------/
`,

		12: `/---\
|   |  v----\
| /-+--+-\  |
| | |  | |  |
\-+-/  \<+--/
  \------/
`,

		13: `/---\
|   |  /----\
| /-+--v-\  |
| | |  | |  |
\-+-/  ^-+--/
  \------/
`,

		14: `/---\
|   |  /----\
| /-+--+-\  |
| | |  X |  |
\-+-/  \-+--/
  \------/
`,
	}

	m := SetupMine("example-collision.txt")

	for i := 1; i <= 14; i++ {
		m.Step()
		o := m.Print()
		out := tests[i]
		if o != out {
			t.Errorf("wrong output for step %v: \nexpected:\n%vgot:\n%v\n", i, out, o)
		}
	}

}

func TestStepUntilCollision(t *testing.T) {
	expect := `/---\
|   |  /----\
| /-+--+-\  |
| | |  X |  |
\-+-/  \-+--/
  \------/
`
	m := SetupMine("example-collision.txt")
	x, y := m.StepUntilCollision()
	ex, ey := 7, 3

	if ex != x || ey != y {
		t.Errorf("wrong collision coorindates; expected <%v, %v>, got <%v, %v>", ex, ey, x, y)
	}

	got := m.Print()
	if got != expect {
		t.Errorf("wrong output, expected:\n%vgot:\n%v\n", expect, got)
	}
}

func TestStepUntilOneCart(t *testing.T) {
	ex, ey := 6, 4
	m := SetupMine("part2-test.txt")
	x, y := m.StepUntilOneCart()

	// spew.Dump(m.carts)

	o := m.Print()
	e := `/---\
|   |
| /-+-\
| | | |
\-+-/ ^
  |   |
  \---/
`
	if ex != x || ey != y {
		t.Errorf("wrong coordinates for last cart, expected <%v, %v>, got <%v, %v>", ex, ey, x, y)
	}

	if o != e {
		// spew.Dump(m.cartPos)
		t.Errorf("wrong mine for last cart.\nexpected:\n%v\ngot:\n%v\n", e, o)
	}
}
