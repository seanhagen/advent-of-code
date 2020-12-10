package day4

import (
	"fmt"
	"testing"
)

const exampleInput = `ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm

iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929

hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm

hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in`

func TestParseFile(t *testing.T) {
	should := 4
	out := ParseFile(exampleInput)

	l := len(out)
	if l != should {
		t.Errorf("wrong number of lines, expected %v got %v", should, l)
	}
}

func TestCreatePassport(t *testing.T) {
	ex1 := `ecl:gry pid:860033327 eyr:2020 hcl:#fffffd byr:1937 iyr:2017 cid:147 hgt:183cm`
	ex2 := `iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884 hcl:#cfa07d byr:1929`
	ex3 := `hcl:#ae17e1 iyr:2013 eyr:2024 ecl:brn pid:760753108 byr:1931 hgt:179cm`
	ex4 := `hcl:#cfa07d eyr:2025 pid:166559648 iyr:2011 ecl:brn hgt:59in`
	ex5 := `ecl:blu iyr:2019 pid:960083875 eyr:2027 hgt:71in hcl:#c0946f byr:1921`

	tests := []struct {
		in    string
		valid bool
	}{
		{"", false},
		{"xxx:111", false},
		{":w", false},
		{"dddd_ddd", false},
		{"iyr:", false},
		{"hcl: ", false},
		{"ecl: ", false},
		{"byr:1937", true},
		{"byr:1937 iyr:2020", true},
		{"eyr:2021", true},
		{ex1, true},
		{ex2, true},
		{ex3, true},
		{ex4, true},
		{ex5, true},
	}

	for i, x := range tests {
		tt := x
		t.Run(fmt.Sprintf("test_%v", i), func(t *testing.T) {
			p, err := NewPassport(tt.in, true)
			if tt.valid {
				if err != nil {
					t.Fatalf("expected Passport, got error: %v", err)
				}
				if p == nil {
					t.Errorf("expected Passport, got nil!")
				}

			} else {
				if err == nil {
					t.Errorf("expected error, got none? input: '%v'", tt.in)
				}
			}
		})
	}
}

func TestNumValid(t *testing.T) {
	ex1 := `ecl:gry pid:860033327 eyr:2020 hcl:#fffffd byr:1937 iyr:2017 cid:147 hgt:183cm`
	ex2 := `iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884 hcl:#cfa07d byr:1929`
	ex3 := `hcl:#ae17e1 iyr:2013 eyr:2024 ecl:brn pid:760753108 byr:1931 hgt:179cm`
	ex4 := `hcl:#cfa07d eyr:2025 pid:166559648 iyr:2011 ecl:brn hgt:59in`
	ex5 := `ecl:blu iyr:2019 pid:960083875 eyr:2027 hgt:71in hcl:#c0946f byr:1921`

	tests := []struct {
		in       string
		numValid int
	}{
		{"iyr:d", 0},
		{"eyr:d", 0},
		{"hgt:d", 0},
		{"hgt:in", 0},
		{"hgt:cm", 0},
		{"pid:d", 1},
		{"cid:d", 0},
		{"byr:1937", 1},
		{"byr:1937 iyr:2020", 2},
		{"eyr:2021", 1},
		{ex1, 8},
		{ex2, 7},
		{ex3, 7},
		{ex4, 6},
		{ex5, 7},
	}

	for i, x := range tests {
		tt := x
		t.Run(fmt.Sprintf("test_%v", i), func(t *testing.T) {
			p, err := NewPassport(tt.in, true)
			if err != nil {
				t.Fatalf("should be valid passport string, unable to create Passport: %v", err)
			}

			n := p.NumValid()
			if n != tt.numValid {
				t.Errorf("wrong number of valid fields; expected %v got %v", tt.numValid, n)
			}

		})
	}
}

func TestIsValid(t *testing.T) {
	ex1 := `ecl:gry pid:860033327 eyr:2020 hcl:#fffffd byr:1937 iyr:2017 cid:147 hgt:183cm`
	ex2 := `iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884 hcl:#cfa07d byr:1929`
	ex3 := `hcl:#ae17e1 iyr:2013 eyr:2024 ecl:brn pid:760753108 byr:1931 hgt:179cm`
	ex4 := `hcl:#cfa07d eyr:2025 pid:166559648 iyr:2011 ecl:brn hgt:59in`
	ex5 := `ecl:blu iyr:2019 pid:960083875 eyr:2027 hgt:71in hcl:#c0946f byr:1921`

	tests := []struct {
		in      string
		isValid bool
	}{
		{"byr:1937", false},
		{"byr:1937 iyr:2020", false},
		{"eyr:2021", false},
		{ex1, true},
		{ex2, false},
		{ex3, true},
		{ex4, false},
		{ex5, true},
	}

	for i, x := range tests {
		tt := x
		t.Run(fmt.Sprintf("test_%v", i), func(t *testing.T) {
			p, err := NewPassport(tt.in, true)
			if err != nil {
				t.Fatalf("should be valid passport string, unable to create Passport: %v", err)
			}
			p.SetStrict(false)

			n := p.IsValid()
			if n != tt.isValid {
				t.Errorf("wrong output, expected '%v' got '%v'", tt.isValid, n)
			}
		})
	}
}

func TestStrictIsValid(t *testing.T) {
	iv1 := "eyr:1972 cid:100 hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926"
	iv2 := "iyr:2019 hcl:#602927 eyr:1967 hgt:170cm ecl:grn pid:012533040 byr:1946"
	iv3 := "hcl:dab227 iyr:2012 ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277"
	iv4 := "hgt:59cm ecl:zzz eyr:2038 hcl:74454a iyr:2023 pid:3556412378 byr:2007"

	iv5 := "cid:80 byr:1936 iyr:2017 hgt:94 hcl:#2e7503 ecl:oth eyr:2030 pid:597284996"

	v1 := "pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980 hcl:#623a2f"
	v2 := "eyr:2029 ecl:blu cid:129 byr:1989 iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm"
	v3 := "hcl:#888785 hgt:164cm byr:2001 iyr:2015 cid:88 pid:545766238 ecl:hzl eyr:2022"
	v4 := "iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719"

	tests := []struct {
		in      string
		isValid bool
	}{
		{"eyr:192", false},
		{"eyr:2029 ecl:blu cid:129 iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm", false},
		{"eyr:2035 pid:189cm hgt:77 iyr:1973 ecl:#dc83d5 hcl:z byr:2004", false},
		{iv1, false},
		{iv2, false},
		{iv3, false},
		{iv4, false},
		{iv5, false},
		{v1, true},
		{v2, true},
		{v3, true},
		{v4, true},
	}

	for i, x := range tests {
		tt := x
		t.Run(fmt.Sprintf("test_%v", i), func(t *testing.T) {
			p, err := NewPassport(tt.in, true)
			if err != nil {
				t.Fatalf("should be valid passport string, unable to create Passport: %v", err)
			}
			p.SetStrict(true)
			n := p.IsValid()
			if n != tt.isValid {
				t.Errorf("wrong output, expected '%v' got '%v' for passport '%v'", tt.isValid, n, tt.in)
			}
		})
	}
}
