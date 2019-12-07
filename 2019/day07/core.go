package day07

import (
	"fmt"

	"github.com/seanhagen/advent-of-code/2019/lib2019"
)

const ModeSimple = 0
const ModeFeedback = 1

// Amp ...
type Amp struct {
	p        *lib2019.Program
	phase    int
	firstRun bool
}

// Run ...
func (a *Amp) Run(input int) (int, error) {
	if a.firstRun {
		a.p.AddInput(a.phase)
		a.firstRun = false
	}

	a.p.AddInput(input)
	err := a.p.Run()
	if err != nil {
		return input, err
	}
	out := a.p.GetOutputs()
	return out[len(out)-1], nil
}

// AmpCircuit
type AmpCircuit struct {
	amps []*Amp
	mode int
}

func NewAmpCircuit(code string, phases []int) (*AmpCircuit, error) {
	if len(phases) != 5 {
		return nil, fmt.Errorf("should be exactly 5 amp phases")
	}

	amps := []*Amp{}
	for _, ph := range phases {
		prog, err := lib2019.FromString(code)
		if err != nil {
			return nil, err
		}
		amps = append(amps, &Amp{p: prog, phase: ph, firstRun: true})
	}

	return &AmpCircuit{amps: amps, mode: ModeSimple}, nil
}

// SetMode ...
func (ac *AmpCircuit) SetMode(i int) {
	switch i {
	case ModeFeedback:
		ac.mode = i
	default:
		ac.mode = ModeSimple
	}
}

// Run ...
func (ac AmpCircuit) Run() (int, error) {
	err := ac.validate()
	if err != nil {
		return 0, err
	}
	//chain input->output, starts at 0
	ci := 0

	switch ac.mode {
	case ModeSimple:
		for _, a := range ac.amps {
			ci, err = a.Run(ci)
			if err != nil {
				return ci, err
			}
		}

	case ModeFeedback:
	feedback:
		for _, a := range ac.amps {
			a.p.SetPauseOnOutput(true)
			ci, err = a.Run(ci)
			if err != nil {
				if err == lib2019.ErrHalt {
					return ci, nil
				}
				return ci, err
			}

			if ci > 139629730 {
				return ci, fmt.Errorf("for test, output is too big now")
			}
		}
		goto feedback

	}

	return ci, nil
}

// validate ...
func (ac AmpCircuit) validate() error {
	for _, amp := range ac.amps {
		switch ac.mode {
		case ModeSimple:
			if amp.phase < 0 || amp.phase > 4 {
				return fmt.Errorf("phase must be between 0 & 5 inclusive in simple mode")
			}

		case ModeFeedback:
			if amp.phase < 5 || amp.phase > 9 {
				return fmt.Errorf("phase must be between 5 & 9 (inclusive) in feedback mode")
			}
		}
	}
	return nil
}

// Permutations calls f with each permutation of a.
func Permutations(a []int, f func([]int)) {
	perm(a, f, 0)
}

// Permute the values at index i to len(a)-1.
func perm(a []int, f func([]int), i int) {
	if i > len(a) {
		f(a)
		return
	}
	perm(a, f, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		perm(a, f, i+1)
		a[i], a[j] = a[j], a[i]
	}
}
