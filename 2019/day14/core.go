package day14

import (
	"math"
	"strconv"
	"strings"
)

const oneTril int64 = 1000000000000

type reaction struct {
	name string

	req     int64
	perProd int64

	inputs  map[string]int64
	outputs map[string]int64
}

// Nanofactory ...
type Nanofactory struct {
	outputs map[string]*reaction
}

// CreateNanofactory ...
func CreateNanofactory(input []string) (*Nanofactory, error) {
	nf := &Nanofactory{outputs: map[string]*reaction{}}
	for _, v := range input {
		err := nf.addReaction(v)
		if err != nil {
			return nil, err
		}
	}

	nf.outputs["FUEL"].req = 1
	nf.outputs["ORE"].perProd = 1

	return nf, nil
}

// addReaction ...
func (n *Nanofactory) addReaction(in string) error {
	bits := strings.Split(in, " => ")
	outBits := strings.Split(bits[1], " ")

	outName := outBits[1]
	o, ok := n.outputs[outName]
	if !ok {
		on, err := strconv.ParseInt(outBits[0], 10, 64)
		if err != nil {
			return err
		}

		o = &reaction{
			name:    outName,
			perProd: on,
			inputs:  map[string]int64{},
			outputs: map[string]int64{},
		}
	} else {
		on, err := strconv.ParseInt(outBits[0], 10, 64)
		if err != nil {
			return err
		}

		o.perProd = on
	}

	inputs := strings.Split(bits[0], ", ")

	for _, v := range inputs {
		r := strings.Split(v, " ")
		c, err := strconv.ParseInt(r[0], 10, 64)
		if err != nil {
			return err
		}

		in := r[1]
		inp, ok := n.outputs[in]
		if !ok {
			inp = &reaction{name: in, perProd: 0, req: 0, inputs: map[string]int64{}, outputs: map[string]int64{}}
		}

		o.inputs[in] = c
		inp.outputs[o.name] = c
		n.outputs[in] = inp
	}

	n.outputs[o.name] = o
	return nil
}

// CalcOreReq ...
func (n Nanofactory) CalcOreReq() int64 {
	return calcOre(n.outputs)
}

func calcOre(graph map[string]*reaction) int64 {
	if _, ok := graph["ORE"]; !ok {
		return 0
	}

	list := revTopSort(graph, "ORE")

	for _, na := range list {
		v := graph[na]
		c := float64(v.req) / float64(v.perProd)
		ri := int64(math.Max(1.0, math.Ceil(c)))
		for name, cost := range v.inputs {
			graph[name].req += ri * cost
		}
	}

	return graph["ORE"].req

}

func revTopSort(graph map[string]*reaction, s string) []string {
	ord := []string{}
	visited := map[string]bool{}

	var fn func(v *reaction, name string)
	fn = func(v *reaction, name string) {
		if len(v.outputs) == 0 {
			ord = append(ord, v.name)
			return
		}

		for ne := range v.outputs {
			if _, ok := visited[ne]; ok {
				continue
			}
			visited[ne] = true
			fn(graph[ne], ne)
		}
		ord = append(ord, name)
	}
	fn(graph[s], s)
	return ord
}

// copy ...
func (n Nanofactory) copy() map[string]*reaction {
	out := map[string]*reaction{}

	for k, v := range n.outputs {
		inputs := map[string]int64{}
		outputs := map[string]int64{}

		for kk, vv := range v.inputs {
			inputs[kk] = vv
		}

		for kk, vv := range v.outputs {
			outputs[kk] = vv
		}

		out[k] = &reaction{
			name:    v.name,
			req:     v.req,
			perProd: v.perProd,
			inputs:  inputs,
			outputs: outputs,
		}
	}

	return out
}

// CalcTrillionOre ...
func (n Nanofactory) CalcTrillionOre() int64 {
	var best int64
	var oreLim = oneTril
	var right = oneTril
	var left int64 = 1

	for left <= right {
		graph := n.copy()
		mid := (left + right) / 2
		graph["FUEL"].req = mid
		reqOre := calcOre(graph)

		if reqOre < oreLim {
			best = int64(math.Max(float64(best), float64(mid)))
			left = mid + 1
		} else if reqOre > oneTril {
			right = mid - 1
		} else {
			return mid
		}
	}

	return best
}
