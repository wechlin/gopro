package main

import "fmt"

type (
	Km = float64
	Kg float64
)

func arrayChunk() {
	var distance Km
	var length float64

	var legs = [...]Km{1100.0, 32.7, 55.5}

	copyLegs := legs

	distance = 1100.0
	length = 981.3

	for n, val := range copyLegs {
		fmt.Println(n, val)
	}

	legs[1]++
	fmt.Println(legs)
	fmt.Println(copyLegs)

	fmt.Println("Length is", len(copyLegs))

	if length == distance {
		fmt.Println("length and distance are equal")
	}

	fmt.Println(distance)
}

func main() {
	arrayChunk()
}
