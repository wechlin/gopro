package main

import "fmt"

type (
	Km = float64
	Kg float64
)

func main() {
	var distance Km
	var length float64

	distance = 1100.0
	length = 981.3

	if length == distance {
		fmt.Println("length and distance are equal")
	}

	fmt.Println(distance)
}
