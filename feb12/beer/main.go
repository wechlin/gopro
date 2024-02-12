// Beer will print the lyrics to "99 Bottles of Beer on the Wall".
// Invoke with -n <number> to start counting from number.
package main

import (
	"flag"
	"fmt"
)

var count = flag.Int("n", 99, "Number to start counting from")

func bottles(count int) string {
	switch count {
	case 1:
		return "1 bottle"
	case 0:
		return "No more bottles"
	default:
		return fmt.Sprintf("%d bottles", count)
	}
}

func main() {
	flag.Parse()
	for n := *count; n >= 1; n-- {
		fmt.Println(bottles(n), "of beer on the wall")
		fmt.Println(bottles(n), "of beer")
		fmt.Println("Take one down")
		fmt.Println("And pass it around")
		fmt.Print(bottles(n-1), " of beer on the wall!\n\n")
	}
}
