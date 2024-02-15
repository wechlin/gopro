package main

import "fmt"

type Name string

type IsExactlyString interface {
	// A bare type means exactly that type
	string
}

type IsString interface {
	// ~ means "all types whose underlying type is the named type"
	~string
}

type IsStringOrInteger interface {
	// A pipe means this type OR that type
	~string | ~int
}

// Named constraints are fine
func Printer[T IsStringOrInteger](obj T) {
	fmt.Println(obj)
}

// Constraints do not need to be named (note dropping the `interface{}` syntax is allowed)
func Print2[T ~string | ~int](obj T) {
	fmt.Println(obj)
}

func main() {
	var hello Name = "Hello"
	var greet = 450

	// Common Interfaces
	// any: Literally anything
	// comparable: Types that can be compared with ==, !=
	// error: interface { Error() string }
	// fmt.Stringer() : interface { String() string }

	Printer(greet)
	Printer(hello)
}
