package ch3e1

import (
	"errors"
	"regexp"
	"strings"
)

func Roman(value uint) (string, error) {
	if value > 3999 {
		return "", errors.New("out of range for Roman numerals")
	}

	var roman string

	mapping := []struct {
		arabic int
		roman  string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}
	for _, v := range mapping {
		if t := int(value) / v.arabic; t > 0 {
			roman += strings.Repeat(v.roman, t)
			value = value % uint(v.arabic)
		}
	}

	return roman, nil
}

var validRomanNumeral *regexp.Regexp

func init() {
	validRomanNumeral = regexp.MustCompile("^M{0,3}(?:CM|CD|D?C{0,3})(?:XC|XL|L?X{0,3})(?:IX|IV|V?I{0,3})$")
}

func HinduArabic(roman string) (uint, error) {
	var val uint

	if !validRomanNumeral.MatchString(roman) {
		return 0, errors.New("not a roman numeral")
	}

	var seenD, seenL, seenV bool
	for i, ch := range roman {
		switch ch {
		case 'M':
			val += 1000
		case 'D':
			if seenD {
				return 0, errors.New("too many D digits")
			}
			val += 500
			seenD = true
		// If this was `case 'C', 'J':` then the fuzzer found it
		case 'C':
			if i+1 < len(roman) {
				switch roman[i+1] {
				case 'D', 'M':
					val -= 100
				case 'C', 'L', 'X', 'V', 'I':
					val += 100
				default:
					return 0, errors.New("invalid subtraction " + roman[i:i+2])
				}
			} else {
				val += 100
			}
		case 'L':
			if seenL {
				return 0, errors.New("too many L digits")
			}
			val += 50
			seenL = true
		case 'X':
			if i+1 < len(roman) {
				switch roman[i+1] {
				case 'L', 'C':
					val -= 10
				case 'X', 'V', 'I':
					val += 10
				default:
					return 0, errors.New("invalid subtraction " + roman[i:i+2])
				}
			} else {
				val += 10
			}
		case 'V':
			if seenV {
				return 0, errors.New("too many V digits")
			}
			val += 5
			seenV = true
		case 'I':
			if i+1 < len(roman) && roman[i+1] != 'I' {
				if roman[i+1] == 'V' || roman[i+1] == 'X' {
					val--
				} else {
					return 0, errors.New("invalid subtraction " + roman[i:i+2])
				}
			} else {
				val++
			}
		default:
			return 0, errors.New("invalid numeral " + string(ch))
		}
	}

	return val, nil
}
