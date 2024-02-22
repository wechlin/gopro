package ch3e1

import (
	"math"
	"testing"
)

func BenchmarkHinduArabic(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, tt := range cases {
			HinduArabic(tt.roman)
		}
	}
}

func FuzzToHinduArabic(f *testing.F) {
	for _, tt := range cases {
		f.Add(tt.roman)
	}

	f.Fuzz(func(t *testing.T, input string) {
		val, err := HinduArabic(input)
		if err == nil {
			roman, err := Roman(val)
			if err != nil {
				t.Errorf("Roman(HinduArabic(%v)) error %v", input, err)
			} else if roman != input {
				t.Errorf("Roman(HinduArabic(%v)) = %v, want %v (inter %v)", input, roman, input, val)
			}
		}
	})
}

func TestToRoman(t *testing.T) {
	for _, tt := range cases {
		val, err := Roman(tt.integer)
		if err != nil {
			t.Errorf("Roman(%v) error %v", tt.integer, err)
		} else if val != tt.roman {
			t.Errorf("Roman(%v) == %v, want %v", tt.integer, val, tt.roman)
		}
	}
}

func TestToHinduArabic(t *testing.T) {

	for _, tt := range cases {
		val, err := HinduArabic(tt.roman)
		if err != nil {
			t.Errorf("HinduArabic(%v) error %v", tt.roman, err)
		} else if val != tt.integer {
			t.Errorf("HinduArabic(%v) == %v, want %v", tt.roman, val, tt.integer)
		}
	}
}

func TestRomanErrors(t *testing.T) {
	romanErrors := []uint{
		4000,
		4001,
		math.MaxUint,
	}

	for _, tt := range romanErrors {
		val, err := Roman(tt)
		if err == nil {
			t.Errorf("Roman(%v) == %v, want Error", tt, val)
		}
	}
}

func TestHinduArabicErrors(t *testing.T) {
	hinduArabicErrors := []string{
		"Q",
	}
	for _, tt := range hinduArabicErrors {
		val, err := HinduArabic(tt)
		if err == nil {
			t.Errorf("HinduArabic(%v) == %v, want Error", tt, val)
		}
	}
}

var cases = []struct {
	roman   string
	integer uint
}{
	{"", 0},
	{"I", 1},
	{"II", 2},
	{"III", 3},
	{"IV", 4},
	{"V", 5},
	{"VI", 6},
	{"VII", 7},
	{"VIII", 8},
	{"IX", 9},
	{"X", 10},
	{"XI", 11},
	{"XIV", 14},
	{"XV", 15},
	{"XVI", 16},
	{"XVII", 17},
	{"XVIII", 18},
	{"XIX", 19},
	{"XX", 20},
	{"XXI", 21},
	{"XXIV", 24},
	{"XXV", 25},
	{"XXVI", 26},
	{"XXIX", 29},
	{"XXX", 30},
	{"XXXIX", 39},
	{"XL", 40},
	{"XLI", 41},
	{"XLIV", 44},
	{"XLV", 45},
	{"XLVI", 46},
	{"XLVIII", 48},
	{"XLIX", 49},
	{"L", 50},
	{"LI", 51},
	{"LII", 52},
	{"LIII", 53},
	{"LIV", 54},
	{"LV", 55},
	{"LVI", 56},
	{"LIX", 59},
	{"LX", 60},
	{"LXI", 61},
	{"LXXII", 72},
	{"LXXXIII", 83},
	{"LXXXVIII", 88},
	{"LXXXIX", 89},
	{"XC", 90},
	{"XCI", 91},
	{"XCII", 92},
	{"XCIV", 94},
	{"XCIV", 94},
	{"XCV", 95},
	{"XCVIII", 98},
	{"XCIX", 99},
	{"C", 100},
	{"CI", 101},
	{"CIV", 104},
	{"CV", 105},
	{"CIX", 109},
	{"CX", 110},
	{"CXI", 111},
	{"CL", 150},
	{"CXC", 190},
	{"CC", 200},
	{"CCC", 300},
	{"CD", 400},
	{"D", 500},
	{"DC", 600},
	{"DCCC", 800},
	{"DCCCLXXXVIII", 888},
	{"MMMDCCCLXXXVIII", 3888},
	{"MMMCMXCIX", 3999},
}
