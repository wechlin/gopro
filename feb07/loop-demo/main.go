package main

import "fmt"

func main() {
	apostles := [...]string{"John", "Paul", "George", "Ringo"}

	// Equivalent to a while-loop from other languages
	countdown := 10
	for countdown > 0 {
		fmt.Println(countdown, "...")
		countdown--
	}
	fmt.Println("Blast off!")

	// For-range loop to iterate over a collection of items
	// First element is index, optional second element is the item
	for i, name := range apostles {
		fmt.Println(name)
		apostles[i] = name + " Beatle"
	}
	fmt.Println(apostles)

	// Bog-standard for loop: initializer, condition, and post-loop instruction
	for n := 1; n < 47; n *= 2 {
		fmt.Println(n)
	}

	// Infinite loop
	for {
		fmt.Println("Inside the inifinite loop!")
		break
	}
}
