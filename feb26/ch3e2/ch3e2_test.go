package ch3e2

import "testing"

func BenchmarkModulo(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = n%2 == 0
	}
}

func BenchmarkBitwiseAnd(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = n&1 == 0
	}
}

func BenchmarkMeasure(b *testing.B) {
	BenchmarkBitwiseAnd(b)
}
