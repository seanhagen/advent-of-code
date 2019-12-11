package facing

import (
	"fmt"
	"testing"
)

func TestTypeCheck(t *testing.T) {
	tests := []struct {
		a    interface{}
		b    NewObjFn
		good bool
	}{
		{"", func() interface{} { return "" }, true},
		{"", func() interface{} { return 1 }, false},
		{Mover{}, func() interface{} { return Mover{} }, true},
		{Mover{}, func() interface{} { return "" }, false},
	}

	for i, tt := range tests {
		x := tt
		t.Run(fmt.Sprintf("test %v", i), func(t *testing.T) {
			v := x.b()
			eq := typeEqual(x.a, v)
			if eq != x.good {
				t.Errorf("expected value (%#v) type %T same as (%#v) type %T to return '%v', got '%v'", x.a, x.a, v, v, x.good, eq)
			}
		})
	}
}
