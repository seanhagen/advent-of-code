package day06

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// LightGrid is our one million light grid
type LightGrid [1000][1000]int

type inst string

const on inst = "on"
const off inst = "off"
const toggle inst = "toggle"

var reOn = regexp.MustCompile(`turn on -?[0-9]+,-?[\d]+ through -?[\d]+,-?[\d]+`)
var reOff = regexp.MustCompile(`turn off -?[0-9]+,-?[\d]+ through -?[\d]+,-?[\d]+`)
var reTog = regexp.MustCompile(`toggle -?[0-9]+,-?[\d]+ through -?[\d]+,-?[\d]+`)

var ref = regexp.MustCompile(`-?[\d]+`)

const numCoords = 4

func validInstruction(in string) bool {
	m1 := reOn.MatchString(in)
	m2 := reOff.MatchString(in)
	m3 := reTog.MatchString(in)

	return m1 || m2 || m3
}

// insToCoords ...
func insToCoords(in string) []int {
	out := make([]int, numCoords)

	if !validInstruction(in) {
		return out
	}

	tmp := ref.FindAllString(in, -1)
	if len(tmp) != numCoords {
		return out
	}

	for i, v := range tmp {
		x, err := strconv.Atoi(v)
		if err != nil {
			return out
		}

		if x > 9999 {
			x = 9999
		}
		if x < 0 {
			x = 0
		}

		out[i] = x
	}

	return out
}

func instruction(in string) inst {
	if !validInstruction(in) {
		return off
	}

	if reTog.MatchString(in) {
		return toggle
	}

	if reOn.MatchString(in) {
		return on
	}

	return off
}

func onOrOff(in string) int {
	bits := strings.Split(in, " ")
	switch bits[1] {
	case "on":
		return 1
	}
	return 0
}

// ProcessInstruction ...
func (lg *LightGrid) ProcessInstruction(in string) error {
	if !validInstruction(in) {
		return fmt.Errorf("instruction doesn't match format")
	}

	coords := insToCoords(in)
	ins := instruction(in)
	startX, endX, startY, endY := coords[0], coords[2], coords[1], coords[3]

	for i := startX; i <= endX; i++ {
		for j := startY; j <= endY; j++ {
			switch ins {
			case on:
				lg[i][j] = 1
			case off:
				lg[i][j] = 0
			case toggle:
				if lg[i][j] == 1 {
					lg[i][j] = 0
				} else {
					lg[i][j] = 1
				}
			}
		}
	}
	return nil
}

// ProcessInstructionsV2 ...
func (lg *LightGrid) ProcessInstructionsV2(in string) error {
	if !validInstruction(in) {
		return fmt.Errorf("instruction doesn't match format")
	}

	coords := insToCoords(in)
	ins := instruction(in)
	startX, endX, startY, endY := coords[0], coords[2], coords[1], coords[3]

	for i := startX; i <= endX; i++ {
		for j := startY; j <= endY; j++ {

			switch ins {
			case on:
				lg[i][j]++
			case off:
				lg[i][j]--
				if lg[i][j] < 0 {
					lg[i][j] = 0
				}
			case toggle:
				lg[i][j] += 2
			}
		}
	}

	return nil
}

// TotalBrightness ...
func (lg LightGrid) TotalBrightness() int {
	b := 0

	for _, row := range lg {
		for _, v := range row {
			b += v
		}
	}

	return b
}

// CountOn ...
func (lg LightGrid) CountOn() int {
	count := 0
	for _, row := range lg {
		for _, v := range row {
			if v > 0 {
				count++
			}
		}
	}
	return count
}

// SetAll ...
func (lg *LightGrid) SetAll(in int) {
	if in < 0 {
		in = 0
	}

	for i, row := range lg {
		for j := range row {
			lg[i][j] = in
		}
	}
}
