// Hello is the hello world program that everyone starts with.
// Nearly every programmer starts their journey with this.
//
// This is a second paragraph.
package main

import "fmt"

// This is the value to count to.

var Value int = 3

const (
	Pi  = 3.14159265358979323846264338
	Tau = 2 * Pi
	Circle
)

const (
	Sunday = 1 << iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {
	// This is a comment

	fmt.Println(Thursday)
	fmt.Println(Tau, Circle)
	/*
	   fmt.Println("\u3046 Go world!", float64(Value))
	   letter := 'e'
	   apostles := [...]string{"John", "Paul", "George", "Ringo"}

	   apostles[0] = "John"
	   apostles[1] = "Paul"

	   fmt.Println(apostles)
	   fmt.Println(string(letter))
	   fmt.Println(`\u3046\tWorld

	   `)
	*/
}
