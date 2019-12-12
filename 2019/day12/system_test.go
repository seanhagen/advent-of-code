package day12

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/andreyvit/diff"
)

func factorial(x int) *big.Int {
	return fac(big.NewInt(int64(x)))
}

func fac(x *big.Int) *big.Int {
	n := big.NewInt(1)
	if x.Cmp(big.NewInt(0)) == 0 {
		return n
	}
	return n.Mul(x, fac(n.Sub(x, n)))
}

func nChooseK(n, k int) int64 {
	nfac := factorial(n)
	kfac := factorial(k)
	nmk := factorial(n - k)

	_ = kfac.Mul(kfac, nmk)
	_ = nfac.Div(nfac, kfac)

	return nfac.Int64()
}

func TestAddPlanets(t *testing.T) {
	tests := []struct {
		input []string
	}{
		{[]string{"<x=-1, y=0, z=2>"}},
		{[]string{"<x=-1, y=0, z=2>", "<x=2, y=-10, z=-7>"}},
		{[]string{"<x=-1, y=0, z=2>", "<x=2, y=-10, z=-7>", "<x=4, y=-8, z=8>"}},
		{[]string{"<x=-1, y=0, z=2>", "<x=2, y=-10, z=-7>", "<x=4, y=-8, z=8>", "<x=3, y=5, z=-1>"}},
	}
	for i, x := range tests {
		tt := x
		t.Run(fmt.Sprintf("test %v", i), func(t *testing.T) {
			sys := NewSystem()
			for _, s := range tt.input {
				sys.AddPlanet(s)
			}
			lp := len(sys.planets)
			if lp != len(tt.input) {
				t.Errorf("not enough planets, expected %v got %v", len(tt.input), lp)
			}
		})
	}
}

func TestCreatePairs(t *testing.T) {
	numChoose := 2
	tests := []struct {
		input []string
	}{
		{[]string{"<x=-1, y=0, z=2>", "<x=2, y=-10, z=-7>"}},
		{[]string{"<x=-1, y=0, z=2>", "<x=2, y=-10, z=-7>", "<x=4, y=-8, z=8>"}},
		{[]string{"<x=-1, y=0, z=2>", "<x=2, y=-10, z=-7>", "<x=4, y=-8, z=8>", "<x=3, y=5, z=-1>"}},
		{[]string{"<x=-1, y=0, z=2>", "<x=2, y=-10, z=-7>", "<x=4, y=-8, z=8>", "<x=3, y=5, z=-1>", "<x=5, y=2, z=3"}},
	}

	for i, x := range tests {
		tt := x
		t.Run(fmt.Sprintf("test %v", i), func(t *testing.T) {
			sys := NewSystem()
			for _, s := range tt.input {
				sys.AddPlanet(s)
			}
			sys.SetupPairs()
			numPairs := nChooseK(len(tt.input), numChoose)
			sp := int64(len(sys.pairs))
			if sp != numPairs {
				t.Errorf("not enough pairs, expected %v got %v", numPairs, sp)
			}
		})
	}
}

func TestSteps(t *testing.T) {
	firstInput := []string{"<x=-1, y=0, z=2>", "<x=2, y=-10, z=-7>", "<x=4, y=-8, z=8>", "<x=3, y=5, z=-1>"}
	secondInput := []string{"<x=-8, y=-10, z=0>", "<x=5, y=5, z=10>", "<x=2, y=-7, z=3>", "<x=9, y=-8, z=-3>"}

	tests := []struct {
		input  []string
		steps  int
		output string
	}{
		{firstInput, 0, `pos=<x= -1, y=  0, z=  2>, vel=<x=  0, y=  0, z=  0>
pos=<x=  2, y=-10, z= -7>, vel=<x=  0, y=  0, z=  0>
pos=<x=  4, y= -8, z=  8>, vel=<x=  0, y=  0, z=  0>
pos=<x=  3, y=  5, z= -1>, vel=<x=  0, y=  0, z=  0>
`},

		{firstInput, 1, `pos=<x=  2, y= -1, z=  1>, vel=<x=  3, y= -1, z= -1>
pos=<x=  3, y= -7, z= -4>, vel=<x=  1, y=  3, z=  3>
pos=<x=  1, y= -7, z=  5>, vel=<x= -3, y=  1, z= -3>
pos=<x=  2, y=  2, z=  0>, vel=<x= -1, y= -3, z=  1>
`},

		{firstInput, 3, `pos=<x=  5, y= -6, z= -1>, vel=<x=  0, y= -3, z=  0>
pos=<x=  0, y=  0, z=  6>, vel=<x= -1, y=  2, z=  4>
pos=<x=  2, y=  1, z= -5>, vel=<x=  1, y=  5, z= -4>
pos=<x=  1, y= -8, z=  2>, vel=<x=  0, y= -4, z=  0>
`},

		{firstInput, 10, `pos=<x=  2, y=  1, z= -3>, vel=<x= -3, y= -2, z=  1>
pos=<x=  1, y= -8, z=  0>, vel=<x= -1, y=  1, z=  3>
pos=<x=  3, y= -6, z=  1>, vel=<x=  3, y=  2, z= -3>
pos=<x=  2, y=  0, z=  4>, vel=<x=  1, y= -1, z= -1>
`},

		{secondInput, 100, `pos=<x=  8, y=-12, z= -9>, vel=<x= -7, y=  3, z=  0>
pos=<x= 13, y= 16, z= -3>, vel=<x=  3, y=-11, z= -5>
pos=<x=-29, y=-11, z= -1>, vel=<x= -3, y=  7, z=  4>
pos=<x= 16, y=-13, z= 23>, vel=<x=  7, y=  1, z=  1>
`},
	}

	for i, x := range tests {
		tt := x
		t.Run(fmt.Sprintf("test %v", i), func(t *testing.T) {
			sys := NewSystem()
			for _, s := range tt.input {
				err := sys.AddPlanet(s)
				if err != nil {
					t.Fatalf("unable to add planet '%v', error: %v", s, err)
				}
			}
			err := sys.SetupPairs()
			if err != nil {
				t.Fatalf("unable to setup pairs: %v", err)
			}

			sys.DoSteps(tt.steps)

			out := sys.Output()
			if out != tt.output {
				t.Errorf("invalid output: \n%v", diff.LineDiff(tt.output, out))
			}
		})
	}
}

func TestSystemTotalEnergy(t *testing.T) {
	tests := []struct {
		input  []string
		steps  int
		energy int
	}{
		{[]string{"<x=-1, y=0, z=2>", "<x=2, y=-10, z=-7>", "<x=4, y=-8, z=8>", "<x=3, y=5, z=-1>"}, 10, 179},
		{[]string{"<x=-8, y=-10, z=0>", "<x=5, y=5, z=10>", "<x=2, y=-7, z=3>", "<x=9, y=-8, z=-3>"}, 100, 1940},
	}

	for i, x := range tests {
		tt := x
		t.Run(fmt.Sprintf("test %v", i), func(t *testing.T) {
			sys := NewSystem()
			for _, s := range tt.input {
				err := sys.AddPlanet(s)
				if err != nil {
					t.Fatalf("unable to add planet '%v', error: %v", s, err)
				}
			}

			err := sys.SetupPairs()
			if err != nil {
				t.Fatalf("unable to setup pairs: %v", err)
			}

			sys.DoSteps(tt.steps)

			if te := sys.CalculateTotalEnergy(); te != tt.energy {
				t.Errorf("wrong total energy, expected %v, got %v", tt.energy, te)
			}
		})
	}
}
