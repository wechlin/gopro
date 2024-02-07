package main

import "fmt"

func fn() int {
	return 5
}

func main() {

	// Conditions do not go in parentheses
	// Additionally, conditions may start with an optional short variable definition
	if value := fn(); value > 10 {
		fmt.Println("Greater than 10")
	} else if value > 5 {
		fmt.Println("Greater than 5")
	} else {
		fmt.Println("Not greater than 10")
	}

	//ERROR, the variabl definition inside the if statement is scoped to that
	// if-block
	//fmt.Println(value)
}
