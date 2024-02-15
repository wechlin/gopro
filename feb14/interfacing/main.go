package main

import (
	"fmt"
	"sort"
)

type Accumulator interface {
	Current() int
	MixIn(int)
}

func Accum100(a Accumulator) {
	for n := 0; n <= 100; n++ {
		a.MixIn(n)
	}
}

type DevNull bool

func (d DevNull) Current() int {
	return 0
}

func (d DevNull) MixIn(val int) {

}

type Median []int

func (m *Median) MixIn(val int) {
	*m = append(*m, val)
}

func (m *Median) Current() int {
	if len(*m) == 0 {
		return 0
	}

	sort.Ints(*m)

	return (*m)[len(*m)/2]
}

type Sum int

func (s *Sum) Current() int {
	return int(*s)
}

func (s *Sum) MixIn(val int) {
	*s += Sum(val)
}

func main() {
	var something *Median = new(Median)

	Accum100(something)

	fmt.Println(something.Current())
}
