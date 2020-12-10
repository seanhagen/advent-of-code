package day4

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

var validEyeColors = []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}

type Passport struct {
	input string

	cidOpt bool
	strict bool

	pid *string
	cid *int

	birthYear  *int
	issueYear  *int
	expireYear *int

	height   *int
	heightCM bool
	heightIN bool

	hairColor *string
	eyeColor  *string

	errors map[string]error
}

// NewPassport ...
func NewPassport(in string, cidOpt bool) (*Passport, error) {
	in = strings.TrimSpace(in)
	if in == "" {
		return nil, fmt.Errorf("invalid input string, can't be empty")
	}

	p := &Passport{input: in, cidOpt: cidOpt}
	errs := map[string]error{}

	bits := strings.Split(in, " ")
	form := "%v %v"
	for _, x := range bits {
		x = strings.Replace(x, ":", " ", 1)
		var f, v string

		_, err := fmt.Sscanf(x, form, &f, &v)
		if err != nil {
			return nil, fmt.Errorf("unable to scan passport value '%v', reason: %v\n", x, err)
		}

		switch f {
		case "byr":
			var y int
			y, err = strconv.Atoi(v)
			p.birthYear = &y

		case "iyr":
			var y int
			y, err = strconv.Atoi(v)
			p.issueYear = &y

		case "eyr":
			var y int
			y, err = strconv.Atoi(v)
			p.expireYear = &y

		case "hgt":
			if strings.Index(v, "cm") > 0 {
				v = strings.Replace(v, "cm", "", 1)
				p.heightCM = true
			}

			if strings.Index(v, "in") > 0 {
				v = strings.Replace(v, "in", "", 1)
				p.heightIN = true
			}

			var y int
			y, err = strconv.Atoi(v)
			p.height = &y

		case "hcl":
			v = strings.TrimSpace(v)
			if v == "" {
				err = fmt.Errorf("hair color can't be blank")
			}
			p.hairColor = &v

		case "ecl":
			v = strings.TrimSpace(v)
			if v == "" {
				err = fmt.Errorf("eye color can't be blank")
			}
			p.eyeColor = &v

		case "pid":
			v = strings.TrimSpace(v)
			if v == "" {
				err = fmt.Errorf("eye color can't be blank")
			}
			p.pid = &v

		case "cid":
			var y int
			y, err = strconv.Atoi(v)
			p.cid = &y

		default:
			return nil, fmt.Errorf("unknown passport data code: %v", f)
		}
		if err != nil {
			errs[f] = err
		}
	}
	p.errors = errs
	return p, nil
}

// SetStrict ...
func (p *Passport) SetStrict(s bool) {
	p.strict = s
}

// NumValid ...
func (p *Passport) NumValid() int {
	nv := 0
	if p.birthYear != nil && *p.birthYear != 0 {
		nv++
	}

	if p.issueYear != nil && *p.issueYear != 0 {
		nv++
	}

	if p.expireYear != nil && *p.expireYear != 0 {
		nv++
	}

	if p.height != nil && *p.height != 0 {
		nv++
	}

	if p.hairColor != nil && *p.hairColor != "" {
		nv++
	}

	if p.eyeColor != nil && *p.eyeColor != "" {
		nv++
	}

	if p.pid != nil && *p.pid != "" {
		nv++
	}

	if p.cid != nil && *p.cid != 0 {
		nv++
	}

	return nv
}

// IsValid ...
func (p *Passport) IsValid() bool {
	if p.strict {
		return p.strictValid()
	}
	return p.nonStrictValid()
}

// NonStrictValid ...
func (p *Passport) nonStrictValid() bool {
	// fmt.Printf("non strict validation for '%v'\n", p.input)
	if p.cidOpt {
		delete(p.errors, "cid")
	}

	if len(p.errors) != 0 {
		fmt.Printf("parse error encountered:")
		spew.Dump(p.errors)
		fmt.Printf("\n")
		return false
	}

	nv := p.NumValid()
	if p.cidOpt && p.cid == nil {
		nv++
	}
	//spew.Dump(p.cidOpt, p.cid, nv == 8)
	// fmt.Printf("number of valid fields: %v\n", nv)

	if nv == 8 {
		return true
	}

	return false
}

// StrictValid ...
func (p *Passport) strictValid() bool {
	valid := true

	if p.birthYear == nil {
		valid = false
		p.errors["birth_year"] = fmt.Errorf("birth year is null")
	} else {
		if *p.birthYear < 1920 || *p.birthYear > 2002 {
			p.errors["birth_year"] = fmt.Errorf("invalid birth year ( %v )", *p.birthYear)
			valid = false
		}
	}

	if p.issueYear == nil {
		valid = false
		p.errors["issue_year"] = fmt.Errorf("issue year is null")
	} else {
		if *p.issueYear < 2010 || *p.issueYear > 2020 {
			p.errors["issue_year"] = fmt.Errorf("invalid issue year ( %v )", *p.issueYear)
			valid = false
		}
	}

	if p.expireYear == nil {
		valid = false
		p.errors["expire_year"] = fmt.Errorf("expire year is null")
	} else {
		if *p.expireYear < 2020 || *p.expireYear > 2030 {
			valid = false
			p.errors["expire_year"] = fmt.Errorf("invalid expire year ( %v )", *p.expireYear)
		}
	}

	if p.height == nil {
		valid = false
		p.errors["height"] = fmt.Errorf("height is null")
	} else {
		if !p.heightCM && !p.heightIN {
			valid = false
			p.errors["height"] = fmt.Errorf("height must be in inches or centimeters")
		}

		if p.heightCM && (*p.height < 150 || *p.height > 193) {
			valid = false
			p.errors["height"] = fmt.Errorf("invalid height in CM: %v", *p.height)
		}
		if p.heightIN && (*p.height < 59 || *p.height > 76) {
			valid = false
			p.errors["height"] = fmt.Errorf("invalid height in IN: %v", *p.height)
		}
	}

	if p.hairColor == nil {
		valid = false
		p.errors["hair_color"] = fmt.Errorf("hair color is null")
	} else {
		if _, err := ParseHexColorFast(*p.hairColor); err != nil {
			valid = false
			p.errors["hair_color"] = fmt.Errorf("invalid hair color: %v", *p.hairColor)
		}
	}

	if p.eyeColor == nil {
		valid = false
		p.errors["eye_color"] = fmt.Errorf("eye color is null")
	} else {
		ec := *p.eyeColor
		v := false
		for _, x := range validEyeColors {
			if ec == x {
				v = true
			}
		}
		if !v {
			p.errors["eye_color"] = fmt.Errorf("invalid eye color: %v", ec)
			valid = false
		}
	}

	if p.pid == nil {
		valid = false
		p.errors["pid"] = fmt.Errorf("PID is null")
	} else {
		v := *p.pid
		l := len(v)
		_, err := strconv.Atoi(v)
		if l != 9 || err != nil {
			valid = false
			p.errors["pid"] = fmt.Errorf("PID is invalid: %v", v)
		}
	}

	return valid

}

// PrintErrors ...
func (p *Passport) PrintErrors() {
	for k, v := range p.errors {
		fmt.Printf("\t\tError parsing field '%v': %v\n", k, v)
	}
}
