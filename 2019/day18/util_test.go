package day18

import (
	"fmt"
	"testing"
)

func TestInputToData(t *testing.T) {
	tests := []struct {
		input  string
		output map[int]map[int]string
	}{
		{`##`, map[int]map[int]string{
			0: map[int]string{
				0: "#",
				1: "#",
			},
		}},
		{``, map[int]map[int]string{}},
		{`##
##`, map[int]map[int]string{
			0: map[int]string{
				0: "#",
				1: "#",
			},
			1: map[int]string{
				0: "#",
				1: "#",
			},
		}},
		{`#########
#b.A.@.a#
#########`, map[int]map[int]string{
			0: map[int]string{
				0: "#",
				1: "#",
				2: "#",
				3: "#",
				4: "#",
				5: "#",
				6: "#",
				7: "#",
				8: "#",
			},
			1: map[int]string{
				0: "#",
				1: "b",
				2: ".",
				3: "A",
				4: ".",
				5: "@",
				6: ".",
				7: "a",
				8: "#",
			},
			2: map[int]string{
				0: "#",
				1: "#",
				2: "#",
				3: "#",
				4: "#",
				5: "#",
				6: "#",
				7: "#",
				8: "#",
			},
		}},
	}

	for i, x := range tests {
		tt := x
		t.Run(fmt.Sprintf("test %v", i), func(t *testing.T) {
			out := inputToData(tt.input)
			if err := mapEq(out, tt.output); err != nil {
				t.Errorf("maps not equal: %v", err)
			}
		})
	}
}

func mapEq(a, b map[int]map[int]string) error {
	if len(a) != len(b) {
		return fmt.Errorf("maps not equal length")
	}

	for k, v := range b {
		w, ok := a[k]
		if !ok {
			return fmt.Errorf("no key '%v' in map a", k)
		}

		for i, vv := range v {
			ww, ok := w[i]
			if !ok {
				return fmt.Errorf("no key '%v' for submap '%v' in map a", i, k)
			}
			if vv != ww {
				return fmt.Errorf("a[%v][%v] (%v) != b[%v][%v] (%v)", i, k, vv, i, k, ww)
			}
		}
	}

	return nil
}
