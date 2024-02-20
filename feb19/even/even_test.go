package even

import "testing"

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
