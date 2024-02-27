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

// Generic function gets square brackets after the name
func sum[T Numeric](a ...T) T {
	var total T
	for _, n := range a {
		total += n
	}
	return total
}

// Generic type gets square brackets after the name
// comparable may only be a constraint
type LinkedList[T comparable] struct {
	Data T
	Next *LinkedList[T]
}

// Generic type's methods just need the name, not the constraint
func (node *LinkedList[T]) String() string {
	if node == nil {
		return ""
	}
	return fmt.Sprintf("%v", node.Data) + node.Next.String()
}

func main() {
	var x int = 4
	var y Km = 5
	fmt.Println(sum(1.1, 2.7, 3.1, 4.159))
	fmt.Println(sum(1, 2, 3, x))
	fmt.Println(sum(y, y, y))

	// Instantiate a function using square brackets
	var adderFunction = sum[int]

	adderFunction(1, 2, 3)

	var phoneNumber LinkedList[uint]
	phoneNumber.Data = 8
	phoneNumber.Next = new(LinkedList[uint])
	phoneNumber.Next.Data = 6

	fmt.Println(phoneNumber)

	// `comparable` means works with == or !=, which slices do not
	//var illegal LinkedList[[]bool]
}
