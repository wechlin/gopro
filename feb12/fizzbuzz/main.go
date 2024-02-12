package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s <number>\n", os.Args[0])
	}

	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("Not a number '%s': %v", os.Args[1], err)
	}

	for n := 1; n <= num; n++ {
		switch {
		case n%15 == 0:
			fmt.Println("FizzBuzz")
		case n%5 == 0:
			fmt.Println("Buzz")
		case n%3 == 0:
			fmt.Println("Fizz")
		default:
			fmt.Println(n)
		}
	}
}
