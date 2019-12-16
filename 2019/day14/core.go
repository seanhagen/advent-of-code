package day14

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

type reaction struct {
	name string

	out     int
	inputs  map[string]int
	outputs map[string]int

	required int
}

func newReaction(in string) (*reaction, error) {
	parts := strings.Split(in, " => ")

	input := strings.TrimSpace(parts[0])
	out := strings.TrimSpace(parts[1])

	parts = strings.Split(out, " ")
	name := parts[1]
	c, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil, err
	}

	req := 0
	if name == "FUEL" {
		req = 1
	}

	r := reaction{
		name:     name,
		out:      c,
		inputs:   map[string]int{},
		outputs:  map[string]int{},
		required: req,
	}

	parts = strings.Split(input, ",")
	for _, v := range parts {
		tmp := strings.Split(strings.TrimSpace(v), " ")
		c, err = strconv.Atoi(tmp[0])
		if err != nil {
			return nil, err
		}
		r.inputs[tmp[1]] = c
	}
	return &r, nil
}

// Nanofactory ...
type Nanofactory struct {
	outputs map[string]*reaction
}

// CreateNanofactory ...
func CreateNanofactory() *Nanofactory {
	return &Nanofactory{outputs: map[string]*reaction{}}
}

// AddReaction ...
func (n *Nanofactory) AddReaction(in string) error {
	r, err := newReaction(in)
	if err != nil {
		return err
	}
	n.outputs[r.name] = r

	return nil
}

// CalcOreReq ...
func (n Nanofactory) CalcOreReq() int {
	reqs := map[string]int{}

	oreOuts := map[int]map[string]int{}

	onlyOreIn := map[string]bool{}

	for _, r := range n.outputs {
		if v, ok := r.inputs["ORE"]; ok && len(r.inputs) == 1 {
			o, ok := oreOuts[v]
			if !ok {
				o = map[string]int{}
			}
			onlyOreIn[r.name] = true

			o[r.name] = r.num
			oreOuts[v] = o
		}

		if _, ok := r.inputs["ORE"]; !ok {
			for k, v := range r.inputs {
				reqs[k] += v
			}
		}
	}

	spew.Dump(reqs, oreOuts, onlyOreIn)

	fmt.Printf("ore outs: \n%v\n\nreqs:\n%v\n\n", spew.Sdump(oreOuts), spew.Sdump(reqs))

	ore := 0
	ores := map[string]int{}

	// for k := range onlyOreIn {
	// 	v := reqs[k]
	// 	fmt.Printf("need %v %v\n", v, k)

	// 	r := n.outputs[k]
	// 	fmt.Printf("thing: %#v\n", r)

	// 	os.Exit(1)
	// }

	for or, r := range oreOuts {
		fmt.Printf("%v ore gives: %#v\n", or, r)
		tmp := 0
		for k, v := range r {
			required := reqs[k]

			// fmt.Printf("have %v ore, need %v, tmp: %v, v: %v\n", ore, required, tmp, v)
			for {
				tmp += v
				if tmp >= required {
					ores[k] = tmp
					break
				}
			}
		}

		ore += tmp
	}

	fmt.Printf("\nore now: %v\n", ore)
	spew.Dump(ores)

	return ore
}
