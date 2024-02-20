package even

import (
	"math"
	"testing"
)

// In VS Code, the "Go: Toggle Test Coverage" flag from the Command Palette
// will enable code test coverage measurement as well as highlight any code
// red/green as to its coverage

func TestEven(t *testing.T) {
	evens := []int{-2, 0, 2, 1024, 8, 12, 16}

	for _, num := range evens {
		if !IsEven(num) {
			t.Errorf("IsEven(%d) = false; want true", num)
		}
	}
}

func TestOdd(t *testing.T) {
	odds := []int{-1, 1, 17, 13, 1023, 65535, -17}

	for _, num := range odds {
		if IsEven(num) {
			t.Errorf("IsEven(%d) = true; want false", num)
		}
	}
}

func TestKmToMiles(t *testing.T) {
	cases := []struct {
		in  Kilometers
		out float64
		tol float64
	}{
		{0, 0, 0.0001},
		{1.6095, 1, 0.0001},
		{1, 0.621371, 0.0001},
		{-10, 0, 0.000000001},
	}

	for _, item := range cases {
		if res := item.in.InMiles(); math.Abs(res-item.out) > item.tol {
			t.Errorf("(%v).InMiles() = %f; want %f", item.in, res, item.out)
		}
	}
}
