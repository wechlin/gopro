package main

import "fmt"

func fn() int {
	return 6
}

func main() {
	// switch may take an short variable initialization
	switch month := 6; month {
	// default case should be first or last
	default:
		fmt.Println("31 days")
	// Apparently a function call can be a case value, but don't
	// expect to be efficient (no jump table :(
	case 9, 4, fn(), 11:
		fmt.Println("30 days")
	case 2:
		fmt.Println("28 days, usually")

	}

	x, y := 5, 7

	switch {
	case x > y:
		fmt.Println("x > y")
	case x == y:
		fmt.Println("x = y")
	case x < y:
		fmt.Println("x < y")
	}
}
