package main

import "fmt"

// Constraints can be defined as interfaces
// ~ means "this type and anything with it as an underlying type"
// | means "or this constraint"
type Numeric interface {
	~int | ~uint | ~float64
}

// Km works because it passes constraint ~uint
type Km uint

func sum[T ~int | ~uint | ~float64](a ...T) T {
	var total T
	for _, n := range a {
		total += n
	}
	return total
}

func main() {
	var x int = 4
	var y Km = 5
	fmt.Println(sum(1.1, 2.7, 3.1, 4.159))
	fmt.Println(sum(1, 2, 3, x))
	fmt.Println(sum(y, y, y))
}
