package main

import "fmt"

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
	for n := 99; n >= 1; n-- {
		fmt.Println(bottles(n), "of beer on the wall")
		fmt.Println(bottles(n), "of beer")
		fmt.Println("Take one down")
		fmt.Println("And pass it around")
		fmt.Print(bottles(n-1), " of beer on the wall!\n\n")
	}
}
