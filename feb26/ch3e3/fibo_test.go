package fibonacci

import (
	"fmt"
	"testing"
)

func BenchmarkFibonacci(b *testing.B) {
	for _, target := range []uint{1, 10, 20} {

		name := fmt.Sprintf("target=%d", target)
		b.Run(name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				Fib(target)
			}
		})
	}
}

func TestFibonacci(t *testing.T) {
	cases := []struct {
		in, out uint
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
	}

	for _, tt := range cases {
		if val := Fib(tt.in); val != tt.out {
			t.Errorf("Fib(%d) = %d; want %d", tt.in, val, tt.out)
		}
	}
}
