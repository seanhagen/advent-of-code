package day08

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

const (
	Black       = 0
	White       = 1
	Transparent = 2
)

type Layer struct {
	id   int
	data [][]int
}

// NumZeroes ...
func (l Layer) NumZeroes() int {
	c := 0
	for _, r := range l.data {
		for _, x := range r {
			if x == 0 {
				c++
			}
		}
	}
	return c
}

// OneByTwo ...
func (l Layer) OneByTwo() int {
	no, nt := 0, 0
	for _, r := range l.data {
		for _, c := range r {
			if c == 1 {
				no++
			}
			if c == 2 {
				nt++
			}
		}
	}
	return no * nt
}

type Image struct {
	width  int
	height int

	layers []*Layer
}

func NewImage(in string, w, h int) (*Image, error) {
	m := w * h
	in = strings.Replace(in, "\n", "", -1)

	if len(in)%m != 0 {
		return nil, fmt.Errorf("length of input not divisible by w*h ( w: %v, h: %v, len: %v )", w, h, len(in))
	}

	idx := 0
	layers := []*Layer{}
	bits := strings.Split(in, "")

	for k := 0; k < len(in)/m; k++ {
		ld := [][]int{}

		for i := 0; i < h; i++ {
			row := []int{}
			for j := 0; j < w; j++ {
				z, err := strconv.Atoi(bits[idx])
				if err != nil {
					return nil, err
				}
				row = append(row, z)
				idx++
			}
			ld = append(ld, row)
		}
		layers = append(layers, &Layer{data: ld, id: k + 1})
	}

	return &Image{
		width:  w,
		height: h,
		layers: layers,
	}, nil
}

// FindSmallestNumZeroLayer ...
func (img Image) FindSmallestNumZeroLayer() *Layer {
	c := img.width * img.height
	var smallest *Layer
	for _, l := range img.layers {
		if nz := l.NumZeroes(); nz < c {
			c = nz
			smallest = l
		}
	}
	return smallest
}

// Output ...
func (img Image) Output() string {
	tmp := make([][]int, img.height)
	for i := 0; i < img.height; i++ {
		x := make([]int, img.width)
		for j := 0; j < img.width; j++ {
			x[j] = Transparent
		}
		tmp[i] = x
	}

	for _, l := range img.layers {
		for i, r := range l.data {
			for j, c := range r {
				if c != Transparent && tmp[i][j] == Transparent {
					// fmt.Printf("layer %v, row %v, col %v not transparent: %v\n", )
					tmp[i][j] = c
				}
			}
		}
	}

	out := bytes.NewBufferString("")
	for x, r := range tmp {
		for _, c := range r {
			fmt.Fprintf(out, "%v", c)
		}
		if x < len(tmp)-1 {
			fmt.Fprintf(out, "\n")
		}
	}

	return out.String()
}
