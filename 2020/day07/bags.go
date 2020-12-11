package day7

import (
	"regexp"
	"strconv"
	"strings"
)

type BagRule struct {
	name    string
	contain map[*BagRule]int
}

type BagRules struct {
	types []string
	bags  []*BagRule
}

func ParseBags(input string) *BagRules {
	re := regexp.MustCompile("(\\d) (.+)")
	lines := strings.Split(input, "\n")
	bags := map[string]map[string]int{}
	types := []string{}

	for _, v := range lines {
		bits := strings.Split(v, " bags contain ")
		name := bits[0]
		types = append(types, name)

		r := strings.Replace(bits[1], ".", "", 1)
		r = strings.Replace(r, " bags", "", -1)
		r = strings.Replace(r, " bag", "", -1)
		rules := strings.Split(r, ", ")

		tmp := map[string]int{}
		for _, vv := range rules {
			if vv == "no other" {
				continue
			}
			o := re.FindStringSubmatch(vv)
			num, _ := strconv.Atoi(o[1])
			t := o[2]
			tmp[t] = num
		}

		bags[name] = tmp
	}

	br := map[string]*BagRule{}

	for name, _ := range bags {
		br[name] = &BagRule{name: name, contain: map[*BagRule]int{}}
	}

	for name, rules := range bags {
		b := br[name]
		for n, num := range rules {
			bt := br[n]
			b.contain[bt] = num
		}
	}

	list := []*BagRule{}
	for _, bag := range br {
		list = append(list, bag)
	}

	return &BagRules{
		types: types,
		bags:  list,
	}
}

/*
shiny gold
  1 dark olive
    3 faded blue
      none

    4 dotted black
      none

  2 vibrant plum
    5 faded blue
      none

    6 dotted black
      none
*/
