package reverse

import (
	"testing"
	"unicode/utf8"
)

func TestReverse(t *testing.T) {
	cases := []struct {
		in, out string
	}{
		{"noon", "noon"},
		{"Hello", "olleH"},
		{"My name is", "si eman yM"},
	}

	for _, tt := range cases {
		if res := Reverse(tt.in); res != tt.out {
			t.Errorf("Reverse(%s) = %s, want %s", tt.in, res, tt.out)
		}
	}
}

func FuzzReverse(f *testing.F) {
	f.Add("noon")
	f.Add("Hello")

	f.Fuzz(func(t *testing.T, s string) {
		if !utf8.ValidString(s) {
			return
		}

		mid := Reverse(s)
		final := Reverse(mid)
		if !utf8.ValidString(final) {
			t.Errorf("Reverse(%s) = %s, Reverse(%s) = %s invalid unicode", s, mid, mid, final)
		} else if final != s {
			t.Errorf("Reverse(Reverse(%s)) = %s, want %s", s, final, s)
		}
	})
}
